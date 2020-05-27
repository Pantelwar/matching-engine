package engine

import (
	"math/big"
)

func Add(x, y string) string {
	xBig, _ := new(big.Float).SetPrec(200).SetString(x)
	yBig, _ := new(big.Float).SetPrec(200).SetString(y)

	// valX, _ := xBig.Float64()
	// valY, _ := yBig.Float64()
	sum := new(big.Float).Add(xBig, yBig)
	// fmt.Println("x, y", valX, valY, sum.String())
	return sum.String()
}

func Sub(x, y string) string {
	xBig, _ := new(big.Float).SetPrec(200).SetString(x)
	yBig, _ := new(big.Float).SetPrec(200).SetString(y)

	// valX, _ := xBig.Float64()
	// valY, _ := yBig.Float64()
	sum := new(big.Float).Sub(xBig, yBig)
	// fmt.Println("x, y", valX, valY, sum.String())
	return sum.String()
}

func GreaterThan(x, y string) bool {
	xBig, _ := new(big.Float).SetPrec(200).SetString(x)
	yBig, _ := new(big.Float).SetPrec(200).SetString(y)

	return xBig.Cmp(yBig) == 1
}

func LessThan(x, y string) bool {
	xBig, _ := new(big.Float).SetPrec(200).SetString(x)
	yBig, _ := new(big.Float).SetPrec(200).SetString(y)

	return xBig.Cmp(yBig) == -1
}

func Equal(x, y string) bool {
	xBig, _ := new(big.Float).SetPrec(200).SetString(x)
	yBig, _ := new(big.Float).SetPrec(200).SetString(y)

	return xBig.Cmp(yBig) == 0
}
