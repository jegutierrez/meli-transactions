package main

import (
	"testing"
)

var rawLog = "[user:Keanu] [type:inversión] [ammount:1115401358389555]"
var parsedLog = &Transaction{"Keanu", "inversión", 1115401358389555}

func TestSingleLogParsing(t *testing.T) {
	ts := &Transaction{}
	ParseLog(ts, rawLog)
	if ts.user != "Keanu" {
		t.Errorf("User was incorrect, got: %s, expected: %s.", ts.user, "Keanu")
	}
	if ts.operation != "inversión" {
		t.Errorf("Transaction type was incorrect, got: %s, expected: %s.", ts.operation, "inversión")
	}
	if ts.amount != 1115401358389555 {
		t.Errorf("Amount was incorrect, got: %v, expected: %v.", ts.amount, 1115401358389555)
	}
}
