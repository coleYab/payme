package repository

import (
	"context"
	"fmt"
	"payme/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(*models.User) error
	Update(*models.User) error
	Delete(string) error
	FindUserByID(string) (*models.User, error)
	FindUsers() ([]*models.User, error)
	FindUserByEmail(string) (*models.User, error)
}

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserMongoRepository{
		collection: collection,
	}
}

func (ur *UserMongoRepository)Save(user *models.User) error {
	if _, err := ur.FindUserByEmail(user.Email); err == nil {
		return fmt.Errorf("email has to be unique")
	}

	if _, err := ur.FindUserByID(user.ID); err == nil {
		return fmt.Errorf("email has to be unique")
	}

	_, err := ur.collection.InsertOne(context.TODO(), user)

	return err
}


func (ur *UserMongoRepository)Update(user *models.User) error {
	_, err := ur.FindUserByID(user.ID)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: user.ID}}
	ur.collection.FindOneAndReplace(context.TODO(), filter, user)
	return nil
}


func (ur *UserMongoRepository)Delete(id string) error {
	if _, err := ur.FindUserByID(id); err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: id}}
	_, err := ur.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (ur *UserMongoRepository)FindUserByID(id string) (*models.User, error) {
	user := &models.User{}
	filter := bson.D{{Key: "id", Value: id}}
	if err := ur.collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserMongoRepository)FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	filter := bson.D{{Key: "email", Value: email}}
	if err := ur.collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}


func (ur *UserMongoRepository)FindUsers() ([]*models.User, error) {
	cur, err := ur.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())
	users := []*models.User{}
	for cur.Next(context.Background()) {
		curUser := &models.User{}
		if err := cur.Decode(&curUser); err != nil {
			return nil, err
		}
		users = append(users, curUser)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
