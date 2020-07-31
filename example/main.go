package main

import (
    lgr "github.com/loerac/ledger"
)

const (
    LEDGER1 string = "notebook1.lgr"
    LEDGER2 string = "notebook2.lgr"
)

func main() {
    /* Get a new init ledger struct */
    ledger := lgr.NewLedger()

    /* Parse the user arguments */
    ledger.ArgumentParser()
}
