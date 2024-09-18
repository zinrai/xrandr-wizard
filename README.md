# xrandr-wizard

xrandr-wizard simplifies the process of configuring multiple displays using xrandr. It provides an interactive interface to select displays and their positions, generating and executing the appropriate xrandr commands.

## Features

- Detect connected displays
- Interactive selection of base display and additional displays to configure
- Easy positioning of displays relative to the base display
- Support for rotation and turning off displays
- Automatic generation and execution of xrandr commands
- Ability to configure multiple displays in a single session

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

### Configuring Multiple Displays

Here's an example of using xrandr-wizard to configure multiple displays:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------

Connected displays:
1. eDP-1
2. DP-1
3. HDMI-2
Select the base display (enter the number): 1

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. DP-1
2. HDMI-2
Select the display to configure (enter the number): 1
Configuring DP-1
Enter position (above, below, left, right, left-rotate, right-rotate, off): above
Do you want to configure another display? (y/n)
y

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. HDMI-2
Select the display to configure (enter the number): 1
Configuring HDMI-2
Enter position (above, below, left, right, left-rotate, right-rotate, off): left
Executing command: xrandr --output eDP-1 --auto --output DP-1 --auto --above eDP-1 --output HDMI-2 --auto --left-of eDP-1
Command executed successfully
Configuration complete. Goodbye!
```

In this example, the user configures DP-1 to be above eDP-1 and HDMI-2 to be left of eDP-1.

### Turning Off Multiple Displays

Here's an example of using xrandr-wizard to turn off multiple displays:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------

Connected displays:
1. eDP-1
2. DP-1
3. HDMI-2
Select the base display (enter the number): 1

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. DP-1
2. HDMI-2
Select the display to configure (enter the number): 1
Configuring DP-1
Enter position (above, below, left, right, left-rotate, right-rotate, off): off
Do you want to configure another display? (y/n)
y

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. HDMI-2
Select the display to configure (enter the number): 1
Configuring HDMI-2
Enter position (above, below, left, right, left-rotate, right-rotate, off): off
Executing command: xrandr --output eDP-1 --auto --output DP-1 --off --output HDMI-2 --off
Command executed successfully
Configuration complete. Goodbye!
```

In this example, the user turns off both DP-1 and HDMI-2 displays.

### Configuring a Single Display

Here's an example of configuring a single display:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------

Connected displays:
1. eDP-1
2. DP-1
3. HDMI-2
Select the base display (enter the number): 1

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. DP-1
2. HDMI-2
Select the display to configure (enter the number): 2
Configuring HDMI-2
Enter position (above, below, left, right, left-rotate, right-rotate, off): left
Do you want to configure another display? (y/n)
n
Executing command: xrandr --output eDP-1 --auto --output HDMI-2 --auto --left-of eDP-1
Command executed successfully
Configuration complete. Goodbye!
```

In this example, the user configures only HDMI-2 to be left of eDP-1.

### Turning Off a Single Display

Here's an example of turning off a single display:

```
$ ./xrandr-wizard
Welcome to xrandr-wizard!
This tool will help you configure your displays using xrandr.
----------------------------------------------------------

Connected displays:
1. eDP-1
2. HDMI-2
Select the base display (enter the number): 1

Configuring display relative to eDP-1 (Base Display)
Remaining displays to configure:
1. HDMI-2
Select the display to configure (enter the number): 1
Configuring HDMI-2
Enter position (above, below, left, right, left-rotate, right-rotate, off): off
Executing command: xrandr --output eDP-1 --auto --output HDMI-2 --off
Command executed successfully
Configuration complete. Goodbye!
```

In this example, the user turns off the HDMI-2 display.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
