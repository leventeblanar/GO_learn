package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//  struct a requesthez
type EchoRequest struct {
	Msg string `json:"msg"`
}

//  struct a responsehoz
type EchoResponse struct {
	Echo string `json:"echo"`
}


//  Sima "/" endpoint - Itt azt néztük meg, hogy lehet egy sima w responsewriterrel requestet fogadni és csak szöveget visszadni
//  Itt a http headert belövi a go automatikusan text/palin-re
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there!")
}



//  Ellenben itt már kezeljük a request metódust, mivel ha nem GET akkor errort küldünk
//  Illetve konkrétan leírjuk a HTTP header típusát
func healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}



//  POST endpoint
//  az msg := r.FormValue elérhetőtév teszi a body-ban az msg cím alatt futó üzenetet (ezt postoljuk)
//  híváskor: curl -X POST -d "msg=hello" http:// localhost:8080/echo
func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post method allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		msg = "empty"
	}

	fmt.Fprintf(w, "you said: %s\n", msg)
}


// JSON POST endpoint
func echojson(w http.ResponseWriter, r * http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	//  Request oldalon decode
	//  Itt a req egy új változó amintek a struktúrása az EchoRequest
	var req EchoRequest
	//  Inline error check - a Decode() egy error-t ad vissza azért deklaráljuk és inicializáljuk az err változót itt
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	//  Response oldalon encode
	resp := EchoResponse{Echo: req.Msg}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}



//  A main functionben a HandleFunc-al tudjuk meghívni a function-öket hogy az endpointok létrejöjjenek
func main() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/echojson", echojson)

	log.Println("Listenning on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}