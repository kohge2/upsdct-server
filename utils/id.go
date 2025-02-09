package utils

import (
	"log"
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	uid, err := ulid.New(ms, entropy)
	if err != nil {
		log.Fatal(err)
	}
	return uid.String()
}
