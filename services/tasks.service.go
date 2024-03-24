package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jonpena/api-task-go/config"
	"github.com/jonpena/api-task-go/models"
	"go.mongodb.org/mongo-driver/bson"
)



func GetTasks() []models.Tarea {
  // Select the collection.
  collection := config.DB.Collection("tareas")

  // Get all documents in the collection.
  cursor, err := collection.Find(context.Background(), bson.M{})

  if err != nil {
    fmt.Println("Error getting documents:", err)
    return []models.Tarea{}
  }

  // Iterate over the cursor and print the documents.
  var tareas []models.Tarea

  for cursor.Next(context.Background()) {
    var tarea models.Tarea
    err := cursor.Decode(&tarea)
    if err != nil {
      fmt.Println("Error decoding document: ", err)
      return []models.Tarea{}
    }
    tareas = append(tareas, tarea)
  }

  if len(tareas) == 0 {
    return []models.Tarea{}
  }

  // Return the tasks.
  return tareas
}



func GetTask(id string) models.Tarea {
  // Select the collection.
  collection := config.DB.Collection("tareas")

  // Create a filter to search for the document by its ID.
  filter := bson.M{"id": id}

  // Search for the document by its ID.
  var tarea models.Tarea
  err := collection.FindOne(context.Background(), filter).Decode(&tarea)
  if err != nil {
    fmt.Println("Error searching for document:", err)
    return models.Tarea{}
  }

  return tarea
}



func CreateTask(body models.Tarea) models.Tarea {
  // Select the collection.
  collection := config.DB.Collection("tareas")

  body.CreatedAt = time.Now()
  body.UpdatedAt = time.Now()

  // Create a document.
  document := bson.M {
  "id": body.Id, 
  "title": body.Title, 
  "description": body.Description, 
  "createdAt": body.CreatedAt,
  "updatedAt": body.UpdatedAt,
  }

  // Insert the document into the collection.
  result, err := collection.InsertOne(context.Background(), document)
  if err != nil {
    fmt.Println("Error inserting document:", err)
    return models.Tarea{}
  }

  // Print the document ID.
  fmt.Println("Document Created successfully: ", result.InsertedID)

  return body;
}



func UpdateTask(id string, body models.Tarea) int64 {
  // Select the collection.
  collection := config.DB.Collection("tareas")

  // Create a filter to search for the document by its ID.
  filter := bson.M{"id": id}

  // Create the update document.
  update := bson.M{
    "$set": bson.M{
      "title": body.Title,
      "description": body.Description,
      "updatedAt": time.Now(),
    },
  }

  // Update the document.
  result, err := collection.UpdateOne(context.Background(), filter, update)
  if err != nil {
    panic(err)
  }
  // Return the number of documents updated.
  return result.ModifiedCount
}



func DeleteTask(id string) int64 {
  // Select the collection.
  collection := config.DB.Collection("tareas")

  // Create a filter to search for the document by its ID.
  filter := bson.M{"id": id}

  // Delete the document by its ID.
  result, err := collection.DeleteOne(context.Background(), filter)
  if err != nil {
    fmt.Println("Error deleting document:", err)
    return 0
  }

  // Print the number of documents deleted.
  fmt.Println("Documents Deleted:", result.DeletedCount)

  return result.DeletedCount
}
