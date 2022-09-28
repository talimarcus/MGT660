package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var fortunes = []string{
	"Take a walk outside",
	"Hug a friend today",
	"Call your sisters to say hi",
	"Trust your decisions",
}

func main() {
	http.HandleFunc("/fortune", FortuneHandler)
	http.HandleFunc("/", HelloHandler) // when you get a request to then "/" url, then HelloHandler should handle the request
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func FortuneHandler(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(fortunes))
	fortune := fortunes[randomIndex]
	name := r.URL.Query().Get("name")
	htmlBody := `
		<html>
			<head>
				<link rel="stylesheet"
          		href="https://fonts.googleapis.com/css?family=Birthstone+Bounce|Roboto+Slab">
				<title>Fortune Teller</title>
				<style>
					body {
						background-color: cornflowerblue;
						color: navy;
						font-family: Roboto Slab;
						font-size: 20px;
						margin: 20 20;
					}
					h1 {
						font-family: Birthstone Bounce;
						font-size: 80px;
					}
					p {
						font-size: 40px;
						padding-left: 10px;
					}
				</style>
			</head>
			<body>
				<h1>Your Fortune</h1>
				<h2>A personalized fortune for %s</h2>
				<p>"%s"</p>
			</body>
	`
	fmt.Fprintf(w, htmlBody, name, fortune)
}
