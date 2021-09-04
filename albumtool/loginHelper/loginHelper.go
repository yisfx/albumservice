package loginHelper

import (
	"albumservice/albumtool/constfield"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/model"
	"encoding/json"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func SaveLoginToken(passwordList map[string]string, ip string) string {
	loginToken := &model.LoginToken{}
	loginToken.Date = utils.DateTime.Now().ToString()
	loginToken.IP = ip
	loginToken.PasswordList = passwordList
	loginToken.Uuid = uuid.NewV4().String()

	redisTool.SetString(constfield.Login_Token_Key, utils.StringTool.SerializerToJson(loginToken))
	return loginToken.Uuid
}

func ValidateLoginStatus(token string) bool {

	if token == "" {
		return false
	}

	loginToken := &model.LoginToken{}
	tokenJson := redisTool.GetString(constfield.Login_Token_Key)
	err := json.Unmarshal([]byte(tokenJson), loginToken)
	if err != nil {
		return false
	}
	if strings.EqualFold(token, loginToken.Uuid) {
		return true
	}
	return false
}
