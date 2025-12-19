package handlers

import (
	"academic-enrollments/data"
	"academic-enrollments/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func EnrollmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(data.Enrollments)
		return
	}

	if r.Method == "POST" {
		var e models.Enrollment
		json.NewDecoder(r.Body).Decode(&e)

		foundStudent := false
		for _, s := range data.Students {
			if s.ID == e.StudentID {
				foundStudent = true
			}
		}
		if !foundStudent {
			http.Error(w, "Estudiante no encontrado", 400)
			return
		}

		foundCourse := false
		max := 0
		for _, c := range data.Courses {
			if c.ID == e.CourseID {
				foundCourse = true
				max = c.MaxQuota
			}
		}
		if !foundCourse {
			http.Error(w, "Curso no encontrado", 400)
			return
		}

		count := 0
		for _, m := range data.Enrollments {
			if m.CourseID == e.CourseID {
				count++
			}
		}
		if count >= max {
			http.Error(w, "Cupo agotado", 409)
			return
		}

		if e.Amount <= 0 {
			http.Error(w, "Valor invalido", 400)
			return
		}

		if e.Status != "pagado" && e.Status != "pendiente" {
			http.Error(w, "Estado invalido", 400)
			return
		}

		e.ID = data.NextID
		data.NextID++
		e.Date = time.Now()

		data.Enrollments = append(data.Enrollments, e)

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(e)
	}
}

func EnrollmentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/enrollments/")
	id, _ := strconv.Atoi(idStr)

	for i, e := range data.Enrollments {
		if e.ID == id {

			if r.Method == "GET" {
				json.NewEncoder(w).Encode(e)
				return
			}

			if r.Method == "PUT" {
				var input models.Enrollment
				json.NewDecoder(r.Body).Decode(&input)

				if input.Status != "pagado" && input.Status != "pendiente" {
					http.Error(w, "Estado invalido", 400)
					return
				}

				data.Enrollments[i].Status = input.Status
				json.NewEncoder(w).Encode(data.Enrollments[i])
				return
			}

			if r.Method == "DELETE" {
				data.Enrollments = append(data.Enrollments[:i], data.Enrollments[i+1:]...)
				w.WriteHeader(204)
				return
			}
		}
	}

	http.Error(w, "Matricula no encontrada", 404)
}
