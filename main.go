package main

import (
	"fmt"
	"time"
	"os"
	"strconv"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	minutes := 0
	ticker := time.NewTicker(time.Minute)

	for {
		<- ticker.C
		minutes++
		fmt.Printf("\rMinutes passed: %v", minutes)

		homeDir, err := os.UserHomeDir()
		check(err)
		f, err := os.OpenFile(homeDir + "/lastminute", os.O_RDWR|os.O_CREATE, 0664)
		check(err)
		_, err = f.Write([]byte(strconv.Itoa(minutes) + "\n"))
		check(err)
	}
}
