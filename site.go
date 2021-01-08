package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	DataPath string `envconfig:"KO_DATA_PATH" default:"/var/run/ko/" required:"true"`
	WWWPath string `envconfig:"WWW_PATH" default:"www" required:"true"`
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
	m.Handle("/", http.FileServer(http.Dir(www)))

	log.Println("Listening on", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), m))
}
