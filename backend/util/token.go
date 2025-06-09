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
