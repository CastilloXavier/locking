package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Test_33"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	user, err := us.ByID(1)
	/*db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})
	*/
	/*var users []User
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}*/
	fmt.Println(user)
	//fmt.Println(users.Orders)
	//createOrder(db, u, 1001, "Fake Description #1")
	//createOrder(db, u, 9999, "Fake Description #2")
	//createOrder(db, u, 100, "Fake Description #3")

	/*db = db.Where("email = ?", "blah@blah.com").First(&u)
	if err := db.Where("email = ?", "blah@blah.com").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found!")
		default:
			panic(err)
		}
	}
	fmt.Println(u)*/
	/*if db.RecordNotFound() {
		fmt.Println("No user found!")
	} else if db.Error != nil {
		panic(db.Error)
	} else {
		fmt.Println(u)
	}*/
}
