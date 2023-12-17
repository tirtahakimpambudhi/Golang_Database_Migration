package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Books struct {
	ID                int        `json:"id" bson:"_id"`
	ISBN              string     `json:"isbn" bson:"isbn"`
	Title             string     `json:"title" bson:"title"`
	Author            string     `json:"author" bson:"author"`
	Status_Borrow     bool       `json:"status_borrow" bson:"status_borrow"`
	Publisher         string     `json:"publisher" bson:"publisher"`
	Publication_Years string     `json:"publication_years" bson:"publication_years"`
	Description       string     `json:"description" bson:"description"`
	CreatedAt         *time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt          *time.Time `json:"updatedAt" bson:"updateAt"`
}
type BooksCollections struct {
	ISBN              string     `json:"isbn" bson:"isbn"`
	Title             string     `json:"title" bson:"title"`
	Author            string     `json:"author" bson:"author"`
	Status_Borrow     bool       `json:"status_borrow" bson:"status_borrow"`
	Publisher         string     `json:"publisher" bson:"publisher"`
	Publication_Years string     `json:"publication_years" bson:"publication_years"`
	Description       string     `json:"description" bson:"description"`
	CreatedAt         *time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt          *time.Time `json:"updatedAt" bson:"updateAt"`
}

type BooksDocument struct {
	ID                primitive.ObjectID `bson:"_id"`
	ISBN              string             `json:"isbn" bson:"isbn"`
	Title             string             `json:"title" bson:"title"`
	Author            string             `json:"author" bson:"author"`
	Status_Borrow     bool               `json:"status_borrow" bson:"status_borrow"`
	Publisher         string             `json:"publisher" bson:"publisher"`
	Publication_Years string             `json:"publication_years" bson:"publication_years"`
	Description       string             `json:"description" bson:"description"`
	CreatedAt         *time.Time         `json:"createdAt" bson:"createdAt"`
	UpdateAt          *time.Time         `json:"updatedAt" bson:"updateAt"`
}

//type BooksCollections struct {
//	ID                string     `json:"id" bson:"_id"`
//	ISBN              string     `json:"isbn" bson:"isbn"`
//	Title             string     `json:"title" bson:"title"`
//	Author            string     `json:"author" bson:"author"`
//	Status_Borrow     bool       `json:"status_borrow" bson:"status_borrow"`
//	Publisher         string     `json:"publisher" bson:"publisher"`
//	Publication_Years string     `json:"publication_years" bson:"publication_years"`
//	Description       string     `json:"description" bson:"description"`
//	CreatedAt         *time.Time `json:"createdAt" bson:"createdAt"`
//	UpdateAt          *time.Time `json:"updatedAt" bson:"updateAt"`
//}
