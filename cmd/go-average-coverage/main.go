package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var covRex *regexp.Regexp = regexp.MustCompile(`coverage: (\d+\.?\d*)% of statements`)

func main() {
	coverages := make([]float64, 0, 16)
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		fmt.Println(in.Text())
		if matches := covRex.FindStringSubmatch(in.Text()); matches != nil {
			coverage, err := strconv.ParseFloat(matches[1], 64)
			if err != nil {
				fmt.Printf("go-average-coverage: %v\n", err)
				return
			}
			coverages = append(coverages, coverage)
		}
	}

	var average float64
	for _, coverage := range coverages {
		average += coverage / float64(len(coverages))
	}

	fmt.Printf("average coverage: %.1f%% of statements\n", average)
}
