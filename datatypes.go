// datatypes
package rets

import (
	//	"log"
	"strconv"
	"time"
)

func ToCharacter(val string) string {
	return val
}

func ToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return intVal
}

func ToDateTime(val string) time.Time {
	dtVal, err := time.Parse(time.RFC3339, val)
	if err == nil {
		return dtVal
	}

	dtVal, err = time.Parse("2006-01-02T15:04:05", val)
	if err == nil {
		return dtVal
	}

	return time.Time{}
}

func ToDate(val string) time.Time {
	dateVal, err := time.Parse(time.RFC1123, val)
	if err == nil {
		return dateVal
	}

	dateVal, err = time.Parse("2006-01-02", val)
	if err == nil {
		return dateVal
	}

	return time.Time{}
}

func ToDecimal(val string) float64 {
	decVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0.0
	}
	return decVal
}

func ToBoolean(val string) bool {
	switch val {
	case "0":
		return false
	case "1":
		return true
	default:
		return false
	}
}
