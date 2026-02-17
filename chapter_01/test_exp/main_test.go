package main

import (
	"log"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	name1 := reverseNameFixed(strings.TrimSpace("William"))
	expected1 := "mailliW"
	if !strings.EqualFold(name1, expected1) {
		t.Errorf("Response from reverseNameFixed is unexpected value. got [%s], expected [%s]", name1, expected1)
	}

	name2 := reverseNameFixed(strings.TrimSpace("Mister :)"))
	expected2 := "): retsiM"
	log.Println(name2, expected2, len(name2), len(expected2))

	if !strings.EqualFold(name2, expected2) {
		t.Errorf("Response from reverseNameFixed is unexpected value. got [%s], expected [%s]", name2, expected2)
	}
}