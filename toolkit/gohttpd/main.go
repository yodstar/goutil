package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		panic(err)
	} else if err = os.Chdir(filepath.Dir(file)); err != nil {
		panic(err)
	} else {
		name = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		path := flag.String("c", fmt.Sprintf("./%s.conf", name), "Config file path")
		flag.Parse()

		data, err := os.ReadFile(*path)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(data, &conf); err != nil {
			panic(err)
		}
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// routes
	rewrite := make(map[string]string)
	for route, path := range conf.Routes {
		if strings.HasSuffix(route, "*") {
			prefix := strings.TrimSuffix(route, "*")
			rewrite[prefix] = path
			http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
				path := rewrite[prefix]
				file := filepath.Join(filepath.Dir(path), strings.TrimPrefix(r.URL.Path, prefix))
				if info, _ := os.Stat(file); info != nil && info.Mode().IsRegular() {
					http.ServeFile(w, r, file)
				} else {
					http.ServeFile(w, r, path)
				}
			})
		} else if strings.HasPrefix(path, "http://") {
			target, err := url.Parse(path)
			if err != nil {
				panic(err.Error())
			}
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ModifyResponse = func(resp *http.Response) error {
				resp.Header.Set("Access-Control-Allow-Origin", "*")
				resp.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				resp.Header.Set("Access-Control-Allow-Headers", "Content-Type")
				return nil
			}
			http.Handle(route, proxy)
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
