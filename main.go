package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	db.AutoMigrate(&User{})

	// newUser := User{
	// 	Name:        "Michał",
	// 	LastName:    "",
	// 	AccessToken: "",
	// }
	// db.Create(&newUser)

	// result := db.Delete(&User{}, 1)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// users := []User{}
	// result := db.Find(&users, "name = ?", "Michał")
	// if result.Error != nil {
	// 	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 		fmt.Println(result.Error)
	// 		return
	// 	}
	// }
	// fmt.Println(users)

	// user := User{}
	// user.ID = 2
	// user.Name = "Michał2"
	// user.LastName = "Nowak"

	// result := db.Model(&user).Updates(user)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
}

type User struct {
	gorm.Model
	Name        string `json:"name";gorm:"name"`
	LastName    string `json:"lastName";gorm:"last_name"`
	AccessToken string `json:"accessToken";gorm:"access_token"`
}
