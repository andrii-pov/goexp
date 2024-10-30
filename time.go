package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // difference between now and then
	p(diff)

	p(diff.Hours())
	p(diff.Seconds())

	p(then.Add(diff))  // here we add the difference between now and then to then - getting current time
	p(then.Add(-diff)) // here we subtract the difference between now and then from then - getting then

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixMicro())
	p(now.UnixNano())

	p(now.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
