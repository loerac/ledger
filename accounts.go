package ledger

import (
    "fmt"
    "log"
    "hash/maphash"
    "os"
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
 * Each account holds the path to the ledger notebook,
 * fullname of the account, and the list of entries in
 * the ledger (divided up by the account number)
 **/
type Account struct {
    Filepath    string
    Fullname    string
    Entry       []EntryItem
}

/**
 * @brief:  Create a new accout for the ledger. The account will be a hash
 *          of the users full name, date of creation, and the initial balance.
 *          The new account will then be added to the map of accounts. A new
 *          ledger notebook will also be created with the initial balance.
 *
 * @arg:    fullname - Full name of the user being added
 * @arg:    fpath - Filepath to create the new ledger notebook
 * @arg:    initBalance - Initial balance of the user.
 *
 * @return: The hash value of the account
 **/
func (lgr *Ledger) CreateAccountHash(fullname, fpath string, initBalance float64) string {
    /* Calculate the hash and then add it to the ledger */
    var h maphash.Hash
    h.WriteString(fullname + GetDate(DATE_TIME) + NumToStr(initBalance))
    acctNum := NumToStr(h.Sum64())
    lgr.Accounts[acctNum] = &Account{fpath, fullname, []EntryItem{}}

    /* Create a new ledger notebook */
    err := os.RemoveAll(fpath)
    CheckErr(err)
    f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    CheckErr(err)

    /* Add the account fullname and account number at the beginning */
    if _, err := f.Write([]byte(fullname + ":" + acctNum + "\n")); err != nil {
        log.Fatal(err)
    }

    /* Add the first entry with the initial balance */
    newEntry := fmt.Sprintf("%s:---:---:Income:0.0:%0.2f", GetDate(DATE_TIME), initBalance)
    logger.Printf("New account created for '%s -- %s': %s\n", fullname, acctNum, newEntry)
    if _, err := f.Write([]byte(newEntry + "\n")); err != nil {
        log.Fatal(err)
    }
    lgr.ReadLedger(fpath)

    return acctNum
}

/**
 * @brief:  Get the account number given the ledger notebook
 *
 * @arg:    filepath - Path to the ledger notebook
 *
 * @return: The account number of the filepath
 **/
func (lgr *Ledger) GetAcctNum(filepath string) string {
    for key, acct := range lgr.Accounts {
        if acct.Filepath == filepath {
            return key
        }
    }

    return ""
}

/**
 * @brief:  Check if the given account number is valid in the ledger
 *
 * @arg:    acctNum - Account number to check if found/valid
 *
 * @return true if valid, else false
 **/
func (lgr Ledger) IsValidAccount(acctNum string) bool {
    if _, ok := lgr.Accounts[acctNum]; !ok {
        fmt.Println("Account number not found:", acctNum)
        return false
    }

    return true
}
