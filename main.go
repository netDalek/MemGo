package main

import "fmt"
import "net/http"
import "log"
import "html"
import "io/ioutil"

var storage map[string]string
const HandlePath = "/memgo/"

func main() {
  storage = make(map[string]string)
  storage["v"] = "MemGo 0.1"
  http.HandleFunc(HandlePath, httpHandler)

  log.Print("start")
  log.Fatal(http.ListenAndServe(":8080", nil))
  log.Print("end")
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
  var url = html.EscapeString(r.URL.Path)
  var phase = url[len(HandlePath):]
  log.Printf("Request %q with method %q", phase, r.Method)
  switch r.Method {
    case "GET":
      fmt.Fprint(w, storage[phase])
    case "POST":
      data, error := ioutil.ReadAll(r.Body)
      if error == nil {
        storage[phase] = string(data)
      } else {
        log.Print(error)
      }
  }
}
