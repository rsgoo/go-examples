package main

import (
	"testing"
	"sdk"
)

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		str    string
		expect bool
	}

	testCaseList := map[string]testCase{
		"simple":    {"abcba", true},
		"chinese":   {"上海自来水来自海上", true},
		"caseFalse": {"aa里贝里", false},
		"caseTrue": {"小大大小a", false},
	}

	for _, caseDetail := range testCaseList {
		got := sdk.IsPalindrome(caseDetail.str)

		if got != caseDetail.expect {
			t.Errorf("excepted:%v, got:%v", caseDetail.expect, got)
		}
	}
}
