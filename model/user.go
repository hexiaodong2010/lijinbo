package model

import "github.com/jinzhu/gorm"

func init() {

}
type User struct {
	Openid string
	gorm.Model
}
