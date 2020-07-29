package ledger

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

const (
    LEDGER_DESC string = "The ledger notebook"
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

func (lgr *Ledger) ArgumentParser() {
    flag.StringVar(&notebook, "l", "", LEDGER_DESC)
    flag.BoolVar(&printPretty, "pp", false, PRINTPRETTY_DESC)
    flag.StringVar(&printTable, "pt", "", PRINTTABLE_DESC)
    flag.BoolVar(&addEntry, "add-entry", false, ADDENTRY_DESC)
    flag.BoolVar(&createAccount, "new-acct", false, NEWACCOUNT_DESC)
    flag.Parse()

    if notebook == "" && !createAccount {
        fmt.Println("Error: Missing ledger notebook")
        flag.PrintDefaults()
        os.Exit(1)
    }

    /* Read the ledger notebook from above */
    lgr.ReadLedger(notebook)
    lgr.CreateAccount()
    lgr.AddToEntry()
    lgr.PrintPretty()
    lgr.PrintTable()
}

func (lgr *Ledger) CreateAccount() {
    /* Create a new account */
    if createAccount {
        fullname := readBuf("Enter the full name of the new account", true)
        balance := StrToFloat(readBuf("Enter initial balance for " + fullname + " account", true))
        notebook = fullname + ".lgr"
        notebook = strings.ToLower(strings.Replace(notebook, " ", "-", -1))
        lgr.CreateAccountHash(fullname, notebook, balance)
    }
}

func (lgr *Ledger) AddToEntry() {
    /* Add to new items to the ledger, one expense and one income type */
    if addEntry {
        transaction = readBuf("Enter area of transaction", true)
        location = readBuf("Enter area of location of transaction", false)
        detail = readBuf("Enter details of transaction", true)
        cost = StrToFloat(readBuf("Enter cost of transaction", true))
        lgr.AddEntry(lgr.GetAcctNum(notebook), transaction, location, detail, cost)
    }
}

func (lgr Ledger) PrintPretty() {
    /* Print the ledger out pretty */
    if printPretty {
        lgr.PrintLedger()
    }
}

func (lgr Ledger) PrintTable() {
    /* Print the information into a markdown table */
    if printTable != "" {
        lgr.PrintToTable(lgr.GetAcctNum(notebook), printTable)
    }
}
