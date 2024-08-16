// Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package logging

import (
	"errors"
	"fmt"
	"log/slog"
	_ "net/http/pprof" // #nosec G108
	"os"
	"path/filepath"

	"github.com/lmittmann/tint"
	slogmulti "github.com/samber/slog-multi"
)

const (
	LevelTrace = slog.Level(-8)
	LevelFatal = slog.Level(12)
)

//nolint:misspell
var LevelNames = map[slog.Leveler]string{
	LevelTrace: "TRACE",
	LevelFatal: "FATAL",
}

//revive:disable:flag-parameter
func New(level string, logFile string) *slog.Logger {
	// Create an slog consoleOpts object.
	consoleOpts := &tint.Options{
		ReplaceAttr: levelReplacer,
	}
	fileOpts := &tint.Options{
		ReplaceAttr: levelReplacer,
		NoColor:     true,
	}

	// Set the log level.
	switch level {
	case "trace":
		consoleOpts.Level = LevelTrace
		consoleOpts.AddSource = true
		fileOpts.Level = LevelTrace
		fileOpts.AddSource = true
	case "debug":
		consoleOpts.Level = slog.LevelDebug
		fileOpts.Level = slog.LevelDebug
	default:
		consoleOpts.Level = slog.LevelInfo
		fileOpts.Level = slog.LevelInfo
	}

	// Set the slog handler
	logHandler := slogmulti.Fanout(
		tint.NewHandler(os.Stdout, consoleOpts),
	)

	// Unless no log file was requested, set up file logging.
	if logFile != "" {
		logFH, err := openLogFile(logFile)
		if err != nil {
			slog.Warn("unable to open log file", "file", logFile, "error", err)
		} else {
			logHandler = slogmulti.Fanout(
				tint.NewHandler(os.Stdout, consoleOpts),
				tint.NewHandler(logFH, fileOpts),
			)
		}
	}

	logger := slog.New(logHandler)

	slog.SetDefault(logger)

	return logger
}

func levelReplacer(_ []string, attr slog.Attr) slog.Attr {
	if attr.Key == slog.LevelKey {
		level, ok := attr.Value.Any().(slog.Level)
		if !ok {
			level = slog.LevelInfo
		}

		levelLabel, exists := LevelNames[level]
		if !exists {
			levelLabel = level.String()
		}

		attr.Value = slog.StringValue(levelLabel)
	}

	return attr
}

// openLogFile will attempt to open the specified log file. It will also attempt
// to create the directory containing the log file if it does not exist.
func openLogFile(logFile string) (*os.File, error) {
	logDir := filepath.Base(logFile)
	// Create the log directory if it does not exist.
	_, err := os.Stat(logDir)

	if err == nil || errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("unable to create log file directory %s: %w", logDir, err)
		}
	}
	// Open the log file.
	logFileHandle, err := os.Create(logFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open log file: %w", err)
	}

	return logFileHandle, nil
}

// Reset will remove the log file.
func Reset(file string) error {
	// If the log file doesn't exist, just exit.
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}
	// Else, remove the file.
	err = os.Remove(file)
	if err != nil {
		return fmt.Errorf("could not remove log file: %w", err)
	}

	return nil
}
