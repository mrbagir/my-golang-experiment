package main

// func TestReadLogs(t *testing.T) {
// 	logInput := `
// {"ts": 1718091000.123, "level": "info", "message": "Service started", "metadata": {"port": 8080}}
// {"ts": 1718091010.456, "level": "error", "message": "Database failed", "metadata": {"retry": true}}
// {"ts": 1718091020.789, "level": "debug", "message": "Internal check"}
// Not a JSON log line
// `

// 	reader := strings.NewReader(logInput)
// 	readLogs(reader)

// 	for _, summary := range allSummaries {
// 		fmt.Println(summary)
// 	}

// 	for _, log := range allRawLogs {
// 		fmt.Println(log)
// 	}
// }
