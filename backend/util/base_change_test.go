package util

import (
	"fmt"
	"testing"
)

func TestTenToThirtySix(t *testing.T) {
	// expected := "0"
	// got := TenToThirtysix(0)
	// if got != expected {
	// 	t.Error("Expected : ", expected, " but got: ", got)
	// }

	// expected = "5"
	// got = TenToThirtysix(5)
	// if got != expected {
	// 	t.Error("Expected : ", expected, " but got: ", got)
	// }

	// expected = "9"
	// got = TenToThirtysix(9)
	// if got != expected {
	// 	t.Error("Expected : ", expected, " but got: ", got)
	// }

	expected := "a"
	got := TenToThirtysix(10)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	fmt.Println("test '10:36'")
	expected = "10"
	got = TenToThirtysix(36)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "11"
	got = TenToThirtysix(37)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "100"
	got = TenToThirtysix(1296)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}
