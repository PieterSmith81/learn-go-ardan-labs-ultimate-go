package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	/* Read a value from a file.
	Here we are simulating reading a process id numerical value from a file,
	and then killing the process (side note: no actual code to kill the process has been written - it's just a theoretical example). */
	err := KillServer("server.pid")

	if err != nil {
		// Basic error output.
		// fmt.Println("ERROR:", err)

		// Check for a specific error.
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("File not found.")
		}

		// Range over all errors.
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func KillServer(pidFile string) error {
	// Idiom: Try to acquire a resource, check for error, defer release.
	file, err := os.Open(pidFile)

	if err != nil {
		return err
	}

	/*
		- Defer happens when a function exits, including when your program panics.
		- Defer works at the function level.
		- Defers are executed in reverse order (stack, LIFO).
	*/
	// defer file.Close()
	defer func() {
		if err := file.Close(); err != nil {
			// We're not failing, only warning.
			slog.Warn("close", "file", pidFile, "error", err)
		}
	}()

	var pid int

	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", pidFile, err)
	}

	slog.Info("Killing", "pid", pid)

	if err := os.Remove(pidFile); err != nil {
		// We're not failing, only warning.
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}
