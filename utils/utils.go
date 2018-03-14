package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// InfoMsg prints an info message
func InfoMsg(message string) { fmt.Println("Info:  " + message) }

// ErrorMsg prints an error message
func ErrorMsg(message string) { fmt.Println("Error: " + message) }

// GenerateSalt generate a 20 chars salt
func GenerateSalt() string {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	salt := []string{}
	for i := 0; i < 20; i++ {
		salt = append(salt, string(alphabet[rand.Intn(len(alphabet))]))
	}
	return strings.Join(salt, "")
}

// GenerateToken generate a md5 string from a password and a salt
func GenerateToken(password string, salt string) string {
	hasher := md5.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Prompt ask you to input details via console
func Prompt(p string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>>>   Please enter your " + p + ": ")
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}
