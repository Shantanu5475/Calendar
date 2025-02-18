package models

type Holiday struct {
	// ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Date    string `json:"date" bson:"date"`
	Country string `json:"country" bson:"country"`
}
