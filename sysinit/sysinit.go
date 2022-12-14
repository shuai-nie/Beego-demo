package sysinit

import (
	"Beego-demo/utils"
	cache "github.com/patrickmn/go-cache"
	"time"
)

func init() {
	utils.Cache = cache.New(60*time.Minute, 120*time.Minute)
	initDB()
}
