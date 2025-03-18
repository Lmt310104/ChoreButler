package schemas

import "go.mongodb.org/mongo-driver/v2/bson"

type Lesson struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	ThumbnailUrl string        `json:"thumbnail_url" bson:"thumbnail_url"`
	Title        string        `json:"title" bson:"title"`
	Description  string        `json:"description" bson:"description"`
	LessonItems  []LessonItem  `json:"lesson_items" bson:"lesson_items"`
	CreatedAt    string        `json:"created_at" bson:"created_at"`
	UpdatedAt    string        `json:"updated_at" bson:"updated_at"`
}

type LessonItem struct {
	ID              string   `json:"id" bson:"_id"`
	LessonID        string   `json:"lesson_id" bson:"lesson_id"`
	Title           string   `json:"title" bson:"title"`
	Order           int      `json:"order" bson:"order"`
	Description     string   `json:"description" bson:"description"`
	VideoUrl        string   `json:"video_url" bson:"video_url"`
	CustomVideoUrls []string `json:"custom_video_urls" bson:"custom_video_urls"`
}
