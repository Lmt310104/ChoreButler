package schemas

import "go.mongodb.org/mongo-driver/v2/bson"

type Progress struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	ChildrenID  string        `json:"children_id" bson:"children_id"`
	IsCompleted bool          `json:"is_completed" bson:"is_completed"`
	LessonID    string        `json:"lesson_id" bson:"lesson_id"`
}
