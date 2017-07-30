package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mrmaveric/commbank/commbanklib"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "[>1 CSV file]")
		os.Exit(1)
	}

	transactions := commbanklib.MakeTransactionList(os.Args[1:]...)

	var total float32
	var count int32

	start, err := time.Parse("2/1/06", "1/7/17")
	finish, err := time.Parse("2/1/06", "1/8/17")

	if err != nil {
		log.Fatal(err)
	}

	for _, t := range transactions.Debits().After(start).Before(finish).Contains("mcd") {
		fmt.Println(t)
		total += t.Ammount
		count++
	}

	fmt.Println("Total: ", total)
	fmt.Println("Transaction Count: ", count)
}
