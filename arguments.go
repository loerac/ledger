package ledger

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

const (
    LEDGER_DESC string      = "The ledger notebook"
    PRINTPRETTY_DESC string = "Print the ledger notebook pretty"
    PRINTTABLE_DESC string  = "Print the ledger notebook to a markdown table\nNote: Output file is required"
    ADDENTRY_DESC string    = "Add a new entry to the ledger notebook\nNote: Ledger notebook is required"
    NEWACCOUNT_DESC string  = "Create a new account"
)

/***
 * @brief:  Reads input from the user
 *
 * @arg:    prompt - String to ask the user
 * @arg:    required - Input is required if true
 *
 * @return: The user response from the prompt
 ***/
func readBuf(prompt string, required bool) string {
        response := ""
        for {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print(prompt + " > ")
            response , _ = reader.ReadString('\n')
            response = strings.Replace(response, "\n", "", -1)

            if required && response == "" {
                fmt.Println("Error: input required, none given")
            } else {
                break
            }
        }

        return response
}

/***
 * @brief:  Checks to see if user has passed any flags.
 *          -l: The ledger notebook
 *          -pp: Print the ledger notebook out pretty
 *          -pt: Save the ledger notebook as a markdown table
 *          -add-entry: Adds an entry to the ledger notebook
 *          -new-acct: Creates a new account and ledger notebook
 ***/
func (lgr *Ledger) ArgumentParser() {
    var (
        notebook    string
        printPretty bool
        printTable  string
        addEntry    bool
        createAccount  bool
    )

    flag.StringVar(&notebook, "l", "", LEDGER_DESC)
    flag.BoolVar(&printPretty, "pp", false, PRINTPRETTY_DESC)
    flag.StringVar(&printTable, "pt", "", PRINTTABLE_DESC)
    flag.BoolVar(&addEntry, "a", false, ADDENTRY_DESC)
    flag.BoolVar(&createAccount, "n", false, NEWACCOUNT_DESC)
    flag.Parse()

    /* Either the ledger or create new account flag needs to be passed, not both */
    if notebook != "" && createAccount {
        fmt.Println("Error: Cannot read ledger notebook and create account.")
        fmt.Println("       Only one can be done at a time")
        os.Exit(1)
    } else if notebook == "" && !createAccount {
        fmt.Println("Error: Missing ledger notebook")
        flag.PrintDefaults()
        os.Exit(1)
    }

    /* Create a new account */
    if createAccount {
        fullname := readBuf("Enter the full name of the new account", true)
        balance := StrToFloat(readBuf("Enter initial balance for " + fullname + " account", true))
        notebook = fullname + ".lgr"
        notebook = strings.ToLower(strings.Replace(notebook, " ", "-", -1))
        lgr.CreateAccountHash(fullname, notebook, balance)
    }

    /* Read the ledger notebook */
    lgr.ReadLedger(notebook)
    acctNum := lgr.GetAcctNum(notebook)

    /* Add new entry to the ledger */
    if addEntry {
        transaction := readBuf("Enter where the transaction took place", true)
        location := readBuf("Enter address of the transaction (optional)", false)
        detail := readBuf("Enter details of transaction", true)
        cost := StrToFloat(readBuf("Enter cost of transaction", true))
        lgr.AddEntry(acctNum, transaction, location, detail, cost)
    }

    /* Print the ledger out pretty */
    if printPretty {
        lgr.PrintLedger()
    }

    /* Print the information into a markdown table */
    if printTable != "" {
        lgr.PrintToTable(acctNum, printTable)
    }
}
