package utils

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var emailValidateTemplate = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsAccount check input flag.
// Return true if flag is account.
// Return false if flag is a domain name.
func IsAccount(flag string) bool {
	if emailValidateTemplate.MatchString(flag) {
		return true
	}
	return false
}

// SplitAccount on login and and domain name.
func SplitAccount(account string) ([]string, error) {
	account = strings.ToLower(account)

	// check email regexp
	if emailValidateTemplate.MatchString(account) {
		return strings.Split(account, "@"), nil
	}
	return nil, errors.New("error: invalid email")
}

// GeneratePassword for email account.
// Accept int - length password.
// Return password string.
func GeneratePassword(n int) string {
	// allowed symbol
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`!@#$%^&*()-_=+,.?/:;{}[]<>")

	rand.Seed(time.Now().UnixNano())

	result := make([]rune, n)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// RandomInt generate and return random string from ints.
// Accept int - length password.
func RandomInt(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intArray := make([]int, n)

	for i := 0; i < len(intArray); i++ {
		intArray[i] = r.Intn(10)
	}

	s := strings.Trim(strings.Replace(fmt.Sprint(intArray), " ", "", -1), "[]")
	return string(s)
}

// ReadStdIn helper for read string.
func ReadStdIn(stringToPrint string) string {
	fmt.Print(stringToPrint)

	scan := bufio.NewReader(os.Stdin)
	str, err := scan.ReadString('\n')
	if err != nil {
		fmt.Printf("read string error: %v", err)
	}
	// cut '\n' from string
	result := strings.Trim(str, "\n")

	return result
}

// ErrorCheck helper for rendered results.
// Checks status in response and error.
func ErrorCheck(status, message string, err error) error {
	if err != nil {
		return err
	} else if status != "ok" {
		return errors.New(message)
	}
	return nil
}
