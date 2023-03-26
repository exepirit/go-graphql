package namegen

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"
)

func NewName() string {
	var seed [6]byte
	_, _ = rand.Reader.Read(seed[:])
	return NewSeededName(seed)
}

func NewSeededName(seed [6]byte) string {
	adjectiveIndex := int(binary.LittleEndian.Uint16(seed[:2])) % len(adjectives)
	adjective := adjectives[adjectiveIndex]

	animals := birds
	animalIndex := int(binary.LittleEndian.Uint16(seed[3:6])) % len(animals)
	animal := animals[animalIndex]

	return fmt.Sprintf("%s-%s", adjective, strings.ToLower(animal))
}
