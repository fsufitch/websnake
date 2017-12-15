package log

import (
	"log"
	"os"
)

// Info is a logger that prints out non-error monitoring information to stdout
var Info = wrappedLogger{log.New(os.Stdout, "[INFO]", log.Ltime|log.LUTC|log.Ldate), true}

// Error is a logger that prints out error information to stderr
var Error = wrappedLogger{log.New(os.Stderr, "[ERROR]", log.Ltime|log.LUTC|log.Ldate), true}

// Debug is a logger that prints out non-essential, detailed debug data to stdout
// Disabled by default
var Debug = wrappedLogger{log.New(os.Stdout, "[DEBUG]", log.Ltime|log.LUTC|log.Ldate), false}
