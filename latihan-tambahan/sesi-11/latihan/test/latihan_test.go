package test

import (
	"fmt"
	"sesi-11/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestTableSoal1 struct {
	Expected bool
	Request  bool
	Message  string
}

type TestTableSoal2 struct {
	Expected int
	Request  int
	Message  string
}

func TestSoal1(t *testing.T) {
	testSoal1 := []TestTableSoal1{
		{
			Request:  helpers.Soal1("pria", 16),
			Expected: false,
			Message:  "should return false",
		},
		{
			Request:  helpers.Soal1("pria", 17),
			Expected: false,
			Message:  "should return false",
		},
		{
			Request:  helpers.Soal1("pria", 50),
			Expected: true,
			Message:  "should return true",
		},
		{
			Request:  helpers.Soal1("pria", 60),
			Expected: false,
			Message:  "should return false",
		},
		{
			Request:  helpers.Soal1("wanita", 16),
			Expected: false,
			Message:  "should return false",
		},
		{
			Request:  helpers.Soal1("wanita", 17),
			Expected: false,
			Message:  "should return false",
		},
		{
			Request:  helpers.Soal1("wanita", 50),
			Expected: true,
			Message:  "should return true",
		},
		{
			Request:  helpers.Soal1("wanita", 60),
			Expected: false,
			Message:  "should return false",
		},
	}

	for i, test := range testSoal1 {
		t.Run(fmt.Sprintf("TestSoal1-%d", i+1), func(t *testing.T) {
			require.Equal(t, test.Expected, test.Request, test.Message)
		})
	}
}

func TestSoal2(t *testing.T) {
	testSoal2 := []TestTableSoal2{
		{
			Request:  helpers.Soal2([]int{5,3,9,8}, []int{2,2,3,1}, 3),
			Expected: 7,
			Message:  "should return 7",
		},
		{
			Request:  helpers.Soal2([]int{2,6,3,9}, []int{2,2,3,5}, 2),
			Expected: 14,
			Message:  "should return 14",
		},
	}

	for i, test := range testSoal2{
		t.Run(fmt.Sprintf("TestSoal2-%d", i+1), func(t *testing.T) {
			require.Equal(t, test.Expected, test.Request, test.Message)
		})
	}
}
