package config

type MySQL struct {
	Login    string
	Pass     string
	DbName   string // 'mysql'
	Ip       string
	Port     string
	Protocol string
}

type Redis struct {
	Addr          string
	Pass          string // Optional
	DB            int
	ExpireSeconds int
}

type Email struct {
	Mail     string
	Password string
	Hostname string
	Port     int
}

type Host struct {
	Port string
}

type Config struct {
	MySQL MySQL
	Redis Redis
	Host  Host
	Smtp  Email
}
