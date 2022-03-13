package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

var products = template.Must(template.New("products").
	Parse(`
<h1>{{len .}} Products</h1>
<table>
<tr style='text-align: left'>
	<th>Name</th>
	<th>Price</th>
</tr>
{{range $item, $price := .}}
<tr>
	<td>{{$item}}</td>
	<td>{{$price}}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := products.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
