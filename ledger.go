package ledger

import (
    "fmt"
    "os"
    "strings"
    "time"
)

type EntryItem struct {
    Date string
    Store string
    Address string
    Detail string
    Exchange string
    Cost float64
    Balance float64
}

type Ledger struct {
    NumEntries uint64
    Filepath string
    Entry []EntryItem
}

const (
    LGR_DATE = iota
    LGR_LOC
    LGR_DETAIL
    LGR_EXCHANGE
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
    lgr.NumEntries = 0
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
    entry := EntryItem{}

    lgr.NumEntries += 1
    entry.Date = split[LGR_DATE]
    loc := strings.Split(split[LGR_LOC], "@")
    entry.Store = loc[0]
    entry.Address = ""
    if 2 == len(loc) {
        entry.Address = loc[1]
    }
    entry.Detail = split[LGR_DETAIL]
    entry.Exchange = split[LGR_EXCHANGE]
    entry.Cost = StrToFloat(split[LGR_COST])
    entry.Balance = StrToFloat(split[LGR_BALANCE])

    lgr.Entry = append(lgr.Entry, entry)
}

/**
 * @brief:  Create a new entry for the ledger. Will append entry to file.
 *
 * @arg:    store - Name of the store
 * @arg:    addr - Location of the store. Optional, can be left blank
 * @arg:    detail - Reasoning or brief detail
 * @arg:    exchange - Either true ("Income") or false ("Expense")
 * @arg:    cost - Gain or expense
 **/
func (lgr *Ledger) AddEntry(store, addr, detail string, isIncome bool, cost float64) {
    entry := EntryItem{}
    currTime := time.Now()
    date := fmt.Sprintf("%d%02d%02dT%02d%d%d",
            currTime.Year(), currTime.Month(), currTime.Day(),
            currTime.Hour(), currTime.Minute(), currTime.Second())
    balance := lgr.Entry[lgr.NumEntries - 1].Balance + cost
    exchange := "Expense"
    if isIncome {
        exchange = "Income"
    }

    lgr.NumEntries += 1
    entry.Date = date
    entry.Store = store
    entry.Address = addr
    entry.Detail = detail
    entry.Exchange = exchange
    entry.Cost = cost
    entry.Balance = balance
    lgr.Entry = append(lgr.Entry, entry)

    exchange += ":"
    if isIncome {
        exchange += "+"
    }

    newEntry := fmt.Sprintf("%s:%s@%s:%s:%s%0.2f:%0.2f", date, store, addr, detail, exchange, cost, balance)
    if "" == addr {
        newEntry = fmt.Sprintf("%s:%s:%s:%s%0.2f:%0.2f", date, store, detail, exchange, cost, balance)
    }

    fmt.Println("Added new ledger entry:", newEntry)
    lgr.PrintLedgerItem(lgr.NumEntries - 1)
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
 * @brief:  Print individual item for ledger. To print the complete ledger,
 *          run `PrintLedger()`
 *
 * @arg:    item - Item to print in the ledger
 **/
func (lgr Ledger) PrintLedgerItem(item uint64) {
    if item > lgr.NumEntries {
        return
    }

    entry := lgr.Entry
    fmt.Println(entry[item].Date, entry[item].Store)

    entitySign := ""
    if entry[item].Exchange == "Income" {
        entitySign = "+"
    }
    fmt.Printf("\t\t%s: %s %s%f\n", entry[item].Exchange, entry[item].Detail, entitySign, entry[item].Cost)
    fmt.Printf("\t\tBalance: %f\n", entry[item].Balance)

    if entry[item].Address != "" {
        fmt.Println("\t\tLocation:", entry[item].Address)
    }
}

/**
 * @brief:  Print the complete ledger. To print individual item of the ledger,
 *          run `PrintLedgerItem()`
 **/
func (lgr Ledger) PrintLedger() {
    var i uint64 = 0
    for ; i < lgr.NumEntries; i++ {
        lgr.PrintLedgerItem(i)
    }
}
