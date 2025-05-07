package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Stock temporaire en mémoire
var stock []Product

type Product struct {
	Fruit  int
	From   string
	Weight float32
	Date   time.Time
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/stock", stockHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Serveur démarré sur http://localhost:8092")
	err := http.ListenAndServe(":8092", nil)
	if err != nil {
		log.Fatal("Erreur serveur :", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/add.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Traitement du formulaire (POST)
	r.ParseForm()

	fruit, _ := strconv.Atoi(r.FormValue("fruit"))
	from := r.FormValue("from")
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 32)

	// Ajout au stock en mémoire
	p := Product{
		Fruit:  fruit,
		From:   from,
		Weight: float32(weight),
		Date:   time.Now(),
	}
	stock = append(stock, p)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func stockHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("stock.html").Funcs(template.FuncMap{
		"index": func(arr []string, i int) string {
			if i >= 0 && i < len(arr) {
				return arr[i]
			}
			return "❓"
		},
	}).ParseFiles("templates/stock.html"))

	data := struct {
		Fruits []string
		Items  []Product
	}{
		Fruits: []string{"🍌 Banane", "🍒 Cerise", "🍎 Pomme", "🍓 Fraise", "🍍 Ananas", "🍑 Abricot", "🥝 Kiwi"},
		Items:  stock,
	}

	tmpl.Execute(w, data)
}
