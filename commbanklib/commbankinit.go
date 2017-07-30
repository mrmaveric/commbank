package commbanklib

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Date        time.Time
	Ammount     float32
	Balance     float32
	Description string
	Debit       bool
}

type TransactionList []Transaction

func MakeTransactionList(source ...string) TransactionList {
	var transactionList TransactionList

	for _, input := range source {

		rawFile, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}

		transactions, err := csv.NewReader(rawFile).ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range transactions {
			entry := Transaction{}

			entry.Description = t[2]

			tempAmmount, err := strconv.ParseFloat(t[1], 32)
			if err != nil {
				log.Fatal(err)
			}
			entry.Ammount = float32(math.Abs(tempAmmount))

			if tempAmmount <= 0 {
				entry.Debit = true
			}

			tempBalance, err := strconv.ParseFloat(t[3], 32)
			if err != nil {
				log.Fatal(err)
			}
			entry.Balance = float32(tempBalance)

			tempDate, err := time.Parse("02/01/2006", t[0])
			if err != nil {
				log.Fatal(err)
			}
			entry.Date = tempDate
			transactionList = append(transactionList, entry)

		}
	}

	return transactionList

}

func (t TransactionList) Debits() TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Debit {
			filteredList = append(filteredList, entry)
		}
	}

	return filteredList
}

func (t TransactionList) Credits() TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Debit == false {
			filteredList = append(filteredList, entry)
		}
	}

	return filteredList
}

func (t TransactionList) Contains(sub string) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if strings.Contains(strings.ToLower(entry.Description), strings.ToLower(sub)) {
			filteredList = append(filteredList, entry)
		}
	}

	return filteredList
}

func (t TransactionList) LessThan(bar float32) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Ammount < bar {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func (t TransactionList) GreaterThan(bar float32) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Ammount > bar {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func (t TransactionList) EqualTo(bar float32) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Ammount == bar {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func (t TransactionList) After(date time.Time) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Date.After(date) {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func (t TransactionList) Before(date time.Time) TransactionList {
	var filteredList TransactionList
	for _, entry := range t {
		if entry.Date.Before(date) {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func (t TransactionList) Count() int {
	return len(t)
}

func (t TransactionList) Total() float32 {
	var total float32
	for _, entry := range t {
		total += entry.Ammount
	}
	return total
}
