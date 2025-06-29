package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	conf := flag.String("c", "./gohttpd.conf", "Config file path")
	flag.Parse()

	// chroot
	if exefile, err := os.Executable(); err != nil {
		panic(err.Error())
	} else if err = os.Chdir(filepath.Dir(exefile)); err != nil {
		panic(err.Error())
	}

	// config
	data, err := ioutil.ReadFile(*conf)
	if err != nil {
		panic(err.Error())
	}

	config := struct {
		Routes map[string]string
		Listen string
	}{}
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}

	// routes
	for route, path := range config.Routes {
		if strings.HasSuffix(route, "*") {
			f := path
			d := filepath.Dir(path)
			p := strings.TrimSuffix(route, "*")
			http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
				file := filepath.Join(d, strings.TrimPrefix(r.URL.Path, p))
				if info, _ := os.Stat(file); info != nil && info.Mode().IsRegular() {
					http.ServeFile(w, r, file)
				} else {
					http.ServeFile(w, r, f)
				}
			})
		} else {
			http.Handle(route, http.StripPrefix(route, http.FileServer(http.Dir(path))))
		}
	}

	// listen
	if err := http.ListenAndServe(config.Listen, nil); err != nil {
		panic(err.Error())
	}
}
