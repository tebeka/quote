package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	var buf bytes.Buffer
	cmd := exec.Command("go", "run", ".", "-h")
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Fatalf("running: %s", err)
	}

	if !strings.Contains(buf.String(), "usage: ") {
		t.Fatalf("unexpected output:\n%s", buf.String())
	}
}

func TestVersion(t *testing.T) {
	var buf bytes.Buffer
	cmd := exec.Command("go", "run", ".", "-version")
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Fatalf("running: %s", err)
	}

	if !strings.Contains(buf.String(), "version") {
		t.Fatalf("unexpected output:\n%s", buf.String())
	}
}

func TestCount(t *testing.T) {
	var buf bytes.Buffer
	count := 7
	cmd := exec.Command("go", "run", ".", "-n", fmt.Sprintf("%d", count))
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Fatalf("running: %s", err)
	}

	matches := regexp.MustCompile(`(?m)^    - .+$`).FindAll(buf.Bytes(), -1)
	if len(matches) != count {
		t.Fatalf("expected %d quotes, got %d", count, len(matches))
	}
}

func TestJSON(t *testing.T) {
	var buf bytes.Buffer
	count := 5
	cmd := exec.Command("go", "run", ".", "-n", fmt.Sprintf("%d", count), "-json")
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		t.Fatalf("running: %s", err)
	}

	var quotes []map[string]string
	if err := json.NewDecoder(&buf).Decode(&quotes); err != nil {
		t.Fatalf("parsing JSON: %s", err)
	}

	if len(quotes) != count {
		t.Fatalf("expected %d quotes, got %d", count, len(quotes))
	}
}
