package main

import (
	"go_with_tests/dependency_injection"
	"go_with_tests/mocking"
	"os"
	"time"
)

func main() {
	dependency_injection.Greet(os.Stdout, "Elodie")
	//log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
