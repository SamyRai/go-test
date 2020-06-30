package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	name := req.URL.Query().Get("name")
	io.WriteString(res, fmt.Sprintf(`<doctype html>
<html>
	<head>
    	<title>Hello World!</title>
	</head>
	<body>
    	Hello %s!
	</body>
</html>`, name))
}

func helloTemplate(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	tmpl, err := template.ParseFiles("templates/greetings.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
		Guest string
	}{"Greetings Page", name}
	tmpl.Execute(res, data)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello_template", helloTemplate)

	http.ListenAndServe(":8085", nil)
}
