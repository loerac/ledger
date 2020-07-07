package main

import (
    "bufio"
    "io"
    "os"

    lgr "github.com/loerac/ledger"
)

const LEDGER string = "notebook.lgr"

func main() {
    /* Get a new init ledger struct */
    ledger := lgr.NewLedger(LEDGER)

    /**
     * Read the ledger and add it to our struct
     * TODO: This should be in the ledger package
     **/
    f, err := os.Open(ledger.Filepath)
    defer f.Close()
    lgr.CheckErr(err)

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        if err == io.EOF {
            break
        }

        lgr.CheckErr(err)
        ledger.ParseLedgerLine(line[:len(line) -1])
    }

    /* Print the ledger out pretty */
    ledger.PrintLedger()

    /* Add to new items to the ledger, one expense and one income type */
    ledger.AddEntry("Store store", "S 456 St., Small Town, Big State", "Shopping", false, -19)
    ledger.AddEntry("Farm Big Lot", "", "Farm equipment", true, 19)
}
