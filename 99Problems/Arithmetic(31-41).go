package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println("#31 determine whether a given integer number is prime")
    fmt.Printf("%v -> %v\n", 1, isPrime(1))
    fmt.Printf("%v -> %v\n\n", 315, isPrime(315))

    fmt.Println("#32 determine the prime factors of a given positive integer")
    fmt.Printf("%v -> %v\n", 1, primeFactors(1))
    fmt.Printf("%v -> %v\n\n", 315, primeFactors(315))

    fmt.Println("#33 determine the prime factors of a given positive integer (2)")
    fmt.Printf("%v -> %v\n", 1, primeFactorsMult(1))
    fmt.Printf("%v -> %v\n\n", 315, primeFactorsMult(315))

    fmt.Println("#34 a list of prime numbers")
    fmt.Printf("%v, %v -> %v\n", 1, 2, primeList(1, 2))
    fmt.Printf("%v, %v -> %v\n", 3, 15, primeList(3, 15))
}

func isPrime(number int) (isPrime bool) {
    if number < 2 {
        return
    }

    for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
         if number % i == 0 {
             return
         }

         i++
     }

    return true
}

func primeFactors(number int) (primeFactors []int) {
    for number % 2 == 0 {
        primeFactors = append(primeFactors, 2)
        number = number / 2
    }

    for i := 3; i * i <= number; i = i + 2 {
		for number % i == 0 {
			primeFactors = append(primeFactors, i)
			number = number / i
		}
	}

    if number > 2 {
        primeFactors = append(primeFactors, number)
    }

    return
}

func primeFactorsMult(number int) (primeFactorsMult [][2]int) {
    primeFactors := primeFactors(number)
    current := -1
    replications := 1
    for _, head := range primeFactors {
        if current != head {
            if replications > 1 {
                primeFactorsMult = append(primeFactorsMult, [2]int{current, replications})
                replications = 1
            } else if current != -1 {
                primeFactorsMult = append(primeFactorsMult, [2]int{current, 1})
            }

            current = head
        } else {
            replications++
        }
    }

    if replications > 1 {
        primeFactorsMult = append(primeFactorsMult, [2]int{current, replications})
    } else if current != -1 {
        primeFactorsMult = append(primeFactorsMult, [2]int{current, 1})
    }

    return
}

func primeList(start, end int) (primes []int) {
    if start < 2 || end < 2 || start > end {
        return
    }

    for start <= end {
        if isPrime(start) {
            primes = append(primes, start)
        }

        start++
    }

    return
}
