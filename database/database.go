package database

import (
	"Hands_On/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = "Hands_On"
var librarian = "librarian"
var student = "student"
var books = "books"

func createConnection(url string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel, err

}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func CreateStudent(user model.Student) error {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal("connection to database failed")
	}
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(student)
	var count int64
	count, err = collection.EstimatedDocumentCount(context)
	if err != nil {
		log.Fatal(err)
	}
	user.Student_ID = int(count) + 1
	_, err = collection.InsertOne(context, user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetBooks() []model.Books {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal("connection to database failed")
	}
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(books)
	res, err := collection.Find(context, bson.D{})
	if err != nil {
		log.Fatal("error fetching")
	}
	var users []model.Books
	if err = res.All(context, &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func FindLibrarian(name string) (model.Librarian, error) {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		return model.Librarian{}, err
	} else {

		defer close(client, context, cancel)
		collection := client.Database(db).Collection(librarian)
		var user model.Librarian
		err = collection.FindOne(context, bson.M{"librarian_no": name}).Decode(&user)
		if err != nil {
			return model.Librarian{}, err
		}
		return user, nil
	}
}
func FindStudent(name string) (model.Student, error) {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		return model.Student{}, err
	} else {

		defer close(client, context, cancel)
		collection := client.Database(db).Collection(student)
		var user model.Student
		err = collection.FindOne(context, bson.M{"student_name": name}).Decode(&user)
		if err != nil {
			return model.Student{}, err
		}
		return user, nil
	}
}

func UpdateBook(update model.Books) error {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal("conection to database failed")
	}
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(books)
	filter := bson.M{"book_id": update.Book_ID}
	query := bson.M{"$set": bson.M{"quantity": update.Quantity}}
	_, err = collection.UpdateOne(context, filter, query)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// func DeleteUser(name string) {
// 	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
// 	CheckError(err)
// 	defer close(client, context, cancel)
// 	collection := client.Database(db).Collection(col)
// 	_, err = collection.DeleteOne(context, bson.M{"name": name})
// 	CheckError(err)
// }

func AddBook(user model.Books) error {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal("connection to database failed")
	}
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(books)
	var count int64
	count, err = collection.EstimatedDocumentCount(context)
	if err != nil {
		log.Fatal(err)
	}
	user.Book_ID = int(count) + 1
	_, err = collection.InsertOne(context, user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Borrow(update model.Borrow) error {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal("conection to database failed")
	}
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(student)
	var user model.Student
	filter := bson.M{"student_id": update.Student.Student_ID}
	err = collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		log.Fatal("error", err)
	}
	user.Borrowed = append(user.Borrowed, update.Borrow_ID)
	query := bson.M{"$set": bson.M{"borrowed": user.Borrowed}}
	_, err = collection.UpdateOne(context, filter, query)
	if err != nil {
		log.Fatal("could not update student borrow list", err)
	}
	books := client.Database(db).Collection(books)
	var newquantity model.Books
	err = books.FindOne(context, bson.M{"book_id": update.Borrow_ID}).Decode(&newquantity)
	if err != nil {
		log.Fatal("could not get boook")
	}
	filter = bson.M{"book_id": update.Borrow_ID}
	query = bson.M{"$set": bson.M{"quantity": newquantity.Quantity - 1}}
	_, err = books.UpdateOne(context, filter, query)
	if err != nil {
		return err
	} else {
		return nil
	}
}
