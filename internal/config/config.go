package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mtsfy/umm/internal/types"
)

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}

	baseDir := filepath.Join(home, ".umm-cli")
	configPath := filepath.Join(baseDir, "config.json")

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", fmt.Errorf("error creating base directory %s: %w", baseDir, err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := types.Config{
			Model: "gpt-4o-mini",
		}
		data, _ := json.MarshalIndent(defaultConfig, "", "  ")
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return "", fmt.Errorf("error creating config file %s: %w", configPath, err)
		}
	}

	return configPath, nil
}

func loadFile(flag int) (*os.File, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
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

func readConfig() types.Config {
	file, err := loadFile(os.O_RDWR | os.O_CREATE)
	if err != nil {
		fmt.Printf("Warning: failed to load config file: %v\n", err)
		return types.Config{Model: "gpt-4o-mini"}
	}
	defer closeFile(file)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Warning: failed to read config file: %v\n", err)
		return types.Config{Model: "gpt-4o-mini"}
	}

	c := types.Config{}
	if len(data) > 0 {
		if err = json.Unmarshal(data, &c); err != nil {
			fmt.Printf("Warning: failed to parse config file: %v\n", err)
			return types.Config{Model: "gpt-4o-mini"}
		}
	}

	return c
}

func writeConfig(config types.Config) error {
	file, err := loadFile(os.O_RDWR | os.O_CREATE)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}
	defer closeFile(file)

	newData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize updated config: %w", err)
	}

	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate config file: %w", err)
	}

	if _, err = file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek config file: %w", err)
	}

	if _, err = file.Write(newData); err != nil {
		return fmt.Errorf("failed to write updated config to file: %w", err)
	}

	return nil
}

func maskKey(apiKey string) string {
	if len(apiKey) < 8 {
		return "***"
	}
	if len(apiKey) < 16 {
		return apiKey[:4] + "***" + apiKey[len(apiKey)-4:]
	}
	return apiKey[:8] + "..." + apiKey[len(apiKey)-8:]
}

func Setup() {
	config := readConfig()
	var response string
	var choice string

	if config.ApiKey != "" {
		fmt.Printf("Current API Key: %s\n", maskKey(config.ApiKey))
		fmt.Print("Do you want to update the API key? (y/n): ")
		fmt.Scanln(&response)

		if strings.ToLower(response) == "y" {
			fmt.Print("Enter new API Key: ")
			fmt.Scanln(&config.ApiKey)
			fmt.Println("✓ API key updated")
		}
	} else {
		fmt.Print("Enter API Key: ")
		fmt.Scanln(&config.ApiKey)
		if config.ApiKey != "" {
			fmt.Println("✓ API key set")
		}
	}

	fmt.Println("\nAvailable models:")
	fmt.Println("1. gpt-4o-mini (fast, cost-effective)")
	fmt.Println("2. gpt-4o (balanced)")
	fmt.Println("3. gpt-4 (high quality)")

	if config.Model != "" {
		fmt.Printf("\nCurrent model: %s\n", config.Model)
		fmt.Print("Do you want to change the model? (y/n): ")
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			fmt.Println("Keeping current model.")
			goto save
		}
	}

	fmt.Print("Select model (1-3): ")
	fmt.Scanln(&choice)

	fmt.Println(choice)
	switch choice {
	case "1":
		config.Model = "gpt-4o-mini"
		fmt.Println("✓ Model set to gpt-4o-mini")
	case "2":
		config.Model = "gpt-4o"
		fmt.Println("✓ Model set to gpt-4o")
	case "3":
		config.Model = "gpt-4"
		fmt.Println("✓ Model set to gpt-4")
	default:
		fmt.Println("Invalid choice.")
		if config.Model == "" {
			config.Model = "gpt-4o-mini"
			fmt.Println("Using default model: gpt-4o-mini")
		} else {
			fmt.Println("Keeping current model.")
		}
	}

save:
	if err := writeConfig(config); err != nil {
		fmt.Printf("\nError saving configuration: %v\n", err)
		return
	}

	fmt.Println("\n✓ Configuration saved successfully!")

	fmt.Println("\nFinal configuration:")
	fmt.Printf("  API Key: %s\n", maskKey(config.ApiKey))
	fmt.Printf("  Model: %s\n", config.Model)
}

func Show() {
	config := readConfig()

	if config.ApiKey != "" {
		fmt.Printf("API Key: %s\n", maskKey(config.ApiKey))
	} else {
		fmt.Println("API Key: Not set")
	}

	if config.Model != "" {
		fmt.Printf("Model: %s\n", config.Model)
	} else {
		fmt.Println("Model: Not set (will use default)")
	}
}
