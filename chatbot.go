//adapted from https://golang.org/doc/articles/wiki/


package main

import (
    "fmt"
    "net/http"
)


func chatFunc(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w,r , "index.html");
	
    //Call to ParseForm makes form fields available.
    err := r.ParseForm()
    if err != nil {
        // Handle error here via logging and then return            
    }

    userInput := r.PostFormValue("userInput")
    fmt.Fprintf(w, "Hello, %s!", userInput)
}


}



func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}