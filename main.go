package main

import (
    "encoding/json"
	"fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lunaimaginea/rest_api_test/db"
    "github.com/lunaimaginea/rest_api_test/user"
)


// our main function
func main() {
    router := mux.NewRouter()
    // router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users", CreateUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
    
}

func CreateUser(w http.ResponseWriter, r *http.Request){
    decoder := json.NewDecoder(r.Body)

    var u user.Users

    err := decoder.Decode(&u)
    if err != nil {
        panic(err)
    }
   
    db_con, err := db.Connection()
    db.InitializeDb(db_con)
    fmt.Println(db_con)
    user.Create(db_con, &u)

    if err != nil {
        log.Fatal(err)
    }

    /*defer db.CloseConnection(db_con)
    db.InitializeDb(db_con)
    
    // fmt.Println(u)
    user.Create(db_con, &u)
    // json.NewEncoder(w).Encode(u)*/
    
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
    /*decoder := json.NewDecoder(r.Body)
    var u user.User
    err := decoder.Decode(&u)
    if err != nil {
        panic(err)
    }
    log.Println(u)
    db_con, err := db.Connection()

    if err != nil {
        log.Fatal(err)
    }

    defer db.CloseConnection(db_con)
    db_con.AutoMigrate(&user.User{})
    // sampleUser := user.User{Name: "Luna"}
    user.Create(db_con, &u)
    fmt.Println(u.ID)
	/*fmt.Println("Hello world!!!")
	fruits := make(map[string]int)
    fruits["Apples"] = 25
    fruits["Oranges"] = 10
    w.Header().Set("Content-Type", "application/json")
    // get a payload
    d := Data{fruits} 
    p := Payload{d}
	
	json.NewEncoder(w).Encode(p)*/
}