package main

import  (
	"html/template"
	"fmt"
	"net/http"
	"path"
//        "gopkg.in/mgo.v2"
//        "gopkg.in/mgo.v2/bson"
        "time"
)
// ****************************************************************
// * DATUMS
// ****************************************************************
type User struct {
        Id string 
        Name string
        Avatar string
}

type Stats struct {
        Viewed int64
        TotalStay time.Time
        AvgStay time.Time
        Offers int64
}

type Visitor struct {
        Start time.Time
        User int64

}
type Offer struct {
        LastAction time.Time
        For int64
        From int64
}
// ****************************************************************
// * CONTROLLERS
// ****************************************************************
func Index(w http.ResponseWriter, r *http.Request){

	r.ParseForm()

	user := User{"0","jim","http://i.imgur.com/C1hvIZ6.jpg"}
	fmt.Println("New User Connected")
	
	//Reder Index
	fp := path.Join("assets", "html", "index.html")
	templ, err := template.ParseFiles(fp)

	//Catch Errors
	err = templ.Execute(w,user)
	if err != nil { fmt.Println("error occurred")}
}
// *****************************************************************
// * MAIN
// *****************************************************************
func main(){
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {fmt.Println("error brother")}
}



