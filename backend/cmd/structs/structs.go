package structs

type Settings struct {
	RedisURL      string `yaml:"redis-url"`
	RedisPort     int    `yaml:"redis-port"`
	RedisPassword string `yaml:"redis-password"`

	ExposedURL string `yaml:"exposed-url"`

	MongoURI string `yaml:"mongo-uri"`

	ZapLevel string `yaml:"zap-level"`

	GinMode string `yaml:"gin-mode"`
}

type ShortLink struct {
	ID   string `bson:"shortid" json:"shortid"`
	Link string `bson:"shortlink" json:"shortlink"`
	URL  string `bson:"url" json:"url"`
}

type LinkPostBody struct {
	URL string `json:"url"`
}
