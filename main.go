package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)


func main()  {
	type Email struct {
		Title string
		Subject string
		To string
		Message string
		Greeting string
	}
		Info:= Email {
		Title:    "Demande Avis",
		Subject:  " Consultant",
		To:       "chker@gmail.com",
		Message:  "votre avis concernant dossier de Mr.Mouhamed Awina ",
		Greeting: "Bonjour",
	}

	templ,err := template.ParseFiles("./hello.html")

	if err != nil {
		fmt.Println(err)
	}
	err1 := templ.Execute(os.Stdout , Info)
	if err != nil {
		log.Fatal(err1)
	}
	


}