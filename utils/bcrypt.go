package utils

import "golang.org/x/crypto/bcrypt"

/*
Function to create a hashed password

Params:

	password string

Returns:

	string
	error
*/
func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(b), err
}

/*
Function to compare a clear text password with a hashed password. returns if its a valid password or not

Params:

	cleartextPassword string
	hashedPassword string

Return:

	bool
*/
func CompareHashAndPassword(cleartextPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(cleartextPassword))

	return err == nil
}
