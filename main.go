package main
 
import (
	    "net/http"
		"fmt"
	   )
 
func main() {
	http.Handle("/", new(testHandler))
    fmt.Println("start")					 
	http.ListenAndServe(":5000", nil)
}
 
type testHandler struct {
		    http.Handler
}
 
func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Server Running"
	w.Write([]byte(str))
}
