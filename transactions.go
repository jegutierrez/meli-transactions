package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// Transaction is used to parse transaction logs.
type Transaction struct {
	user      string
	operation string
	amount    float64
}

// OperationResult is used to store the averages by operation type.
type OperationResult struct {
	name  string
	count int
	avg   float64
}

// OperationUserResult is used to store the user with greather amount of operation.
type OperationUserResult struct {
	name  string
	user  string
	count int
}

// MeanByOperationType receive a list of transactions.
// Then calculates the average amount by operation type
// This implementation uses a linear algorithm iterating the list once
// And saving the partial averages in a map to have constant look up
func MeanByOperationType(ts *[]Transaction) map[string]OperationResult {
	rs := make(map[string]OperationResult)
	for _, t := range *ts {
		var op OperationResult
		if _, ok := rs[t.operation]; ok {
			op = rs[t.operation]
		} else {
			op = OperationResult{t.operation, 0, 0}
		}
		op.count++
		op.avg += (t.amount - op.avg) / float64(op.count)
		rs[t.operation] = op
	}
	return rs
}

func printMean(rs map[string]OperationResult) {
	fmt.Println("Average amount by operation type:")
	for _, v := range rs {
		fmt.Printf("OperationType: %s - [avg : %f]\n", v.name, v.avg)
	}
}

// LeadingUserByOperationType receive a list of transactions.
// Then the user with more transaction by operation type
// This implementation uses a linear algorithm iterating the list once
// And saving the partial counts in a map to have constant look up
func LeadingUserByOperationType(ts *[]Transaction) map[string]OperationUserResult {
	rs := make(map[string]OperationUserResult)
	userOpCount := make(map[string]map[string]int)
	for _, t := range *ts {
		var op OperationUserResult
		if _, ok := rs[t.operation]; ok {
			op = rs[t.operation]
		} else {
			op = OperationUserResult{t.operation, t.user, 0}
		}
		var usr = make(map[string]int)
		if v, ok := userOpCount[t.user]; ok {
			usr = v
			if _, ok := usr[t.operation]; ok {
				usr[t.operation]++
			} else {
				usr[t.operation] = 1
			}
			userOpCount[t.user] = usr
		} else {
			userOpCount[t.user] = map[string]int{t.operation: 1}
		}
		if t.user == op.user {
			op.count++
		} else {
			if op.count < usr[t.operation] {
				op.user = t.user
				op.count = usr[t.operation]
			}
		}
		rs[t.operation] = op
	}
	return rs
}

func printMainUsers(rs map[string]OperationUserResult) {
	fmt.Println("Users with more transactions:")
	for _, v := range rs {
		fmt.Printf("OperationType: %s - [user : %s]\n", v.name, v.user)
	}
}

// Sort transactions using quick sort algorithm.
func Sort(ts *[]Transaction) {
	if len(*ts) < 2 {
		return
	}
	left, right := 0, len(*ts)-1
	pivotIndex := rand.Int() % len(*ts)
	(*ts)[pivotIndex], (*ts)[right] = (*ts)[right], (*ts)[pivotIndex]
	for i := range *ts {
		if (*ts)[i].amount < (*ts)[right].amount {
			(*ts)[i], (*ts)[left] = (*ts)[left], (*ts)[i]
			left++
		}
	}
	(*ts)[left], (*ts)[right] = (*ts)[right], (*ts)[left]
	l := (*ts)[:left]
	r := (*ts)[left+1:]
	Sort(&l)
	Sort(&r)
	return
}

// Percentile receive a transaction list.
// Then return amount indicanting x% of the values are below it.
func Percentile(ts []Transaction, percent float64) float64 {
	Sort(&ts)
	i := int(math.Floor(float64(len(ts)) * percent))
	value := ts[i].amount
	return value
}

func printPerceltile(percentile float64, value float64) {
	fmt.Printf("%v of the values are below %f\n", percentile, value)
}

func main() {
	if len(os.Args) < 2 {
		panic("You must provide the filename as param")
	}
	fileName := os.Args[1]
	ts := GetTransactionsFromFile(fileName)
	rs := MeanByOperationType(&ts)
	printMean(rs)
	fmt.Println("----------------")
	usr := LeadingUserByOperationType(&ts)
	printMainUsers(usr)
	fmt.Println("----------------")
	perc95 := Percentile(ts, 0.95)
	printPerceltile(0.95, perc95)
}
