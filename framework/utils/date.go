package utils

import "fmt"

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Min    int
	Second int
}

func (this *Date) ToString() string {
	return fmt.Sprintf("{0}-{1}-{2} {3}:{4}:{5}", this.Year, this.Month, this.Day, this.Hour, this.Min, this.Second)
}
