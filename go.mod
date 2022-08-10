module github.com/all-skeleton/gin-skeleton

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-redis/redis/v8 v8.11.6-0.20220405070650-99c79f7041fc
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/putyy/gokv v0.0.0-20220303073116-12b2d275431c
	github.com/qiniu/go-sdk/v7 v7.13.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/silenceper/wechat/v2 v2.1.3
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.0.0-20220622184535-263ec571b305 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220622161953-175b2fd9d664 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)

replace (
	github.com/all-skeleton/gin-skeleton/app => ./app
	github.com/all-skeleton/gin-skeleton/config => ./config
	github.com/all-skeleton/gin-skeleton/router => ./router
)
