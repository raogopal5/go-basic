package details

import "os"

func getHostName() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}
