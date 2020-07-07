# ledger
A quick dirty and easy ledger. The `ledger.go` package is used to read (not done yet) and write into a ledger notebook. An example can be seen in the <a href="https://github.com/loerac/ledger/tree/trunk/example" target="_blank">example directory</a>.

### Ledger Meta Data
The ledger has a specific format, as seen below. This addresses the needed information for each item in the ledger.
The `<ADDRESS>` data doesn't have to be specified, this can be left blank with empty quotes

##### Complete meta data
*  `<YYYYMMDDTHHMMSS>:<STORE>@<ADDRESS>:<DETAILS>:<COST-ENTITY>:<COST>:<BALANCE>`

##### No address
* `<YYYYMMDDTHHMMSS>:<STORE>:<DETAILS>:<COST-ENTITY>:<COST>:<BALANCE>`

##### Definitions:
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
`20200706T121409:Apple Store:Surface Pro 7:Expense:15.99:+6969.69`
