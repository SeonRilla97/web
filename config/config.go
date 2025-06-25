package config

type Config struct {
	Server struct {
		Port string `json:"port"`
		Info string `json:"info"`
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err == nil {
		panic(err)
	}else if err := json.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	}else {
	return c
	}
}
