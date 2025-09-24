package history

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/mtsfy/umm/internal/types"
)

func getHistoryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}

	baseDir := filepath.Join(home, ".umm-cli")
	historyPath := filepath.Join(baseDir, "history.json")

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", fmt.Errorf("error creating base directory %s: %w", baseDir, err)
	}

	if _, err := os.Stat(historyPath); os.IsNotExist(err) {
		emptyHistory := "{\"interactions\": []}\n"
		if err := os.WriteFile(historyPath, []byte(emptyHistory), 0644); err != nil {
			return "", fmt.Errorf("error creating history file %s: %w", historyPath, err)
		}
	}

	return historyPath, nil
}

func loadFile(flag int) (*os.File, error) {
	path, err := getHistoryPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get history path: %w", err)
	}

	f, err := os.OpenFile(path, flag, 0644)
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

func Save(interaction types.Interaction) error {
	history := readHistory()
	history.Interactions = append([]types.Interaction{interaction}, history.Interactions...)
	return writeHistory(history)
}

func GetLatest() types.Interaction {
	history := readHistory()
	if len(history.Interactions) == 0 {
		return types.Interaction{}
	}
	return history.Interactions[0]
}

func GetByID(id int) types.Interaction {
	history := readHistory()

	if id < 1 || id > len(history.Interactions) {
		return types.Interaction{}
	}

	return history.Interactions[id-1]
}

func AllHistory() {
	history := readHistory()

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "No.\tUser Query\tSuggested Command\tDate")

	for i, item := range history.Interactions {
		date := item.Date.Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, item.UserInput, item.AIResponse.Command, date)
	}
}

func PaginatedHistory(page, size int) {
	history := readHistory()
	total := len(history.Interactions)

	start := (page - 1) * size
	if start >= total {
		fmt.Println("Page out of range.")
		return
	}

	end := start + size
	if end > total {
		end = total
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "No.\tUser Query\tSuggested Command\tDate")

	for i, item := range history.Interactions {
		date := item.Date.Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, item.UserInput, item.AIResponse.Command, date)
	}

	fmt.Fprintf(w, "\nShowing page %d (entries %d to %d of %d)\n", page, start+1, end, total)
}

func FilterHistory(keyword string) {
	history := readHistory()

	var filtered []struct {
		index       int
		interaction types.Interaction
	}
	for i, inter := range history.Interactions {
		if strings.Contains(strings.ToLower(inter.UserInput), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(inter.AIResponse.Command), strings.ToLower(keyword)) {
			filtered = append(filtered, struct {
				index       int
				interaction types.Interaction
			}{
				index:       i + 1,
				interaction: inter,
			})
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "No.\tUser Query\tSuggested Command\tDate")

	for _, item := range filtered {
		date := item.interaction.Date.Format("2006-01-02 15:04:05")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", item.index, item.interaction.UserInput, item.interaction.AIResponse.Command, date)
	}

}

func DeleteAllHistory() error {
	emptyHistory := types.History{Interactions: []types.Interaction{}}
	return writeHistory(emptyHistory)
}

func DeleteHistory(id int) error {
	history := readHistory()

	if id < 1 || id > len(history.Interactions) {
		return fmt.Errorf("invalid history ID: %d. Valid range is 1-%d", id, len(history.Interactions))
	}

	index := id - 1
	history.Interactions = append(history.Interactions[:index], history.Interactions[index+1:]...)

	return writeHistory(history)
}
