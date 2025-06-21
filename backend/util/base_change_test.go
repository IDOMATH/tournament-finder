package util

import (
	"testing"
)

func TestTenToThirtySix(t *testing.T) {
	expected := "0"
	got := TenToThirtysix(0)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "5"
	got = TenToThirtysix(5)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "9"
	got = TenToThirtysix(9)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "a"
	got = TenToThirtysix(10)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "2s"
	got = TenToThirtysix(100)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

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

func TestSr(t *testing.T) {
	expected := "a"
	got := sr(10)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = "z"
	got = sr(35)
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}

func TestThirtysixToTen(t *testing.T) {
	var expected int64 = 10
	got := ThirtysixToTen("a")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 35
	got = ThirtysixToTen("z")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 8
	got = ThirtysixToTen("8")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 36
	got = ThirtysixToTen("10")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1295
	got = ThirtysixToTen("zz")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 109
	got = ThirtysixToTen("31")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1296
	got = ThirtysixToTen("100")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}

	expected = 1260
	got = ThirtysixToTen("z0")
	if got != expected {
		t.Error("Expected : ", expected, " but got: ", got)
	}
}
