package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mang0kitty/website/src/helpers"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Handle() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", BookHandler)
	r.HandleFunc("/images/{name}", ImageHandler)
	return r
}

func BookHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Successfully Opened books.json")

	enableCors(w)
	fptr, err := os.Open("src/store/books.json")
	helpers.CheckError(err)
	fmt.Println("Successfully Opened books.json")
	defer fptr.Close()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.Copy(w, fptr)
}

func ImageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	imageName := vars["name"]

	imagePath := fmt.Sprintf("src/store/images/%s", imageName)

	fptr, err := os.Open(imagePath)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	fmt.Println("Successfully Opened /images/", imageName)
	defer fptr.Close()

	io.Copy(w, fptr)
}
