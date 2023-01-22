package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/prajapatiomkar/todo-app-in-golang/config"
	"github.com/prajapatiomkar/todo-app-in-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var USERNAME string = config.Config("USER_NAME")
var PASSWORD string = config.Config("PASSWORD")
var DATABASE string = config.Config("DATABASE")
var COLLECTION string = config.Config("COLLECTION")

// Database Connection String
var CONNECTION_STRING = fmt.Sprintf("mongodb+srv://%v:%v@cluster0.bbxjqad.mongodb.net/?retryWrites=true&w=majority", USERNAME, PASSWORD)

// Collection Object/Instance
var collection *mongo.Collection

// Create Connection With MongoDB
func init() {
	clientOption := options.Client().ApplyURI(CONNECTION_STRING)

	client, err := mongo.Connect(context.TODO(), clientOption)
	HandleErr(err)

	// Check The Connection
	err = client.Ping(context.TODO(), nil)
	HandleErr(err)

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(DATABASE).Collection(COLLECTION)
	fmt.Println("Collection instance created!")
}

// Get All Task
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Using Define Model
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	payload := getAllTask()

	json.NewEncoder(w).Encode(payload)
}

// Create Task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

// Task Complete
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask
func UndoTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteAllTask
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "")

	count := deletedAllTask()
	json.NewEncoder(w).Encode(count)
}

// -----------------------------------------------------------------------------------------------------
// GETTING DATA FORM DB

// Get All Task Form The DB And Return It
func getAllTask() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	HandleErr(err)

	var results []primitive.M

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		HandleErr(err)

		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		HandleErr(err)
	}
	defer cursor.Close(context.Background())
	return results
}

// Insert One Task In The DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)
	HandleErr(err)
	fmt.Println("Inserted a single record", insertResult.InsertedID)
}

// Task Complete Method, Update Task’s Status To True
func taskComplete(task string) {
	fmt.Println(task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"status": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	HandleErr(err)

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// Task Undo Method, Update Task's status To False
func undoTask(task string) {
	fmt.Println(task)

	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"status": false}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	HandleErr(err)
	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// Delete One Task From The DB, Delete By Id
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)

	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	HandleErr(err)
	fmt.Println("Deleted Document ", deleted.DeletedCount)
}

// Delete All The Tasks From The DB
func deletedAllTask() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	HandleErr(err)

	fmt.Println("Deleted Documents ", deleted.DeletedCount)
	return deleted.DeletedCount
}

// Handling ERR Function
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
