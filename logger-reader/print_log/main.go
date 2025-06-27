package main

import (
	"fmt"
	"time"
)

func main() {
	log := []string{
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
		`{"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}`,
		`{"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}`,
		`{"ts": 1718091020.789, "level": "debug", "message": "Internal check"}`,
	}

	for _, entry := range log {
		time.Sleep(1 * time.Second)
		fmt.Println(entry)
	}

	for {
	}
}
