package main

import (
    "fmt"

    lgr "github.com/loerac/ledger"
)

const (
    LEDGER1 string = "notebook1.lgr"
    LEDGER2 string = "notebook2.lgr"
)

func main() {
    /* Get a new init ledger struct */
    ledger := lgr.NewLedger()

    /* Read the ledger notebook from above */
    ledger.ReadLedger(LEDGER1, LEDGER2)

    /* Print the ledger out pretty */
    ledger.PrintLedger()

    /* Add to new items to the ledger, one expense and one income type */
    ledger.AddEntry("5524c5d66aeee973", "Store store", "S 456 St., Small Town, Big State", "Shopping", -19)
    ledger.AddEntry("936e1204e7b8c686", "Farm Big Lot", "", "Farm equipment", 19)

    /* Print the information on an account */
    ledger.PrintLedgerAccount("5524c5d66aeee973")

    /* Print the information into a markdown table */
    ledger.PrintToTable("936e1204e7b8c686", "")

    /* Add a new account */
    fullname := "Jimmy Johns"
    acctNum := ledger.CreateAccountHash(fullname, "notebook3.lgr", 12.90)
    fmt.Println("New account number for", fullname, "-", acctNum)
    ledger.AddEntry(acctNum, "Subway", "Subway 456 Rd., Subs City, Subs State", "Subs", -6.12)
}
