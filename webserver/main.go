package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",webPage)
	http.ListenAndServe(":8086",nil)
}

func webPage(wt http.ResponseWriter,rt *http.Request){
	io.WriteString(wt, "<h2>The Web service on port 8086 has been changed!</h2>")
}