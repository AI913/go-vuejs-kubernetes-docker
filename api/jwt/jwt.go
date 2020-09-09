package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	Users "github.com/dennischiu/govuekuber/api/models/users"
	"github.com/dgrijalva/jwt-go"
	"xorm.io/xorm"
)

var signKey string

func init() {
	signKey = os.Getenv("SECRET")
	if len(signKey) == 0 {
		log.Println("WARNING!! Secret not set for JWT!")
	}
}

// CreateToken - this creates a jwt token from a user id
func CreateToken(id int64) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = id

	token.Claims = claims

	tokenString, err = token.SignedString([]byte(signKey))

	return
}

func ParseToken(val string) (id int64, err error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}
		id = int64(claims["id"].(float64))

		return
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token has expired")
		default:
			log.Println(vErr)
			return 0, errors.New("Error parsing token or no token")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}

}

//GetUserFromToekn - grab user data from the database via the token data

func GetUserFromToken(db *xorm.Engine, tokenVal string) (user Users.User, err error) {
	if len(tokenVal) == 0 {
		err = errors.New("No token present")
		return
	}

	userID, err := ParseToken(tokenVal)
	if err != nil {
		err = errors.New("Token is invalid")
		return
	}

	if userID < 1 {
		err = errors.New("Token is missing required data")
		return
	}

	user.ID = userID
	users, err := Users.Index(db, &user)

	if err != nil || len(users) == 0 {
		err = errors.New("Unable to get user from token")
		return
	}

	user = users[0]
	user.Password = ""

	return
}
