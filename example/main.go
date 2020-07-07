package main

import (
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
    ledger.AddEntry("Store store", "S 456 St., Small Town, Big State", "Shopping", false, -19)
    ledger.AddEntry("Farm Big Lot", "", "Farm equipment", true, 19)
}
