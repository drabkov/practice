package infrastructure

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myserver/domain"
)

// SERVER the DB server
const SERVER = "mongodb://localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "bookstore"

// DOCNAME the name of the document
const DOCNAME = "books"

// ConnectDB : This is helper function to connect mongoDB
func connectDB() *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI(SERVER)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database(DBNAME).Collection(DOCNAME)

	return collection
}

func GetBooks() ([]domain.Book, error) {

	// we created Book array
	var books []domain.Book

	//Connection mongoDB with helper class
	collection := connectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return books, err
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book domain.Book
		// & character returns the memory address of the following variable.
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return books, nil
}

func GetBookByID(id primitive.ObjectID) (*domain.Book, error) {

	var book domain.Book

	collection := connectDB()

	// We create filter.
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		return &book, err
	}

	return &book, nil
}

func CreateBook(book *domain.Book) (*mongo.InsertOneResult, error) {
	// connect db
	collection := connectDB()

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateBookByID(id primitive.ObjectID, book *domain.Book) error {
	collection := connectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := connectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
