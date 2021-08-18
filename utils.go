package ledger

import (
    "fmt"
    "log"
    "strconv"
    "time"
)

const (
    DATE_TIME   string = "20060102T150405"
    DATE_ONLY   string = "2006-01-02"
)

/**
 * @brief:  Check if any errors occured, log the error if there are any
 *
 * @arg:    e - Error
 **/
func CheckErr(e error) {
    if nil != e {
        log.Fatal(e)
    }
}

/**
 * @brief:  Converts a string to an uint64.
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

/***
 * @brief:  Convert a number (int, float, etc) to a string
 *
 * @arg:    number - Number that is to be converted
 *
 * @return: string value of the number
 ***/
func NumToStr(number interface{}) string {
    return fmt.Sprintf("%v", number)
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
 * @agr:    The format of which to put the date
 *
 * @return: The date and/or time in the format of which the user has entered
 **/
func GetDate(format string) string {
    currTime := time.Now()
    return currTime.Format(format)
}

func FormatDate(date string) string{
    /* Date */
    fmtDate := date[:4] + "/" + date[4:6] + "/" + date[6:8] + " "

    /* Time */
    if len(date) > 19 {
        fmtDate += date[12:13] + ":" + date[15:16] + ":" + date[18:19]
    }

    return fmtDate
}
