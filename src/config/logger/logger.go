package logger

var (
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init () {
	logConfig := zap.NewConfig(){
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getgetLevelLogs()),
		Encoding: zapcore.EncoderConfig{
			Levelkey:    "level",
			TimeKey:     "time",
			MessageKey: "message",
			EncondeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
			
		}
	}
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}