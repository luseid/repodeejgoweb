package main


//Por falta de tiempo, no puedo modulizar, pero considero una ejercitación sobre el traslado de datos JSON a un slice

import(
	"net/http"
	"github.com/go-chi/chi/v5"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Product struct {
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var products []Product

func jsonReadProducts(slice *[]Product) {
	//Read the file and save it in "file", that is a bytes slice
	file, err := ioutil.ReadFile("/Users/lseidman/code/Repository/goWeb/clase2/ej1/products.json")
	if err != nil {
		fmt.Println("1Error al leer el archivo JSON:", err)
		return
	}
	//Change the JSON format to bytes, and save it in products
	err = json.Unmarshal(file, &slice)
	if err != nil {
		fmt.Println("2Error al leer el archivo JSON:", err)
		return
	}
	return
}

func bienvenidaHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Bienvenido a mi API WEB !")}}
	

func main(){
	
	jsonReadProducts(&products)
	
	//router
	rt:=chi.NewRouter()
	rt.Get("/", bienvenidaHandler())
	
	//endpoints
	http.ListenAndServe(":8080",rt)
}

