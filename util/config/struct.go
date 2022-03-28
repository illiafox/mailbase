package config

type MySQL struct {
	Login    string
	Pass     string
	DbName   string // 'mysql'
	IP       string
	Port     string
	Protocol string
}

type Redis struct {
	Addr string
	Pass string // Optional
	DB   int
}

type Email struct {
	Mail     string
	Password string
	Hostname string
	Port     int
}

type Host struct {
	Port string
	HTTP bool   // Start server in http mode
	Key  string // Path to TLS key
	Cert string // Path to TLS key
}

type Config struct {
	MySQL MySQL
	Redis Redis
	Host  Host
	SMTP  Email
}
