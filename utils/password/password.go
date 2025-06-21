package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(hashed)
}

func CheckPassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
