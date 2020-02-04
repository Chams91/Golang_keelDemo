package main
import (  
		  "fmt"
		  "log"
		  "net/http"
		 )
var version =1
func handler(w http.ResponseWriter, r *http.Request) {
	        fmt.Fprintln(w, "Hello, 世界.. From V %s", version ) }
func main() {
			http.HandleFunc("/", handler)
			log.Fatal(http.ListenAndServe(":8888", nil)) 
			} 