package aguid

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"strings"
)

// Aguid takes a string and provides a deterministic reproducible UUID hash
// It will return a random UUID if the input string is empty ""
func Aguid(str string) (string, error) {
	if str == "" {
		res, err := uuid.NewV4()
		if err != nil {
			return "", err
		}
		return res.String(), nil
	}

	var replacer = strings.NewReplacer("+", "")
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hash := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
	hash = replacer.Replace(hash)
	hash = string(hash[0:36])
	hash = string(hash[:8]) + string("-") + string(hash[8+1:])
	hash = string(hash[:13]) + string("-") + string(hash[13+1:])
	hash = string(hash[:14]) + string("4") + string(hash[14+1:])
	hash = string(hash[:18]) + string("-") + string(hash[18+1:])
	hash = string(hash[:19]) + string("8") + string(hash[19+1:])
	hash = string(hash[:20]) + string("4") + string(hash[20+1:])
	hash = string(hash[:23]) + string("-") + string(hash[23+1:])

	return hash, nil
}
