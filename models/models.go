package models

type Student struct {
	StudentID int              `bson:studentid`
	Name      string           `bson:name`
	Age       int              `bson:age`
	Gender    string           `bson:gender`
	Grade     string           `bson:grade`
	Marks     []map[string]int `bson:marks`
}
