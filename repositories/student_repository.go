package repositories

import (
	"context"
	"fmt"

	"student-management-system/models"
	"student-management-system/utils"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	collection *mongo.Collection
}

func NewStudentRepository(db *mongo.Database) *StudentRepository {
	return &StudentRepository{
		collection: db.Collection("students"),
	}
}

func (r *StudentRepository) Create(ctx context.Context, student *models.Student) error {
	_, err := r.collection.InsertOne(ctx, student)
	return err
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]*models.Student, error) {
	cur, err := r.collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var students []*models.Student
	for cur.Next(ctx) {
		var student models.Student
		if err := cur.Decode(&student); err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepository) GetByID(ctx context.Context, id string) (*models.Student, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var student models.Student
	err = r.collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Update(ctx context.Context, id string, student *models.Student) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.LogInfo2("provided ID is not a valid student ID")

		return fmt.Errorf("provided ID is not a valid student ID")
	}

	update := bson.M{
		"$set": bson.M{
			"first_name": student.FirstName,
			"last_name":  student.LastName,
			"course":     student.Course,
			"grade":      student.Grade,
			"updated_by": student.UpdatedBy,
			"updated_on": student.UpdatedOn.UTC(), // Ensure UTC format
		},
	}

	_, err = r.collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return err
	}
	return nil
}
func (r *StudentRepository) Delete(ctx context.Context, idStr string) error {
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		utils.LogInfo2("invalid student ID")

		return fmt.Errorf("invalid student ID: %w", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("error deleting student: %w", err)
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	utils.LogInfo("Student deleted successfully", idStr)

	return nil
}
