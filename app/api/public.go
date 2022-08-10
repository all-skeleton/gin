package api

import (
	"fmt"
	"github.com/all-skeleton/gin-skeleton/app/library"
	"github.com/all-skeleton/gin-skeleton/config"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"strings"
)

func Scene(c *gin.Context) {
	ResponseSuccess(c, library.GetSceneAll())
}

func GetQiNiuWebToken(c *gin.Context) {
	scene := c.Query("scene")
	ext := c.Query("ext")
	if len(scene) <= 0 {
		ResponseError(c, "scene 必须", "")
		return
	}

	extL := len(ext)
	if extL <= 0 || extL > 6 {
		ResponseError(c, "ext 错误", "")
		return
	}

	sceneRes := strings.Split(scene, "-")
	var data = make([]interface{}, len(sceneRes))
	for i, v := range sceneRes {
		// 需要覆盖的文件名
		key, err1 := library.MakeFileName(v)
		if err1 != nil {
			ResponseError(c, "error", "")
			return
		}
		key = key + "." + ext
		var bucket, domain string
		switch ext {
		case "png", "jpg", "jpeg":
			bucket = config.Qiniu.ImageBucket
			domain = config.Qiniu.ImageDomain
		case "mp4", "mp3":
			fallthrough
		default:
			bucket = config.Qiniu.MediaBucket
			domain = config.Qiniu.MediaDomain
		}
		putPolicy := storage.PutPolicy{
			Scope:      fmt.Sprintf("%s:%s", bucket, key),
			ReturnBody: "{\"key\": \"" + domain + "/" + key + "\", \"hash\": $(etag)}",
		}
		mac := qbox.NewMac(config.Qiniu.AccessKey, config.Qiniu.SecretKey)
		data[i] = map[string]string{
			"scene":  v,
			"key":    key,
			"token":  putPolicy.UploadToken(mac),
			"domain": domain,
			"up_serve": "https://upload-z2.qiniup.com",
		}
	}
	ResponseSuccess(c, data)
}