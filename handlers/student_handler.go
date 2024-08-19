package handlers

import (
	"encoding/json"
	"net/http"
	"student-management-system/models"
	"student-management-system/services"
	"student-management-system/utils"

	"go.mongodb.org/mongo-driver/mongo"

	"errors"
	"time"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	ctx := r.Context()

	username, _ := utils.GetUsernameFromContext(ctx)

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(ctx, &student, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.LogInfo3("Student created successfully", student.FirstName)

}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}

	if len(students) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No students found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Indent JSON for better readability
	encoder.Encode(students)    // Encode without extra error handling
	utils.LogInfo2("All students fetched successfully")

}

func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	student, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(student, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	utils.LogInfo("Student fetched successfully ", id)

	w.Write(prettyJSON)
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var student models.Student
	ctx := r.Context()
	username, _ := utils.GetUsernameFromContext(ctx)

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.UpdatedOn = time.Now().UTC()
	student.UpdatedBy = username

	err := h.service.Update(ctx, id, &student)
	if err != nil {
		if err.Error() == "provided ID is not a valid student ID" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	utils.LogInfo("Student updated successfully", id)

	w.WriteHeader(http.StatusNoContent)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	err := h.service.Delete(r.Context(), idStr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Student not found", http.StatusNotFound)
		} else {
			http.Error(w, "Student with that ID is not present", http.StatusInternalServerError)
		}
		return
	}
	utils.LogInfo("Student deleted successfully", idStr)

	w.WriteHeader(http.StatusNoContent)
}
