package passwords

import (
	"strings"
	"testing"
)

func TestGenerateStdPassword(t *testing.T) {
	pwd := GenerateStdPassword()
	hasLen12 := len(pwd) == 12
	has1UpperChar := strings.ContainsAny(pwd, upperCharSet)
	has1SpecialChar := strings.ContainsAny(pwd, specialCharSet)
	has1NumberChar := strings.ContainsAny(pwd, numberSet)
	if !(hasLen12 && has1SpecialChar && has1NumberChar && has1UpperChar) {
		t.Fatal("generate std password failed")
	}
}
