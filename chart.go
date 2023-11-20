package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
	"github.com/guptarohit/asciigraph"
)

/*

C:\src\minion\go\chart [main ↑1]> $env:chart
|23|20|15|12|12|15|15|6|5|0

C:\src\minion\go\chart [main ↑1]> chart
[0 0 0 23 20 15 12 12 15 15 6 5 0]
Chart scale is 5
 4.00 ┼  ╭─╮
 3.00 ┤  │ ╰╮ ╭─╮
 2.00 ┤  │  ╰─╯ │
 1.00 ┤  │      ╰─╮
 0.00 ┼──╯        ╰

*/

func GetLatest(datastring string, max int) (res []string) {
	strData := strings.Split(datastring, "|")
	if len(strData) < max {
		return strData
	}
	return strData[len(strData)-max:]
}

func GetChartData(fileName string, envVar string, max int) (result []string) {

	// Set some default data.
	datastring := "|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15"
	datastring = datastring + "|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15"
	datastring = datastring + "|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15"

	// Pull data from a file
	if(fileName!=""){ 
		dat, err := ioutil.ReadFile(fileName)
		if err != nil { panic(err) }
		datastring = string(dat)
	}

	// Pull data from an environment variable.
	if(envVar!="") {
		datastring = os.Getenv(envVar)
	}

	// Return the allowed amount of data.
	return GetLatest(datastring, max)
}

func main() {
	envVar := flag.String("var", "chart", "Environment variable to chart from. (default: chart)")
	fileName := flag.String("file", "chart.csv", "File to chart from. (default: chart.csv)")
	flag.Parse()

	data := []float64{0, 0}

	strData := GetChartData(*fileName, *envVar, 73)

	// Turn it into numbers.
	var item string
	var bigItem float64 = 0
	for _, item = range strData {
		item = strings.TrimSpace(item)
		dataItem, _ := strconv.ParseFloat(item, 32)
		if dataItem > bigItem {
			bigItem = dataItem
		}
		data = append(data, dataItem)
	}

	// Scale the chart for available space.
	sizedData := []float64{}
	var scaleFactor float64 = 1
	if bigItem > 50 {
		scaleFactor = 10
	} else if bigItem > 20 {
		scaleFactor = 5
	} else if bigItem > 10 {
		scaleFactor = 2
	} else {
		scaleFactor = 1
	}
	fmt.Println("Chart scale is", scaleFactor)
	for _, dataItem := range data {
		sizedData = append(sizedData, math.Floor(dataItem/scaleFactor))
	}

	graph := asciigraph.Plot(sizedData)
	fmt.Println(graph)
}
