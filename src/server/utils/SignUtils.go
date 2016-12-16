package utils

import (
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
)

func CheckSignature(signature, timestamp, nonce string) bool {
	signatureGen := str2sha1(timestamp, nonce)
	if signatureGen != signature {
		return false
	}
	return true
}

func str2sha1(timestamp, nonce string) string {
	token := Token
	arr := []string{token, timestamp, nonce}
	sort.Strings(arr)
	s := sha1.New()
	io.WriteString(s, strings.Join(arr, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}