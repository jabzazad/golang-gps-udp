package model

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Coordinates struct {
	Imei      string
	Altitude  float64
	Latitude  float64
	Longitude float64
	Time      string
}

type Message struct {
	Msg    string
	Size   int
	Source string
}

func parse_ll(s string, n int, is_positive bool) (f float64) {
	a, err := strconv.ParseFloat(s[0:n+1], 64)
	if err != nil {
		fmt.Printf("error %s", err)
		log.Println("There was an error:", err)
	}
	d, err := strconv.ParseFloat(s[n+1:], 64)
	if err != nil {
		fmt.Printf("error %s", err)
		log.Println("There was an error:", err)
	}
	d /= 60.0
	res := a + d
	if is_positive {
		return res
	} else {
		return -res
	}
}

func parse_message(msg string) (c Coordinates) {
	fields := strings.Split(msg, ",")
	c.Imei = strings.Split(fields[0], ":")[1]
	c.Time = fields[2]
	//still not clear on what the format for this is. it doesn't look like meters.
	c.Altitude, _ = strconv.ParseFloat(fields[5], 64)
	c.Latitude = parse_ll(fields[7], 1, "N" == fields[8])
	c.Longitude = parse_ll(fields[9], 2, "E" == fields[10])
	return c
}
