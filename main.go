package main

import (
    "encoding/json"
	"fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
type Payload struct {
    Stuff Data
}
type Data struct {
    Fruit Fruits
}
type Fruits map[string]int

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/list", GetList).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!!!")
	fruits := make(map[string]int)
    fruits["Apples"] = 25
    fruits["Oranges"] = 10
    w.Header().Set("Content-Type", "application/json")
    // get a payload
    d := Data{fruits} 
    p := Payload{d}
	
	json.NewEncoder(w).Encode(p)
}