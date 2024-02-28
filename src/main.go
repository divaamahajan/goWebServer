	package main

	import (
		"fmt"
		"html/template"
		"log"
		"net/http"
	)

	// Define a struct to hold the data to be passed to the template
	type HelloData struct {
		Name    string
		Address string
	}


	var (
		name    string
		address string
	)

	func submitHandler(w http.ResponseWriter, r *http.Request){
		// w io.Writer
		if r.Method != http.MethodPost {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("ParseForm error: %v", err), http.StatusBadRequest)
			return
		}

		// fmt.Fprintf(w, "POST request successful!\n")
		name = r.FormValue("name")
		address = r.FormValue("address")

		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)
		
		// call hello.html and pass name and address to it
		
		// Create an instance of the HelloData struct with the name and address
		data := HelloData{
			Name:    name,
			Address: address,
		}

		// Parse the hello.html template file
		tmpl, err := template.ParseFiles("static/hello.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
			return
		}

		// Execute the template, passing the data to it
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
			return
		}
	}

	func formHandler(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/form" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
			return
		}

		http.ServeFile(w, r, "static/form.html")	
	}

	func helloHandler(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintf(w, "Hello, %s! You are from %s", name, address)
	}

	func main() {
		fileserver := http.FileServer(http.Dir("static"))
		
		http.Handle("/", fileserver)
		http.HandleFunc("/submit", submitHandler)
		http.HandleFunc("/form", formHandler)
		http.HandleFunc("/hello", helloHandler)

		port := 8080
		portStr := fmt.Sprintf(":%d", port)

		fmt.Println("Starting Server at port http://localhost"+portStr)

		// ListenAndServe(addr string, handler Handler) error
		if err := http.ListenAndServe(portStr, nil); err != nil {
			log.Fatal(err)
		}
	}
