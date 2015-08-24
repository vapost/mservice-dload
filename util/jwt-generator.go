package jwtutil

import(
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// JWT
type JwtGenerator struct {
	Key string
	Consumer string
	TTL string
}

func (j *JwtGenerator) GetToken() string {

	mySigningKey := []byte(j.Key)
	
	TTL, _ := strconv.Atoi(j.TTL)

	consumer, _ := strconv.Atoi(j.Consumer)

	now := time.Now()

 	// Create the token
    token := jwt.New(jwt.SigningMethodHS256)
    // Set some claims
    token.Claims["consumer"] = consumer
    token.Claims["exp"] = now.Add(time.Second * time.Duration(TTL)).Unix()
    token.Claims["iss"] = now.Unix()
    token.Claims["nbf"] = now.Unix()
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
		log.Fatal(err)
		return ""
	}

	
	return tokenString
}