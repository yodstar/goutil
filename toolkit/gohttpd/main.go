package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var name string
	conf := struct {
		Routes map[string]string
		Listen string
	}{}

	//chroot
	if file, err := os.Executable(); err != nil {
		panic(err.Error())
	} else if err = os.Chdir(filepath.Dir(file)); err != nil {
		panic(err.Error())
	} else {
		name = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		path := flag.String("c", fmt.Sprintf("./%s.conf", name), "Config file path")
		flag.Parse()

		data, err := os.ReadFile(*path)
		if err != nil {
			panic(err.Error())
		}

		if err := json.Unmarshal(data, &conf); err != nil {
			panic(err.Error())
		}
	}

	// routes
	for route, path := range conf.Routes {
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
	log.Printf("%s listen on %s", name, conf.Listen)
	if err := http.ListenAndServe(conf.Listen, nil); err != nil {
		panic(err.Error())
	}
}
