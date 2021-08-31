package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FindReturn struct {
	Librarian Librarian
	Student   Student
	Books     Books
}

type Librarian struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Librarian_No string             `json:"librarian_no"`
	Password     string             `json:"password"`
}

type Student struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Student_Name string             `json:"student_name"`
	Student_ID   int                `json:"student_id"`
	Password     string             `json:"password"`
	Borrowed     []int              `json:"borrowed"`
}

type Books struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Book_Title  string             `json:"book_title"`
	Book_ID     int                `json:"book_id"`
	Author      string             `json:"author"`
	Quantity    int                `json:"quantity"`
	Description string             `json:"description"`
}
type Borrow struct {
	Student   Student `json:"student"`
	Borrow_ID int     `json:"borrow_id"`
}
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

var SaveToken = make(map[string]string)
