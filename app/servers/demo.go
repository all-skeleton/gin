package servers

import (
	"github.com/all-skeleton/gin-skeleton/app/models"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

func UserInfo(m map[string]interface{}, field string) (res models.User) {
	models.Db().Unscoped().Where(m).Select(field).Find(&res)
	return
}

func UserRegister(wxUser auth.ResCode2Session, fromUid int) (int, error) {
	// todo
	return 666, nil
}
