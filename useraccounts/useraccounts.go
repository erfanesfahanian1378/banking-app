package useraccounts

import (
	"booking-app/database"
	"booking-app/helpers"
	"booking-app/interfaces"
	"booking-app/transactions"
	"strconv"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	// didnt need a int() ?
	responseAcc.Balance = account.Balance
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	// userIdString := fmt.Strint(userId)
	userIdString := strconv.Itoa(int(userId))
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "Your not the owner"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}

		updateAccount := updateAccount(from, int(fromAccount.Balance)-amount)

		// updateAccount2 := updateAccount(from, int(fromAccount.Balance)-amount)
		// updateAccount2 := updateAccount(to, int(fromAccount.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = updateAccount
		return response
	} else {
		return map[string]interface{}{"message": "not valid tokens"}
	}
}
