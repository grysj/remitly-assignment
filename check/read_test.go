package main

import (
	"testing"

	read "github.com/grysj/remitly-assignment/check/read"
)

func TestIncorrectJSON1(t *testing.T) {

	want := "invalid json"
	_, got := read.ReadFile("testJson/wrongjson1.json")

	if got.Error() != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestIncorrectJSON2(t *testing.T) {

	want := "invalid json"
	_, got := read.ReadFile("testJson/wrongjson2.json")

	if got.Error() != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestIncorrectJSON3(t *testing.T) {

	want := "invalid json"
	_, got := read.ReadFile("testJson/wrongjson3.json")

	if got.Error() != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestIncorrectJSON4(t *testing.T) {

	want := "invalid json"
	_, got := read.ReadFile("testJson/wrongjson4.json")

	if got.Error() != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCorrectJSON1(t *testing.T) {
	_, got := read.ReadFile("testJson/correctjson1.json")

	if got != nil {
		t.Errorf("got %s, want nil", got)
	}
}

func TestCorrectJSON2(t *testing.T) {
	_, got := read.ReadFile("testJson/correctjson2.json")

	if got != nil {
		t.Errorf("got %s, want nil", got)
	}
}

func TestCorrectByte(t *testing.T) {

	byteValue := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "*"
	
				}
			]
		}
	}`)

	_, got := read.ReadByte(byteValue)

	if got != nil {
		t.Errorf("got %s, want nil", got)
	}
}

func TestIncorrectByte(t *testing.T) {

	byteValue := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "*"
				}
			]
		}
	`)

	_, got := read.ReadByte(byteValue)

	want := "invalid json"

	if got.Error() != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
