package api

import (
	"github.com/all-skeleton/gin-skeleton/app/library"
	"github.com/all-skeleton/gin-skeleton/app/servers"
	config2 "github.com/all-skeleton/gin-skeleton/config"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"strconv"
)

// 微信小程序登录
func WxMiNiLogin(c *gin.Context) {
	redisOpts := &cache.RedisOpts{
		Host:        config2.Redis.Host,
		Database:    config2.Redis.Database,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}

	redisCache := cache.NewRedis(c, redisOpts)

	cfg := &config.Config{
		AppID:     config2.Wechat.AppID,
		AppSecret: config2.Wechat.AppSecret,
	}

	wc := wechat.NewWechat()
	wc.SetCache(redisCache)
	mini := wc.GetMiniProgram(cfg)
	wxUser, err := mini.GetAuth().Code2Session(c.Query("code"))
	if err != nil || wxUser.ErrCode != 0 {
		ResponseError(c, "登录失败", "")
		return
	}

	maps := make(map[string]interface{})
	maps["unionid"] = wxUser.UnionID
	user := servers.UserInfo(maps, "*")
	if user.ID <= 0 {
		fromUid, _ := strconv.Atoi(c.DefaultQuery("from_uid", "1"))
		uid, err := servers.UserRegister(wxUser, fromUid)
		if err != nil {
			ResponseError(c, "注册失败", "")
			return
		}
		user.ID = uid
	}

	if user.DeletedAt.Valid == true {
		ResponseError(c, "信息错误", "")
		return
	}

	token, err := library.GenerateApiToken(library.ApiClaims{
		Uid: user.ID,
	})
	if err != nil {
		ResponseError(c, "授权失败", "")
		return
	}

	data := make(map[string]interface{})
	data["uid"] = user.ID
	data["token"] = token
	ResponseSuccess(c, data)
	return
}
