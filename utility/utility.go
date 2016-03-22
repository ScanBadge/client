package utility

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net"
)

// ReadData reads input from the specified file path.
func ReadData(filename string) ([]byte, error) {
	if filename == "" {
		return nil, fmt.Errorf("no filename specified")
	}

	return ioutil.ReadFile(filename)
}

// HashPassword hashes a password with the bcrypt algorithm using a cost of 10.
func HashPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

// IsPortOpen verifies whether or not the provided port on the provided host name is open or closed by dialing.
// Returns true on success; false on failure.
func IsPortOpen(proto, host string, port int) bool {
	conn, err := net.Dial(proto, fmt.Sprintf("%s:%v", host, port))

	if err != nil {
		log.Println("Connection error:", err)
		return false
	}

	defer conn.Close()
	return true
}
