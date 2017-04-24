package cis193stats

import (
	"errors"
	"math"
)

// Median gets the median number in a slice of numbers
func Median(inp data) (median float64, err error) {

	c := sortedCopy(inp)

	l := len(c)
	if l == 0 {
		return math.NaN(), errors.New("Empty Set")
	} else if l%2 == 0 {
		median, _ = Mean(c[l/2-1 : l/2+1])
	} else {
		median = float64(c[l/2])
	}

	return median, nil
}

func Mode(inp data) (mode []float64, err error) {
	l := inp.Len()
	if l == 1 {
		return inp, nil
	} else if l == 0 {
		return nil, errors.New("Empty Set")
	}

	c := sortedCopyDif(inp)

	mode = make([]float64, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		if c[i] == c[i-1] {
			cnt++
		} else if cnt == maxCnt && maxCnt != 1 {
			mode = append(mode, c[i-1])
			cnt = 1
		} else if cnt > maxCnt {
			mode = append(mode[:0], c[i-1])
			maxCnt, cnt = cnt, 1
		} else {
			cnt = 1
		}
	}

	if cnt == maxCnt {
		mode = append(mode, c[l-1])
	} else if cnt > maxCnt {
		mode = append(mode[:0], c[l-1])
		maxCnt = cnt
	}

	if maxCnt == 1 {
		return data{}, nil
	}

	return mode, nil
}
