package main

import (
    "fmt"

    lgr "github.com/loerac/ledger"
)

const LEDGER string = "notebook.lgr"

func main() {
    /* Get a new init ledger struct */
    ledger := lgr.NewLedger(LEDGER)

    /* Read the ledger notebook from above */
    ledger.ReadLedger()

    /* Print the ledger out pretty */
    ledger.PrintLedger()

    /* Add to new items to the ledger, one expense and one income type */
    ledger.AddEntry("5524c5d66aeee973", "Store store", "S 456 St., Small Town, Big State", "Shopping", -19)
    ledger.AddEntry("936e1204e7b8c686", "Farm Big Lot", "", "Farm equipment", 19)

    /* Print the information on an account */
    ledger.PrintLedgerAccount("5524c5d66aeee973")
    ledger.PrintLedgerAccount("936e1204e7b8c686")

    /* Add a new account */
    fullname := "Christian Loera"
    fmt.Println("New account number for", fullname, "-", ledger.CreateAccountHash(fullname, 12.90))
}
