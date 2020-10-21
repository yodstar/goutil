package restful

import (
	"reflect"
	"strings"
)

// router
var router map[string]reflect.Type = make(map[string]reflect.Type)

// Router
func Router(controller IController, path ...string) {
	rtype := reflect.Indirect(reflect.ValueOf(controller)).Type()
	if len(path) > 0 {
		router[strings.ToLower(path[0])] = rtype
	} else {
		router[strings.ToLower(strings.TrimSuffix(rtype.Name(), "Controller"))] = rtype
	}
}
