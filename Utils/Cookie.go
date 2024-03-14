package Utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GenCookie(uName string, pwdSummary string) string {
	cookie := md5.Sum([]byte(uName + pwdSummary + strconv.FormatInt(time.Now().Unix(), 10)))
	return hex.EncodeToString(cookie[:])
}
