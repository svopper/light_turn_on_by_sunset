package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kelvins/sunrisesunset"
	"github.com/stianeikeland/go-rpio"
)

var (
	pin               = rpio.Pin(4)
	calculationParams = sunrisesunset.Parameters{
		Latitude:  55.566935,
		Longitude: 12.245406,
		UtcOffset: 2.0,
		Date:      time.Now(),
	}
)

func main() {
	rpioErr := rpio.Open()
	if rpioErr != nil {
		fmt.Println(rpioErr)
		os.Exit(1)
	}

	defer rpio.Close()
	pin.Output()

	for {

		sunrise, sunset, err := calculationParams.GetSunriseSunset()

		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
		// now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 14, 12, 13, 0, time.UTC)
		tdySunset := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, time.UTC)
		tmrSunrise := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, time.UTC)
		if err == nil {
			fmt.Print("Now: ")
			fmt.Println(now)
			fmt.Println(tdySunset)
			fmt.Println(tmrSunrise)
			fmt.Printf("\n")

			if (now.After(tdySunset) && now.Before(tmrSunrise.Add(time.Hour*24))) || (now.Before(tdySunset.Add(time.Hour*24)) && now.Before(tmrSunrise)) {
				pin.High()
				fmt.Println("Light is on")
			} else {
				fmt.Println("Light is off")
				pin.Low()
			}

		} else {
			fmt.Println(err)
			os.Exit(1)
		}
		time.Sleep(time.Minute * 5)

	}
}
