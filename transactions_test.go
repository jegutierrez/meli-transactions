package main

import (
	"fmt"
	"testing"
)

var data = []Transaction{
	{"Juan", "descuento", 1},
	{"Gustavo", "pago", 1},
	{"Dalí", "inversión", 2},
	{"Kung", "descuento", 2},
	{"Dalí", "descuento", 3},
	{"Kung", "cobro", 1},
	{"Gustavo", "pago", 2},
	{"Keanu", "inversión", 3},
	{"Juan", "descuento", 4},
	{"Gustavo", "pago", 3},
	{"Dalí", "inversión", 4},
	{"Kung", "descuento", 5},
	{"Dalí", "descuento", 6},
	{"Kung", "cobro", 2},
	{"Gustavo", "pago", 4},
	{"Keanu", "inversión", 5},
	{"Juan", "descuento", 7},
	{"Gustavo", "pago", 5},
	{"Keanu", "inversión", 6},
	{"Keanu", "inversión", 1},
}

var unsortedData = []Transaction{
	{"Gustavo", "pago", 2},
	{"Keanu", "inversión", 6},
	{"Dalí", "inversión", 4},
	{"Kung", "cobro", 1},
}

var sortedData = []Transaction{
	{"Kung", "cobro", 1},
	{"Gustavo", "pago", 2},
	{"Dalí", "inversión", 4},
	{"Keanu", "inversión", 6},
}

var avgRs = map[string]OperationResult{
	"inversión": {"inversión", 6, 3.5},
	"descuento": {"descuento", 7, 4},
	"pago":      {"pago", 5, 3},
	"cobro":     {"cobro", 2, 1.5},
}

var userOpRs = map[string]OperationUserResult{
	"inversión": {"inversión", "Keanu", 4},
	"descuento": {"descuento", "Juan", 3},
	"pago":      {"pago", "Gustavo", 5},
	"cobro":     {"cobro", "Kung", 2},
}

func areSlicesEqual(a, b []Transaction) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].amount != b[i].amount {
			return false
		}
	}
	return true
}

func TestMeanByOperationType(t *testing.T) {
	means := MeanByOperationType(&data)
	if len(means) != 4 {
		t.Errorf("Amount of Operation Types was incorrect, got: %d, expected: %d.", len(means), 4)
	}
	for _, v := range means {
		op, ok := avgRs[v.name]
		if !ok {
			t.Errorf("Opertation %s not calculated", op.name)
		} else {
			if op.avg != v.avg {
				t.Errorf("Average for %s was incorrect, got: %f, expected: %f.", v.name, v.avg, op.avg)
			}
			if op.count != v.count {
				t.Errorf("Count for %s was incorrect, got: %d, expected: %d.", v.name, v.count, op.count)
			}
		}
	}
	if len(means) != 4 {
		t.Errorf("Amount of Operation Types was incorrect, got: %v, expected: %v.", len(means), 4)
	}
}

func TestSortHelper(t *testing.T) {
	Sort(&unsortedData)
	if !areSlicesEqual(unsortedData, sortedData) {
		t.Error("slice is not sorted by amount got:", unsortedData)
	}
}

func TestUserWithGreaterAmountByOperationType(t *testing.T) {
	rs := LeadingUserByOperationType(&data)
	if len(rs) != 4 {
		t.Errorf("Amount of Operation Types was incorrect, got: %v, expected: %v.", len(rs), 4)
	}
	for _, v := range rs {
		op, ok := userOpRs[v.name]
		if !ok {
			t.Errorf("Opertation %s not registered", op.name)
		} else {
			if op.user != v.user {
				t.Errorf("User with more %s was incorrect, got: %s, expected: %s.", v.name, v.user, op.user)
			}
			if op.count != v.count {
				t.Errorf("Count for %s with more %s was incorrect, got: %d, expected: %d.", v.user, v.name, v.count, op.count)
			}
		}
	}
}

func printPerceltilea(percentile float64, value float64) {
	fmt.Printf("%v of the values are below %f\n", percentile, value)
}

func Test50Percentile(t *testing.T) {
	p := Percentile(data, 0.5)
	if p != 3 {
		t.Errorf("50 percent of values should be below %v, got: %v.", 3, p)
	}
}

func Test95Percentile(t *testing.T) {
	p := Percentile(data, 0.95)
	if p != 7 {
		t.Errorf("95 percent of values should be below %v, got: %v.", 3, p)
	}
}
