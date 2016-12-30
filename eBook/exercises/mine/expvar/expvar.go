package main

import "fmt"
import "time"
import "net/http"

//import _ "expvar"
import "expvar"
import "runtime"

/*
Test golang performance monitoring.
See http://blog.ralch.com/tutorial/golang-metrics-with-expvar/
*/
func main() {

	go func() {
		fmt.Println("starting")
		http.ListenAndServe(":8080", http.DefaultServeMux)
	}()

	for i := 0; i < 999; i++ {
		fmt.Printf("The counter is at %d\n", i)
		time.Sleep(2 * time.Second)
		//		time.Sleep(1e9)

		//fmt.Printf("%#v", expvar)

		memstatsFunc := expvar.Get("memstats").(expvar.Func)
		memstats := memstatsFunc().(runtime.MemStats)
		fmt.Println("Alloc", memstats.Alloc)

		//		expvar.Do(func(variable expvar.KeyValue) {
		//			fmt.Printf("expvar.Key: %s expvar.Value: %s", variable.Key, variable.Value)
		//		})

		timeMetrics.Add(1 * time.Minute)

	}
}

type TimeVar struct {
	value time.Time
}

// Sets a time.Time as time metrics value
func (v *TimeVar) Set(date time.Time) {
	v.value = date
}

// Adds a time.Duration to current time metrics value
func (v *TimeVar) Add(duration time.Duration) {
	v.value = v.value.Add(duration)
}

// Converts the TimeVar metrics to string
func (v *TimeVar) String() string {
	return v.value.Format(time.UnixDate)
}

var (
	timeMetrics *TimeVar
)

func init() {
	timeMetrics = &TimeVar{value: time.Now()}
	expvar.Publish("Time", timeMetrics)
}
