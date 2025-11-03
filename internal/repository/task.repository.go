
package repository

import (
	"context"
	"payme/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Save(*models.Task) error
	Update(*models.Task) error
	Delete(string) error
	FindTaskByID(string) (*models.Task, error)
	FindTasks() ([]*models.Task, error)
	FindPublicTasks() ([]*models.Task, error)
}

type TaskMongoRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &TaskMongoRepository{
		collection: collection,
	}
}

func (ur *TaskMongoRepository)Save(task *models.Task) error {
	_, err := ur.collection.InsertOne(context.TODO(), task)
	return err
}


func (ur *TaskMongoRepository)Update(task *models.Task) error {
	_, err := ur.FindTaskByID(task.ID)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: task.ID}}
	ur.collection.FindOneAndReplace(context.TODO(), filter, task)
	return nil
}


func (ur *TaskMongoRepository)Delete(id string) error {
	if _, err := ur.FindTaskByID(id); err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: id}}
	_, err := ur.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (ur *TaskMongoRepository)FindTaskByID(id string) (*models.Task, error) {
	task := &models.Task{}
	filter := bson.D{{Key: "id", Value: id}}
	if err := ur.collection.FindOne(context.TODO(), filter).Decode(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (ur *TaskMongoRepository)FindPublicTasks() ([]*models.Task, error) {
	filters := bson.D{{Key: "status", Value: "public"}}
	cur, err := ur.collection.Find(context.Background(), filters)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())
	tasks := []*models.Task{}
	for cur.Next(context.Background()) {
		curTask := &models.Task{}
		if err := cur.Decode(&curTask); err != nil {
			return nil, err
		}
		tasks = append(tasks, curTask)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ur *TaskMongoRepository)FindTasks() ([]*models.Task, error) {
	cur, err := ur.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())
	tasks := []*models.Task{}
	for cur.Next(context.Background()) {
		curTask := &models.Task{}
		if err := cur.Decode(&curTask); err != nil {
			return nil, err
		}
		tasks = append(tasks, curTask)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
