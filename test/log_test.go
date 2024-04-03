package test

import (
	"log"
	"testing"
)

func TestLogPrefix(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("TEST\n")
}
