package main

import (
	"log"
	"time"
)

var j = 0

func value() int {
	result := j
	j++
	return result
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")(j)
	j++
	time.Sleep(10 * time.Second)
}

func trace(msg string) func(k int) {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func(k int) {
		log.Printf("exit %s (%s) with j = %d", msg, time.Since(start), k)
	}
}

func main() {
	bigSlowOperation()
}
