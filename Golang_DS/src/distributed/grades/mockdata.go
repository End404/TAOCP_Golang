package grades

func init() {
	studetns = []Student{
		{
			ID: 1, FirstName: "Nick", LastName: "Carter",
			Grades: []Grade{
				{Title: "Quiz 1", Type: GradeQuiz, Score: 85},
				{Title: "Final Exam", Type: GradeQuiz, Score: 95},
				{Title: "Quiz 2", Type: GradeQuiz, Score: 75},
			},
		},
	}
}
