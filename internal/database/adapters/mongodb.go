package adapters

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDbDatabase[T database.RecordType] struct {
	db *mongo.Database
}

func NewMongoDBClient() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set the 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	dbName := "streaming_platform"

	return client.Database(dbName)
}

func AdaptTypeWithMongoDb[T database.RecordType](db *mongo.Database) *mongoDbDatabase[T] {
	return &mongoDbDatabase[T]{db: db}
}

func (mdba *mongoDbDatabase[T]) Add(collectionName string, record T) (string, error) {
	collection := mdba.db.Collection(collectionName)

	documentInsertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return "", err
	}

	switch v := documentInsertResult.InsertedID.(type) {
	case primitive.ObjectID:
		return v.Hex(), nil
	case string:
		return v, nil
	default:
		return "", errors.New("id is neither string nor objectID")
	}
}

func (mdba *mongoDbDatabase[T]) Get(collectionName string, id string) (T, error) {
	var value T
	collection := mdba.db.Collection(collectionName)

	var recordId any
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		recordId = id
	}

	doc := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: recordId}})
	if doc.Err() != nil {
		return value, doc.Err()
	}

	err = doc.Decode(&value)
	if err != nil {
		println(err.Error())
		return value, fmt.Errorf("unable to find record of %T with id: %s in collection: %s", value, id, collectionName)
	}

	return value, nil
}

func (mdba *mongoDbDatabase[T]) Update(collectionName string, id string, updatedRecord T) error {
	collection := mdba.db.Collection(collectionName)

	var recordId any
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		recordId = id
	}

	updateResult, err := collection.UpdateByID(context.TODO(), recordId, bson.M{"$set": updatedRecord})
	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("no document found with id: %s in collection: %s", id, collectionName)
	}

	return nil
}

func (mdba *mongoDbDatabase[T]) Delete(collectionName string, id string) error {
	collection := mdba.db.Collection(collectionName)

	var recordId any
	recordId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		recordId = id
	}

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: recordId}})
	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("no document found with id: %s in collection: %s", id, collectionName)
	}

	return nil
}

func (mdba *mongoDbDatabase[T]) GetWithPagination(collectionName string, page int, pageSize int) ([]T, error) {
	collection := mdba.db.Collection(collectionName)

	opts := options.Find().SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))
	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []T
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
