package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"


	verify "github.com/GhvstCode/vonage/utils"
)

type Message struct {
	Phone string
	Id 	  string
}

func formHandler (w http.ResponseWriter, r *http.Request){
		if err := r.ParseForm(); err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		Phone :=  r.FormValue("phone")
		Id := verify.VerStart(Phone)
		msg := &Message{
			Id: Id,
			Phone: Phone,
		}

		render(w, "./static/form.html", msg)
}
//interface. 
func confirmHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	Id := r.FormValue("requestId")
	phone := r.FormValue("phone")

	codeChan := make(chan string)
	go func() {
		codeChan <- r.FormValue("confirmation")
	}()
	confirmation := <- codeChan

	response := verify.VerCheck(Id, confirmation)
	if response != "0" {
		fmt.Fprint(w,"Verification failed! Input the correct code sent to ", phone)
		return
	}

	fmt.Fprint(w,"🎉 Success! 🎉")
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/confirm", confirmHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
