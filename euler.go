package main

import "fmt"

func main() {

	var number int  // holds the count
	var pudding int // where the proof is

	for {

		for i := 20; i > 0; i-- {
			if number%i != 0 {
				break
			}

			if i == 1 {
				pudding = number
				break
			}

			continue
		}

		if pudding > 0 {
			fmt.Println("Proof: ", pudding)
			break
		} else {
			number++
		}
	}
}

// For Mr. Todd McLeod! - You Rock!
// https://projecteuler.net/problem=5

// 2520 is the smallest number that can be divided by each of the
// numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by
// all of the numbers from 1 to 20?
