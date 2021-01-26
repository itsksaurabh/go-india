// Package provides functions to generate the map of India and print it in Tricolor.
// Inspired from https://www.geeksforgeeks.org/code-to-generate-the-map-of-india-with-explanation/
// Author: LinkedIn: http://linkedin.com/in/itsksaurabh , Github: https://github.com/itsksaurabh/

package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	// ANSI Color Codes for
	saffronColor = "\033[38;5;214m%s\033[0m\n"
	greenColor   = "\033[1;32m%s\033[0m\n"
)

func main() {
	printMapInTricolor(getIndianMap())
}

// getIndianMap returns map of India as a string
func getIndianMap() string {
	const (
		// String with which the function will determine how many spaces or exclamation marks to draw the flag
		s = "TFy!QJu ROo TNn(ROo)SLq SLq ULo+UHs UJq TNn*RPn/QP,\n" +
			"bEWS_JSWQAIJO^NBELPeHBFHT}TnALVlBLOFAkHFOuFETpHCStHAUFAgcEAelc,\n" +
			"lcn^r^r\\tZvYxXyT|S~Pn SPm SOn TNn ULo0ULo#ULo-WHq!WFs XDt!"
	)

	var (
		a    = 10
		b    = 0
		c    = 10
		flag string
	)

	for a != 0 {
		if b < 170 {
			a = int(s[b])
			b++
			for a > 64 {
				a--
				c++

				if c == 90 {
					c = c / 9
					flag += string(c)
				} else {
					flag += string(33 ^ (b & 0x01))
				}

			}
		} else {
			break
		}
	}
	return flag
}

// printMapInTricolor prints the map string in tricolors.
// It breaks the map string 'm' into 3 parts and prints
// each part in different colors as per the Indian Flag.
func printMapInTricolor(m string) {
	i := 0
	lines := strings.Count(m, "\n")

	scanner := bufio.NewScanner(strings.NewReader(m))
	for scanner.Scan() {
		switch {
		case i <= lines/3:
			fmt.Printf(saffronColor, scanner.Text())
		case i > lines-lines/3:
			fmt.Printf(greenColor, scanner.Text())
		default:
			fmt.Println(scanner.Text())
		}

		i++
		time.Sleep(50 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
