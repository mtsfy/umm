package history

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/mtsfy/umm/internal/types"
)

func Save(interaction types.Interaction) error {
	history := readHistory()
	history.Interactions = append(history.Interactions, interaction)
	return writeHistory(history)
}

func loadFile(flag int) (*os.File, error) {
	f, err := os.OpenFile("data/umm-history.json", flag, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open history file: %w", err)
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("failed to acquire file lock: %w", err)
	}
	return f, nil
}

func closeFile(f *os.File) error {
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN); err != nil {
		return fmt.Errorf("failed to release file lock: %w", err)
	}
	return f.Close()
}

func readHistory() types.History {
	file, err := loadFile(os.O_RDWR | os.O_CREATE)
	if err != nil {
		panic(fmt.Errorf("failed to load history file: %w", err))
	}
	defer closeFile(file)

	data, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed to read history file: %w", err))
	}

	history := types.History{}
	if len(data) > 0 {
		err = json.Unmarshal(data, &history)
		if err != nil {
			panic(fmt.Errorf("failed to parse history file: %w", err))
		}
	}
	return history
}

func writeHistory(history types.History) error {
	file, err := loadFile(os.O_RDWR | os.O_CREATE)
	if err != nil {
		return fmt.Errorf("failed to load history file: %w", err)
	}
	defer closeFile(file)

	newData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize updated history: %w", err)
	}

	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate history file: %w", err)
	}

	if _, err = file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek history file: %w", err)
	}

	if _, err = file.Write(newData); err != nil {
		return fmt.Errorf("failed to write updated history to file: %w", err)
	}

	return nil
}

func GetLatest() types.Interaction {
	history := readHistory()
	return history.Interactions[len(history.Interactions)-1]
}
