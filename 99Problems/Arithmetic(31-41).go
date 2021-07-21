package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
    fmt.Println("#31 determine whether a given integer number is prime")
    fmt.Printf("%v -> %v\n", 1, isPrime(1))
    fmt.Printf("%v -> %v\n\n", 31, isPrime(31))

    fmt.Println("#32 determine the prime factors of a given positive integer")
    fmt.Printf("%v -> %v\n", 1, primeFactors(1))
    fmt.Printf("%v -> %v\n\n", 315, primeFactors(315))

    fmt.Println("#33 determine the prime factors of a given positive integer (2)")
    fmt.Printf("%v -> %v\n", 1, primeFactorsMult(1))
    fmt.Printf("%v -> %v\n\n", 315, primeFactorsMult(315))

    fmt.Println("#34 a list of prime numbers")
    fmt.Printf("%v, %v -> %v\n", 1, 2, primeList(1, 2))
    fmt.Printf("%v, %v -> %v\n\n", 3, 15, primeList(3, 15))

    fmt.Println("#35 goldbach's conjecture")
    a, b := goldbach(28)
    fmt.Printf("%v -> (%v, %v)\n", 28, a, b)
    a, b = goldbach(256)
    fmt.Printf("%v -> (%v, %v)\n\n", 256, a, b)

    fmt.Println("#36 a list of Goldbach compositions")
    fmt.Printf("%v, %v -> %v\n", 1, 3, goldbachList(1, 3))
    fmt.Printf("%v, %v -> %v\n\n", 4, 16, goldbachList(4, 16))

    fmt.Println("#37 determine the greatest common divisor of two positive integer numbers")
    fmt.Printf("%v, %v -> %v\n", 7, 23, greatestCommonDivisor(7, 23))
    fmt.Printf("%v, %v -> %v\n\n", 36, 63, greatestCommonDivisor(36, 63))

    fmt.Println("#38 determine whether two positive integer numbers are coprime")
    fmt.Printf("%v, %v -> %v\n", 36, 63, coprime(36, 63))
    fmt.Printf("%v, %v -> %v\n\n", 7, 23, coprime(7, 23))

    fmt.Println("#39 calculate Euler's totient function phi(m)")
    fmt.Printf("%v -> %v\n", 1, totient(1))
    fmt.Printf("%v -> %v\n\n", 10, totient(10))

    fmt.Println("#40 calculate Euler's totient function phi(m) (2)")
    fmt.Printf("%v -> %v\n", 1, totientEfficient(1))
    fmt.Printf("%v -> %v\n\n", 10, totientEfficient(10))

    fmt.Println("#41 compare the two methods of calculating Euler's totient function")
    start := time.Now()
    initialImplementation := totient(10090)
    elapsed := time.Since(start)
    fmt.Printf("initial implementation with %v -> %v took %v\n", 10090, initialImplementation, elapsed)

    start = time.Now()
    imporvedImplementation := totientEfficient(10090)
    elapsed = time.Since(start)
    fmt.Printf("improved implementation with %v -> %v took %v\n", 10090, imporvedImplementation, elapsed)
}

func isPrime(number int) (isPrime bool) {
    if number < 2 {
        return
    }

    for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
         if number % i == 0 {
             return
         }
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
    multiplicity := 1
    for _, head := range primeFactors {
        if current != head {
            if multiplicity > 1 {
                primeFactorsMult = append(primeFactorsMult, [2]int{current, multiplicity})
                multiplicity = 1
            } else if current != -1 {
                primeFactorsMult = append(primeFactorsMult, [2]int{current, 1})
            }

            current = head
        } else {
            multiplicity++
        }
    }

    if multiplicity > 1 {
        primeFactorsMult = append(primeFactorsMult, [2]int{current, multiplicity})
    } else if current != -1 {
        primeFactorsMult = append(primeFactorsMult, [2]int{current, 1})
    }

    return
}

func primeList(start, end int) (primesList []int) {
    if start < 2 || end < 2 || start > end {
        return
    }

    for start <= end {
        if isPrime(start) {
            primesList = append(primesList, start)
        }

        start++
    }

    return
}

func goldbach(number int) (a, b int) {
    if number < 3 || number % 2 == 1 {
        return
    }

    primes := primeList(2, number)
    differenceMap := map[int]int{}
    for _, prime := range primes {
        if difference, found := differenceMap[number - prime]; found {
            return number - prime, difference
        } else if prime + prime == number {
            return prime, prime
        }

        differenceMap[prime] = number - prime
    }

    return
}

func goldbachList(start, end int) (goldbachList [][3]int) {
    if start < 3 || end < 3 || start > end {
        return
    }

    for start <= end {
        if start % 2 == 0 {
            a, b := goldbach(start)
            goldbachList = append(goldbachList, [3]int{start, a, b})
        }

        start++
    }

    return
}

func greatestCommonDivisor(a, b int) int {
      for b != 0 {
        a, b = b, a % b
      }

      return a
}

func coprime(a, b int) bool {
    return greatestCommonDivisor(a, b) == 1
}

func totient(number int) (totient int) {
    if number == 1 {
        return 1
    }

    for i := 1; i < number; i++ {
        if coprime(number, i) {
            totient++
        }
    }

    return
}

func totientEfficient(number int) (totient int) {
    totient = 1
    primeFactorsMult := primeFactorsMult(number)
    for _, primeFactorMult := range primeFactorsMult {
        primeFactor := primeFactorMult[0]
        multiplicity := primeFactorMult[1]
        totient *= (primeFactor - 1) * int(math.Pow(float64(primeFactor), float64(multiplicity - 1)))
    }

    return
}
