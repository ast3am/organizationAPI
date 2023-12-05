package models

type Config struct {
	ListenPort        string            `yaml:"listen_port"`
	SqlConfig         SqlConfig         `yaml:"sql_config"`
	EmailDealerConfig EmailDealerConfig `yaml:"email_dealer_config"`
	LogLevel          string            `yaml:"log_level"`
}

type SqlConfig struct {
	UsernameDB string `yaml:"username_db"`
	PasswordDB string `yaml:"password_db"`
	HostDB     string `yaml:"host_db"`
	PortDB     string `yaml:"port_db"`
	DBName     string `yaml:"db_name"`
	DelayTime  int    `yaml:"delay_time"`
}

type EmailDealerConfig struct {
	Host     string `yaml:"host_email"`
	Port     int    `yaml:"port_email"`
	Username string `yaml:"username_email"`
	Password string `yaml:"password_email"`
}
