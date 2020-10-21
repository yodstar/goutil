# restful
RESTful Framework for Golang

## Example:

- main.go

```
import (
	...

	"github.com/yodstar/goutil/restful"
)

func main() {
	restful.Router(&controller.IndexController{})

	if CONF.Certificate == "" || CONF.PrivateKey == "" {
		restful.Listen(CONF.RootDir, CONF.Listen)
		return
	}

	restful.ListenTLS(CONF.RootDir, CONF.Listen, CONF.Certificate, CONF.PrivateKey)
}


```

- controller/index.go

```
package controller

import (
	...

	"github.com/yodstar/goutil/restful"
)


type IndexController struct {
	restful.Controller
}

func (c *IndexController) IndexAction() {
	c.WriteString("[Hello RESTful]")
}


```