package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/IDOMATH/CheetahMath/formulas"
)

func IntifyId(token []byte) int {
	id, err := strconv.Atoi(string(token[:]))
	if err != nil {
		fmt.Println("Error intifying token")
		return 0
	}
	return id
}

func MakeToken(id int) string {
	id64 := int64(id)

	t := time.Now().UnixMilli()*formulas.IntPow(10, int64(formulas.GetDigits(id64))) + id64
	return TenToThirtysix(t)

}

func GetUserIdFromToken(token []byte) int64 {
	// token -> base 10
	var baseTen int64 = 12345423452654
	// remove leading unixMilli()
	leftOver := int64(formulas.GetDigits(baseTen) - 13)
	return baseTen % formulas.IntPow(10, leftOver)

	// should be left with Id
}
