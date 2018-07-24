package cmd

import (
	"os/exec"
	"testing"
)

func TestAdd(t *testing.T) {
	cmd := exec.Command("task", "add", "go to office")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("error at add command")
	}

}

func TestDo(t *testing.T) {
	cmd := exec.Command("task", "do", "1", "2")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("error at do command")
	}

}

func TestList(t *testing.T) {
	cmd := exec.Command("task", "list")
	_, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("error at list command")
	}

}
