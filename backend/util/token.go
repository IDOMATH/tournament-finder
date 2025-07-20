package util

import (
	"time"

	"github.com/IDOMATH/CheetahMath/formulas"
)

// time.Now().UnixMilli() won't add another digit for about 260 years
// so this is very safe because nobody will be using this when that happens.
const unixMilliDigits = 13

// This will work as long as we have less than 1,000,000 users.
// We can celebrate if we ever have to figure out how to work around that
func MakeToken(id int) string {
	id64 := int64(id)

	t := time.Now().UnixMilli()*formulas.IntPow(10, int64(formulas.GetDigits(id64))) + id64
	return TenToThirtysix(t)

}

func GetUserIdFromToken(token string) int {
	// token -> base 10
	var baseTen int64 = ThirtysixToTen(token)
	// remove leading unixMilli()
	leftOver := int64(formulas.GetDigits(baseTen) - unixMilliDigits)
	id := baseTen % formulas.IntPow(10, leftOver)
	return int(id)
}
