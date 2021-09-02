package loginHelper

import (
	"albumservice/albumtool"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/model"
	"encoding/json"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func SaveLoginToken(passwordList map[string]string, ip string) string {
	loginToken := &model.LoginToken{}
	loginToken.Date = utils.Now().ToString()
	loginToken.IP = ip
	loginToken.PasswordList = passwordList
	loginToken.Uuid = uuid.NewV4().String()
	token, err := utils.SerializerToJson(loginToken)
	if err != nil {
		return ""
	}

	redisTool.SetString(albumtool.Login_Token_Key, token)
	return loginToken.Uuid
}

func ValidateLoginStatus(token string) bool {

	if token == "" {
		return false
	}

	loginToken := &model.LoginToken{}
	tokenJson := redisTool.GetString(albumtool.Login_Token_Key)
	err := json.Unmarshal([]byte(tokenJson), loginToken)
	if err != nil {
		return false
	}
	if strings.EqualFold(token, loginToken.Uuid) {
		return true
	}
	return false
}
