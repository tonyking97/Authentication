package core

import (
	"../models"
	"context"
	"encoding/json"
	"github.com/mssola/user_agent"
	"github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Contains logical methods for the authentication service
 */

func Login(requestUser *models.User,userAgent string) (int, string) {

	authBackend := InitJWTAuthenticationBackend()
	valid, err, ff_id := authBackend.Authenticate(requestUser)
	fs_id, _ := uuid.NewV4()

	if valid {
		token, err, timeNow, expTime := authBackend.Generate(ff_id, fs_id.String())
		if err != nil {
			//return http.StatusInternalServerError, []byte("")
			return http.StatusInternalServerError, "Internal Server Error Occured. Please try again later."
		} else {
			ua := user_agent.New(userAgent)
			browserName, browserVersion := ua.Browser()
			engineName, engineVersion := ua.Engine()
			userDetails := models.SessionDetails{
				ff_id,
				fs_id.String(),
				token,
				"active",
				timeNow,
				expTime,
				requestUser.Ip_details,
				//r.Referer(),
				"",
				ua.Bot(),
				ua.Mobile(),
				browserName,
				browserVersion,
				ua.Mozilla(),
				engineName,
				engineVersion,
				requestUser.Localization,
				ua.OSInfo().FullName,
				ua.OSInfo().Name,
				ua.OSInfo().Version,
				ua.Platform(),
				"",
				"",
				userAgent,
			}
			sessionCollection := SessionCollectionInit()
			insertResult, err := sessionCollection.InsertOne(context.TODO(), userDetails)
			if err != nil {
				log.Println(err)
				//response, _ := json.Marshal(models.ErrorMessage{"Internal Server Error Occured. Please try again later."})
				response :="Internal Server Error Occured. Please try again later."
				return http.StatusInternalServerError, response
			}
			log.Println("New session created for user "+ff_id+" with id : ", insertResult.InsertedID)

			//response, _ := json.Marshal(models.TokenAuthentication{token})
			return http.StatusOK, token
		}
	} else if !valid && err == 1 {
		//response, _ := json.Marshal(models.ErrorMessage{"User not found"})
		response := "User not found"
		return http.StatusUnauthorized, response
	} else if !valid && err == 2 {
		//response, _ := json.Marshal(models.ErrorMessage{"Invalid password"})
		response := "Invalid password"
		return http.StatusUnauthorized, response
	} else {
		//response, _ := json.Marshal(models.ErrorMessage{"Error Occured. Please contact administrator."})
		response := "Error Occured. Please contact administrator."
		return http.StatusInternalServerError, response
	}
}

//Returns true and avatar name(First letters of firstname and lastname) while username exist
func CheckUsername(requestUser *models.Username) (bool, string, string) {
	var result models.CheckUsernameAvatar
	usersCollection := UsersCollectionInit()
	_ = usersCollection.FindOne(context.TODO(), requestUser).Decode(&result)

	if result.Username != "" {
		return true, result.Username, (string(result.Firstname[0]) + string(result.Lastname[0]))
	} else {
		return false, "", ""
	}
}

//Returns true and avatar name(First letters of firstname and lastname) while username exist
func GetUserDetails(requestUser *models.Ff_id_user) (bool, models.UserNameDetails) {
	var result models.UserNameDetails

	usersCollection := UsersCollectionInit()
	_ = usersCollection.FindOne(context.TODO(), requestUser).Decode(&result)

	if result.Username != "" {
		return true, result
	} else {
		return false, result
	}
}

//Returns true while email exist
func CheckEmail(requestUser *models.Email) bool {
	var result models.Email
	usersCollection := UsersCollectionInit()
	_ = usersCollection.FindOne(context.TODO(), requestUser).Decode(&result)

	if result.Email != "" {
		return true
	} else {
		return false
	}
}

func GetActiveSessionList(requestUser *models.Ff_id) []models.SessionDetails {
	var results []models.SessionDetails
	sessionCollection := SessionCollectionInit()
	filter := bson.D{{"ff_id", requestUser.Ff_id}, {"status", "active"}}
	cur, err := sessionCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.SessionDetails
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())
	return results
}

func GetInActiveSessionList(requestUser *models.Ff_id) []models.SessionDetails {
	var results []models.SessionDetails
	sessionCollection := SessionCollectionInit()
	filter := bson.D{{"ff_id", requestUser.Ff_id}, {"status", "inactive"}}
	cur, err := sessionCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.SessionDetails
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())

	return results
}

func GetSessionList(ff_id string) []models.SessionDetails_limited {
	var results []models.SessionDetails_limited
	sessionCollection := SessionCollectionInit()

	option := options.Find()
	option.SetSort(bson.D{{"logged_Time", -1}})
	filter := bson.D{{"ff_id", ff_id}}
	cur, err := sessionCollection.Find(context.TODO(), filter, option)
	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.SessionDetails_limited
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elem.LoggedTimeUTC = time.Unix(elem.LoggedTime, 0).String()
		elem.ExpiryTimeUTC = time.Unix(elem.ExpiryTime, 0).String()
		elem.Token = (elem.Token[0:5] + "...." + elem.Token[len(elem.Token)-7:len(elem.Token)-1])
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.TODO())
	return results
}

func GetCurrentSessionList(Fs_id string) models.CurrentSessionDetails {
	var result models.CurrentSessionDetails
	sessionCollection := SessionCollectionInit()
	filter := bson.D{{"fs_id", Fs_id}, {"status", "active"}}
	err := sessionCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func GetActiveSessionsCount(Ff_id string) int64 {
	sessionCollection := SessionCollectionInit()
	filter := bson.D{{"ff_id", Ff_id}, {"status", "active"}}
	count, err := sessionCollection.CountDocuments(context.TODO(), filter)

	if err != nil {
		log.Println(err)
	}

	return count
}

func GetActiveTokensCount(Ff_id string) int64 {
	sessionCollection := SessionCollectionInit()
	filter := bson.D{{"ff_id", Ff_id}, {"status", "active"}}
	count, err := sessionCollection.CountDocuments(context.TODO(), filter)

	if err != nil {
		log.Println(err)
	}

	return count
}

func Signup(requestUser *models.SignUpDetails) bool {
	u4, err := uuid.NewV4()
	requestUser.Id = u4.String()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(requestUser.Password), 10)
	requestUser.Password = string(hashedPassword)
	refTime := time.Now().UTC().Unix()
	requestUser.CreatedTime = refTime
	requestUser.UpdatedTime = refTime

	usersCollection := UsersCollectionInit()
	insertResult, err := usersCollection.InsertOne(context.TODO(), requestUser)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("Created a new user with id: ", insertResult.InsertedID)
	return true
}

func RefreshToken(requestUser *models.TokenAuthentication) []byte {
	authBackend := InitJWTAuthenticationBackend()

	ff_id, fs_id, err := CheckToken(requestUser.Token)

	if err != nil {
		response, _ := json.Marshal(models.ErrorMessage{err.Error()})
		return response
	}

	sessionCollection := SessionCollectionInit()

	token, err, _, ExpTime := authBackend.Generate(ff_id, fs_id)

	if err != nil {
		log.Println(err)
		response, _ := json.Marshal(models.ErrorMessage{"Error in genrating Token"})
		return response
	}

	filter := bson.D{{"fs_id", fs_id}, {"ff_id", ff_id}}
	update := bson.D{
		{"$set", bson.D{
			{"exipry_Time", ExpTime},
		}},
		{"$set", bson.D{
			{"token", token},
		}},
	}

	updateResult, err := sessionCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		response, _ := json.Marshal(models.ErrorMessage{"Internel server error"})
		return response
	} else if updateResult.ModifiedCount == 0 {
		log.Println("No data found")
		response, _ := json.Marshal(models.ErrorMessage{Message: "token invalid"})
		return response
	} else {
		log.Println("Token refreshed for user -> ", ff_id, updateResult.ModifiedCount)
	}

	response, err := json.Marshal(models.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func CheckToken(token string) (string, string, error) {
	authBackend := InitJWTAuthenticationBackend()

	ff_id, fs_id, err := authBackend.ExtractToken(token)

	if err != nil {
		log.Println("------>>>>>", err)
		return "", "", err
	}

	sessionCollection := SessionCollectionInit()

	filter := bson.D{{"ff_id", ff_id}, {"fs_id", fs_id}, {"token", token}}

	var result models.SessionDetails
	err = sessionCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return "", "", errors.New("Token is expired")
	}

	if result.Status == "active" && (result.ExpiryTime-time.Now().Unix()) > 0 {
		return ff_id, fs_id, nil
	}

	return "", "", errors.New("invalid token")
}

func Logout(ff_id, fs_id string) error {

	sessionCollection := SessionCollectionInit()

	filter := bson.D{{"ff_id", ff_id}, {"fs_id", fs_id}}
	update := bson.D{
		{"$set", bson.D{
			{"expiry_Time", time.Now().Unix()},
		}},
		{"$set", bson.D{
			{"status", "inactive"},
		}},
	}

	updateResult, err := sessionCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	} else if updateResult.ModifiedCount == 0 {
		log.Println("No data found")
		return errors.New("invalid token")
	} else {
		return nil
	}
}
