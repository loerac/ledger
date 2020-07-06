package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"

    ledger "package-project/ledger"
)

func check(e error) {
    if nil != e {
        panic(e)
    }
}

func main() {
    ledger := Ledger{}
    f, err := os.Open("notebook.lgr")
    defer f.Close()
    check(err)

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        if err == io.EOF {
            break
        }

        check(err)
        ledger.ParseLedgerLine(line[:len(line) -1])
    }

    ledger.PrintLedger()
}
