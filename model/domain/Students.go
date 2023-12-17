package domain

type Student struct {
	ID      int    `json:"id" bson:"_id"`
	NIS     int    `json:"nis" bson:"nis"`
	Name    string `json:"name" bson:"name"`
	Jurusan string `json:"jurusan" bson:"jurusan"`
}
