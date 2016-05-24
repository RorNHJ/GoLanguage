package main
 
import (
    "fmt"
    "html/template"

    "net/http"
    "gopkg.in/mgo.v2"

)
var mainTemplate, _ = template.ParseFiles("html/test.html")

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}
 
func mainPage(w http.ResponseWriter, r *http.Request) {
    mainTemplate.Execute(w, nil)
}
 
func main() {
    http.HandleFunc("/login", login)
    http.HandleFunc("/", mainPage)
    http.ListenAndServe(":8080", nil) 
  
  }



