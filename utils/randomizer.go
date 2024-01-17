package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func RandomOrderID() string {
	n := 10
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomUUIDString() string {
	uuidp := uuid.New()
	return uuidp.String()
}

var (
	mu           sync.Mutex
	lastUnixTime int64
	sequence     int
)

func GenerateID() int64 {
	mu.Lock()
	defer mu.Unlock()

	currentTime := time.Now().UnixNano()

	// If the current time is the same as the last time, increment the sequence number
	if currentTime <= lastUnixTime {
		sequence++
	} else {
		sequence = 0
		lastUnixTime = currentTime
	}

	// Shift the timestamp to the left by 10 bits to make room for the sequence number
	id := (currentTime << 10) + int64(sequence)

	return id
}

func GenerateStringID() string {
	mu.Lock()
	defer mu.Unlock()

	currentTime := time.Now().UnixNano()

	// If the current time is the same as the last time, increment the sequence number
	if currentTime <= lastUnixTime {
		sequence++
	} else {
		sequence = 0
		lastUnixTime = currentTime
	}

	// Shift the timestamp to the left by 10 bits to make room for the sequence number
	id := (currentTime << 10) + int64(sequence)

	return strconv.FormatInt(id, 10)
}
