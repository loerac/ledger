package ledger

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/* Currently, the ledger holds a map of the accounts. */
type Ledger struct {
    Accounts    map[string]*Account
}

/**
 * Enum index value of the metadata
 * Account number starts at index 0
 **/
const (
    LGR_DATE = iota
    LGR_LOC
    LGR_DETAIL
    LGR_EXCHANGE
    LGR_COST
    LGR_BALANCE

    /* Length of the ledger metadata */
    METADATA_LEDGER_LEN
)

const (
    ACCT_NAME = iota
    ACCT_NUMB

    /* Length of account metadata */
    METADATA_ACCT_LEN
)

var logger *LogFile

/**
 * @brief:  Create a new ledger struct object
 *
 * @return: A new created ledger struct with the file path
 **/
func NewLedger() Ledger {
    lgr := Ledger{}
    lgr.Accounts = make(map[string]*Account)
    logger = NewLog("ledger.log", "ledger - ")

    return lgr
}

/**
 * @brief:  Parse the meta data of each line in the ledger, as seen below.
 *          Each data will be added to the ledger.
 *          <YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
 *
 * @arg:    data - meta data
 * @arg:    acctNum - Account Number of the entry
 **/
func (lgr *Ledger) ParseLedgerLine(data, acctNum string) {
    split := strings.Split(data, ":")
    if (METADATA_LEDGER_LEN != len(split)) {
        fmt.Println("Malformed entry:", data)
        return
    }

    date := split[LGR_DATE]
    loc := strings.Split(split[LGR_LOC], "@")
    store := loc[0]
    address := ""
    if 2 == len(loc) {
        address = loc[1]
    }
    detail := split[LGR_DETAIL]
    exchange := split[LGR_EXCHANGE]
    cost := StrToFloat(split[LGR_COST])
    balance := StrToFloat(split[LGR_BALANCE])

    lgr.Accounts[acctNum].Entry =
        append(
            lgr.Accounts[acctNum].Entry,
            EntryItem{date, store, address, detail, exchange, cost, balance},
        )
}

/**
 * @brief:  Create a new entry for the ledger. Will append entry to file.
 *
 * @arg:    acctNum - Account Number of the entry
 * @arg:    store - Name of the store
 * @arg:    addr - Location of the store. Optional, can be left blank
 * @arg:    detail - Reasoning or brief detail
 * @arg:    cost - Gain or expense
 **/
func (lgr *Ledger) AddEntry(acctNum, store, addr, detail string, cost float64) {
    if !lgr.IsValidAccount(acctNum) {
        return
    }

    date := GetDate()
    balance := lgr.Accounts[acctNum].Entry[len(lgr.Accounts[acctNum].Entry) - 1].Balance + cost
    exchange := ternary(cost < 0.00, "Expense", "Income").(string)

    lgr.Accounts[acctNum].Entry =
        append(
            lgr.Accounts[acctNum].Entry,
            EntryItem{date, store, addr, detail, exchange, cost, balance},
        )

    newEntry := fmt.Sprintf("%s:%s@%s:%s:%s:%0.2f:%0.2f", date, store, addr, detail, exchange, cost, balance)
    if "" == addr {
        newEntry = fmt.Sprintf("%s:%s:%s:%s:%0.2f:%0.2f", date, store, detail, exchange, cost, balance)
    }

    logger.Printf("Added new ledger entry for %s: %s\n", lgr.Accounts[acctNum].Fullname, newEntry)
    fmt.Println("Added new ledger entry:")
    lgr.PrintLedgerItem(lgr.Accounts[acctNum].Entry[len(lgr.Accounts[acctNum].Entry) - 1])
    fmt.Println()
    lgr.AddToLedger(newEntry, acctNum)
}

/**
 * @brief:  Add the new entry to the ledger file
 *
 * @arg:    entry - Entry in the meta data format
 *          <YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
 * @arg:    acctNum - Account Number of the entry
 **/
func (lgr Ledger) AddToLedger(entry, acctNum string) {
    f, err := os.OpenFile(lgr.Accounts[acctNum].Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    CheckErr(err)

    if _, err := f.Write([]byte(entry + "\n")); err != nil {
        f.Close()
        panic(err)
    }
}

/***
 * @brief:  Read the ledger that was supplied, and add the entries
 *          to the EntryItem struct
 *
 * @arg:    fpaths - variadic filepaths of the ledger notebook
 ***/
func (lgr *Ledger) ReadLedger(fpaths ...string) {
    for _, fpath := range fpaths {
        f, err := os.Open(fpath)
        defer f.Close()
        CheckErr(err)

        rd := bufio.NewReader(f)

        /* Get the user name and account number */
        line, err := rd.ReadString('\n')
        split := strings.Split(line[:len(line) -1], ":")
        if METADATA_ACCT_LEN != len(split) {
            fmt.Println("Malformed ledger:", fpath)
            return
        }
        lgr.Accounts[split[ACCT_NUMB]] = &Account{fpath, split[ACCT_NAME], []EntryItem{}}

        for {
            line, err := rd.ReadString('\n')
            if err == io.EOF {
                break
            }

            CheckErr(err)
            lgr.ParseLedgerLine(line[:len(line) -1], split[ACCT_NUMB])
        }
    }
}

/**
 * @brief:  Print individual item for ledger. To print the complete ledger,
 *          run `PrintLedger()`
 *
 * @arg:    entry - Entry in the meta data format
 **/
func (lgr Ledger) PrintLedgerItem(entry EntryItem) {
    fmt.Println(entry.Date, entry.Store)

    fmt.Printf("\t\t%s: %s $%0.2f\n", entry.Exchange, entry.Detail, entry.Cost)
    fmt.Printf("\t\tBalance: $%0.2f\n", entry.Balance)

    if entry.Address != "" {
        fmt.Println("\t\tLocation:", entry.Address)
    }
}

/**
 * @brief:  Print the entire ledger for an account
 *
 * @arg:    acctNum - Account Number of the entry
 **/
func (lgr Ledger) PrintLedgerAccount(acctNum string) {
    if !lgr.IsValidAccount(acctNum) {
        return
    }

    fmt.Println("Account Name:", lgr.Accounts[acctNum].Fullname)
    fmt.Println("Account Number:", acctNum)
    fmt.Println("================================")
    for _, v := range lgr.Accounts[acctNum].Entry {
        lgr.PrintLedgerItem(v)
    }
    fmt.Println()
}

/**
 * @brief:  Print the complete ledger. To print individual item of the ledger,
 *          run `PrintLedgerItem()`
 **/
func (lgr Ledger) PrintLedger() {
    for key := range lgr.Accounts {
        lgr.PrintLedgerAccount(key)
    }
}
