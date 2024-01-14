package app

type User struct {
	Name     string `json:"name" form:"name" xml:"name"`
	Username string `json:"username,omitempty" form:"username" xml:"username"`
	Email    string `json:"email,omitempty" xml:"email"`
	Password string `json:"password,omitempty" form:"password" xml:"password"`
}
