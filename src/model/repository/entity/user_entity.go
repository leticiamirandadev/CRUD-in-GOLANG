package entity
type UserEntity struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Name     string `bson:"name,omitempty"`
	Age      int8   `bson:"age,omitempty"`
}
