package routes

import (
	"student-management-system/handlers"
	"student-management-system/middlewares"
	"student-management-system/repositories"
	"student-management-system/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) *mux.Router {
	studentRepo := repositories.NewStudentRepository(db)
	studentService := services.NewStudentService(studentRepo)
	studentHandler := handlers.NewStudentHandler(studentService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/register", authHandler.Register).Methods("POST")

	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTAuth)

	api.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	api.HandleFunc("/students", studentHandler.GetAllStudents).Methods("GET")
	api.HandleFunc("/students/{id}", studentHandler.GetStudentByID).Methods("GET")
	api.HandleFunc("/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	api.HandleFunc("/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")

	return router
}
