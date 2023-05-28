package commands

import (
	"dan/go-shell/pkg/commands"
	"os"
	"testing"
)

func TestCd(t *testing.T) {
	// Save the current working directory.
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Change to the parent directory.
	commands.Cd([]string{".."})

	// Check that the current working directory has changed.
	newDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	if newDir == oldDir {
		t.Fatalf("Directory did not change: was %s, now %s", oldDir, newDir)
	}

	// Change back to the original directory.
	os.Chdir(oldDir)
}
