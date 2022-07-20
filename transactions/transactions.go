package transactions

import (
	"booking-app/database"
	"booking-app/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	transactions := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	database.DB.Create(&transactions)

}
