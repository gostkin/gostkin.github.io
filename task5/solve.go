package main

import "encoding/json"
import "io/ioutil"
import "log"
import "net/http"
import "strconv"

import "github.com/gorilla/mux"

var urls map[string]string

func nextKey() string {
	return strconv.Itoa(len(urls))
}

func getRedirect(writer http.ResponseWriter, request *http.Request) {
	key := mux.Vars(request)["key"]
	url, res := urls[key]
	
	if !res {
		http.NotFound(writer, request)
		return
	}
	
	http.Redirect(writer, request, url, 301)
}

func addLink(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	
	var body []byte
	body, _ = ioutil.ReadAll(request.Body)
	
	var rjson map[string]string
	if error := json.Unmarshal(body, &rjson); error != nil {
		panic(error)
	}
	
	url := rjson["url"]
	
	key := nextKey()
	urls[key] = url

	resjson := map[string]string {"key": key}
	result, _ := json.Marshal(resjson)
	writer.Write(result)
}

func main() {
	urls = make(map[string]string)

	r := mux.NewRouter()
  
	r.HandleFunc("/{key}", getRedirect).Methods("GET")
	r.HandleFunc("/", addLink).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", r))
}
