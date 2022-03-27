package config

var MONGO = map[string]string{
	"URI":            "mongodb://mongodb:27017",
	"DATABASE":       "bitaksi",
	"COLLECTION":     "driver",
	"TESTCOLLECTION": "test",
}

// func init() {
// 	file, _ := os.Open("config/conf.json")
// 	defer file.Close()
// 	decoder := json.NewDecoder(file)
// 	config := configuration{}
// 	err := decoder.Decode(&config)
// 	if err != nil {
// 		log.Fatalln("Error reading configuration.", err.Error())
// 	}

// 	MONGO = config.MONGO
// }
