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


/**
 * @brief:  Initialize a log file
 *
 * @arg:    fname - Log file file name
 * @arg:    prefix - Name for the logs
 *
 * @return: LogFile struct pointer
 **/
func NewLog(fname, prefix string) *LogFile {
    var lf   *LogFile
    var once sync.Once

    once.Do(func() {
        lf = CreateLogFile(fname, prefix)
    })

    return lf
}

/**
 * @brief:  Creates a global log file
 *
 * @arg:    fname - Log file file name
 * @arg:    prefix - Name for the logs
 *
 * @return: LogFile struct pointer
 **/
func CreateLogFile(fname, prefix string) *LogFile {
    filename := ternary(fname != "", fname, GetDate(DATE_TIME) + ".log").(string)
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    CheckErr(err)

    return &LogFile{
        Filename:   filename,
        Logger:     log.New(f, prefix, log.LstdFlags|log.Lmsgprefix),
    }
}

/**
 * @brief:  Save the ledger data into a table, the data will
 *          be saved into a markdown table. To print a all
 *          the table, see `PrintAllToTable()`
 *
 * @arg:    fname - File name for the markdown file, optional
 * @arg:    acctNum - Account name to save the data if fname is not present
 **/
func (lgr Ledger) PrintToTable(acctNum, fname string) {
    filename := GetDate(DATE_TIME)
    if fname != "" {
       filename = fname
    } else if (lgr.IsValidAccount(acctNum)) {
        filename = acctNum
    }
    filename += ".md"

    err := os.RemoveAll(filename)
    CheckErr(err)
    f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    CheckErr(err)

    lf := log.New(f, "", 0)
    /* Section below is used for Markdown in Go-Hugo
    lf.Println("---")
    lf.Println("title:", lgr.Accounts[acctNum].Fullname)
    lf.Println("date:", GetDate(DATE_ONLY))
    lf.Printf("description: Last transaction - '%s $%0.2f'", lgr.Accounts[acctNum].Entry[len(lgr.Accounts[acctNum].Entry) - 1].Detail, lgr.Accounts[acctNum].Entry[len(lgr.Accounts[acctNum].Entry) - 1].Cost)
    lf.Println("---")
    */

    lf.Println("##### Account Name:", lgr.Accounts[acctNum].Fullname)
    lf.Println("##### Account Number:", acctNum)
    lf.Println("| Date | Transfer To | Description | Cost | Balance |")
    lf.Println("|---|---|---|---|---|")

    total_entries := len(lgr.Accounts[acctNum].Entry) - 1
    for i := total_entries; i >= 0; i-- {
        entry := lgr.Accounts[acctNum].Entry[i]

        ent := fmt.Sprintf ("|%s|%s|%s|%s|%s|",
            FormatDate(entry.Date),
            entry.Store,
            entry.Detail,
            fmt.Sprintf("$%0.2f", entry.Cost),
            fmt.Sprintf("$%0.2f", entry.Balance))
        if entry.Address != "" {
            ent = fmt.Sprintf ("|%s|%s<br>*@%s*|%s|%s|%s|",
                FormatDate(entry.Date),
                entry.Store,
                entry.Address,
                entry.Detail,
                fmt.Sprintf("$%0.2f", entry.Cost),
                fmt.Sprintf("$%0.2f",
                entry.Balance))
        }
        lf.Println(ent)
    }
}

/**
 * @brief:  Prints all the ledgers into a markdown table
 *          To print a single table, see `PrintToTable()`
 **/
func (lgr Ledger) PrintAllToTable() {
    for key := range lgr.Accounts {
        lgr.PrintToTable(key, "")
    }
}
