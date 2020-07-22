# ledger
A quick dirty and easy ledger. The `ledger.go` package is used to read and write into a ledger notebook. An example can be seen in the [example directory](https://github.com/loerac/ledger/tree/trunk/example)

## Ledger
### Account Meta Data
At the start of the ledger, there should be the account name and number. Both are required. See example below.

### Entry Meta Data
The ledger has a specific format, as seen below. This addresses the needed information for each item in the ledger.
The `<ADDRESS>` data doesn't have to be specified, this can be left blank with empty quotes

##### Complete meta data
```
<ACCT-NAME>:<ACCT-NUMBER>
<YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
```

##### No address
```
<ACCT-NAME>:<ACCT-NUMBER>
<YYYYMMDDTHHMMSS>:<STORE>:<DETAILS>:<EXCHANGE-TYPE>:<COST>:<BALANCE>
```

##### Definitions:
* `<ACCT-NAME>`
  * Name of the ledger
* `<ACCT-NUMBER>`
  * Account number for the ledger
* `<YYYYMMDDTHHMMSS>`
  * Year, month, day, (24) hour, minute, seconds
  * The 'T' splits the date from the time
* `<STORE>`
  * Name of the store
* `<ADDRESS>`
  * Location of the store
  * Can be left blank, ""
* `<DETAILS>`
  * Reasoning of exchange
* `<EXCHANGE-TYPE>`
  * Either expense or income
* `<COST>`
  * Gain or loss of exchange
* `<BALANCE>`
  * Updated balance from exchange

##### Example:
```
Rick Kicks:936e1204e7b8c686
20200706T121409:Apple Store:Surface Pro 7:Expense:15.99:+6969.69
```
