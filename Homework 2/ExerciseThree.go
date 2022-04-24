package main

import (
	"fmt"
	"errors"
)

func perimeter(length,width float64) (float64,error){
	const twoSidesIndex float64 = 2.0

	if(length < 0 && width < 0){
		return 0,errors.New(
			fmt.Sprintf("A width of [%.1f] and" + 
				" length of [%.1f] are invalid",width,length))
	}else if length < 0{
		return 0,errors.New(
			fmt.Sprintf("A length of [%.1f] is invalid",length))
	}else if width < 0{
		return 0,errors.New(
			fmt.Sprintf("A width of [%.1f] is invalid",width))
	}

	return (twoSidesIndex * length) + (twoSidesIndex * width), nil  
}

func main(){
	var lengthPlotA, widthPlotA float64 = -8.2,10
	var lengthPlotB, widthPlotB float64 = 5,-5.2
	var lengthPlotC, widthPlotC float64 = 6.2,4.5

	perimeterPlotA, errorPlotA := perimeter(lengthPlotA,widthPlotA)
	perimeterPlotB, errorPlotB := perimeter(lengthPlotB,widthPlotB)
	perimeterPlotC, errorPlotC := perimeter(lengthPlotC,widthPlotC)
	totalPerimeter := perimeterPlotA + perimeterPlotB + perimeterPlotC
	

	if errorPlotA != nil {
		fmt.Printf("\nFor the plot A: %v", errorPlotA.Error())
	}else{
		fmt.Printf("\nThe perimeter of plot A is: %.1f meters\n",perimeterPlotA)
	}

	if errorPlotB != nil {
		fmt.Printf("\nFor the plot B: %v", errorPlotB.Error())
	}else{
		fmt.Printf("\nThe perimeter of plot B is: %.1f meters\n",perimeterPlotB)
	}

	if errorPlotC != nil {
		fmt.Printf("\nFor the plot C: %v", errorPlotC.Error())
	}else{
		fmt.Printf("\nThe perimeter of plot C is: %.1f meters",perimeterPlotC)
	}

	fmt.Printf("\nTotal length is %.1f meters\n",totalPerimeter)

}