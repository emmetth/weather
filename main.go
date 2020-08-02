package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func render(f Forecast) {
	fmt.Println(f.Hourly.Summary)
	fmt.Println(f.Daily.Summary)

	fmt.Println(strings.Repeat("=", 55))
	fmt.Println("Today\t\t\t|\tTomorrow")
	fmt.Println("Time\tTemp\tRain %\t|\tTime\tTemp\tRain %")
	fmt.Println(strings.Repeat("=", 55))

	for i := 0; i < 24; i++ {
		today := time.Unix(f.Hourly.Data[i].Time, 0)
		tomorrow := time.Unix(f.Hourly.Data[i+24].Time, 0)
		fmt.Printf("%02d:%02d\t%0.f\t%0.f%%\t|\t%02d:%02d\t%0.f\t%0.f%%\n",
			today.Hour(), today.Minute(), f.Hourly.Data[i].Temperature, f.Hourly.Data[i].PrecipProbability*100,
			tomorrow.Hour(), tomorrow.Minute(), f.Hourly.Data[i+24].Temperature, f.Hourly.Data[i+24].PrecipProbability*100)
		if today.Hour() == 0 && i > 0 && i < 24 {
			fmt.Println(strings.Repeat("=", 55))
		}

	}

	for _, alert := range f.Alerts {
		fmt.Println(strings.Repeat("=", 55))
		fmt.Println(alert.Title)
		fmt.Println(alert.Description)
	}
}

func main() {
	envLat, latOk := os.LookupEnv("LATITUDE")
	envLong, longOk := os.LookupEnv("LONGITUDE")
	if !latOk {
		fmt.Println("Error: Enivronment variable LATITUDE is not set.")
	}
	if !longOk {
		fmt.Println("Error: Enivronment variable LONGITUDE is not set.")
	}
	if !latOk || !longOk {
		os.Exit(1)
	}

	lat, latErr := strconv.ParseFloat(envLat, 64)
	if latErr != nil {
		fmt.Println("Error: Invalid Latitude (float64)")
	}
	long, longErr := strconv.ParseFloat(envLong, 64)
	if longErr != nil {
		fmt.Println("Error: Invalid Longitude (float64)")
	}
	if latErr != nil || longErr != nil {
		os.Exit(1)
	}

	f, err := getForecast(lat, long)
	if err != nil {
		fmt.Println(err)
		return
	}
	render(f)
}
