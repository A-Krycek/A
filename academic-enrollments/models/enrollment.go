package models

import "time"

type Enrollment struct {
	ID        int
	StudentID int
	CourseID  int
	Amount    float64
	Status    string
	Date      time.Time
}
