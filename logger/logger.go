package logger

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

var logger *slog.Logger

func Init(logLevel slog.Level, writers ...io.Writer) {
	lvl := LevelTrace
	if logLevel >= LevelTrace && logLevel <= LevelFatal {
		lvl = logLevel
	}

	var writer io.Writer
	switch len(writers) {
	case 0:
		writer = os.Stdout
	case 1:
		writer = writers[0]
	default:
		writer = io.MultiWriter(writers...)
	}

	logHandler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				// a.Key = "LEVEL"
				level := a.Value.Any().(slog.Level)

				switch {
				case level < LevelDebug:
					a.Value = slog.StringValue("TRACE")
				case level < LevelInfo:
					a.Value = slog.StringValue("DEBUG")
				case level < LevelWarn:
					a.Value = slog.StringValue("INFO")
				case level < LevelError:
					a.Value = slog.StringValue("WARNING")
				case level < LevelFatal:
					a.Value = slog.StringValue("ERROR")
				default:
					a.Value = slog.StringValue("FATAL")
				}
			}

			return a
		},
	})
	logger = slog.New(logHandler)
	slog.SetDefault(logger)
}

func Debug(msg string, args ...any) {
	if len(args) == 0 {
		logger.Debug(msg, slog.String("source", getCaller()))
		return
	}

	logger.Debug(msg, genSlogAttrs(args)...)
}

func Info(msg string, args ...any) {
	if len(args) == 0 {
		logger.Info(msg, slog.String("source", getCaller()))
		return
	}

	logger.Info(msg, genSlogAttrs(args)...)
}

func Warn(msg string, args ...any) {
	if len(args) == 0 {
		logger.Warn(msg, slog.String("source", getCaller()))
		return
	}

	logger.Warn(msg, genSlogAttrs(args)...)
}

func Error(err error, msg string, args ...any) {
	if len(args) == 0 {
		logger.Error(msg, slog.String("error", err.Error()), slog.String("source", getCaller()))
		return
	}

	args = append(args, "error", err.Error())
	logger.Error(msg, genSlogAttrs(args)...)
}

func Fatal(msg string, args ...any) {
	if len(args) == 0 {
		logger.Log(context.Background(), LevelFatal, msg, slog.String("source", getCaller()))
		panic(msg)
	}

	logger.Log(context.Background(), LevelFatal, msg, genSlogAttrs(args)...)
	panic(msg)
}

func genSlogAttrs(args []any) (retval []any) {
	argsLen := len(args)
	slogAttrLen := argsLen / 2
	const orphanKey = "NO-KEY"

	var orphan any
	if argsLen%2 != 0 {
		log.Println("uneven number of arguments, placing last provided value under the key '" + orphanKey + "'")
		orphan = args[argsLen-1]
		args = args[:argsLen-1]
		slogAttrLen += 1
	}

	attrs := make([]slog.Attr, slogAttrLen)
	attrs = append(attrs, slog.String("source", getCaller()))

	for i := 0; i < argsLen-1; i += 2 {
		attrs = append(attrs, slog.Any(args[i].(string), args[i+1]))
	}

	if orphan != nil {
		attrs = append(attrs, slog.Any(orphanKey, orphan))
	}

	for i := range attrs {
		retval = append(retval, attrs[i].Key, attrs[i].Value)
	}

	return
}

func getCaller() string {
	_, file, line, ok := runtime.Caller(runtimeCallerDepth)
	if !ok {
		return "unknown source"
	}
	return file + ":" + strconv.Itoa(line)
}
