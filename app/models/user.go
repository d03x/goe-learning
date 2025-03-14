package models

type User struct {
	Id             uint   `db:"id"`
	Name           string `db:"name"`
	Age            uint   `db:"age"`
	Email          string `db:"email"`
	Address        string `db:"address"`
	Password       string `db:"password"`
	IsAdmin        bool   `db:"is_admin"`
	SchoolId       string `db:"school_id"`
	ProfilePicture string `db:"profile_picture"`
}
