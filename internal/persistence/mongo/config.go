package mongo

type Config struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	UseMock  bool   `yaml:"useMock"`
}
