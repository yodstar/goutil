package restful

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

// HookHandler
var hookHandler http.HandlerFunc = nil

// Handle
func Handle(hook http.HandlerFunc) {
	hookHandler = hook
}

// httpHandler
type httpHandler struct {
	root string
	pool sync.Pool
}

// ServeHTTP implements the http.Handler interface for a HttpServe
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hookHandler != nil {
		hookHandler(w, r)
	}

	var controllerName, actionName string
	pathInfo := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathInfo) > 0 {
		controllerName = strings.ToLower(pathInfo[0])
	}
	if controllerName == "" {
		controllerName = "index"
	}

	controllerT, ok := router[controllerName]
	if !ok {
		h.ServeFile(w, r)
		return
	}

	// Action
	if len(pathInfo) > 1 {
		a := strings.Split(pathInfo[1], "-")
		for i, v := range a {
			a[i] = strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
		}
		actionName = strings.Join(a, "")
	}
	if actionName == "" {
		actionName = "Index"
	}

	refV := reflect.New(controllerT)
	action := refV.MethodByName(actionName + "Action")
	if !action.IsValid() {
		if action = refV.MethodByName("ErrorAction"); !action.IsValid() {
			h.ServeFile(w, r)
			return
		}
	}

	// Context
	ctx := h.pool.Get().(*Context)
	defer h.pool.Put(ctx)
	ctx.Init(w, r, h.root)

	ctx.Controller = controllerName
	ctx.Action = actionName

	// Params/Handle
	if len(pathInfo) > 2 {
		count := len(pathInfo)
		for i := 2; i < count; i += 2 {
			if i+1 < count {
				ctx.Params[pathInfo[i]] = pathInfo[i+1]
			} else {
				ctx.Params[pathInfo[i]] = ""
			}
		}
		ctx.Handle = pathInfo[2]
		ctx.PathVar = strings.Join(pathInfo[2:], "/")
	}

	// Controller
	controller := refV.Interface().(IController)
	controller.New(ctx)
	controller.Init()
	if controller.OnBefore() {
		action.Call(nil)
	}
	controller.OnAfter()
	controller.Destroy()
}

// ServeFile
func (h *httpHandler) ServeFile(w http.ResponseWriter, r *http.Request) {
	if path.Ext(r.URL.Path) != "" {
		http.ServeFile(w, r, h.root+r.URL.Path)
	} else {
		file := fmt.Sprintf("%s/%s/index.html", h.root, strings.Trim(r.URL.Path, "/"))
		if _, err := os.Stat(file); err == nil {
			http.ServeFile(w, r, file)
		} else {
			http.NotFound(w, r)
		}
	}
}

// ListenTLS
func ListenTLS(root, addr, cert, key string) {
	var err error
	// Handler
	h := &httpHandler{}
	if h.root, err = filepath.Abs(root); err != nil {
		log.Fatal(err.Error())
	}
	h.pool.New = func() interface{} {
		return &Context{}
	}
	// Server
	s := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	log.Fatal(s.ListenAndServeTLS(cert, key))
}

// Listen
func Listen(root, addr string) {
	var err error
	// Handler
	h := &httpHandler{}
	if h.root, err = filepath.Abs(root); err != nil {
		log.Fatal(err.Error())
	}
	h.pool.New = func() interface{} {
		return &Context{}
	}
	// Server
	s := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	log.Fatal(s.ListenAndServe())
}
