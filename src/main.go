package main

import (
	"main/src/controllers/productcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)

	http.ListenAndServe(":8080", nil)
}
