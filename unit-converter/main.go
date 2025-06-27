package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"unit-converter/unit-converter/converters"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "length")
	})
	http.HandleFunc("/length", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "length")
	})
	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "weight")
	})
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "temperature")
	})

	http.HandleFunc("/convert/length", handleLengthConversion)
	http.HandleFunc("/convert/weight", handleWeightConversion)
	http.HandleFunc("/convert/temperature", handleTemperatureConversion)

	println("Server is running: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func render(w http.ResponseWriter, r *http.Request, page string) {
	base := filepath.Join("templates", "base.html")
	content := filepath.Join("templates", page+".html")

	tmpl, err := template.ParseFiles(base, content)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"ActiveTab": page,
	}

	if r.Header.Get("HX-Request") == "true" {
		tmpl.ExecuteTemplate(w, "content", data)
	} else {
		tmpl.ExecuteTemplate(w, "base", data)
	}
}

func handleLengthConversion(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	result, err := converters.ConvertLength(value, from, to)
	if err != nil {
		w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + err.Error() + `">`))
		return
	}

	w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + result + `">`))
}

func handleWeightConversion(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	result, err := converters.ConvertWeight(value, from, to)
	if err != nil {
		w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + err.Error() + `">`))
		return
	}

	w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + result + `">`))
}

func handleTemperatureConversion(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	result, err := converters.ConvertTemperature(value, from, to)
	if err != nil {
		w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + err.Error() + `">`))
		return
	}

	w.Write([]byte(`<input type="text" id="result" name="result" readonly class="result-field" value="` + result + `">`))
}
