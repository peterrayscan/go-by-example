package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	p := fmt.Println

	p(time.Now())
	p(time.Date(2023, 12, 18, 19, 15, 00, 000, time.UTC))
	p(time.Date(2023, 12, 18, 19, 15, 00, 000, time.Local))
	p()

	p(time.Now().Year())
	p(time.Now().Month())
	p(time.Now().Day())
	p(time.Now().Hour())
	p(time.Now().Minute())
	p(time.Now().Second())
	p(time.Now().Nanosecond())
	p(time.Now().Weekday())
	p()

	p(time.Now().Location())
	p()

	now := time.Now()
	later := now.Add(1 * time.Hour)
	p(now.Before(later))
	p(now.Equal(later))
	p(now.After(later))
	p()

	more := later.Sub(now)
	p(more)
	p(more.Hours())
	p(more.Minutes())
	p(more.Seconds())
	p(more.Nanoseconds())

	p(now.Add(-1 * time.Hour))
}
