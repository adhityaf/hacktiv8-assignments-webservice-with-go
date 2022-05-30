package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", WeatherTemplate)
	fmt.Println("Running in port 8080")
	http.ListenAndServe(":8080", nil)
}

func WeatherTemplate(rw http.ResponseWriter, req *http.Request) {
	tmp := template.Must(template.ParseFiles("./resources/index.gohtml"))

	rand.Seed(time.Now().UnixNano())
	water := rand.Intn(10)
	waterStatus := ""
	if water <= 5 {
		waterStatus = "Aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Siaga"
	} else if water >= 9 {
		waterStatus = "Bahaya"
	}

	rand.Seed(time.Now().UnixNano())
	wind := rand.Intn(20)
	windStatus := ""
	if wind <= 6 {
		windStatus = "Aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Siaga"
	} else if wind > 15 {
		windStatus = "Bahaya"
	}

	data := map[string]interface{}{
		"Wind": map[string]interface{}{
			"Value":  wind,
			"Status": windStatus,
		},
		"Water": map[string]interface{}{
			"Value":  water,
			"Status": waterStatus,
		},
	}

	err := tmp.ExecuteTemplate(rw, "index.gohtml", data)
	if err != nil {
		panic(err)
	}
}
