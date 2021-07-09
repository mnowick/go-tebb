package examples

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func Error(sq float64) {
	f, err := sqrt(sq)
	if err != nil {
		fmt.Println(err)
		return
	}

	var sb strings.Builder
	sqs := fmt.Sprintf("%f", sq)
	fs := fmt.Sprintf("%f", f)
	sb.WriteString("The square root of ")
	sb.WriteString(sqs)
	sb.WriteString(" is ")
	sb.WriteString(fs)
	fmt.Println(sb.String())
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("input must be greater than 0, imaginary numbers are silly")
	}

	return math.Sqrt(f), nil
}
