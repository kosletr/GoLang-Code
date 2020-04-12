package main

import (
	"errors"
	"fmt"
	"math"
)

// Complex Structure

type complex struct {
	real float64
	imag float64
}

// Methods

func (z complex) conjComplex() complex {
	return complex{z.real, -z.imag}
}

func (z complex) magComplex() float64 {
	return math.Sqrt(math.Pow(z.real, 2) + math.Pow(z.imag, 2))
}

func (z complex) formulaComplex() string {

	if z.real == 0 && z.imag == 1 {
		return fmt.Sprint("i")
	} else if z.real == 0 && z.imag == -1 {
		return fmt.Sprint("-i")

	} else if z.imag == 1 {
		return fmt.Sprint(z.real, "+i")
	} else if z.imag == -1 {
		return fmt.Sprint(z.real, "-i")

	} else if z.imag == 0 {
		return fmt.Sprint(z.real, "")
	} else if z.real == 0 {
		return fmt.Sprint(z.imag, "i")

	} else if z.imag < 0 {
		return fmt.Sprint(z.real, "", z.imag, "i")
	} else {
		return fmt.Sprint(z.real, "+", z.imag, "i")
	}
}

// Main function

func main() {

	// Get a complex number as input from user
	z_1, e1 := readComplex()

	if e1 != nil {
		fmt.Println(e1)
		return
	}

	// Get another one
	z_2, e2 := readComplex()
	if e2 != nil {
		fmt.Println(e2)
		return
	}

	// Apply Operations
	z_3 := addComplex(z_1, z_2)
	z_4 := subComplex(z_1, z_2)
	z_5 := multComplex(z_1, z_2)
	z_6, e_dev := devComplex(z_1, z_2)
	z_7 := z_1.magComplex()
	z_8 := z_1.conjComplex()

	// Check for division by zero errors
	if e_dev != nil {
		fmt.Println(e_dev)
		return
	}

	// Print Results
	fmt.Println("z1 = " + z_1.formulaComplex())
	fmt.Println("z2 = " + z_2.formulaComplex())
	fmt.Println("z1 + z2 = " + z_3.formulaComplex())
	fmt.Println("z1 - z2 = " + z_4.formulaComplex())
	fmt.Println("z1 * z2 = " + z_5.formulaComplex())
	fmt.Println("z1 / z2 = " + z_6.formulaComplex())
	fmt.Println("|z1| = " + fmt.Sprint(z_7, ""))
	fmt.Println(" z1* = " + z_8.formulaComplex())

}

// Functions

func readComplex() (complex, error) {

	var re float64
	var im float64

	fmt.Print("Enter a complex number ( a+bi or a-bi ): ")
	_, e := fmt.Scanf("%f%fi\n", &re, &im)

	if e == nil {
		return complex{re, im}, nil
	} else {
		return complex{0, 0}, errors.New("\nInvalid Input\n")
	}

}

func addComplex(z1 complex, z2 complex) complex {

	return complex{z1.real + z2.real, z1.imag + z2.imag}
}
func subComplex(z1 complex, z2 complex) complex {

	return complex{z1.real - z2.real, z1.imag - z2.imag}
}
func multComplex(z1 complex, z2 complex) complex {

	return complex{z1.real*z2.real - z1.imag*z2.imag, z1.real*z2.imag + z1.imag*z2.real}
}
func devComplex(z1 complex, z2 complex) (complex, error) {

	sqr_mag_z2 := math.Pow(z2.magComplex(), 2)

	if sqr_mag_z2 == 0 {
		return complex{0, 0}, errors.New("\nCan't devide by zero\n")
	}

	z := multComplex(z1, z2.conjComplex())
	z.real /= sqr_mag_z2
	z.imag /= sqr_mag_z2
	return z, nil
}
