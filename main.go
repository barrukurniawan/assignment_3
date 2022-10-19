package main

import (
	"assigment3/services"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

/* Assignment 3
   B class
   GLN038MNC007
   barru.kurniawan@gmail.com
   Barru Kurniawan */

func main() {
	go services.UpdateWeather()
	http.HandleFunc("/", dataCuaca)
	http.ListenAndServe(":9000", nil)
}

func dataCuaca(w http.ResponseWriter, r *http.Request) {
	dataWeather, err := ioutil.ReadFile("cuaca.json")
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	statusWeather := services.StatusWeather{}
	errUn := json.Unmarshal(dataWeather, &statusWeather)
	if errUn != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	water := statusWeather.Status.Water
	wind := statusWeather.Status.Wind
	var status_water string
	var status_wind string

	//check water level
	if water <= 5 {
		status_water = "Aman"
	} else if water >= 6 && water <= 8 {
		status_water = "Siaga"
	} else {
		status_water = "Bahaya"
	}

	//check wind level
	if wind <= 6 {
		status_wind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		status_wind = "Siaga"
	} else {
		status_wind = "Bahaya"
	}

	resultWeather := services.ResultWeather{}
	resultWeather.Water = water
	resultWeather.Wind = wind
	resultWeather.status_water = status_water
	resultWeather.status_wind = status_wind

	tpl, errTmpl := template.ParseFiles("static/index.html")
	if errTmpl != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": errTmpl.Error(),
		})
		return
	}
	tpl.Execute(w, resultWeather)

}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
