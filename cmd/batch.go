package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"rprov/internal/batch"
	"rprov/soap/admin"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const maxBatchLineSize = 1024 * 1024

type savedFlag struct {
	flag    *pflag.Flag
	value   string
	changed bool
}

func runBatch(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var saved []savedFlag
	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		saved = append(saved, savedFlag{flag: f, value: f.Value.String(), changed: f.Changed})
	})

	rootCmd.SilenceUsage = true

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), maxBatchLineSize)

	lineNumber := 0
	total := 0
	failed := 0

	for scanner.Scan() {
		lineNumber++

		line := strings.TrimSuffix(scanner.Text(), "\r")

		if lineNumber == 1 {
			line = strings.TrimPrefix(line, "\uFEFF")
		}

		trimmed := strings.TrimSpace(line)

		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		total++

		if err := runBatchLine(line, saved); err != nil {
			var authErr *admin.AuthError
			if errors.As(err, &authErr) {
				return fmt.Errorf("line %d: %w; stopping to avoid repeated login attempts", lineNumber, err)
			}

			failed++
			fmt.Fprintf(os.Stderr, "ERROR: line %d: %v\n", lineNumber, err)
		}
	}

	if err := scanner.Err(); err != nil {
		if errors.Is(err, bufio.ErrTooLong) {
			return fmt.Errorf("line %d is longer than %d bytes", lineNumber+1, maxBatchLineSize)
		}

		return err
	}

	if failed > 0 {
		return fmt.Errorf("%d of %d commands failed", failed, total)
	}

	return nil
}

func runBatchLine(line string, saved []savedFlag) error {
	if err := restoreFlags(saved); err != nil {
		return err
	}

	args, err := batch.Split(line)
	if err != nil {
		return err
	}

	cmd, _, err := rootCmd.Find(args)
	if err != nil {
		return err
	}

	if cmd == rootCmd {
		return fmt.Errorf("unknown command %q", args[0])
	}

	rootCmd.SetArgs(args)

	_, err = rootCmd.ExecuteC()

	return err
}

func restoreFlags(saved []savedFlag) error {
	if err := resetLocalFlags(rootCmd); err != nil {
		return err
	}

	for _, s := range saved {
		if err := s.flag.Value.Set(s.value); err != nil {
			return err
		}

		s.flag.Changed = s.changed
	}

	return nil
}

func resetLocalFlags(c *cobra.Command) error {
	var err error

	c.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag) {
		if err != nil {
			return
		}

		if slice, ok := f.Value.(pflag.SliceValue); ok {
			err = slice.Replace(parseSliceDefault(f.DefValue))
		} else {
			err = f.Value.Set(f.DefValue)
		}

		f.Changed = false
	})

	if err != nil {
		return err
	}

	for _, sub := range c.Commands() {
		if err := resetLocalFlags(sub); err != nil {
			return err
		}
	}

	return nil
}

func parseSliceDefault(def string) []string {
	def = strings.TrimSuffix(strings.TrimPrefix(def, "["), "]")

	if def == "" {
		return []string{}
	}

	return strings.Split(def, ",")
}
