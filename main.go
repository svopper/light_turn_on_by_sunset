package main

import (
	"fmt"
	"time"

	"github.com/kelvins/sunrisesunset"
)

func main() {
	isOn := false
	for {
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
				fmt.Printf("Is on?: %t\n", isOn)
			} else {
				isOn = false
				fmt.Printf("Is on?: %t\n", isOn)
			}

		} else {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Minute * 5)

	}
}
