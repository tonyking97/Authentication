package core

import (
	"../models"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/packr"
	//"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Create and manage the token and Authenticate the user
 */

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend = nil

//Initialize the private and public key
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

//generate Encrypted Token
func (backend *JWTAuthenticationBackend) Generate(ff_id, fs_id string) (string, error, int64, int64) {

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	timeNow := time.Now().Unix()
	//expTime := time.Now().Add(time.Hour * time.Duration(72)).Unix();
	expTime := time.Now().Add(time.Hour).Unix()
	token.Claims = jwt.MapClaims{
		"exp":   expTime,
		"iat":   timeNow,
		"ff_id": ff_id,
		"fs_id": fs_id,
	}
	tokenString, err := token.SignedString([]byte("thisissecretkeynum001707"))
	if err != nil {
		log.Println(err)
		return "", err, 0, 0
	}

	hash := sha512.New()
	encToken, err := rsa.EncryptOAEP(hash, rand.Reader, authBackendInstance.PublicKey, []byte(tokenString), nil)
	if err != nil {
		log.Println(err)
		return "", err, 0, 0
	}
	encTokenString := base64.StdEncoding.EncodeToString(encToken)

	return encTokenString, nil, timeNow, expTime
}

//Extract Details from the token
func (backend *JWTAuthenticationBackend) ExtractToken(token string) (string, string, error) {
	encToken, _ := base64.StdEncoding.DecodeString(token)
	hash := sha512.New()
	decToken, err := rsa.DecryptOAEP(hash, rand.Reader, authBackendInstance.privateKey, encToken, nil)
	if err != nil {
		return "", "", errors.New("invalid token")
	}

	ExtToken, err := jwt.Parse(string(decToken), func(token *jwt.Token) (interface{}, error) {
		return []byte("thisissecretkeynum001707"), nil
	})

	if err != nil {
		return "", "", errors.New(err.Error())
	}

	if claims, ok := ExtToken.Claims.(jwt.MapClaims); ok && ExtToken.Valid {
		ff_id := claims["ff_id"].(string)
		fs_id := claims["fs_id"].(string)

		return ff_id, fs_id, nil
	} else {
		return "", "", errors.New("Token is expired")
	}
}

//Authenticate the user
func (backend *JWTAuthenticationBackend) Authenticate(user *models.User) (bool, int, string) {

	usersCollection := UsersCollectionInit()
	var userDetails *models.SignUpDetails

	filter := bson.D{{"username", user.Username}}

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&userDetails)
	if err != nil {
		log.Println(err)
	}
	if userDetails == nil {
		return false, 1, ""
	}

	if user.Username == userDetails.Username && bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password)) == nil {
		return true, 0, userDetails.Id
	} else {
		return false, 2, ""
	}
}

//Retrieve the Private Key
func getPrivateKey() *rsa.PrivateKey {

	privateKeyBox := packr.NewBox("../config/keys")
	//privateKeyBox := packr.New("privateKey","../config/keys")
	pembytes := privateKeyBox.Bytes("private_key")

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

//Retrieve the Public Key
func getPublicKey() *rsa.PublicKey {

	publicKeyBox := packr.NewBox("../config/keys")
	//publicKeyBox := packr.New("publicKey","../config/keys")
	pembytes := publicKeyBox.Bytes("public_key.pub")

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
