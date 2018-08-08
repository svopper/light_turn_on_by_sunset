package main

import (
	"fmt"
	"time"

	"github.com/kelvins/sunrisesunset"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	rpioErr := rpio.Open()
	pin := rpio.Pin(2)
	pin.Output()
	isOn := false
	for {
		if rpioErr == nil {
			fmt.Println(rpioErr)
			rpio.Close()
			break
		}

		param := sunrisesunset.Parameters{
			Latitude:  55.566935,
			Longitude: 12.245406,
			UtcOffset: 2.0,
			Date:      time.Now(),
		}

		sunrise, sunset, err := param.GetSunriseSunset()
		now := time.Date(1, 1, 1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
		tdySunset := time.Date(1, 1, 1, sunset.Hour(), sunset.Minute(), sunset.Second(), 0, time.UTC)
		tmrSunrise := time.Date(1, 1, 2, sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, time.UTC)
		if err == nil {
			fmt.Printf("Now: %s\n", now.Format("15:04:05"))

			fmt.Printf("Is now after sunset?: %t\n", now.After(tdySunset))

			fmt.Printf("Is now before sunrise?: %t\n", now.Before(tmrSunrise))

			if now.After(tdySunset) && now.Before(tmrSunrise) {
				isOn = true
				pin.High()
				fmt.Printf("Is on?: %t\n", isOn)
			} else {
				isOn = false
				pin.Low()
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
