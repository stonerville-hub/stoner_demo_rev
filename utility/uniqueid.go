package utils

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
)

type cryptSrc struct{}

func (s cryptSrc) Seed(seed int64) {}

func (s cryptSrc) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptSrc) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

// Generate a truly random number
func RandomInt64(i int64) int64 {
	var src cryptSrc
	rnd := rand.New(src)
	return rnd.Int63n(i)
}
