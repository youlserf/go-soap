package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type AddRequest struct {
	XMLName xml.Name `xml:"AddRequest"`
	A       int      `xml:"a"`
	B       int      `xml:"b"`
}

type AddResponse struct {
	XMLName xml.Name `xml:"AddResponse"`
	Result  int      `xml:"result"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	err := xml.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := AddResponse{Result: req.A + req.B}

	w.Header().Set("Content-Type", "text/xml")
	xml.NewEncoder(w).Encode(&resp)
}

func main() {
	http.HandleFunc("/soap", addHandler)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
