package main

import (
	"fmt"
	"time"

	"github.com/rnowt/period"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const layout = "02-01-2006 15:04"

func main() {
	p, err := period.NewPeriod(
		time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local),
		time.Date(2021, time.July, 1, 0, 0, 0, 0, time.Local),
	)
	check(err)

	periods := p.Split(6)
	for _, subP := range periods {
		fmt.Printf("%s - %s\n", subP.Start.Format(layout), subP.End.Format(layout))
	}
}
