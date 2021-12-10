package main

import (
	"os"
	"reflect"
	"testing"
)

func TestGetLatest(t *testing.T) {
	var input string = "|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15"
	expected := []string{"11", "12", "13", "14", "15"}

	result := GetLatest(input, 5)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

func TestGetLatestShort(t *testing.T) {
	var input string = "|1|2|3|4|5"
	expected := []string{"", "1", "2", "3", "4", "5"}

	result := GetLatest(input, 10)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

func TestUseFile(t *testing.T) {
	var fileName string = "example.csv"

	expected := []string{"11", "12", "13", "14", "15"}

	result := GetChartData(fileName, "", 5)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}

func TestUseEnvVar(t *testing.T) {

	expected := []string{"5", "6", "7", "8"}

	os.Setenv("chart", "1|2|3|4|5|6|7|8")
	result := GetChartData("", "chart", 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %q, wanted %q", result, expected)
	}
}
