package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type pageData struct {
	Title     string
	IsResult  bool
	BaseValue string
	Result    string
	FromUnit  string
	ToUnit    string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("length.html"))
		pageData := pageData{
			Title:    "Length",
			IsResult: false,
		}

		if r.Method == "POST" {
			baseValue, _ := strconv.ParseFloat(r.FormValue("length"), 64)
			result := baseValue * UnitDB["length"][r.FormValue("from_unit")][r.FormValue("to_unit")]

			pageData.IsResult = true
			pageData.BaseValue = r.FormValue("length")
			pageData.Result = fmt.Sprintf("%v", result)
			pageData.FromUnit = r.FormValue("from_unit")
			pageData.ToUnit = r.FormValue("to_unit")
		}

		tmpl.Execute(w, pageData)
	})

	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("weight.html"))
		pageData := pageData{
			Title:    "Weight",
			IsResult: false,
		}

		if r.Method == "POST" {
			baseValue, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
			result := baseValue * UnitDB["weight"][r.FormValue("from_unit")][r.FormValue("to_unit")]

			pageData.IsResult = true
			pageData.BaseValue = r.FormValue("weight")
			pageData.Result = fmt.Sprintf("%v", result)
			pageData.FromUnit = r.FormValue("from_unit")
			pageData.ToUnit = r.FormValue("to_unit")
		}

		tmpl.Execute(w, pageData)
	})

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("temperature.html"))
		pageData := pageData{
			Title:    "Temperature",
			IsResult: false,
		}

		if r.Method == "POST" {
			baseValue, _ := strconv.ParseFloat(r.FormValue("temperature"), 64)
			result := 0.0

			if r.FormValue("from_unit") == "celcius" {
				if r.FormValue("to_unit") == "celcius" {
					result = baseValue
				} else if r.FormValue("to_unit") == "fahrenheit" {
					result = (baseValue * 9 / 5) + 32
				} else {
					result = baseValue + 273.15
				}
			} else if r.FormValue("from_unit") == "fahrenheit" {
				if r.FormValue("to_unit") == "celcius" {
					result = (baseValue - 32) * 5 / 9
				} else if r.FormValue("to_unit") == "fahrenheit" {
					result = baseValue
				} else {
					result = (baseValue-32)*5/9 + 273.15
				}
			} else {
				if r.FormValue("to_unit") == "celcius" {
					result = baseValue - 273.15
				} else if r.FormValue("to_unit") == "fahrenheit" {
					result = (baseValue-273.15)*9/5 + 32
				} else {
					result = baseValue
				}
			}

			pageData.IsResult = true
			pageData.BaseValue = r.FormValue("temperature")
			pageData.Result = fmt.Sprintf("%v", result)
			pageData.FromUnit = r.FormValue("from_unit")
			pageData.ToUnit = r.FormValue("to_unit")
		}

		tmpl.Execute(w, pageData)
	})
	http.ListenAndServe(":8000", nil)
}
