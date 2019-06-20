package math

import "errors"

type Calculator interface{

	Add(int, int) int
	Mult(int, int) int
	Sub(int, int) int
	Div(int, int) (float64, error)

}

func NewCalculator() Calculator {
	return &calc{}
}

type calc struct {

}

func (c calc) Add(a int, b int)int{
	return a+b
}

func (c calc) Mult(a int, b int)int{
	return a*b
}

func (c calc) Sub(a int, b int)int{
	return a-b
}

func (c calc) Div(a int, b int)(float64, error){
	f := float64(a)
	d := float64(b)

	if d == 0{
		return 0, errors.New("Undefined")
	}

	return f/d, nil
}
