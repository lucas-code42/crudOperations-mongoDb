package config

type mongoConfig struct {
	Uri               string
	DefaultDb         string
	DefaultCollection string
}

func GetMongoConfig() *mongoConfig {
	return &mongoConfig{
		Uri:               "mongodb://root:example@mongo:27017/",
		DefaultDb:         "db",
		DefaultCollection: "books",
	}
}
