package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Calculator struct {
	firstNumber  float64
	secondNumber float64
	operator     rune
}

func (calc *Calculator) Operate() (float64, error) {
	switch calc.operator {
	case '+':
		return calc.add(), nil
	case '-':
		return calc.substract(), nil
	case '*':
		return calc.multiply(), nil
	case '/':
		return calc.divide()
	case '^':
		return calc.power(), nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func (calc *Calculator) GetArguments(argv []string) error {
	// First case
	if len(argv) != 3 {
		return errors.New("usage: calculator <firstNumber> <operator> <secondNumber>")
	}
	var err error
	// Second case
	calc.firstNumber, err = strconv.ParseFloat(argv[0], 64)
	if err != nil {
		return fmt.Errorf("invalid first number: %s", err.Error())
	}
	// Third case
	if len(argv[1]) != 1 {
		return errors.New("valid operators:\n1. +\n2. -\n3. *\n4. /\n5. ^\n")
	}
	calc.operator = rune(argv[1][0])
	// Fourth case
	calc.secondNumber, err = strconv.ParseFloat(argv[2], 64)
	if err != nil {
		return fmt.Errorf("invalid second number: %s", err.Error())
	}
	// End
	return nil
}

func (calc *Calculator) add() float64 {
	return calc.firstNumber + calc.secondNumber
}

func (calc *Calculator) substract() float64 {
	return calc.firstNumber - calc.secondNumber
}

func (calc *Calculator) multiply() float64 {
	return calc.firstNumber * calc.secondNumber
}

func (calc *Calculator) divide() (float64, error) {
	if calc.secondNumber == 0 {
		return math.NaN(), errors.New("division by zero")
	}
	return calc.firstNumber / calc.secondNumber, nil
}

func (calc *Calculator) power() float64 {
	return math.Pow(calc.firstNumber, calc.secondNumber)
}
