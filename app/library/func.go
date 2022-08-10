package library

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/all-skeleton/gin-skeleton/config"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 获取图片资源地址
func GetImgUrl(src string) string {
	if src == "" {
		return config.App.DefaultImg
	}
	if strings.HasPrefix(src, "http") {
		return src
	}
	return config.Qiniu.ImageDomain + "/" + src
}

// 获取文件资源地址
func GetFileUrl(src string) string {
	if src == "" {
		return config.App.DefaultImg
	}
	if strings.HasPrefix(src, "http") {
		return src
	}
	return config.Qiniu.MediaDomain + "/" + src
}

func GetSaveUrl(src string) string {
	if src == "" {
		return ""
	}

	p, err := url.Parse(src)
	if err != nil {
		return ""
	}

	return strings.TrimLeft(p.Path, "/")
}

func En100(i int) int {
	return i * 100
}

func De100(i int) int {
	return i / 100
}

// 创建文件名称
func MakeFileName(scene string) (string, error) {
	sceneInfo := GetUploadSceneInfo(scene)
	v, ok := sceneInfo["value"]
	if ok == false {
		return "", errors.New("无效scene配置")
	}
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := time.Now().Format("20060102") + fmt.Sprintf("%x", b)
	return "upload/" + v + "/" + s, nil
}

// 生成订单号
func MakeOrderNo(uid int) (string, error) {
	b := make([]byte, 3)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	suid := fmt.Sprintf("%07s%s", strconv.Itoa(uid), "")[:7]
	s := time.Now().Format("20060102") + fmt.Sprintf("%x", b) + suid
	return s, nil
}
