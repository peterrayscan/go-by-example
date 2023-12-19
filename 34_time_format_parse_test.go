package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_time_format_parse(t *testing.T) {
	p := fmt.Println

	now := time.Now()

	p(now.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	p(now.Format("2006年01月02日 15时04分05.999999999秒"))

	t2, _ := time.Parse("3 04 PM", "8 41 PM")
	p(t2)

	fmt.Printf(
		"%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	)

	_, e := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
	p(e)
}
