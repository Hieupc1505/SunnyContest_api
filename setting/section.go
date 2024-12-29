package setting

type Config struct {
	Server ServerSetting   `mapstructure:"server"`
	PgDb   PostgresSetting `mapstructure:"postgres"`
	Logger LoggerSetting   `mapstructure:"logger"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LoggerSetting struct {
	Log_Level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type PostgresSetting struct {
	DbDriver string `mapstructure:"dbDriver"`
	DbSource string `mapstructure:"dbSource"`
}
