package engine

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"}) && ok
	ok = reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"}) && ok
	if !ok {
		t.Fatal("TestParsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("GET", "/hello/soarkey")
	if n == nil {
		t.Fatal("n is nil")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("match /hello/:name failed")
	}
	if params["name"] != "soarkey" {
		t.Fatal("name not equals soarkey")
	}
	fmt.Printf("matched path: %s, params['name'] = %s\n", n.pattern, params["name"])
}
