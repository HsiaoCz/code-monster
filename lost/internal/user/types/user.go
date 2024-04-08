package types

type User struct {
	ID        string `bson:"_id" json:"id"`
	Firstname string `bson:"first_name" json:"first_name"`
	Lastname  string `bson:"last_name" json:"last_name"`
	Passwrod  string `bson:"password" json:"-"`
	Email     string `bson:"email" json:"email"`
	IsAdmin   bool   `bson:"isAdmin" json:"isAdmin"`
}
