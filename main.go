package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.Handle("/minecraft/notify", NewEmail())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Email struct {
}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var msg Notice
	err := json.NewDecoder(req.Body).Decode(&msg)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	cmd := exec.Command("mutt", "-F", "/home/jeremys/.httpemaild.muttrc", msg.Recipient)
	cmd.Stdin = strings.NewReader(msg.Message)
	err = cmd.Run()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

type Notice struct {
	Recipient string
	Message   string
}
