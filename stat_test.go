package main

import "testing"

func TestGetDayString(t *testing.T) {
	ans := getDayString("2022-01-01 00:00:00")
	if ans != "2022-01-01" {
		t.Error("Error")
	}
}
