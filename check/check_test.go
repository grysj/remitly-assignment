package main

import (
	"fmt"
	"testing"
)

func TestChecking1(t *testing.T) {

	got, err := Check("testCheck/correct1.json")

	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
	}

	want := true

	if got != want {
		fmt.Printf("want: %v, got: %v\n", want, got)
		t.Fail()
	}
}

func TestChecking2(t *testing.T) {

	got, err := Check("testCheck/correct2.json")

	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
	}

	want := false

	if got != want {
		fmt.Printf("want: %v, got: %v\n", want, got)
		t.Fail()
	}
}

func TestChecking3(t *testing.T) {

	got, err := Check("testCheck/correct3.json")

	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
	}

	want := false

	if got != want {
		fmt.Printf("want: %v, got: %v\n", want, got)
		t.Fail()
	}
}
