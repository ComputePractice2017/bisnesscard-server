package model

import (
	"math/rand"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
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
	ID    string `json: "id", gorethink:"id"`
	Login string `json:"login", gorethink:"Login"`
	Pass  string `json:"pass", gorethink:"Pass"`
	Token string `json:"token", gorethink:"Token"`
}

func Login(login string, pass string) (string, error) {
	res, err := r.DB("business_card").Table("user").Filter(func(term r.Term) r.Term {
		return r.And(term.Field("Login").Eq(login), term.Field("Pass").Eq(pass))
	}).Run(session)
	if err != nil {
		return "", err
	}

	var token User
	err = res.One(&token)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

func RegisterUser(login string, pass string) (bool, error) {

	res, err := r.UUID().Run(session)
	if err != nil {
		return false, err
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return false, err
	}
	var new_user User
	new_user.ID = UUID
	new_user.Login = login
	new_user.Pass = pass
	new_user.Token = generateToken(16)

	_, err = r.DB("business_card").Table("user").Insert(new_user).Run(session)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ValidToken(token string) (string, error) {

	res, err := r.DB("business_card").Table("user").Filter(func(term r.Term) r.Term {
		return term.Field("Token").Eq(token)
	}).Run(session)
	if err != nil {
		return "", err
	}

	var id User
	err = res.One(&id)
	if err != nil {
		return "", err
	}
	return id.ID, nil

}

type User_info struct {
	id        string `gorethink:"id"`
	User_ID   string `json:"id", gorethink:"User_ID"`
	Name      string `json:"name", gorethink:"name"`
	Desc      string `json:"desc", gorethink:"desc"`
	Link      string `json:"link", gorethink:"link"`
	Mail      string `json:"mail", gorethink:"mail"`
	Site      string `json:"site", gorethink:"site"`
	Messenger string `json:"messenger", gorethink:"messenger"`
	Telegram  string `json:"telegaram", gorethink:"telegram"`
	VK        string `json:"vk", gorethink:"vk"`
	Whatsapp  string `json:"whatsapp", gorethink:"whatsapp"`
	Viber     string `json:"viber", gorethink:"viber"`
	Skype     string `json:"skype", gorethink:"skype"`
}

func GetInfoById(id string) (User_info, error) {
	res, err := r.DB("business_card").Table("user_info").Filter(func(term r.Term) r.Term {
		return term.Field("User_ID").Eq(id)
	}).Run(session)
	if err != nil {
		return User_info{}, err
	}

	var info User_info
	err = res.One(&info)
	if err != nil {
		return User_info{}, err
	}
	return info, nil
}

func GetInfoByLink(link string) (User_info, error) {
	res, err := r.DB("business_card").Table("user_info").Filter(func(term r.Term) r.Term {
		return term.Field("Link").Eq(link)
	}).Run(session)
	if err != nil {
		return User_info{}, err
	}

	var info User_info
	err = res.One(&info)
	if err != nil {
		return User_info{}, err
	}
	return info, nil
}

func CreateCard(user_id string) error {
	var new_card User_info
	new_card.User_ID = user_id

	_, err := r.DB("business_card").Table("user_info").Insert(new_card).RunWrite(session)
	if err != nil {
		return err
	}

	return err
}

func Update(id string, newInfo User_info) error {
	res, err := r.DB("business_card").Table("user_info").Filter(func(term r.Term) r.Term {
		return term.Field("User_ID").Eq(id)
	}).Run(session)

	var info User_info
	err = res.One(&info)

	if err != nil {
		return err
	}
	if newInfo.Name != "" {
		info.Name = newInfo.Name
	}
	if newInfo.Desc != "" {
		info.Desc = newInfo.Desc
	}
	if newInfo.Link != "" {
		info.Link = newInfo.Link
	}
	if newInfo.Mail != "" {
		info.Mail = newInfo.Mail
	}
	if newInfo.Messenger != "" {
		info.Messenger = newInfo.Messenger
	}
	if newInfo.Name != "" {
		info.Name = newInfo.Name
	}
	if newInfo.Site != "" {
		info.Site = newInfo.Site
	}
	if newInfo.Skype != "" {
		info.Skype = newInfo.Skype
	}
	if newInfo.Telegram != "" {
		info.Telegram = newInfo.Telegram
	}
	if newInfo.VK != "" {
		info.VK = newInfo.VK
	}
	if newInfo.Viber != "" {
		info.Viber = newInfo.Viber
	}
	if newInfo.Whatsapp != "" {
		info.Whatsapp = newInfo.Whatsapp
	}

	_, err = r.DB("business_card").Table("user_info").Replace(info).RunWrite(session)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteUserInfo(id string) error {
	_, err := r.DB("business_card").Table("user_info").Filter(func(term r.Term) r.Term {
		return term.Field("id").Eq(id)
	}).Delete().RunWrite(session)
	return err
}
