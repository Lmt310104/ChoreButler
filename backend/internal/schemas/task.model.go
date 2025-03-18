package schemas

import "go.mongodb.org/mongo-driver/v2/bson"

type Task struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	ParentID     string        `json:"parent_id" bson:"parent_id"`
	AssigneeID   string        `json:"assignee_id" bson:"assignee_id"`
	Title        string        `json:"title" bson:"title"`
	NumberOfStar int           `json:"number_of_star" bson:"number_of_star"`
	Evidence     string        `json:"evidence_image" bson:"evidence_image"`
	Description  string        `json:"description" bson:"description"`
	DueDate      string        `json:"due_date" bson:"due_date"`
	CreatedAt    string        `json:"created_at" bson:"created_at"`
	UpdatedAt    string        `json:"updated_at" bson:"updated_at"`
	Review       Review        `json:"review" bson:"review"`
}

type Review struct {
	ID      bson.ObjectID `json:"id" bson:"_id"`
	TaskID  string        `json:"task_id" bson:"task_id"`
	Comment string        `json:"comment" bson:"comment"`
	Rating  int           `json:"rating" bson:"rating"`
}
