package config

import (
	"flag"
)

// Доступные флаги командной строки
var (
	// flagServerHost    string // адрес хоста на котором запущен проект
	// flagServerPort    string // номер порта на котором запущен проект
	flagServerSocket    string
	flagServerBaseURL   string // базовый URL проекта
	flagFileStoragePath string // путь до файла с данными проекта
	// flagDatabaseDSN       string // строка для подключения к базе данных
	// flagBaseProfilePath   string
	// flagResultProfilePath string
	// flagPackageName       string
)

func init() {
	// flag.StringVar(&flagServerHost, "server-host", "", "host of target HTTP address")
	// flag.StringVar(&flagServerPort, "server-port", "", "port of target HTTP address")
	flag.StringVar(&flagServerSocket, "a", ":8080", "host:port of target HTTP address")
	flag.StringVar(&flagServerBaseURL, "b", "http://localhost:8080", "base URL of target HTTP address")
	flag.StringVar(&flagFileStoragePath, "f", "/tmp/short-url-db.json", "path to persistent file storage")
	// flag.StringVar(&flagDatabaseDSN, "database-dsn", "", "connection string to database")
	// flag.StringVar(&flagBaseProfilePath, "base-profile-path", "", "path to base pprof profile")
	// flag.StringVar(&flagResultProfilePath, "result-profile-path", "", "path to result pprof profile")
	// flag.StringVar(&flagPackageName, "package-name", "", "name of package to be tested")
}
