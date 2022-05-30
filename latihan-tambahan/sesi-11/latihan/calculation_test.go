package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestTable struct {
	Expected int
	Request  int
	Message  string
}

func TestSum(t *testing.T) {
	testsequal := []TestTable{
		{
			Request:  Sum(1, 2, 41, 5),
			Expected: 49,
			Message:  "should return 49",
		},
		{
			Request:  Sum(59, 2),
			Expected: 61,
			Message:  "should return 61",
		},
		{
			Request:  Sum(12, 2),
			Expected: 14,
			Message:  "should return 14",
		},
	}

	testsnotequal := []TestTable{
		{
			Request:  Sum(1, 2, 41, 5),
			Expected: 0,
			Message:  "should not return 49",
		},
		{
			Request:  Sum(59, 2),
			Expected: 61,
			Message:  "should not return 61",
		},
		{
			Request:  Sum(12, 2),
			Expected: 0,
			Message:  "should not return 14",
		},
	}

	for i, test := range testsequal {
		t.Run(fmt.Sprintf("TestEqual-%d", i), func(t *testing.T) {
			require.Equal(t, test.Expected, test.Request, test.Message)
		})
	}

	for i, test := range testsnotequal {
		t.Run(fmt.Sprintf("TestNotEqual-%d", i+1), func(t *testing.T) {
			require.NotEqual(t, test.Expected, test.Request, test.Message)
		})
	}
}
