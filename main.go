package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
  
	 var studentDetails []student
    
	studentFile, err := os.Open("grades.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer studentFile.Close()

	

	

	fileReader := csv.NewReader(studentFile)
    studentData,err := fileReader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	var singleStudentData student

	for index,record := range studentData {
		
		// array of string
		if index > 0 {
			singleStudentData.firstName = record[0]
			singleStudentData.lastName = record[1]
			singleStudentData.university = record[2]
			singleStudentData.test1Score,err = strconv.Atoi(record[3])
			singleStudentData.test2Score,err = strconv.Atoi(record[4])
			singleStudentData.test3Score,err = strconv.Atoi(record[5])
			singleStudentData.test4Score,err = strconv.Atoi(record[6])

			studentDetails = append(studentDetails, singleStudentData)
			
		}
        
     
	}


	return studentDetails
}

func calculateGrade(students []student) []studentStat {
	var studentGrades []studentStat

	var finalScore float32
	var grade Grade

	for _,studentData := range students {
       
		finalScore = float32(studentData.test1Score + studentData.test2Score + studentData.test3Score + studentData.test4Score)/float32(4)
	    if finalScore < 35 {
			grade = F
		} else if finalScore >= 35 && finalScore < 50 {
			grade = C 
		} else if finalScore >= 50 && finalScore < 70 {
			grade = B
		} else {
			grade = A 
		}
		studentGrades = append(studentGrades,studentStat{student:studentData,finalScore: float32(finalScore),grade: grade})
	}
	return studentGrades
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	
	
	maxScore := 0.0

	var topperStudent studentStat

	for _,gradedStudent := range gradedStudents {
		
		if gradedStudent.finalScore > float32(maxScore) {
			topperStudent = gradedStudent
			maxScore = float64(gradedStudent.finalScore)
		}

	}
	return topperStudent
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	topperUniversityWise := make(map[string]studentStat)

	universityWiseMaxScore := make(map[string]float32)

	for _,gradedStudent := range gs {

		   if gradedStudent.finalScore > universityWiseMaxScore[gradedStudent.student.university] {
			 topperUniversityWise[gradedStudent.student.university] = gradedStudent
			 universityWiseMaxScore[gradedStudent.student.university] = gradedStudent.finalScore
		   }
	}

	return topperUniversityWise
}
