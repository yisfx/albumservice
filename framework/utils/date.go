package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	timeformate = "2006-01-02 15:04:05"
)

var DateTime dateTime

type dateTime struct {
}

func init() {
	DateTime = dateTime{}
}

type Date struct {
	Year   int `json:"y"`
	Month  int `json:"m"`
	Day    int `json:"d"`
	Hour   int `json:"h"`
	Min    int `json:"min"`
	Second int `json:"s"`
}

func formateInt(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%v", num)
	}
	return fmt.Sprintf("%v", num)
}

func (d *Date) ToString() string {
	return fmt.Sprintf("%v-%v-%v %v:%v:%v", formateInt(d.Year), formateInt(d.Month), formateInt(d.Day), formateInt(d.Hour), formateInt(d.Min), formateInt(d.Second))
}
func (d *Date) IsValid() bool {
	_, err := time.Parse(timeformate, d.ToString())
	return err == nil
}

/*

 */
func (dt *dateTime) Parse(str string) *Date {
	d := &Date{}
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
			return nil
		}
		d.Year = year
	}
	if len(str1) > 1 {
		month, err := strconv.Atoi(str1[1])
		if err != nil {
			return nil
		}
		d.Month = month
	}
	if len(str1) > 2 {
		day, err := strconv.Atoi(str1[2])
		if err != nil {
			return nil
		}
		d.Day = day
	}
	if len(str2) > 0 {
		hour, err := strconv.Atoi(str2[0])
		if err != nil {
			return nil
		}
		d.Hour = hour
	}
	if len(str2) > 1 {
		min, err := strconv.Atoi(str2[1])
		if err != nil {
			return nil
		}
		d.Min = min
	}
	if len(str2) > 2 {
		sec, err := strconv.Atoi(str2[2])
		if err != nil {
			return nil
		}
		d.Second = sec
	}
	return d

}

func (d *dateTime) Now() *Date {
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
