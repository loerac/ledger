package ledger

import (
    "fmt"
    "hash/maphash"
)

func (lgr Ledger) CreateAccountHash(fullname string, initBalance float64) string {
    var h maphash.Hash
    h.WriteString(fullname + GetDate() + fmt.Sprintf("%f", initBalance))
    return fmt.Sprintf("%x", h.Sum64())
}

func (lgr Ledger) IsValidAccount(acctNum string) bool {
    if _, ok := lgr.AccountNum[acctNum]; !ok {
        fmt.Println("Account number not found:", acctNum)
        return false
    }

    return true
}
