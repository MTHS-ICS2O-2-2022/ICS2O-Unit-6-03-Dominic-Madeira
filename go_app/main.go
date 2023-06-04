// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M
// Created on: June 2023
// This program shows you the temperature outside

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Icon string `json:"icon"`
	} `json:"weather"`
}

func main() {

	var jsonData WeatherData

	urlAddress := "https://api.openweathermap.org/data/2.5/weather?lat=45.4211435&lon=-75.6900574&appid=fe1d80e1e103cff8c6afd190cad23fa5"

	response, err := http.Get(urlAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	temperature := jsonData.Main.Temp
	temperatureCelcius := temperature - 273.15
	temperatureFormatted := fmt.Sprintf("%.2f", temperatureCelcius)

	fmt.Println()
	fmt.Print("The temperature outside is ", temperatureFormatted, "°C!")
	fmt.Println()
	fmt.Println("\nDone.")
}
