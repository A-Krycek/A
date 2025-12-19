package data

import (
	"academic-enrollments/models"
	"time"
)

var Students = []models.Student{
	{1, "Juan", "juan@mail.com"},
	{2, "Maria", "maria@mail.com"},
	{3, "Carlos", "carlos@mail.com"},
	{4, "Ana", "ana@mail.com"},
	{5, "Luis", "luis@mail.com"},
}

var Courses = []models.Course{
	{1, "Go", 2},
	{2, "Bases de Datos", 3},
	{3, "Redes", 2},
	{4, "SO", 2},
	{5, "Seguridad", 3},
}

var Enrollments = []models.Enrollment{
	{1, 1, 1, 100, "pagado", time.Now()},
	{2, 2, 1, 100, "pendiente", time.Now()},
	{3, 3, 2, 120, "pagado", time.Now()},
	{4, 4, 3, 90, "pagado", time.Now()},
	{5, 5, 4, 110, "pendiente", time.Now()},
}

var NextID = 6
