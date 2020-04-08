package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Test_33"
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var users []User
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}
	fmt.Println(users)
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

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
