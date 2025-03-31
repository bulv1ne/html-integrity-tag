package internal

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
