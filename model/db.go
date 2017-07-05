package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"time"
	"math/rand"
)

var session *r.Session

func InitSession() error {

	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
var src = rand.NewSource(time.Now().UnixNano())

func generateToken(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
type User struct {
	ID string `json: "id", gorethink:"id"`
	Login string `json:"login", gorethink:"login"`
	Pass string `json:"pass", gorethink:"pass"`
	Token string `json:"token", gorethink:"token"`
}

func Login(login string, pass string) (string, error) {
	res, err := r.DB("business_card").Table("user").Field("token").Filter(map[string]interface{} {
		"login" : login,
		"pass": pass,
	}).Run(session)
	if err != nil {
		return nil, err
	}

	var token string
	err = res.One(&token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func RegisterUser(login string, pass string) bool {

	res, err := r.UUID().Run(session)
	if err != nil {
		return false
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return false
	}
	var new_user User
	new_user.ID = UUID
	new_user.Login = login
	new_user.Pass = pass
	new_user.Token = generateToken(16);

	_, err = r.DB("business_card").Table("user").Insert(new_user).Run(session)
	if err != nil {
		return false
	}

	return true
}

func ValidToken(token string) (string, error) {

	res, err := r.DB("business_card").Table("user").Field("id").Filter(map[string]interface{} {
		"token" : token,
	}).Run(session)
	if err != nil {
		return nil, err
	}

	var id string
	err = res.One(&id)
	if err != nil {
		return nil, err
	}
	return id, nil

}

type User_info struct{

	ID string `json:"id", gorethink:"id"`
	Name string `json:"name", gorethink:"name"`
	Desc string `json:"desc", gorethink:"desc"`
	Link string `json:"link", gorethink:"link"`
	Mail string `json:"mail", gorethink:"mail"`
	Site string `json:"site", gorethink:"site"`
	Messenger string `json:"messenger", gorethink:"messenger"`
	Telegaram string `json:"telegaram", gorethink:"telegeram"`
	VK string `json:"vk", gorethink:"vk"`
	Whatsapp string `json:"whatsapp", gorethink:"whatsapp"`
	Viber string `json:"viber", gorethink:"viber"`
	Skype string `json:"skype", gorethink:"skype"`

}

func GetInfoById(id string) (User_info, error) {
	res, err := r.DB("business_card").Table("user_info").Filter(map[string]interface{} {
		"id": id,
	}).Run(session)
	if err != nil {
		return nil, err
	}

	var info User_info
	err = res.One(&info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GetInfoByLink(link string) (User_info, error) {
	res, err := r.DB("business_card").Table("user_info").Filter(map[string]interface{} {
		"link": link,
	}).Run(session)
	if err != nil {
		return nil, err
	}

	var info User_info
	err = res.One(&info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func CreateCard(user_id string) error  {
	var new_card User_info
	new_card.ID = user_id

	_, err := r.DB("business_card").Table("user_info").Insert(new_card).Run(session)
	if err != nil {
		return err
	}

	return err
}

func Update(id string, newInfo map[string]string) error {
	_, err := r.DB("business_card").Table("user_info").Filter(map[string]interface{} {
		"id": id,
	}).Update(newInfo).RunWrite(session)
	return err
}

func DeleteUserInfo(id int) error {
	_, err := r.DB("business_card").Table("user_info").Filter(map[string]interface{}{
		"id": id,
	}).Delete().RunWrite(session)
	return err
}

