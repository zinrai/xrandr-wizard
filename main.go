package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Display struct {
	Name     string
	Status   string
	Position string
}

type Configuration struct {
	BaseDisplay Display
	Others      []Display
}

func main() {
	if _, err := exec.LookPath("xrandr"); err != nil {
		fmt.Println("xrandr is not installed or not in PATH")
		os.Exit(1)
	}

	fmt.Println("Welcome to xrandr-wizard!")
	printVersion()
	fmt.Println("This tool will help you configure your displays using xrandr.")
	fmt.Println("----------------------------------------------------------")

	displays := getConnectedDisplays()
	if len(displays) == 0 {
		fmt.Println("No displays connected.")
		return
	}

	config := configureDisplays(displays)
	executeCommand(config)

	fmt.Println("Configuration complete. Goodbye!")
}

func getConnectedDisplays() []Display {
	cmd := exec.Command("xrandr")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing xrandr:", err)
		return nil
	}

	var displays []Display
	for _, line := range strings.Split(string(output), "\n") {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		if parts[1] != "connected" {
			continue
		}
		displays = append(displays, Display{Name: parts[0], Status: strings.Join(parts[2:], " ")})
	}
	return displays
}

func configureDisplays(displays []Display) Configuration {
	fmt.Println("\nConnected displays:")
	for i, display := range displays {
		fmt.Printf("%d. %s\n", i+1, display.Name)
	}

	baseIndex := promptForNumber("Select the base display (enter the number): ", 1, len(displays), 1) - 1
	config := Configuration{BaseDisplay: displays[baseIndex]}

	remainingDisplays := removeIndex(displays, baseIndex)
	for len(remainingDisplays) > 0 {
		fmt.Printf("\nConfiguring display relative to %s (Base Display)\n", config.BaseDisplay.Name)
		fmt.Println("Remaining displays to configure:")
		for i, display := range remainingDisplays {
			fmt.Printf("%d. %s\n", i+1, display.Name)
		}
		displayIndex := promptForNumber("Select the display to configure (enter the number): ", 1, len(remainingDisplays), 1) - 1
		display := remainingDisplays[displayIndex]
		fmt.Printf("Configuring %s\n", display.Name)
		display.Position = promptForPosition()
		config.Others = append(config.Others, display)

		remainingDisplays = removeIndex(remainingDisplays, displayIndex)

		if len(remainingDisplays) == 0 {
			break
		}
		fmt.Println("Do you want to configure another display? (y/n)")
		if !confirmContinue() {
			break
		}
	}

	return config
}

func removeIndex(s []Display, i int) []Display {
	result := make([]Display, 0, len(s)-1)
	result = append(result, s[:i]...)
	result = append(result, s[i+1:]...)
	return result
}

func promptForNumber(prompt string, min, max int, defaultValue int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s[%d]: ", prompt, defaultValue)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return defaultValue
		}

		num := 0
		_, err := fmt.Sscanf(input, "%d", &num)
		if err == nil && num >= min && num <= max {
			return num
		}
		fmt.Printf("Invalid input. Please enter a number between %d and %d.\n", min, max)
	}
}

func promptForPosition() string {
	reader := bufio.NewReader(os.Stdin)
	validPositions := map[string]bool{
		"above": true, "below": true, "left": true, "right": true,
		"left-rotate": true, "right-rotate": true, "off": true,
	}

	for {
		fmt.Print("Enter position (above, below, left, right, left-rotate, right-rotate, off): ")
		input, _ := reader.ReadString('\n')
		position := strings.TrimSpace(input)
		if validPositions[position] {
			return position
		}
		fmt.Println("Invalid position. Please try again.")
	}
}

func buildDisplayArgs(display Display, baseName string) []string {
	switch display.Position {
	case "above":
		return []string{"--output", display.Name, "--auto", "--above", baseName}
	case "below":
		return []string{"--output", display.Name, "--auto", "--below", baseName}
	case "left":
		return []string{"--output", display.Name, "--auto", "--left-of", baseName}
	case "right":
		return []string{"--output", display.Name, "--auto", "--right-of", baseName}
	case "left-rotate":
		return []string{"--output", display.Name, "--auto", "--left-of", baseName, "--rotate", "left"}
	case "right-rotate":
		return []string{"--output", display.Name, "--auto", "--right-of", baseName, "--rotate", "right"}
	case "off":
		return []string{"--output", display.Name, "--off"}
	}
	return nil
}

func executeCommand(config Configuration) {
	args := []string{"--output", config.BaseDisplay.Name, "--auto"}
	for _, d := range config.Others {
		args = append(args, buildDisplayArgs(d, config.BaseDisplay.Name)...)
	}
	fmt.Println("Executing command: xrandr", strings.Join(args, " "))
	cmd := exec.Command("xrandr", args...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Command executed successfully")
}

func confirmContinue() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}
