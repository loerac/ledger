package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"

    lgr "github.com/loerac/ledger"
)

const (
    LEDGER1 string = "notebook1.lgr"
    LEDGER2 string = "notebook2.lgr"
    NOTEBOOK_DESC string = "The ledger notebook"
    PRINTPRETTY_DESC string = "Print the ledger notebook pretty"
    PRINTTABLE_DESC string = "Print the ledger notebook to a markdown table\nNote: You can pass an output markdown file name"
    ADDENTRY_DESC string = "Add a new entry to the ledger notebook\nNote: Ledger notebook is required"
    NEWACCOUNT_DESC string = "Create a new account"
)

var (
    notebook    string
    printPretty bool
    printTable  string
    addEntry    bool
    createAccount  bool
    transaction string
    location    string
    detail      string
    cost        float64
)

func readBuf(prompt string, required bool) string {
        text := ""
        for {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print(prompt + " > ")
            text , _ = reader.ReadString('\n')
            text = strings.Replace(text, "\n", "", -1)

            if required && text == "" {
                fmt.Println("Error: input required, none given")
            } else {
                break
            }
        }

        return text
}

func main() {
    flag.StringVar(&notebook, "ledger", "", NOTEBOOK_DESC)
    flag.StringVar(&notebook, "l", "", NOTEBOOK_DESC + " (Short-hand)")
    flag.BoolVar(&printPretty, "print-pretty", false, PRINTPRETTY_DESC)
    flag.BoolVar(&printPretty, "pp", false, PRINTPRETTY_DESC + " (Short-hand)")
    flag.StringVar(&printTable, "print-table", "", PRINTTABLE_DESC)
    flag.StringVar(&printTable, "pt", "", PRINTTABLE_DESC + " (Short-hand)")
    flag.BoolVar(&addEntry, "entry", false, ADDENTRY_DESC)
    flag.BoolVar(&addEntry, "e", false, ADDENTRY_DESC + " (Short-hand)")
    flag.BoolVar(&createAccount, "account", false, NEWACCOUNT_DESC)
    flag.BoolVar(&createAccount, "a", false, NEWACCOUNT_DESC + " (Short-hand)")
    flag.Parse()

    if notebook == "" && !createAccount {
        fmt.Println("Error: Missing ledger notebook")
        flag.PrintDefaults()
        os.Exit(1)
    }

    /* Get a new init ledger struct */
    ledger := lgr.NewLedger()

    /* Create a new account */
    if createAccount {
        fullname := readBuf("Enter the full name of the new account", true)
        balance := lgr.StrToFloat(readBuf("Enter initial balance for " + fullname + " account", true))
        notebook = fullname + ".lgr"
        notebook = strings.ToLower(strings.Replace(notebook, " ", "-", -1))
        ledger.CreateAccountHash(fullname, notebook, balance)
    }

    /* Read the ledger notebook from above */
    ledger.ReadLedger(notebook)
    acctNum := ledger.GetAcctNum(notebook)

    /* Add to new items to the ledger, one expense and one income type */
    if addEntry {
        transaction = readBuf("Enter area of transaction", true)
        location = readBuf("Enter area of location of transaction", false)
        detail = readBuf("Enter details of transaction", true)
        cost = lgr.StrToFloat(readBuf("Enter cost of transaction", true))
        ledger.AddEntry(acctNum, transaction, location, detail, cost)
    }

    /* Print the ledger out pretty */
    if printPretty {
        ledger.PrintLedger()
    }

    /* Print the information into a markdown table */
    if printTable != "" {
        ledger.PrintToTable(acctNum, printTable)
    }
}
