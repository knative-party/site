package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/knative-party/site/rotation"
)

type envConfig struct {
	DataPath string `envconfig:"KO_DATA_PATH" default:"/var/run/ko/" required:"true"`
	WWWPath  string `envconfig:"WWW_PATH" default:"www" required:"true"`
	Port     int    `envconfig:"PORT" default:"8080" required:"true"`
}

var env envConfig

func main() {
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("failed to process env var: %s", err)
	}

	www := path.Join(env.DataPath, env.WWWPath)
	if !strings.HasSuffix(www, "/") {
		www = www + "/"
	}

	m := http.NewServeMux()
	m.HandleFunc("/now", func(w http.ResponseWriter, r *http.Request) {
		out := json.NewEncoder(w)
		out.Encode(getNow())
	})
	m.Handle("/", http.FileServer(http.Dir(www)))

	log.Println("Listening on", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), m))
}

// TODO: add a role button for a job description.

func getNow() Now {
	now := Now{
		// TODO: use rotations for events, too.
		Events: []event{{
			Title:        "ToC Working Group Update", // https://docs.google.com/document/d/1LzOUbTMkMEsCRfwjYm5TKZUWfyXpO589-r9K2rXlHfk/edit#heading=h.jlesqjgc1ij3
			WorkingGroup: "Networking WG",            // https://github.com/knative/community/blob/master/working-groups/WORKING-GROUPS.md
			When:         "March 4, 2021 @ 8:30 – 9:15am PST",
		}},
	}

	// Add from rotations
	rotDir := filepath.Join(env.DataPath, "rotations")
	files, err := ioutil.ReadDir(rotDir)
	if err != nil {
		log.Printf("Unable to open %q: %s", rotDir, err)
	}
	for _, f := range files {
		path := filepath.Join(rotDir, f.Name())
		f, err := os.Open(path)
		if err != nil {
			log.Printf("Unable to read %q: %s", path, err)
			continue
		}
		r, err := rotation.ReadFile(f)
		if err != nil {
			log.Printf("Unable to read %q: %s", path, err)
		}
		entry := r.At(time.Now())

		now.Tiers = append(now.Tiers, tier{
			Title: r.Metadata["title"],
			OnCall: onCall{
				Name:           entry.Data[0],
				Start:          entry.Start.Format("March 2, 2006"),
				End:            entry.End.Format("March 2, 2006"),
				Github:         "https://github.com/" + entry.Data[0],
				Questions:      r.Metadata["slack"],
				QuestionsSlack: r.Metadata["slacklink"],
			},
		})
	}

	return now
}

type tier struct {
	Title  string `json:"title"`
	OnCall onCall `json:"onCall"`
}

type onCall struct {
	Name           string `json:"name"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Github         string `json:"github"`
	Questions      string `json:"questions"`
	QuestionsSlack string `json:"questionsSlack"`
}

type event struct {
	Title        string `json:"title"`
	WorkingGroup string `json:"wg"`
	When         string `json:"when"`
}

type Now struct {
	Tiers  []tier  `json:"support"`
	Events []event `json:"events"`
}
