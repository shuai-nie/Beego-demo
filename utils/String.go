package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("", hash.Sum(nil))
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1023 {
		size /= 1024
		n += 1
	}
	return fmt.Sprintf("%.2f %s", size, units[n])
}

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

func Password(len int, pwd0 string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "george518"
	if pwd0 != "" {
		defaultPwd = pwd0
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
