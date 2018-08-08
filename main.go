package main

import (
	"fmt"
	"time"

	"github.com/kelvins/sunrisesunset"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	// rpioErr := rpio.Open()
	// pin := rpio.Pin(2)
	// pin.Output()
	isOn := false
	for {
		// if rpioErr == nil {
		// 	fmt.Println(rpioErr)
		// 	rpio.Close()
		// 	break
		// }

		param := sunrisesunset.Parameters{
			Latitude:  55.566935,
			Longitude: 12.245406,
			UtcOffset: 2.0,
			Date:      time.Now(),
		}

		sunrise, sunset, err := param.GetSunriseSunset()
		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
		// now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 14, 12, 13, 0, time.UTC)
		tdySunset := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, time.UTC)
		tmrSunrise := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, time.UTC)
		if err == nil {
			fmt.Print("Now: ")
			fmt.Println(now)
			fmt.Println(tdySunset)
			fmt.Println(tmrSunrise)

			if (now.After(tdySunset) && now.Before(tmrSunrise.Add(time.Hour*24))) || (now.Before(tdySunset.Add(time.Hour*24)) && now.Before(tmrSunrise)) {
				isOn = true
				fmt.Println("1")
				// pin.High()
				fmt.Printf("Is on?: %t\n", isOn)
			} else {
				fmt.Println("3")
				isOn = false
				// pin.Low()
				fmt.Printf("Is on?: %t\n", isOn)
			}

		} else {
			fmt.Println(err)
			rpio.Close()
			break
		}
		time.Sleep(time.Minute * 5)

	}
}
