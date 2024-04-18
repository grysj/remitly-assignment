package main

import (
	"fmt"
	"testing"

	"github.com/grysj/check/parse"

	"github.com/grysj/check/read"
)

func equalMaps(map1 map[string]bool, map2 map[string]bool) bool {

	if len(map1) != len(map2) {
		return false
	}

	for key1, val1 := range map1 {
		val2, contains := map2[key1]
		if !contains || val1 != val2 {
			fmt.Printf("%v", key1)
			return false
		}

	}

	return true

}

func TestCorrectParse1(t *testing.T) {

	jn, _ := read.ReadFile("testParse/correctparse1.json")

	want := map[string]bool{"*": true}

	got, err := parse.GetResourseList(jn)

	if err != nil {
		t.Errorf("got %s, want nil", err.Error())
	}

	if !equalMaps(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestCorrectParse2(t *testing.T) {

	jn, _ := read.ReadFile("testParse/correctparse2.json")

	want := map[string]bool{"*": true, "arn:aws:s3:::confidential-data": true,
		"arn:aws:s3:::confidential-data/*": true}

	got, err := parse.GetResourseList(jn)

	if err != nil {
		t.Fail()
	}

	if !equalMaps(want, got) {

		t.Errorf("got %v, want %v", got, want)
	}

}

func TestCorrectParse3(t *testing.T) {

	jn, _ := read.ReadFile("testParse/correctparse3.json")

	want := map[string]bool{"*": true, "arn:aws:s3:::confidential-data": true}

	got, err := parse.GetResourseList(jn)

	if err != nil {
		t.Fail()
	}

	if !equalMaps(want, got) {

		t.Errorf("got %v, want %v", got, want)
	}

}

func TestNoPolicyDoc(t *testing.T) {

	jn, _ := read.ReadFile("testParse/noPolicyDocument.json")

	_, err := parse.GetResourseList(jn)

	if err.Error() != "PolicyDocument not in json" {
		t.Fail()
	}

}

func TestNoStatement(t *testing.T) {

	jn, _ := read.ReadFile("testParse/noStatement.json")

	_, err := parse.GetResourseList(jn)
	if err.Error() != "Statement not in json" {
		t.Fail()
	}

}

func TestNoResource1(t *testing.T) {

	jn, _ := read.ReadFile("testParse/noResource1.json")

	_, err := parse.GetResourseList(jn)

	if err.Error() != "Resource not in json" {
		t.Fail()
	}

}

func TestNoResource2(t *testing.T) {

	jn, _ := read.ReadFile("testParse/noResource2.json")

	_, err := parse.GetResourseList(jn)

	if err.Error() != "Resource not in json" {
		t.Fail()
	}

}
