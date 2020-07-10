package ledger

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/**
 * Each entry must have the date that the entry was created,
 * location that the exchange was done, a reason for the exchange,
 * exchange type (income or expense), cost of exchange, and
 * current balance.
 * The address is optional, this could be left blank.
 **/
type EntryItem struct {
    Date        string
    Store       string
    Address     string
    Detail      string
    Exchange    string
    Cost        float64
    Balance     float64
}

/**
 * Ledger holds the filepath to the ledger notebook,
 * and the list of entries in the ledger (divided up by
 * the account number)
 **/
type Ledger struct {
    Filepath    string
    AccountNum  map[string][]EntryItem
}

const (
    /* Length of the metadata */
    METADATA_LEN int = 7
)

/**
 * Enum index value of the metadata
 * Account number starts at index 0
 **/
const (
    LGR_ACCOUNT_NUM = iota
    LGR_DATE
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
    lgr.Filepath = lgrfp
    lgr.AccountNum = make(map[string][]EntryItem)

    return lgr
}

/**
 * @brief:  Parse the meta data of each line in the ledger, as seen below.
 *          Each data will be added to the ledger.
 *          <ACCOUNT-NUM>:<YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
 *
 * @arg:    data - meta data
 **/
func (lgr *Ledger) ParseLedgerLine(data string) {
    split := strings.Split(data, ":")
    if (METADATA_LEN != len(split)) {
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

    lgr.AccountNum[split[LGR_ACCOUNT_NUM]] =
        append(
            lgr.AccountNum[split[LGR_ACCOUNT_NUM]],
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
 * @arg:    exchange - Either true ("Income") or false ("Expense")
 * @arg:    cost - Gain or expense
 **/
func (lgr *Ledger) AddEntry(acctNum, store, addr, detail string, isIncome bool, cost float64) {
    if !lgr.IsValidAccount(acctNum) {
        return
    }

    date := GetDate()
    balance := lgr.AccountNum[acctNum][len(lgr.AccountNum) - 1].Balance + cost
    exchange := "Expense"
    if isIncome {
        exchange = "Income"
    }

    lgr.AccountNum[acctNum] =
        append(
            lgr.AccountNum[acctNum],
            EntryItem{date, store, addr, detail, exchange, cost, balance},
        )

    newEntry := fmt.Sprintf("%s:%s:%s@%s:%s:%s%0.2f:%0.2f", acctNum, date, store, addr, detail, exchange, cost, balance)
    if "" == addr {
        newEntry = fmt.Sprintf("%s:%s:%s:%s:%s%0.2f:%0.2f", acctNum, date, store, detail, exchange, cost, balance)
    }

    fmt.Println("Added new ledger entry:")
    lgr.PrintLedgerItem(lgr.AccountNum[acctNum][len(lgr.AccountNum) - 1])
    fmt.Println()
    lgr.AddToLedger(newEntry)
}

/**
 * @brief:  Add the new entry to the ledger file
 *
 * @arg:    entry - Entry in the meta data format
 *          <ACCOUNT-NUM>:<YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
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

/***
 * @brief:  Read the ledger that was supplied, and add the entries
 *          to the EntryItem struct
 ***/
func (lgr *Ledger) ReadLedger() {
    f, err := os.Open(lgr.Filepath)
    defer f.Close()
    CheckErr(err)

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        if err == io.EOF {
            break
        }

        CheckErr(err)
        lgr.ParseLedgerLine(line[:len(line) -1])
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

    fmt.Println("Account Number:", acctNum)
    fmt.Println("================================")
    for _, v := range lgr.AccountNum[acctNum] {
        lgr.PrintLedgerItem(v)
    }
    fmt.Println()
}

/**
 * @brief:  Print the complete ledger. To print individual item of the ledger,
 *          run `PrintLedgerItem()`
 **/
func (lgr Ledger) PrintLedger() {
    for key := range lgr.AccountNum {
        lgr.PrintLedgerAccount(key)
    }
}
