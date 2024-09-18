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
	fmt.Println("Welcome to xrandr-wizard!")
	fmt.Println("This tool will help you configure your displays using xrandr.")
	fmt.Println("----------------------------------------------------------")

	displays := getConnectedDisplays()
	if len(displays) == 0 {
		fmt.Println("No displays connected.")
		return
	}

	config := configureDisplays(displays)
	command := generateXrandrCommand(config)
	executeCommand(command)

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
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, " connected") {
			parts := strings.Fields(line)
			displays = append(displays, Display{Name: parts[0], Status: strings.Join(parts[2:], " ")})
		}
	}
	return displays
}

func configureDisplays(displays []Display) Configuration {
	fmt.Println("\nConnected displays:")
	for i, display := range displays {
		fmt.Printf("%d. %s\n", i+1, display.Name)
	}

	baseIndex := promptForNumber("Select the base display (enter the number): ", 1, len(displays)) - 1
	config := Configuration{BaseDisplay: displays[baseIndex]}

	remainingDisplays := append(displays[:baseIndex], displays[baseIndex+1:]...)
	for len(remainingDisplays) > 0 {
		fmt.Printf("\nConfiguring display relative to %s (Base Display)\n", config.BaseDisplay.Name)
		fmt.Println("Remaining displays to configure:")
		for i, display := range remainingDisplays {
			fmt.Printf("%d. %s\n", i+1, display.Name)
		}
		displayIndex := promptForNumber("Select the display to configure (enter the number): ", 1, len(remainingDisplays)) - 1
		display := remainingDisplays[displayIndex]
		fmt.Printf("Configuring %s\n", display.Name)
		display.Position = promptForPosition()
		config.Others = append(config.Others, display)

		remainingDisplays = append(remainingDisplays[:displayIndex], remainingDisplays[displayIndex+1:]...)

		if len(remainingDisplays) > 0 {
			fmt.Println("Do you want to configure another display? (y/n)")
			if !confirmContinue() {
				break
			}
		}
	}

	return config
}

func promptForNumber(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
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

func generateXrandrCommand(config Configuration) string {
	command := fmt.Sprintf("xrandr --output %s --auto", config.BaseDisplay.Name)
	for _, display := range config.Others {
		switch display.Position {
		case "above":
			command += fmt.Sprintf(" --output %s --auto --above %s", display.Name, config.BaseDisplay.Name)
		case "below":
			command += fmt.Sprintf(" --output %s --auto --below %s", display.Name, config.BaseDisplay.Name)
		case "left":
			command += fmt.Sprintf(" --output %s --auto --left-of %s", display.Name, config.BaseDisplay.Name)
		case "right":
			command += fmt.Sprintf(" --output %s --auto --right-of %s", display.Name, config.BaseDisplay.Name)
		case "left-rotate":
			command += fmt.Sprintf(" --output %s --auto --left-of %s --rotate left", display.Name, config.BaseDisplay.Name)
		case "right-rotate":
			command += fmt.Sprintf(" --output %s --auto --right-of %s --rotate left", display.Name, config.BaseDisplay.Name)
		case "off":
			command += fmt.Sprintf(" --output %s --off", display.Name)
		}
	}
	return command
}

func executeCommand(command string) {
	fmt.Println("Executing command:", command)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Command executed successfully")
	}
}

func confirmContinue() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}
