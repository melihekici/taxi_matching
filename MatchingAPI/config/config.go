package config

import "os"

var SECRET = os.Getenv("SECRET")
