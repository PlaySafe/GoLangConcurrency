package main

import (
	"os"
	"strings"
	"testing"
)

const file string = "w.txt"
const total int = 20

func TestWriteNNumbersToFile(t *testing.T){
	n, err := write(file, total)

	if err != nil {
		t.Error(err)
	}

	if n != total {
		t.Error("Expect write ")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	length := len(string(data))
	if total != length {
		t.Errorf("Expect %v characters, but %v", total, length)
	}

	os.Remove(file)
}

func TestReadNumbersFromFile(t *testing.T){
	defer os.Remove(file)
	n, err0 := write(file, total)

	if err0 != nil {
		t.Error(err0)
	}

	if n != total {
		t.Errorf("Expect return %v, but %v", total, n)
	}

	mockStdOut := func(s string){
		expect := func(num byte) string {
			if (num - '0') % 2 == 0 {
				return string(num) + " is even number"
			} else {
				return string(num) + " is odd number"
			}
		}(s[0])
		cmp := strings.Compare(s, expect)
		if cmp != 0 {
			t.Errorf("Expect '%v', but '%v'", expect, s)
		}
	}

	length, err1 := read(file, mockStdOut, 0, total)
	if err1 != nil {
		t.Error(err1)
	}

	data, err2 := os.ReadFile(file)
	if err2 != nil {
		t.Error(err2)
	}

	expectLength := len(string(data))
	if expectLength != length {
		t.Errorf("Expect %v characters, but got %v", expectLength, length)
	}
}
