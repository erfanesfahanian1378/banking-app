// package migrations

// import(
// 	"duomly.com/go-bank-backend/helpers"
// 	"github.com/jinzhu/gorm"
// 	_"github.com/jinzhu/gorm/dialects/postgres"
// )

// type User struct {
// 	gorm.Model
// 	Username string
// 	Email string
// 	Password string
// }

// type Account struct {
// 	gorm.Model
// 	Type string
// 	Name string
// 	Balance string
// 	UserID uint
// }

// func connectDB() *gorm.DB {
// 	db , err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=erfan dbname=bankapp password=181352
// 	sslmode=disable")
// 	return db
// }

package migrations

import (
	"booking-app/database"
	"booking-app/helpers"
	"booking-app/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// This is correct way of creating password
// func HashAndSalt(pass []byte) string {
// 	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
// 	helpers.HandleErr(err)

// 	return string(hashed)
// }

func createAccounts() {

	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

func Migrate() {
	// User := &interfaces.User{}
	// Account := &interfaces.Account{}
	// db := helpers.ConnectDB()
	// db.AutoMigrate(&User{}, &Account{})
	// defer db.Close()

	// createAccounts()
}

func MigrteTransactions() {
	// Transactions := &interfaces.Transaction{}

	// db := helpers.ConnectDB()
	// db.AutoMigrate(&Transactions)
	// defer db.Close()
}
