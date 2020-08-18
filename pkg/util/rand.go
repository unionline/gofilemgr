/*
@Time : 2020/4/14 23:51
@Author : Justin
@Description :
@File : rand
@Software: GoLand
*/
package util

import (
	"fmt"
	"github.com/rs/xid"
	"math/rand"
	"strconv"
	"time"
)

func GetSalt() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int63n(1000000)
	num_str := strconv.FormatInt(num, 10)
	return num_str
}

// 获取一个XID
func GetXID() (id string) {
	guid := xid.New()
	id = guid.String()

	return
}

func GetRandSalt() string {

	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

}
