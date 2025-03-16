package models

type User struct {
	_id          string     `json:"id" bson:"_id"`
	email        string     `json:"email" bson:"email"`
	phone_number string     `json:"phone_number" bson:"phone_number"`
	password     string     `json:"password" bson:"password"`
	full_name    string     `json:"full_name" bson:"full_name"`
	role         string     `json:"role" bson:"role"`
	created_at   string     `json:"created_at" bson:"created_at"`
	updated_at   string     `json:"updated_at" bson:"updated_at"`
	Children     []Children `json:"children" bson:"children"`
}

type Admin struct {
	_id       string `json:"id" bson:"_id"`
	username  string `json:"username" bson:"username"`
	password  string `json:"password" bson:"password"`
	full_name string `json:"full_Name" bson:"full_name"`
	role      string `json:"role" bson:"role"`
}

type Children struct {
	_id        string `json:"id" bson:"_id"`
	parent_id  string `json:"parent_id" bson:"parent_id"`
	full_name  string `json:"full_name" bson:"full_name"`
	age        int    `json:"age" bson:"age"`
	created_at string `json:"created_at" bson:"created_at"`
	updated_at string `json:"updated_at" bson:"updated_at"`
}
