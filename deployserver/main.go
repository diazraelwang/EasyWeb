package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)


func main(){
	http.HandleFunc("/",webPage)
	http.ListenAndServe(":6000",nil)
}

func rebootweb(){
	cmd := exec.Command("sh","./deploy.sh") //在源码中调用deploy.sh
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func webPage(wt http.ResponseWriter,rt *http.Request){
	io.WriteString(wt, "<h1>Deploy server on port 6000 has been started!</h1>")
	rebootweb()
}
