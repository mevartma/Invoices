package Invoices

import (
	"net/http"
	"Invoices/controllers"
)

func main() {
	http.ListenAndServe(":9090",controllers.NewMux())
}
