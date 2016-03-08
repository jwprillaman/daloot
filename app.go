package main

import  (
	"html/template"
	"fmt"
	"net/http"
	"path"
)
type User struct {
	id int
	name string
}
func Index(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user := User{0,"jim"}
	fmt.Println("New User Connected")
	fp := path.Join("assets", "html", "index.html")
	templ, err := template.ParseFiles(fp)
	err = templ.Execute(w,user)
	if err != nil { fmt.Println("error occurred")}
}

func main(){
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
