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

	t := int(time.Now().UnixMilli())*formulas.IntPow(10, formulas.GetDigits(id)) + id
	return TenToThirtysix(t)

}

func GetUserIdFromToken(token string) int {
	// token -> base 10
	baseTen := 12345423452654
	// remove leading unixMilli()
	leftOver := formulas.GetDigits(baseTen) - 13
	return baseTen % formulas.IntPow(10, leftOver)

	// should be left with Id
}
