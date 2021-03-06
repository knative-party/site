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
	"gopkg.in/yaml.v2"

	"github.com/knative-party/site/rotation"
)

type envConfig struct {
	DataPath string `envconfig:"KO_DATA_PATH" default:"/var/run/ko/" required:"true"`
	WWWPath  string `envconfig:"WWW_PATH" default:"www" required:"true"`
	Port     int    `envconfig:"PORT" default:"8080" required:"true"`
}

var env envConfig

func bestDate(r *http.Request) time.Time {
	on := time.Now()
	keys, ok := r.URL.Query()["on"]
	if ok && len(keys) == 1 {
		if then, err := time.Parse(time.RFC3339, keys[0]); err != nil {
			// try simple date.
			if then, err := time.Parse("2006-01-02", keys[0]); err != nil {
				log.Println("invalid date:", keys[0])
			} else {
				on = then
				on = on.Add(time.Hour * (7 + 9)) // 9am PST
			}
		} else {
			on = then
		}
	}

	// it is always monday.
	switch on.Weekday() {
	case time.Sunday:
		on = on.Add(time.Hour * 12 * 1)
	case time.Monday:
		// ok
	case time.Tuesday:
		on = on.Add(-1 * time.Hour * 12 * 1)
	case time.Wednesday:
		on = on.Add(-1 * time.Hour * 12 * 2)
	case time.Thursday:
		on = on.Add(-1 * time.Hour * 12 * 3)
	case time.Friday:
		on = on.Add(-1 * time.Hour * 12 * 4)
	case time.Saturday:
		on = on.Add(time.Hour * 12 * 2)
	}

	log.Println("loading date:", on.String())
	return on
}

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
		on := bestDate(r)
		out := json.NewEncoder(w)
		out.Encode(getNow(on))
	})
	m.Handle("/", http.FileServer(http.Dir(www)))

	log.Println("Listening on", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), m))
}

func getNow(on time.Time) (now Now) {
	configPath := filepath.Join(env.DataPath, "config.yaml")
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("unable to read file %q: %s", configPath, err)
		return
	}
	c := config{}
	if err := yaml.Unmarshal(file, &c); err != nil {
		log.Printf("unable to parse %q: %s", configPath, err)
	}
	eC := make([]chan event, 0, len(c.Events))
	tC := make([]chan tier, 0, len(c.Rotations))

	for _, url := range c.Events {
		c := make(chan event)
		eC = append(eC, c)
		go loadEvent(url, on, c)
	}
	for _, url := range c.Rotations {
		c := make(chan tier)
		tC = append(tC, c)
		go loadTier(url, on, c)
	}

	for _, c := range eC {
		for e := range c {
			now.Events = append(now.Events, e)
		}
	}
	for _, c := range tC {
		for t := range c {
			now.Tiers = append(now.Tiers, t)
		}
	}

	return now
}

func loadEvent(url string, on time.Time, c chan event) {
	defer close(c)
	r, err := rotation.FromURL(url)
	if err != nil {
		log.Printf("Unable to read %q: %s", url, err)
		return
	}
	duration, err := time.ParseDuration(r.Metadata["duration"])
	if err != nil {
		log.Printf("Unable to parse duration from %q: %s", url, err)
		duration = 1 * time.Hour
	}
	entry := r.Next(on)
	end := entry.Start.Add(duration)
	c <- event{
		Title:        r.Metadata["title"],
		WorkingGroup: strings.Join(entry.Data, " "),
		When:         entry.Start.Format("Jan 2, 2006 @ 15:04") + " - " + end.Format("15:04 MST"),
	}
}

func loadTier(url string, on time.Time, c chan tier) {
	defer close(c)
	r, err := rotation.FromURL(url)
	if err != nil {
		log.Printf("Unable to read %q: %s", url, err)
		return
	}
	entry := r.At(on)
	c <- tier{
		Title: r.Metadata["title"],
		OnCall: onCall{
			Name:           entry.Data[0],
			Start:          entry.Start.Format("Jan 2, 2006"),
			End:            entry.End.Format("Jan 2, 2006"),
			Github:         "https://github.com/" + entry.Data[0],
			Questions:      r.Metadata["slack"],
			QuestionsSlack: r.Metadata["slacklink"],
		},
	}
}

// TODO: add a role button for a job description.

type config struct {
	Events    []string
	Rotations []string
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
