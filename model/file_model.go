package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FileName         string             `bson:"file_name" json:"file_name" validate:"required"`
	FileType         string             `bson:"file_type" json:"file_type" validate:"required"`
	Size             int64              `bson:"size" json:"size" validate:"gte=0"`
	UploadDate       time.Time          `bson:"upload_date" json:"upload_date"`
	LastModifiedDate time.Time          `bson:"last_modified_date" json:"last_modified_date"`
	Content          []byte             `bson:"content,omitempty" json:"content,omitempty"`
	ContentType      string             `bson:"content_type" json:"content_type" validate:"required"`
}
