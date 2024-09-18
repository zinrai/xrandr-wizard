package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Display struct {
	Name   string
	Status string
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

	fmt.Println("Connected displays:")
	for i, display := range displays {
		fmt.Printf("%d. %s\n", i+1, display.Name)
	}

	selectedDisplay := promptForDisplay(displays, "Select the display to configure")
	referenceDisplay := promptForDisplay(displays, "Select the reference display (base display)")
	position := promptForPosition()

	command := generateXrandrCommand(selectedDisplay, referenceDisplay, position)
	executeCommand(command)
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
			displays = append(displays, Display{Name: parts[0], Status: "connected"})
		}
	}
	return displays
}

func promptForDisplay(displays []Display, prompt string) Display {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s (enter the number): ", prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		index := 0
		_, err := fmt.Sscanf(input, "%d", &index)
		if err == nil && index > 0 && index <= len(displays) {
			return displays[index-1]
		}
		fmt.Println("Invalid selection. Please try again.")
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

func generateXrandrCommand(display, reference Display, position string) string {
	baseCommand := fmt.Sprintf("xrandr --output %s", display.Name)
	switch position {
	case "above":
		return fmt.Sprintf("%s --auto --above %s", baseCommand, reference.Name)
	case "below":
		return fmt.Sprintf("%s --auto --below %s", baseCommand, reference.Name)
	case "left":
		return fmt.Sprintf("%s --auto --left-of %s", baseCommand, reference.Name)
	case "right":
		return fmt.Sprintf("%s --auto --right-of %s", baseCommand, reference.Name)
	case "left-rotate":
		return fmt.Sprintf("%s --auto --left-of %s --rotate left", baseCommand, reference.Name)
	case "right-rotate":
		return fmt.Sprintf("%s --auto --right-of %s --rotate left", baseCommand, reference.Name)
	case "off":
		return fmt.Sprintf("%s --off", baseCommand)
	default:
		return ""
	}
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
