package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var reLog = regexp.MustCompile(`(?m)\[[^\[\]]*:(.*?)\]`)

// ParseLog receive a log line and convert it to a transaction struct.
func ParseLog(t *Transaction, log string) {
	match := reLog.FindAllStringSubmatch(log, -1)
	var user, operation, amount string
	if len(match) > 0 && len(match[0]) > 1 {
		user = match[0][1]
	}
	if len(match) > 1 && len(match[1]) > 1 {
		operation = match[1][1]
	}
	if len(match) > 2 && len(match[2]) > 1 {
		amount = match[2][1]
	}
	t.user = user
	t.operation = operation
	v, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		t.amount = 0
	} else {
		t.amount = v
	}
}

// ParseMultilineLog receive a transaction slice and a reader.
// Then return each line parsed to a transaction struct
func ParseMultilineLog(ts *[]Transaction, r bufio.Reader) {
	line, err := r.ReadString('\n')
	for err == nil {
		t := &Transaction{}
		ParseLog(t, line)
		if t.operation != "" && t.user != "" && t.amount != 0 {
			*ts = append(*ts, *t)
		}
		line, err = r.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

// GetTransactionsFromFile parse all transactions from a given file.
func GetTransactionsFromFile(fileName string) []Transaction {
	ts := []Transaction{}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return ts
	}
	defer f.Close()
	r := bufio.NewReader(f)
	ParseMultilineLog(&ts, *r)
	return ts
}
