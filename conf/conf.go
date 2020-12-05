package conf

const (
	Debug        = true
	Addr         = ":7717"
	ReadTimeout  = 5000
	WriteTimeout = 5000
	WithTimeout  = 5000
)

type AppConfig struct {
	Debug        bool
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	WithTimeout  int
}

func DefaultAppConfig() *AppConfig {
	return &AppConfig{
		Debug:        Debug,
		Addr:         Addr,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		WithTimeout:  WithTimeout,
	}
}
