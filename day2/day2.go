package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CubeBounds struct {
	colour string
	max int
}

type CubeMins struct {
	colour string
	min int
}

func main() {
	readFile, err := os.Open("day2.txt")
  
    if err != nil {
        fmt.Println(err)
    }

	cubeBounds := []CubeBounds{
		CubeBounds{"red", 12},
		CubeBounds{"green", 13},
		CubeBounds{"blue", 14},
	}

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

	sum1 := 0
	sum2 := 0
	game := 0

    for fileScanner.Scan() {
		game++;

		cubeMins := &[3]CubeMins{
			CubeMins{"red", 0},
			CubeMins{"green", 0},
			CubeMins{"blue", 0},
		}

        line := fileScanner.Text()

		parts := strings.Split(line, ":")
		pulls := strings.Split(parts[1], ";")

		valid := true
		for _, pull := range pulls {
			cubes := strings.Split(pull, ",")

			for _, cube := range cubes {
				vals := strings.Split(cube, " ")
				var number int
				fmt.Sscanf(vals[1], "%d", &number) 

				for _, cubeBound := range cubeBounds {
					if cubeBound.colour == vals[2] {
						if number > cubeBound.max {
							valid = false
						}
					}
				}

				for i, cubeMin := range cubeMins[:] {
					if cubeMin.colour == vals[2] {
						cubeMins[i].min = max(cubeMins[i].min, number)
					}
				}
			}
		}

		if valid {
			sum1 += game
		}

		power := 1
		for _, cubeMin := range cubeMins {
			power *= cubeMin.min
		}

		sum2 += power
   }

	fmt.Printf("Pt 1 sum is %d\n", sum1)
	fmt.Printf("Pt 2 sum is %d\n", sum2)
  
    readFile.Close()
}