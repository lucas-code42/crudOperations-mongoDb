package controller

import "net/http"

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	// panic("not implemented")
	w.Write([]byte("hello from GET"))
}
