package internal

func GetConfig() Config {
	return cfg
}

type Config struct {
	GameDeckSize int
}

var cfg = Config{GameDeckSize: 20}
