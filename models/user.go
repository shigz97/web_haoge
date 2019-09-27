package models

import (
	orm "web_haoge/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

//添加
func (user User) Insert() (id int64, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	return
}

//删除
func (user *User) Destroy(id int64) (Res User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Res = *user
	return
}

//查询
func (user *User) Finduser() (Res User, err error) {
	err = orm.Eloquent.Where(" username= ?", user.Username).First(&Res).Error
	if err != nil {
		return
	}
	return
}
