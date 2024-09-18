# xrandr-wizard

Simplifies the process of configuring multiple displays using xrandr. It provides an interactive interface to select displays and their positions, generating and executing the appropriate xrandr commands.

## Features

- Detect connected displays
- Interactive selection of displays to configure
- Easy positioning of displays relative to each other
- Support for rotation and turning off displays
- Automatic generation and execution of xrandr commands

## Installation

Build the tool:

```
$ go build
```

## Usage

Run the tool by executing:

```
./xrandr-wizard
```

Follow the interactive prompts to configure your displays.

## Examples

### Positioning a Display

Here's an example of using xrandr-wizard to position a display:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------
Connected displays:
1. eDP-1
2. HDMI-2
Select the display to configure (enter the number): 2
Select the reference display (base display) (enter the number): 1
Enter position (above, below, left, right, left-rotate, right-rotate, off): left
Executing command: xrandr --output HDMI-2 --auto --left-of eDP-1
Command executed successfully
```

In this example, the user configures the HDMI-2 display to be positioned to the left of the eDP-1 display.

### Turning Off a Display

Here's an example of using xrandr-wizard to turn off a display:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------
Connected displays:
1. eDP-1
2. HDMI-2
Select the display to configure (enter the number): 2
Select the reference display (base display) (enter the number): 1
Enter position (above, below, left, right, left-rotate, right-rotate, off): off
Executing command: xrandr --output HDMI-2 --off
Command executed successfully
```

In this example, the user turns off the HDMI-2 display.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
