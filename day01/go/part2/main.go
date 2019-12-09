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
		fmt.Printf("%.0f\t: %f\t : %f\n", num, fuel, totalFuel)
	}

	fmt.Printf("%f", totalFuel)
}

func calcFuel(mass float64) float64 {

	var fuel float64
	for math.Floor(mass/3)-2 > 0 {
		mass = math.Floor(mass/3) - 2
		fuel += mass
	}
	return fuel
}
