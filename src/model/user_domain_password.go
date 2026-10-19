package model

import (
	"crypto/sha256"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassword() {
	hash := sha256.New()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
