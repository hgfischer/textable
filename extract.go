package textable

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	fltRgxStr   = `([+-]?\d+(?:[\.eE]\d+)?)`
	fltRgx      = regexp.MustCompile(fltRgxStr)
	intRgxStr   = `([+-]?(?:(?:0[xX][0-9A-Fa-f]+)|(?:0[0-7]+)|(?:\d+)))`
	intRgx      = regexp.MustCompile(intRgxStr)
	uintRgxStr  = `((?:0[xX][0-9A-Fa-f]+)|(?:0[0-7]+)|(?:\d+))`
	uintRgx     = regexp.MustCompile(uintRgxStr)
	isNumRgxStr = `^(` + fltRgxStr + `|` + intRgxStr + `|` + uintRgxStr + `)$`
	isNumRgx    = regexp.MustCompile(isNumRgxStr)
)

// Return true if the string is only of numeric nature
func IsNumeric(str string) bool {
	str = strings.TrimSpace(str)
	if str == "" {
		return false
	}
	return isNumRgx.MatchString(str)
}

func ExtractInt(str string) (intg int64, err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, errors.New("Empty string")
	}
	for _, match := range intRgx.FindAllString(str, -1) {
		// First try conversion with base guessing
		if intg, err = strconv.ParseInt(match, 0, 64); err == nil {
			return
		}
		// If I got here, some base guessing have failed. Force base 10.
		if intg, err = strconv.ParseInt(match, 10, 64); err == nil {
			return
		}
	}
	if err == nil {
		err = errors.New("Could not convert string: " + strconv.Quote(str))
	}
	return 0, err
}

func ExtractUint(str string) (uintg uint64, err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, errors.New("Empty string")
	}
	for _, match := range uintRgx.FindAllString(str, -1) {
		// First try conversion with base guessing
		if uintg, err = strconv.ParseUint(match, 0, 64); err == nil {
			return
		}
		// If I got here, some base guessing have failed. Force base 10.
		if uintg, err = strconv.ParseUint(match, 10, 64); err == nil {
			return
		}
	}
	if err == nil {
		err = errors.New("Could not convert string: " + strconv.Quote(str))
	}
	return 0, err
}

func ExtractFloat(str string) (flt float64, err error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, errors.New("Empty string")
	}
	if fltRgx.MatchString(str) {
		if flt, err = strconv.ParseFloat(fltRgx.FindString(str), 64); err == nil {
			return flt, nil
		}
	}
	if err == nil {
		err = errors.New("Could not extract float from string: " + strconv.Quote(str))
	}
	return
}
