package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data := csv.NewReader(file)
	records, err := data.ReadAll()
	if err != nil {
		log.Println(err)
	}

	var totalFuel float64

	for k := range records {
		num, err := strconv.ParseFloat(records[k][0], 64)
		if err != nil {
			log.Println(err)
		}
		fuel := calcFuel(num)
		totalFuel += fuel
		fmt.Printf("%v : %v : %v\n", num, fuel, totalFuel)
	}

	fmt.Println(totalFuel)
}

func calcFuel(mass float64) float64 {
	fuel := math.Floor(mass/3) - 2

	return fuel
}
