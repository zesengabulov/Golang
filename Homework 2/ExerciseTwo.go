package main

import "fmt"

func perimeter(length,width float64) float64{
	const twoSidesIndex float64 = 2.0
	return (twoSidesIndex * length) + (twoSidesIndex * width)  
}

func main(){
	var lengthPlotA, widthPlotA float64 = 8.2,10
	var lengthPlotB, widthPlotB float64 = 5,5.2
	var lengthPlotC, widthPlotC float64 = 6.2,4.5
	var (
		perimeterPlotA float64 = perimeter(lengthPlotA,widthPlotA)
		perimeterPlotB float64 = perimeter(lengthPlotB,widthPlotB)
		perimeterPlotC float64 = perimeter(lengthPlotC,widthPlotC)
		totalPerimeter float64 = perimeterPlotA + perimeterPlotB + perimeterPlotC
	)

	fmt.Printf(
		"\nThe perimeter for the Plot A is %.1f meters." +
		"\nFor Plot B perimeter is %.1f meters." +  
		"\nFor Plot C perimeter is %.1f meters." +
		"\nTotal length is %.1f meters.\n",
		perimeterPlotA,
		perimeterPlotB,
		perimeterPlotC,
		totalPerimeter,
	)

}