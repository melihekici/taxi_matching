package config

import (
	"fmt"
	"os"
)

// var MONGO = map[string]string{
// 	"URI":            "mongodb://localhost:27017",
// 	"DATABASE":       "bitaksi",
// 	"COLLECTION":     "driver",
// 	"TESTCOLLECTION": "test",
// }

// var SECRET = "^S+%R&YUI&G/H(J/H)("

var MONGO = map[string]string{
	"URI":            fmt.Sprintf("mongodb://%s:27017", os.Getenv("MONGO_HOST")),
	"DATABASE":       os.Getenv("DATABASE"),
	"COLLECTION":     os.Getenv("COLLECTION"),
	"TESTCOLLECTION": os.Getenv("TESTCOLLECTION"),
}

var SECRET = os.Getenv("SECRET")
