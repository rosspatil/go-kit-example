package circuitbreaker

import (
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

var Cb *gobreaker.CircuitBreaker

func init() {
	st := gobreaker.Settings{}
	st.Interval = time.Second * 30
	st.Name = "go-kit-example-circuitbreaker"
	st.MaxRequests = 1
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		fmt.Println(counts)
		return counts.ConsecutiveFailures > 2
	}
	st.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		fmt.Println(name, from.String(), to.String())

	}
	Cb = gobreaker.NewCircuitBreaker(st)
}
