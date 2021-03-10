package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	now := Now{}

	eventDir := filepath.Join(env.DataPath, "events")
	events, err := ioutil.ReadDir(eventDir)
	if err != nil {
		log.Printf("Unable to open %q: %s", eventDir, err)
	}
	for _, f := range events {
		path := filepath.Join(eventDir, f.Name())
		r, err := rotation.RotationFromFile(path)
		if err != nil {
			log.Printf("Unable to read %q: %s", path, err)
		}
		duration, err := time.ParseDuration(r.Metadata["duration"])
		if err != nil {
			log.Printf("Unable to parse duration from %q: %s", path, err)
			duration = 1 * time.Hour
		}
		entry := r.Next(time.Now())
		end := entry.Start.Add(duration)
		now.Events = append(now.Events, event{
			Title:        r.Metadata["title"],
			WorkingGroup: strings.Join(entry.Data, " "),
			When:         entry.Start.Format("March 2, 2006 @ 15:04") + " - " + end.Format("15:04 MST"),
		})
	}

	// Add from rotations
	rotDir := filepath.Join(env.DataPath, "rotations")
	rotations, err := ioutil.ReadDir(rotDir)
	if err != nil {
		log.Printf("Unable to open %q: %s", rotDir, err)
	}
	for _, f := range rotations {
		path := filepath.Join(rotDir, f.Name())

		r, err := rotation.RotationFromFile(path)
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
