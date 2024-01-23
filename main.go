package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Nastavení cesty a handleru
	http.HandleFunc("/", helloHandler)

	// Nastavení portu a spuštění serveru
	port := 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Server běží na adrese http://localhost:%d\n", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Chyba při spouštění serveru:", err)
	}
}

// Handler pro cestu "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Odpověď s textem "ahoj"
	fmt.Fprint(w, "ahoj")
}
