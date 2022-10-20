package util

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
)

const (
	pwdPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\\$%\\^&\\*]).{8,}$"
	emlPattern = "^(([^<>()\\[\\]\\\\.,;:\\s@\"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$"
)

func Sha256(text string) string {
	encode := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", encode)
}

func IsEmail(email string) bool {
	matched, _ := regexp.MatchString(emlPattern, email)
	return matched
}

func IsPassword(password string) bool {
	matched, _ := regexp.MatchString(pwdPattern, password)
	return matched
}

func StripAccents(name string) string {
	source := []string{"À", "Á", "Â", "Ã", "È", "É",
		"Ê", "Ì", "Í", "Ò", "Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "à", "á", "â",
		"ã", "è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ", "ù", "ú", "ý",
		"Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ", "ũ", "Ơ", "ơ", "Ư", "ư", "Ạ",
		"ạ", "Ả", "ả", "Ấ", "ấ", "Ầ", "ầ", "Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ",
		"Ắ", "ắ", "Ằ", "ằ", "Ẳ", "ẳ", "Ẵ", "ẵ", "Ặ", "ặ", "Ẹ", "ẹ", "Ẻ",
		"ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề", "Ể", "ể", "Ễ", "ễ", "Ệ", "ệ",
		"Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ", "Ỏ", "ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ",
		"ổ", "Ỗ", "ỗ", "Ộ", "ộ", "Ớ", "ớ", "Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ",
		"Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ", "ứ", "Ừ", "ừ", "Ử", "ử", "Ữ",
		"ữ", "Ự", "ự", "ý", "ỳ", "ỷ", "ỹ", "ỵ", "Ý", "Ỳ", "Ỷ", "Ỹ", "Ỵ"}

	dist := []string{"A", "A", "A", "A", "E",
		"E", "E", "I", "I", "O", "O", "O", "O", "U", "U", "Y", "a", "a",
		"a", "a", "e", "e", "e", "i", "i", "o", "o", "o", "o", "u", "u",
		"y", "A", "a", "D", "d", "I", "i", "U", "u", "O", "o", "U", "u",
		"A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A",
		"a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "E", "e",
		"E", "e", "E", "e", "E", "e", "E", "e", "E", "e", "E", "e", "E",
		"e", "I", "i", "I", "i", "O", "o", "O", "o", "O", "o", "O", "o",
		"O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O",
		"o", "O", "o", "U", "u", "U", "u", "U", "u", "U", "u", "U", "u",
		"U", "u", "U", "u", "y", "y", "y", "y", "y", "Y", "Y", "Y", "Y", "Y"}

	for index, char := range source {
		name = strings.Replace(name, char, dist[index], -1)
	}

	name = strings.Replace(name, "'", " ", 10)

	return strings.ToUpper(name)
}
