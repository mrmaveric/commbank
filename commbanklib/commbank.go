package commbanklib

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Date        time.Time
	Ammount     float32
	Balance     float32
	Description string
}

func MakeTransactionList(source ...string) []Transaction {
	var transactionList []Transaction

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
			entry.Ammount = float32(tempAmmount)

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
