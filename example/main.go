package main

import (
    "fmt"

    lgr "github.com/loerac/ledger"
)

const LEDGER string = "notebook.lgr"

func main() {
    /* Get a new init ledger struct */
    ledger := lgr.NewLedger(LEDGER, "")

    /* Read the ledger notebook from above */
    ledger.ReadLedger()

    /* Print the ledger out pretty */
    ledger.PrintLedger()

    /* Add to new items to the ledger, one expense and one income type */
    ledger.AddEntry("5524c5d66aeee973", "Store store", "S 456 St., Small Town, Big State", "Shopping", -19)
    ledger.AddEntry("936e1204e7b8c686", "Farm Big Lot", "", "Farm equipment", 19)

    /* Print the information on an account */
    ledger.PrintLedgerAccount("5524c5d66aeee973")

    /* Print the information into a markdown table */
    ledger.PrintToTable("936e1204e7b8c686", "ledger-table")

    /* Add a new account */
    fullname := "Christian Loera"
    hash := ledger.CreateAccountHash(fullname, 12.90)
    ledger.AddEntry(hash, "Amazon", "", "Bright White Ties", -2.90)
    fmt.Println("New account number for", fullname, "-", hash)
}
