package utils

import (
	"stori/src/types"
	"strconv"
	"strings"
	"time"
)

var client types.TClient

func Summary(data chan types.TTransaction) types.TClient {
	client = types.TClient{}

	var transactions []types.TMovements = nil

	for r := range data {
		var row types.TMovements
		row.Id = r.Id
		row.Date = dateFormat(r.Date)
		row.Transaction = r.Transaction

		transactions = append(transactions, row)
	}

	groupByDate(transactions)

	return client

}

func dateFormat(date string) time.Time {

	dateStr := strings.Split(date, "/")
	month, _ := strconv.Atoi(dateStr[0])
	day, _ := strconv.Atoi(dateStr[1])

	newDate := time.Date(
		time.Now().Year(),
		time.Month(month),
		day, 0, 0, 0, 0, time.UTC)
	return newDate
}

func groupByDate(trx []types.TMovements) {
	creditAmount := 0.0
	debitAmount := 0.0

	var numCreditTrx int = 0
	var numDebitTrx int = 0
	var totalBalance float64 = 0.0

	var rows []types.TSummary

	for _, t := range trx {
		exists := monthExists(t.Date.Month().String(), rows)

		if strings.ContainsAny(t.Transaction, "-") {
			t.CreditOrDebit = "Debit"
			convertToFloat, _ := strconv.ParseFloat(strings.Replace(t.Transaction, "-", "", -1), 64)
			debitAmount = debitAmount + convertToFloat
			totalBalance = totalBalance - convertToFloat
			numDebitTrx++
		} else {
			t.CreditOrDebit = "Credit"
			convertToFloat, _ := strconv.ParseFloat(strings.Replace(t.Transaction, "-", "", -1), 64)
			creditAmount = creditAmount + convertToFloat
			totalBalance = totalBalance + convertToFloat
			numCreditTrx++
		}

		if exists != -1 {
			rows[exists].Trxs = append(rows[exists].Trxs, t)
		} else {
			var newMonth types.TSummary
			newMonth.Month = t.Date.Month().String()

			newMonth.Trxs = append(newMonth.Trxs, t)
			rows = append(rows, newMonth)
		}
	}

	client.AvgDebit = debitAmount / float64(numDebitTrx)
	client.AvgCredit = creditAmount / float64(numCreditTrx)
	client.GroupedBy = rows
	client.Transactions = trx
	client.TotalBalance = totalBalance

}

func monthExists(month string, summary []types.TSummary) int {
	for key, exists := range summary {
		if exists.Month == month {
			return key
		}
	}
	return -1
}
