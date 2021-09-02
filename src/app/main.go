package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

import imports "github.com/rocketlaunchr/dataframe-go/imports"
import dataframe "github.com/rocketlaunchr/dataframe-go"

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	ctx := context.Background()
	dataCSV, err := readCSVFromUrl("https://raw.githubusercontent.com/fantasydatapros/data/master/yearly/2020.csv")
	if err != nil {
		panic(err)
	}
	var csvString = ""
	i := 0
	for i < len(dataCSV) {
		var row = strings.Join(dataCSV[i], ",")
		row += "\n"
		csvString += row
		i += 1
	}
	df, err := imports.LoadFromCSV(ctx, strings.NewReader(csvString))
	sks := []dataframe.SortKey{
		{Key: "FantasyPoints", Desc: true},
	}
	df.Sort(ctx, sks)
	fmt.Println(df.Table())
}
