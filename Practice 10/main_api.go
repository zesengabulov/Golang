package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	email     string  `json:"email"`
	GPA       float32 `json:"gpa"`
}

var students []Student = []Student{
	{1, "Temirbolat", "Maratuly", "t_maratuly@kbtu.kz", 3.9},
	{2, "Anuar", "Sarsengaliev", "a_sarsengaliev@kbtu.kz", 3.0},
	{3, "Tamerlan", "Kuankash", "t_kuankash@kbtu.kz", 3.2},
	{4, "Asanali", "Moldash", "a_moldash@kbtu.kz", 3.3},
	{5, "Andew", "Lincoln", "a_lincoln@kbtu.kz", 2.9},
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentByID(c *gin.Context) {
	var foundId string = c.Param("id")
	idInt, err := strconv.ParseInt(foundId, 10, 32)
	if err != nil {
		return
	}

	for _, stud := range students {
		if stud.ID == int(idInt) {
			c.IndentedJSON(http.StatusOK, stud)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func deleteStudentByID(c *gin.Context) {
	var foundId string = c.Param("id")
	idInt, err := strconv.ParseInt(foundId, 10, 32)
	if err != nil {
		return
	}
	for i, stud := range students {
		if stud.ID == int(idInt) {
			students = append(students[:i], students[i+1:]...)
			c.IndentedJSON(http.StatusOK, students)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/postStudent", postStudents)
	router.GET("/student/:id", getStudentByID)
	router.DELETE("/deleteStudent/:id", deleteStudentByID)

	router.Run("localhost:8081")
}
