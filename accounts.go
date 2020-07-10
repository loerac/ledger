package ledger

import (
    "fmt"
    "hash/maphash"
)

/**
 * @brief:  Create a new accout for the ledger. The account will be a hash
 *          of the users full name, date of creation, and the initial balance.
 *
 * @arg:    fullname - Full name of the user being added
 * @arg:    initBalance - Initial balance of the user.
 *
 * @return: The hash value of the account
 **/
func (lgr Ledger) CreateAccountHash(fullname string, initBalance float64) string {
    var h maphash.Hash
    h.WriteString(fullname + GetDate() + fmt.Sprintf("%f", initBalance))
    return fmt.Sprintf("%x", h.Sum64())
}

/**
 * @brief:  Check if the given account number is valid in the ledger
 *
 * @arg:    acctNum - Account number to check if found/valid
 *
 * @return true if valid, else false
 **/
func (lgr Ledger) IsValidAccount(acctNum string) bool {
    if _, ok := lgr.AccountNum[acctNum]; !ok {
        fmt.Println("Account number not found:", acctNum)
        return false
    }

    return true
}
