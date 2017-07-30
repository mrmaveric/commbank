package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrmaveric/commbank/commbanklib"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please provied CSV files")
	}

	transactions := commbanklib.MakeTransactionList(os.Args[1:]...)

	fmt.Println(transactions[0].Date, transactions[0].Ammount, transactions[0].Description, transactions[0].Balance)

	var total float32

	for _, t := range transactions {
		total += t.Ammount
	}

	//total = total / float32(tally)

	fmt.Println("Total: ", total)
}
