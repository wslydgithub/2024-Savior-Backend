package model

type User struct {
	Name     string
	Password string
}

// 添加用户
func (user *User) Useradd(user1 Usersl) {
	user.Name = user1.Email
	user.Password = user1.Password
}

// 将用户信息存入数据库
func Adduser() {
	var a User
	a.Useradd(Usersign)
	DB.Create(a)
}

/*// 将用户的数据存入数据库中
func Createuser(user1 Usersl) {
	var user User
	user.Useradd(user1)
	DB.Create(&user)
}
*/
