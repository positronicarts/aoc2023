package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("day1.txt")
  
    if err != nil {
        fmt.Println(err)
    }

	digits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	sum1 := 0
	sum2 := 0

    for fileScanner.Scan() {
        line := fileScanner.Text()

		found := false
		start := 0
		end := 0
		startIndex := len(line)
		endIndex := 0
		for index, character := range line {
			if character > '0' && character <= '9' {
				if !found {
					found = true
					start = (int(character) - int('0'))
					startIndex = index
				}

				end = (int(character) - int('0'))
				endIndex = index
			}
		}

		sum1 += start * 10 + end

		for digitIndex, digit := range digits {
			for scanIndex := 0; scanIndex < min(startIndex, len(line) - len(digit)); scanIndex++ {
				if digit == line[scanIndex:scanIndex + len(digit)] {
					start = digitIndex + 1
					startIndex = scanIndex
				}
			}

			for scanIndex := endIndex; scanIndex < len(line) - len(digit) + 1; scanIndex++ {
				if digit == line[scanIndex:scanIndex + len(digit)] {
					end = digitIndex + 1
					endIndex = scanIndex
				}
			}
		}

		sum2 += start * 10 + end
    }

	fmt.Printf("Pt 1 sum is %d\n", sum1)
	fmt.Printf("Pt 2 sum is %d\n", sum2)
  
    readFile.Close()
}