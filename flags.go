package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/term"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	authorities    []int
	date           time.Time
	iftttEventName string
	iftttKey       string
)

func parseFlags() error {
	authoritiesFlag := flag.String("authorities", "", "comma-separated list of authorities")
	dateFlag := flag.String("date", "", "date in ISO format, e.g. 2021-06-30")
	iftttEventNameFlag := flag.String("ifttt-event-name", "", "IFTTT event name of webhook")
	flag.Parse()

	if *authoritiesFlag == "" {
		return errors.New("authorities flag missing")
	}

	if *dateFlag == "" {
		return errors.New("date flag missing")
	}

	if *iftttEventNameFlag == "" {
		return errors.New("ifttt-event-name flag missing")
	}

	err := parseAuthoritiesFlag(authoritiesFlag)
	if err != nil {
		return fmt.Errorf("parse authorities failed: %w", err)
	}

	err = parseDateFlag(dateFlag)
	if err != nil {
		return fmt.Errorf("parse date failed: %w", err)
	}

	parseIFTTTEventName(iftttEventNameFlag)

	err = promptIFTTTKey()
	if err != nil {
		return fmt.Errorf("prompt for ifttt key failed: %w", err)
	}

	return nil
}

func parseAuthoritiesFlag(authoritiesFlag *string) error {
	for _, authority := range strings.Split(*authoritiesFlag, ",") {
		atoi, err := strconv.Atoi(authority)
		if err != nil {
			return fmt.Errorf("convert authority string to int failed: %w", err)
		}
		authorities = append(authorities, atoi)
	}

	return nil
}

func parseDateFlag(dateFlag *string) error {
	var err error
	date, err = time.Parse("2006-01-02", strings.TrimSpace(*dateFlag))
	if err != nil {
		return fmt.Errorf("parse date failed: %w", err)
	}

	return nil
}

func parseIFTTTEventName(iftttEventNameFlag *string) {
	iftttEventName = strings.TrimSpace(*iftttEventNameFlag)
}

func promptIFTTTKey() error {
	fmt.Println("Enter IFTTT Key: ")
	iftttKeyBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("read password prompt failed: %w", err)
	}

	iftttKey = string(iftttKeyBytes)
	return nil
}
