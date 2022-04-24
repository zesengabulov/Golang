package main

import "fmt"

func scoreSummary(studentName string,testScoreA,testScoreB,testScoreC float64){
	const scoreNumber float64 = 3.0
	var totalScores float64 = testScoreA + testScoreB + testScoreC 
	var averageScore float64 = totalScores/scoreNumber
	fmt.Printf("\n%10s |%8.2f |%8.2f |%8.2f |%8.2f",studentName,testScoreA,testScoreB,testScoreC,averageScore)

}

func main(){
	fmt.Printf("%10s |%8s |%8s |%8s |%8s\n","Name","Score 1","Score 2","Score 3","Average")
	for curIndex :=0 ;curIndex < 50;curIndex++{
		fmt.Print("-")
	}
	scoreSummary("Temirbolat",95.40,82.30,74.60)
	scoreSummary("Jessie",79.30,99.10,82.50)
	scoreSummary("Lamar",82.20,95.40,77.60)
	fmt.Println("")
}