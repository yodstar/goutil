package restful

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

const (
	ControllerName = "index"
	ActionName     = "Index"
)

// Context
type Context struct {
	Controller     string
	Action         string
	Params         map[string]string
	Handle         string
	PathVar        string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	RootDir        string
	QueryValues    url.Values
}

// Context.Init
func (c *Context) Init(w http.ResponseWriter, r *http.Request, d string) {
	c.Controller = ControllerName
	c.Action = ActionName
	c.Params = make(map[string]string)
	c.Handle = ""
	c.PathVar = ""
	c.ResponseWriter = w
	c.Request = r
	c.RootDir = d
	c.QueryValues, _ = url.ParseQuery(r.URL.RawQuery)
}

// IController
type IController interface {
	New(ctx *Context)
	Init()
	OnBefore() bool
	OnAfter()
	Destroy()
}

// Controller
type Controller struct {
	Ctx  *Context
	Data map[string]interface{}
	Func template.FuncMap
}

// Controller.New
func (c *Controller) New(ctx *Context) {
	c.Ctx = ctx
	c.Data = make(map[string]interface{})
	c.Func = make(template.FuncMap)
}

// Controller.Init
func (c *Controller) Init() {

}

// Controller.OnBefore
func (c *Controller) OnBefore() bool {
	return true
}

// Controller.OnAfter
func (c *Controller) OnAfter() {

}

// Controller.Destroy
func (c *Controller) Destroy() {

}

// Controller.Get
func (c *Controller) Get(name string) string {
	return c.Ctx.QueryValues.Get(name)
}

// Controller.Post
func (c *Controller) Post(name string) string {
	return c.Ctx.Request.PostFormValue(name)
}

// Controller.Param
func (c *Controller) Param(name string) string {
	if value, ok := c.Ctx.Params[name]; ok {
		return value
	}
	return ""
}

// Controller.NotFound
func (c *Controller) NotFound() {
	http.NotFound(c.Ctx.ResponseWriter, c.Ctx.Request)
}

// Controller.Redirect
func (c *Controller) Redirect(uri string) {
	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, uri, http.StatusFound)
}

// Controller.SetHeader
func (c *Controller) SetHeader(key, value string) {
	c.Ctx.ResponseWriter.Header().Set(key, value)
}

// Controller.WriteString
func (c *Controller) WriteString(data string) (err error) {
	_, err = c.Ctx.ResponseWriter.Write([]byte(data))
	return err
}

// Controller.WriteJSON
func (c *Controller) WriteJSON(v interface{}) (err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.SetHeader("Content-Length", fmt.Sprintf("%d", len(data)))
	c.SetHeader("Content-Type", "application/json;charset=UTF-8")
	_, err = c.Ctx.ResponseWriter.Write(data)
	return err
}

// Controller.WriteHTML
func (c *Controller) WriteHTML(html ...string) (err error) {
	data := []byte(strings.Join(html, ""))
	c.SetHeader("Content-Length", fmt.Sprintf("%d", len(data)))
	c.SetHeader("Content-Type", "text/html;charset=UTF-8")
	_, err = c.Ctx.ResponseWriter.Write(data)
	return err
}

// Controller.Write
func (c *Controller) Write(data []byte) (err error) {
	_, err = c.Ctx.ResponseWriter.Write(data)
	return err
}

// Controller.Display
func (c *Controller) Display(tpls ...string) {
	if len(tpls) == 0 {
		return
	}

	name := filepath.Base(tpls[0])
	for i, v := range tpls {
		var path string
		if i == 0 {
			path = c.Ctx.RootDir + "/" + c.Ctx.Controller + "/" + v
		} else {
			path = c.Ctx.RootDir + "/" + v
		}
		tpls[i] = strings.Replace(path, "\\", "/", -1)
	}
	tmpl := template.Must(template.New(c.Ctx.Controller).Funcs(c.Func).ParseFiles(tpls...))
	if err := tmpl.ExecuteTemplate(c.Ctx.ResponseWriter, name, c.Data); err != nil {
		log.Println(err.Error())
	}
}
