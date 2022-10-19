package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type StatusWeather struct {
	Status Weather
}

type Weather struct {
	Water int
	Wind  int
}

type ResultWeather struct {
	Water       int
	Wind        int
	StatusWater string
	StatusWind  string
}

func UpdateWeather() {
	max := 100
	min := 1
	for {
		fmt.Println("Running Update Weather per 15 Second")
		rand.Seed(time.Now().UnixNano())
		statusWeather := StatusWeather{}
		statusWeather.Status.Water = rand.Intn(max-min) + min
		statusWeather.Status.Wind = rand.Intn(max-min) + min

		dataWeather, err := json.Marshal(statusWeather)
		if err != nil {
			fmt.Println(err)
		}

		errs := ioutil.WriteFile("data.json", dataWeather, 0644)
		if errs != nil {
			fmt.Println(errs)
		}
		time.Sleep(15 * time.Second)
	}
}
