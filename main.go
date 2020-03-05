package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func commandHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		cmd, err := exec.Command("/bin/sh", "-c", "mysql -u root -ppassword -e 'use corn;drop table corn_table'").Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, string(cmd))
	}
	fmt.Fprintf(w, "Test")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/commands", commandHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
