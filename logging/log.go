package logging

import (
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Debug(msg string, v ...any) {
	log.SetPrefix("[DEBUG]")
	if len(v) < 1 {
		log.Printf(msg + "\n")
	}
	log.Printf(msg+"\n", v)
}

func Info(msg string, v ...any) {
	log.SetPrefix("[INFO]")
	if len(v) < 1 {
		log.Printf(msg+"\n", v)
	}
	log.Printf(msg+"\n", v)
}

func Warn(msg string, v ...any) {
	log.SetPrefix("[WARN]")
	if len(v) < 1 {
		log.Printf(msg+"\n", v)
	}
	log.Printf(msg+"\n", v)
}

func Error(msg string, v ...any) {
	log.SetPrefix("[ERROR]")
	if len(v) < 1 {
		log.Fatalf(msg+"\n", v)
	}
}
