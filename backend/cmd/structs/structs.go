package structs

type Settings struct {
	RedisURL      string `yaml:"redis-url"`
	RedisPort     int    `yaml:"redis-port"`
	RedisPassword string `yaml:"redis-password"`

	MongoURI string `yaml:"mongo-uri"`
}

type Version struct {
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Version string `bson:"version,omitempty" json:"version,omitempty"`
}
