package models

type Task struct {
	_id            string `json:"id" bson:"_id"`
	parent_id      string `json:"parent_id" bson:"parent_id"`
	assignee_id    string `json:"assignee_id" bson:"assignee_id"`
	title          string `json:"title" bson:"title"`
	number_of_star int    `json:"number_of_star" bson:"number_of_star"`
	evidence_image string `json:"evedent_image" bson:"evedent_image"`
	description    string `json:"description" bson:"description"`
	due_date       string `json:"due_date" bson:"due_date"`
	created_at     string `json:"created_at" bson:"created_at"`
	updated_at     string `json:"updated_at" bson:"updated_at"`
	review         Review `json:"review" bson:"review"`
}

type Review struct {
	_id     string `json:"id" bson:"_id"`
	task_id string `json:"task_id" bson:"task_id"`
	comment string `json:"comment" bson:"comment"`
	rating  int    `json:"rating" bson:"rating"`
}
