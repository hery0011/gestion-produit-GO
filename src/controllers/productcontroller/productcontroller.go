package productcontroller

import (
	"html/template"
	"main/src/models"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var ProductModel models.ProductModel
	products, _ := ProductModel.FindAll()
	data := map[string]interface{}{
		"products": products,
	}

	tmp, _ := template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)
}
