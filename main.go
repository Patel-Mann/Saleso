package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/patel-mann/saleo/internal/utils"
)

func main(){
	
	utils.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello from saleo") 
	})


	log.Println("Saleo server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
