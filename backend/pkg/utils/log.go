package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func ConfigLogger(config Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fileLogger := &lumberjack.Logger{
		Filename:   config.LogFilename,   // Specify the log file name
		MaxSize:    config.LogMaxSize,    // Max size in megabytes before rotation
		MaxBackups: config.LogMaxBackups, // Max number of backup files to keep
		MaxAge:     config.LogMaxAge,     // Max number of days to retain old log files
		Compress:   config.LogCompress,   // Whether to compress old log files
	}

	multiWrite := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stdout},
		fileLogger,
	)
	log.Logger = zerolog.New(multiWrite).With().Timestamp().Caller().Logger()
}
