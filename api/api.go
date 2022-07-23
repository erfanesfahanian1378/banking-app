package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"booking-app/helpers"
	"booking-app/useraccounts"
	"booking-app/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type TransactionBody struct {
	UserId uint
	From   uint
	To     uint
	Amount int
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "well done !" {
		resp := call
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusForbidden)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	apiResponse(login, w)

}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register

	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	apiResponse(register, w)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("TOSHE")
	vars := mux.Vars(r)
	userId := vars["id"]
	// auth := r.Header.Get("Authorization")

	user := users.GetUser(userId)

	fmt.Println("--------------------------------------")
	fmt.Println(user)
	fmt.Println("--------------------------------------")

	apiResponse(user, w)
}

func transaction(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	auth := r.Header.Get("Authorization")
	var formattedBody TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, auth)
	apiResponse(transaction, w)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	// router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.Handle("/user/{id}", helpers.Middleware(http.HandlerFunc(getUser)))
	fmt.Println("App is running on port : 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
