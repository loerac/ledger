package ledger

import (
    "fmt"
    "strconv"
    "time"
)

/**
 * @brief:  Check if any errors occured, panic if so
 *
 * @arg:    e - Error
 **/
func CheckErr(e error) {
    if nil != e {
        panic(e)
    }
}

/**
 * @brief:  Converts a string to an uint64.
 *          Panic if errors
 *
 * @arg:    str - String that is to be converted to uint64
 *
 * @return: uint64 val of string
 **/
func StrToUint(str string) uint64 {
    u, err := strconv.ParseUint(str, 10, 64)
    CheckErr(err)

    return u
}

/**
 * @brief:  Converts a string to an float64.
 *          Panic if errors
 *
 * @arg:    str - String that is to be converted to float64
 *
 * @return: float64 val of string
 **/
func StrToFloat(str string) float64 {
    f, err := strconv.ParseFloat(str, 64)
    CheckErr(err)

    return f
}

/**
 * @brief:  A ternary functions that checks for the condition and
 *          return a value. Type cast the function to get the
 *          data type that is expecting
 *
 * @arg:    condition - Condition to check for validity
 * @arg:    valid - Val being returned if the condition is true
 * @arg:    invalid - Val being returned if the condition is false
 *
 * @return: Depending on the condition, an interface value is returned.
 **/
func ternary(condition bool, valid, invalid interface{}) interface{} {
    if condition {
        return valid
    }

    return invalid
}

/**
 * @brief:  Get the current date and time
 *
 * @return: The date and time in the format of <YYYYMMDD>T<HHMMSS>
 **/
func GetDate() string {
    currTime := time.Now()
    return fmt.Sprintf("%d%02d%02dT%02d%d%d",
            currTime.Year(), currTime.Month(), currTime.Day(),
            currTime.Hour(), currTime.Minute(), currTime.Second())
}
