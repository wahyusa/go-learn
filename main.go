package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

// response weatheresponse extend weather khusus buat response
type WeatherResponse struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
	HumanStatus string `json:"humanstatus"`
}

func main() {
	// go routine
	go updateWeather()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html") // Set to HTML
			http.ServeFile(w, r, "index.html")
			return
		} else if r.URL.Path == "/weather" {
			weather := readWeatherFromFile()
			humanStatus := computeStatus(weather.Status.Water, weather.Status.Wind)
			response := WeatherResponse{
				Status: struct {
					Water int `json:"water"`
					Wind  int `json:"wind"`
				}{
					Water: weather.Status.Water,
					Wind:  weather.Status.Wind,
				},
				HumanStatus: humanStatus,
			}
			responseJSON, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
		} else if r.URL.Path == "/tailwind.css" {
			w.Header().Set("Content-Type", "text/css") // Set to CSS
			http.ServeFile(w, r, "tailwind.css")
			return
		} else if r.URL.Path == "/weather.json" {
			w.Header().Set("Content-Type", "application/json")
			http.ServeFile(w, r, "weather.json")
			return
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":6969", nil)
}

func updateWeather() {
	// loop yang akan dijalankan go routine
	for {
		weather := generateRandomWeather()
		writeJSONToFile(weather)
		time.Sleep(15 * time.Second)
	}
}

func generateRandomWeather() Weather {
	return Weather{
		Status: struct {
			Water int `json:"water"`
			Wind  int `json:"wind"`
		}{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		},
	}
}

func writeJSONToFile(weather Weather) {
	// kasih indent ketika bikin JSONnya biar rapih, default 1 tab = 4 spasi
	data, err := json.MarshalIndent(weather, "", "    ")
	if err != nil {
		fmt.Println("Marshal FAIL:", err)
		return
	}
	err = os.WriteFile("weather.json", data, 0644)
	if err != nil {
		fmt.Println("Writing FAIL:", err)
		return
	}
}

func readWeatherFromFile() Weather {
	file, err := os.ReadFile("weather.json")
	if err != nil {
		fmt.Println("Read FAIL:", err)
		return Weather{}
	}
	var weather Weather
	err = json.Unmarshal(file, &weather)
	if err != nil {
		fmt.Println("Unmarshaling FAIL:", err)
		return Weather{}
	}
	return weather
}

func computeStatus(water, wind int) string {
	if water < 5 || wind < 6 {
		return "Aman"
	} else if water >= 6 && water <= 8 || wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		// water = > 8
		// wind = > 15
		return "Bahaya"
	}
}
