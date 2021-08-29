package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year   int `json:"y"`
	Month  int `json:"m"`
	Day    int `json:"d"`
	Hour   int `json:"h"`
	Min    int `json:"min"`
	Second int `json:"s"`
}

func (this *Date) ToString() string {
	return fmt.Sprintf("{0}-{1}-{2} {3}:{4}:{5}", this.Year, this.Month, this.Day, this.Hour, this.Min, this.Second)
}

/*

 */
func (this *Date) Parse(str string) error {
	str0 := strings.Split(str, " ")
	var str2, str1 []string
	if len(str0) > 0 {
		str1 = strings.Split(str0[0], "-")
	}
	if len(str0) > 1 {
		str2 = strings.Split(str0[1], ":")
	}
	if len(str1) > 0 {
		year, err := strconv.Atoi(str1[0])
		if err != nil {
			return err
		}
		this.Year = year
	}
	if len(str1) > 1 {
		month, err := strconv.Atoi(str1[1])
		if err != nil {
			return err
		}
		this.Month = month
	}
	if len(str1) > 2 {
		day, err := strconv.Atoi(str1[2])
		if err != nil {
			return err
		}
		this.Day = day
	}
	if len(str2) > 0 {
		hour, err := strconv.Atoi(str2[0])
		if err != nil {
			return err
		}
		this.Hour = hour
	}
	if len(str2) > 1 {
		min, err := strconv.Atoi(str2[1])
		if err != nil {
			return err
		}
		this.Min = min
	}
	if len(str2) > 2 {
		sec, err := strconv.Atoi(str2[2])
		if err != nil {
			return err
		}
		this.Second = sec
	}
	return nil

}
func Now() *Date {
	t := time.Now()
	return &Date{
		Year:   t.Year(),
		Month:  int(t.Month()),
		Day:    t.Day(),
		Hour:   t.Hour(),
		Min:    t.Minute(),
		Second: t.Second(),
	}
}
