package main

import  (
	"html/template"
	"fmt"
	"net/http"
	"path"
	"labix.org/v2/mgo.v2.a"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id 'bson:"_id,omitempty"'
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

type Doc struct {
	ID bson.ObjectId 'bson:"_id,omitempty"'
	Type string
	Timestamp time.Time
	Owner int64
	Watchers []int64
	Offers Offer
	Stats Stats
}

func Index(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user := User{0,"anon","http://www.imgur.com/Ot13hl"}
	fp := path.Join("assets", "html", "index.html")
	templ, err := template.ParseFiles(fp)
	err = templ.Execute(w,user)
	if err != nil { fmt.Println("error occurred")}
}

func main(){
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
