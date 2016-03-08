package main

import  (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hiya")
}

func main(){
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
