package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type ParsedTransaction struct {
	Amount        int
	Currency      string
	StatusCode    int
	Order         string
	TransactionId string
	Provider      string
}

var tagsList []ParsedTransaction

func getAllTransaction() []ParsedTransaction {

	transactionsA := parseFileA()

	for _, details := range transactionsA.TransactionsA {

		tagsList = append(tagsList, ParsedTransaction{
			Amount:        details.Amount,
			Currency:      details.Currency,
			StatusCode:    details.StatusCode,
			Order:         details.OrderReference,
			TransactionId: details.TransactionId,
			Provider:      "flypayA",
		})
	}

	transactionsB := parseFileB()

	for _, details := range transactionsB.TransactionsB {

		tagsList = append(tagsList, ParsedTransaction{
			Amount:        details.Value,
			Currency:      details.TransactionCurrency,
			StatusCode:    details.StatusCode,
			Order:         details.OrderInfo,
			TransactionId: details.PaymentId,
			Provider:      "flypayB",
		})
	}

	return tagsList

}

/**
*
**/
func get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	transactions := getAllTransaction()

	queries := r.URL.Query()

	if val, ok := queries["provider"]; ok {
		transactions = searchByProvider(val[0], transactions)
	}

	if val, ok := queries["statusCode"]; ok {
		transactions = searchByStatusCode(val[0], transactions)
	}

	if val, ok := queries["currency"]; ok {
		transactions = searchByCurrency(val[0], transactions)
	}

	if minVal, ok := queries["amountMin"]; ok {
		if maxVal, ok := queries["amountMin"]; ok {
			fmt.Println(queries["amountMin"])
			transactions = searchByAmount(minVal[0], maxVal[0], transactions)
		}

	}

	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(http.StatusOK)
}

func searchByProvider(val string, transactions []ParsedTransaction) []ParsedTransaction {
	var result []ParsedTransaction

	for _, item := range transactions {
		if item.Provider == val {
			result = append(result, item)

		}
	}
	return result
}

func searchByCurrency(val string, transactions []ParsedTransaction) []ParsedTransaction {
	var result []ParsedTransaction

	for _, item := range transactions {
		if item.Currency == val {
			result = append(result, item)

		}
	}
	return result
}

func searchByStatusCode(val string, transactions []ParsedTransaction) []ParsedTransaction {
	var result []ParsedTransaction

	fmt.Println(reflect.TypeOf(val[0]))
	fmt.Println(reflect.TypeOf("authorised"))

	for _, item := range transactions {

		if val == "authorised" {
			if item.StatusCode == 1 || item.StatusCode == 100 {
				result = append(result, item)

			}
		}
		if val == "decline" {
			if item.StatusCode == 2 || item.StatusCode == 200 {
				result = append(result, item)

			}
		}
		if val == "refunded" {
			if item.StatusCode == 3 || item.StatusCode == 300 {
				result = append(result, item)

			}
		}
	}

	return result
}

func searchByAmount(valMin string, valMax string, transactions []ParsedTransaction) []ParsedTransaction {
	fmt.Println(reflect.TypeOf(valMin))
	var result []ParsedTransaction

	//	for _, item := range transactions {
	//		if item.Amount >= valMin && item.Amount <= valMax {
	//			result = append(result, item)

	//		}
	//	}
	return result
}
