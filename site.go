package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	DataPath string `envconfig:"KO_DATA_PATH" default:"/var/run/ko/" required:"true"`
	WWWPath  string `envconfig:"WWW_PATH" default:"www" required:"true"`
	Port     int    `envconfig:"PORT" default:"8080" required:"true"`
}

func main() {
	var env envConfig
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
	return Now{
		Tiers: []tier{{
			Title: "Serving", // https://github.com/knative/serving/blob/master/support/COMMUNITY_CONTACTS.md
			OnCall: onCall{
				Name:           "@dprotaso",
				Start:          "March 1, 2021",
				End:            "March 5, 2021",
				Github:         "https://github.com/dprotaso",
				Questions:      "#serving-questions",
				QuestionsSlack: "https://knative.slack.com/archives/C0186KU7STW",
			},
		}, {
			Title: "Eventing", // https://github.com/knative/eventing/blob/master/support/COMMUNITY_CONTACTS.md
			OnCall: onCall{
				Name:           "@pierDipi",
				Start:          "March 1, 2021",
				End:            "March 5, 2021",
				Github:         "https://github.com/pierDipi",
				Questions:      "#eventing-questions",
				QuestionsSlack: "https://knative.slack.com/archives/C017X0PFC0P",
			},
		}},
		Events: []event{{
			Title:        "ToC Working Group Update", // https://docs.google.com/document/d/1LzOUbTMkMEsCRfwjYm5TKZUWfyXpO589-r9K2rXlHfk/edit#heading=h.jlesqjgc1ij3
			WorkingGroup: "Networking WG",            // https://github.com/knative/community/blob/master/working-groups/WORKING-GROUPS.md
			When:         "March 4, 2021 @ 8:30 – 9:15am PST",
		}},
	}
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
