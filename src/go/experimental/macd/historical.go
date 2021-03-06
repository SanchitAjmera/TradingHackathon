package main

import (
//	"fmt"
	"github.com/luno/luno-go/decimal"
	"github.com/tealeg/xlsx"
)

/*
TODO:
- change colNum far ask and bid
*/

//global variable for historicalData
var historicalData [][][]string

// function to process the csv file and return a 3d array of strings
// historicalData is of the form: [sheetNum][rowNum][colNum]
func parseXlsx() {
	fileSlice, err := xlsx.FileToSlice("../../ticker/data_7to8_July/tickerData09072020.xlsx")
	if err != nil {
		panic(err)
	}
	historicalData = fileSlice
}

// function to get the bid price from a given row in the excel spreadsheet
func getOfflineBid(currRow int64) decimal.Decimal {
	currPrice := historicalData[0][int(currRow)][1] //Change this
	// if data is non applicable skip this row
	if (currPrice == "NaN") {
		return getOfflineBid(currRow - 1)
	}

	currPriceDecimal, err := decimal.NewFromString(currPrice)

	if err != nil {
		panic(err)
	}
	return currPriceDecimal
}

// function to get the ask price from a given row in the excel spreadsheet
func getOfflineAsk(currRow int64) decimal.Decimal {
	currPrice := historicalData[0][currRow][4] //Change this
	// if data is non applicable skip this row
	if (currPrice == "NaN") {
		return getOfflineAsk(currRow - 1)
	}

	currPriceDecimal, err := decimal.NewFromString(currPrice)

	if err != nil {
		panic(err)
	}

	return currPriceDecimal
}
