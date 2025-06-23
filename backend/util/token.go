package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/IDOMATH/CheetahMath/formulas"
)

// time.Now().UnixMicro() won't add another digit for about 260 years
// so this is very safe because nobody will be using this when that happens.
const unixMicroDigits = 16

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

	t := time.Now().UnixMicro()*formulas.IntPow(10, int64(formulas.GetDigits(id64))) + id64
	return TenToThirtysix(t)

}

func GetUserIdFromToken(token []byte) int64 {
	// token -> base 10
	var baseTen int64 = ThirtysixToTen(string(token[:]))
	// remove leading unixMilli()
	leftOver := int64(formulas.GetDigits(baseTen) - unixMicroDigits)
	time.Now().UnixMicro()
	return baseTen % formulas.IntPow(10, leftOver)

	// should be left with Id
}
