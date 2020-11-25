package routerhelper

import (
	"goshop/admin-api/pkg/core/ctl"
	"goshop/admin-api/pkg/utils"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Ctl ctl.ControllerInterface
	*gin.RouterGroup
}

func NewGroupRouter(groupName string, c ctl.ControllerInterface, g *gin.Engine, middleware ...gin.HandlerFunc) *Router {
	group := g.Group(groupName)
	r := &Router{
		Ctl:         c,
		RouterGroup: group,
	}

	if len(middleware) > 0 {
		group.Handlers = append(group.Handlers, middleware...)
	}

	return r
}

func NewRouter(c ctl.ControllerInterface, g *gin.Engine, middleware ...gin.HandlerFunc) *Router {
	r := &Router{
		Ctl:         c,
		RouterGroup: &g.RouterGroup,
	}

	if len(middleware) > 0 {
		r.Handlers = append(r.Handlers, middleware...)
	}

	return r
}

func (r *Router) getMethodName(name string) string {
	var newName string
	nameList := strings.Split(name, "/")
	if strings.Contains(nameList[len(nameList)-1], "-") {
		list := strings.Split(nameList[len(nameList)-1], "-")
		for k := range list {
			newName += utils.FirstLitterToUpper(list[k])
		}
	} else {
		newName = utils.FirstLitterToUpper(nameList[len(nameList)-1])
	}

	return newName
}

func (r *Router) getRealName(url string, method ...string) string {
	var methodName string
	if len(method) == 0 {
		methodName = r.getMethodName(url)
	} else {
		methodName = method[0]
	}

	return methodName
}

func (r *Router) Get(url string, method ...string) {
	r.GET(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Post(url string, method ...string) {
	r.POST(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Put(url string, method ...string) {
	r.PUT(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Patch(url string, method ...string) {
	r.PATCH(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Head(url string, method ...string) {
	r.HEAD(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Delete(url string, method ...string) {
	r.DELETE(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) Options(url string, method ...string) {
	r.OPTIONS(url, r.bindMethod(r.getRealName(url, method...)))
}

func (r *Router) bindMethod(methodName string) gin.HandlerFunc {
	if reflect.TypeOf(r.Ctl).Kind() != reflect.Ptr {
		panic("controller is not ptr type")
	}

	reflectValue := reflect.ValueOf(r.Ctl)
	method := reflectValue.MethodByName(methodName)
	if !method.IsValid() {
		log.Panicf("method name does not exist，controller: %s, method: %s", reflectValue.String(), methodName)
	}

	return func(c *gin.Context) {
		reflectValue = reflect.New(reflect.ValueOf(r.Ctl).Elem().Type())
		execController, ok := reflectValue.Interface().(ctl.ControllerInterface)
		if !ok {
			panic("controller is not ControllerInterface")
		}

		execController.Init(execController, c)
		execController.Prepare()
		reflectValue.MethodByName(methodName).Call(nil)
		execController.Finish()
	}
}
