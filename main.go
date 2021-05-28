/*
   Copyright 2021 Robert Burghart

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server.")
	defer log.Println("Stopping server.")

	http.Handle("/", http.FileServer(http.Dir("content")))
	http.HandleFunc("/template", renderTemplate)
	http.HandleFunc("/api/thing", renderThing)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var (
	siteTemplates = template.Must(template.ParseFiles("templates/test_001.template", "templates/test_002.template"))
)

func renderTemplate(w http.ResponseWriter, _ *http.Request) {
	if err := siteTemplates.ExecuteTemplate(w, "test_002.template", nil); err != nil {
		log.Println(err)
	}
}

func renderThing(w http.ResponseWriter, _ *http.Request) {
	if _, err := io.WriteString(w, `{ "test": 12 }`); err != nil {
		log.Println(err)
	}
}
