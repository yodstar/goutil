package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	file := flag.String("c", "./gohttpd.conf", "Config file path")
	flag.Parse()

	// chroot
	if exefile, err := os.Executable(); err != nil {
		panic(err.Error())
	} else if err = os.Chdir(filepath.Dir(exefile)); err != nil {
		panic(err.Error())
	}

	// config
	data, err := ioutil.ReadFile(*file)
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
		http.Handle(route, http.StripPrefix(route, http.FileServer(http.Dir(path))))
	}

	// listen
	if err := http.ListenAndServe(config.Listen, nil); err != nil {
		panic(err.Error())
	}
}
