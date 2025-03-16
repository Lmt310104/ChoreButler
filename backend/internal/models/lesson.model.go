package models

type Lesson struct {
	_id           string       `json:"id" bson:"_id"`
	thumbnail_url string       `json:"thumbnail_url" bson:"thumbnail_url"`
	title         string       `json:"title" bson:"title"`
	description   string       `json:"description" bson:"description"`
	lesson_items  []LessonItem `json:"lesson_items" bson:"lesson_items"`
	created_at    string       `json:"created_at" bson:"created_at"`
	updated_at    string       `json:"updated_at" bson:"updated_at"`
}

type LessonItem struct {
	_id               string   `json:"id" bson:"_id"`
	lesson_id         string   `json:"lesson_id" bson:"lesson_id"`
	title             string   `json:"title" bson:"title"`
	order             int      `json:"order" bson:"order"`
	description       string   `json:"description" bson:"description"`
	video_url         string   `json:"video_url" bson:"video_url"`
	custom_video_urls []string `json:"custom_video_urls" bson:"custom_video_urls"`
}
