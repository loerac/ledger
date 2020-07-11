package ledger

import (
    "fmt"
    "log"
    "os"
    "sync"
)

type LogFile struct {
    Filename string
    *log.Logger
}

func NewLog(fname, prefix string) *LogFile {
    var lf   *LogFile
    var once sync.Once

    once.Do(func() {
        lf = CreateLogFile(fname, prefix)
    })

    return lf
}

func CreateLogFile(fname, prefix string) *LogFile {
    f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    CheckErr(err)

    return &LogFile{
        Filename:   fname,
        Logger:     log.New(f, prefix, log.Lshortfile|log.Lmsgprefix),
    }
}

func (lgr Ledger) OutputAccount(acctNum, fname string) {
    filename := GetDate()
    if fname != "" {
       filename = fname
    } else if (lgr.IsValidAccount(acctNum)) {
        filename = acctNum
    }
    filename += ".md"

    f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    CheckErr(err)

    lf := log.New(f, "", 0)

    lf.Println("##### Account Number:", acctNum)
    lf.Println("#")
    lf.Println("| Date                | Transfer To                                     | Description                   | Cost     | Balance  |")
    lf.Println("|---------------------|-------------------------------------------------|-------------------------------|----------|----------|")
    for _, entry := range lgr.AccountNum[acctNum] {
        ent := "|" + entry.Date + "|" + entry.Store
        if entry.Address != "" {
            ent += "<br>" + entry.Address
        }
        ent += "|" + entry.Detail + "|" + fmt.Sprintf("%0.2f", entry.Cost) + "|" + fmt.Sprintf("%0.2f", entry.Balance) + "|"
        lf.Println(ent)
    }
}
