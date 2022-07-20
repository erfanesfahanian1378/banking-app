package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"booking-app/interfaces"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// func ConnectDB() *gorm.DB {
// 	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=erfan dbname=bankapp password=181352 sslmode=disable")
// 	HandleErr(err)
// 	return db
// }

func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}

	}
	return true
}

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)
				resp := interfaces.ErrResponse{Message: "internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func ValidateToken(id string, jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	HandleErr(err)

	var userId, _ = strconv.ParseFloat(id, 8)
	if token.Valid && tokenData["user_id"] == userId {
		return true
	} else {
		return false
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		vars := mux.Vars(r)
		userId := vars["id"]
		isValid := ValidateToken(userId, auth)
		if isValid {
			fmt.Print("VALID")
			next.ServeHTTP(w, r)
		}
	})
}

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
// 		if len(authHeader) != 2 {
// 			fmt.Println("Malformed token")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Malformed Token"))
// 		} else {
// 			jwtToken := authHeader[1]
// 			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 				}
// 				return []byte(SECRETKEY), nil
// 			})

// 			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 				ctx := context.WithValue(r.Context(), "props", claims)
// 				// Access context values in handlers like this
// 				// props, _ := r.Context().Value("props").(jwt.MapClaims)
// 				next.ServeHTTP(w, r.WithContext(ctx))
// 			} else {
// 				fmt.Println(err)
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("Unauthorized"))
// 			}
// 		}
// 	})
// }
