package ledger

import (
    "fmt"
    "os"
    "strings"
    "time"
)

/**
 * TODO: Instead of each item being an array,
 *       the Ledger should be used as an array
 **/
type Ledger struct {
    Entries uint64
    Date []string
    Store []string
    Address []string
    Detail []string
    Exchange []string
    Cost []float64
    Balance []float64

    Filepath string
}

const (
    LGR_DATE = iota
    LGR_LOC
    LGR_DETAIL
    LGR_ENTITY
    LGR_COST
    LGR_BALANCE
)

/**
 * @brief:  Create a new ledger struct object
 *
 * @arg:    lgrfp - Ledger filepath
 *
 * @return: A new created ledger struct with the file path
 **/
func NewLedger(lgrfp string) Ledger {
    lgr := Ledger{}
    lgr.Entries = 0
    lgr.Filepath = lgrfp
    fmt.Println("Filepath:", lgr.Filepath)

    return lgr
}

/**
 * @brief:  Parse the meta data of each line in the ledger, as seen below.
 *          Each data will be added to the ledger.
 *          <YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
 *
 * @args:   data - meta data
 **/
func (lgr *Ledger) ParseLedgerLine(data string) {
    split := strings.Split(data, ":")

    lgr.Entries += 1
    lgr.Date = append(lgr.Date, split[LGR_DATE])
    loc := strings.Split(split[LGR_LOC], "@")
    lgr.Store = append(lgr.Store, loc[0])
    lgr.Address = append(lgr.Address, "")
    if 2 == len(loc) {
        lgr.Address = append(lgr.Address[:len(lgr.Address) - 1], loc[1])
    }
    lgr.Detail = append(lgr.Detail, split[LGR_DETAIL])
    lgr.Exchange = append(lgr.Exchange, split[LGR_ENTITY])
    lgr.Cost = append(lgr.Cost, StrToFloat(split[LGR_COST]))
    lgr.Balance = append(lgr.Balance, StrToFloat(split[LGR_BALANCE]))
}

/**
 * @brief:  Create a new entry for the ledger. Will append entry to file.
 *
 * @arg:    store - Name of the store
 * @arg:    addr - Location of the store. Optional, can be left blank
 * @arg:    detail - Reasoning or brief detail
 * @arg:    exchange - Either "Expense" or "Income"
 * @arg:    cost - Gain or expense
 **/
func (lgr *Ledger) AddEntry(store, addr, detail string, isIncome bool, cost float64) {
    currTime := time.Now()
    date := fmt.Sprintf("%d%02d%02dT%02d%d%d",
            currTime.Year(), currTime.Month(), currTime.Day(),
            currTime.Hour(), currTime.Minute(), currTime.Second())
    balance := lgr.Balance[lgr.Entries - 1] + cost
    exchange := "Expense"
    if isIncome {
        exchange = "Income"
    }

    lgr.Entries += 1
    lgr.Date = append(lgr.Date, date)
    lgr.Store = append(lgr.Store, store)
    lgr.Address = append(lgr.Address, addr)
    lgr.Detail = append(lgr.Detail, detail)
    lgr.Exchange = append(lgr.Exchange, exchange)
    lgr.Cost = append(lgr.Cost, cost)
    lgr.Balance = append(lgr.Balance, balance)

    exchange += ":"
    if isIncome {
        exchange += "+"
    }

    newEntry := fmt.Sprintf("%s:%s@%s:%s:%s%0.2f:%0.2f", date, store, addr, detail, exchange, cost, balance)
    if "" == addr {
        newEntry = fmt.Sprintf("%s:%s:%s:%s%0.2f:%0.2f", date, store, detail, exchange, cost, balance)
    }

    fmt.Println("Added new ledger entry:", newEntry)
    lgr.PrintLedgerItem(lgr.Entries - 1)
    lgr.AddToLedger(newEntry)
}

/**
 * @brief:  Add the new entry to the ledger file
 *
 * @arg:    entry - Entry in the meta data format
 *          <YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
 **/
func (lgr Ledger) AddToLedger(entry string) {
    f, err := os.OpenFile(lgr.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    CheckErr(err)

    if _, err := f.Write([]byte(entry + "\n")); err != nil {
        f.Close()
        panic(err)
    }
}

/**
 * @brief:  Print individual item for ledger. To print the complet ledger,
 *          run `PrintLedger()`
 *
 * @arg:    item - Item to print in the ledger
 **/
func (lgr Ledger) PrintLedgerItem(item uint64) {
    if item > lgr.Entries {
        return
    }

    fmt.Println(lgr.Date[item], lgr.Store[item])

    entitySign := ""
    if lgr.Exchange[item] == "Income" {
        entitySign = "+"
    }
    fmt.Printf("\t\t%s: %s %s%f\n", lgr.Exchange[item], lgr.Detail[item], entitySign, lgr.Cost[item])
    fmt.Printf("\t\tBalance: %f\n", lgr.Balance[item])

    if lgr.Address[item] != "" {
        fmt.Println("\t\tLocation:", lgr.Address[item])
    }
}

/**
 * @brief:  Print the complete ledger. To print individual item of the ledger,
 *          run `PrintLedgerItem()`
 **/
func (lgr Ledger) PrintLedger() {
    var i uint64 = 0
    for ; i < lgr.Entries; i++ {
        lgr.PrintLedgerItem(i)
    }
}
