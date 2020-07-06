package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)


type Ledger struct {
    Date []string
    Store []string
    Address []string
    Data []string
    Entity []string
    Cost []string
    Balance []string
}

const (
    LGR_DATE = iota
    LGR_LOC
    LGR_DATA
    LGR_ENTITY
    LGR_COST
    LGR_BALANCE
)

func check(e error) {
    if nil != e {
        panic(e)
    }
}

func (lgr *Ledger) ParseLedgerLine(data string) {
    split := strings.Split(data, ":")

    lgr.Date = append(lgr.Date, split[LGR_DATE])
    loc := strings.Split(split[LGR_LOC], "@")
    lgr.Store = append(lgr.Store, loc[0])
    lgr.Address = append(lgr.Address, "NULL")
    if 2 == len(loc) {
        lgr.Address = append(lgr.Address[:len(lgr.Address) - 1], loc[1])
    }
    lgr.Data = append(lgr.Data, split[LGR_DATA])
    lgr.Entity = append(lgr.Entity, split[LGR_ENTITY])
    lgr.Cost = append(lgr.Cost, split[LGR_COST])
    lgr.Balance = append(lgr.Balance, split[LGR_BALANCE])
}

func (lgr Ledger) PrintLedger() {
    for key := range lgr.Date {
        fmt.Println(lgr.Date[key], lgr.Store[key])

        entitySign := "-"
        if lgr.Entity[key] == "Income" {
            entitySign = "+"
        }
        fmt.Printf("\t\t%s: %s %s%s\n", lgr.Entity[key], lgr.Data[key], entitySign, lgr.Cost[key])
        fmt.Printf("\t\tBalance: %s\n", lgr.Balance[key])

        if lgr.Address[key] != "NULL" {
            fmt.Println("\t\tLocation:", lgr.Address[key])
        }
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
