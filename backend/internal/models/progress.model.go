package models

type Progress struct {
	_id          string `json:"id" bson:"_id"`
	children_id  string `json:"children_id" bson:"children_id"`
	is_completed bool   `json:"is_completed" bson:"is_completed"`
	lesson_id    string `json:"lesson_id" bson:"lesson_id"`
}
