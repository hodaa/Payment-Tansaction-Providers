package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TransactionsA struct {
	TransactionsA []Transaction `json:"transactions"`
}
type TransactionsB struct {
	TransactionsB []TransactionB `json:"transactions"`
}

// User struct which contains a name
// a type and a list of social links
type Transaction struct {
	Amount         int
	Currency       string
	StatusCode     int
	OrderReference string
	TransactionId  string
}

type TransactionB struct {
	Value               int
	TransactionCurrency string
	StatusCode          int
	OrderInfo           string
	PaymentId           string
}

func readFile(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func parseFileA() TransactionsA {

	byteValueA := readFile("json/flypayA.json")

	var transactionsA TransactionsA

	json.Unmarshal(byteValueA, &transactionsA)

	return transactionsA

}
func parseFileB() TransactionsB {

	byteValueB := readFile("json/flypayB.json")

	var transactionsB TransactionsB
	json.Unmarshal(byteValueB, &transactionsB)

	return transactionsB

}
