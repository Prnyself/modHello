package modHello

import "fmt"

func Hi(name string) string {
	if name == "" {
		return "Hello world!"
	}
	return fmt.Sprintf("Hello %s", name)
}
