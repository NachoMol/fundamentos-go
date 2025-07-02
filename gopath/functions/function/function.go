package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Display(myValue int) {
	display(myValue)
}

func display(myValue int) {
	fmt.Println()
	fmt.Printf("Value: %d\n", myValue)
	fmt.Println()
}

func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, fmt.Errorf("unknown operation: %d", op)
	}
}

func Split(v int) (x, y int) {

	x = v * 4 / 9
	y = v - x
	return

}

func MSum(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("no values provided")
	}

	result := values[0]
	for _, value := range values[1:] {
		var err error
		result, err = Calc(op, result, value)
		if err != nil {
			return 0, err
		}
	}
	return result, nil
}
