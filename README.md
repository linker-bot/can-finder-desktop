# CAN Finder Desktop

Base on Project: https://github.com/linker-bot/can-finder

## Project Overview

**CAN Finder** is a cross-platform desktop application specifically designed for LinkerHand devices, providing network device auto-discovery and management functionalities. Built using Go and the Wails framework, CAN Finder listens for UDP broadcasts from devices, displaying essential information in real-time, including device name, IP address, MAC address, model, and software version.

## Features

* **Real-Time Device Discovery**: Automatically discovers CAN devices in real-time through UDP broadcasts.
* **Intuitive Desktop Interface**: Offers an intuitive and user-friendly interface to display detailed device information, including first discovery and last active timestamps.
* **Real-Time Updates**: Utilizes event-driven architecture provided by Wails for seamless real-time frontend-backend communication, eliminating the need for manual refreshes.

## System Architecture

CAN Finder comprises two main parts:

1. **Backend Service**: Developed in Go, leveraging Wails for backend logic, including UDP broadcast handling and event-driven data communication.
2. **Frontend Display**: Created using native HTML, CSS, and JavaScript, seamlessly integrated with the backend via Wails' event-based APIs.

## Technology Stack

* Go Language
* Wails Framework
* HTML, CSS, and JavaScript
* UDP Broadcast Protocol

## Installation and Usage

### Build and Run

Ensure Go and Wails CLI are installed:

```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Verify installation
wails doctor

# Build and run
git clone https://github.com/your-username/can-finder.git
cd can-finder
wails build
```

### Accessing the Application

After building, run the executable located in the `build/bin` directory:

* **Windows**: `CAN Finder.exe`
* **macOS**: `CAN Finder.app`
* **Linux**: `CAN Finder`

The desktop application will start automatically upon launching the executable.

## Project Structure

```
can-finder/
├── main.go                 # Main entry point of the Wails app
├── backend
│   └── app.go              # Backend logic, handling UDP broadcasts
└── frontend
    └── index.html          # Frontend display and logic
```

## Configuration Details

* UDP broadcast listening port is set to `9999` (modifiable in backend source code).

## Example Display Information

The desktop interface dynamically updates and displays:

* Device Name
* IP Address
* MAC Address
* Device Model
* Software Version (with links to GitHub Releases)
* First Seen Time
* Last Active Time

## Error Handling and Logging

Comprehensive logging is implemented to assist in troubleshooting device discovery and communication issues. Logs can be viewed directly within the Wails runtime console.

## Dependencies

Managed automatically via Go modules and Wails:

* github.com/wailsapp/wails/v2

Use `go mod tidy` to manage dependencies.

## Automated Building (GitHub Actions)

CAN Finder leverages GitHub Actions for automatic builds across multiple platforms (Windows, macOS, Linux) upon tagged releases or manual triggers. See `.github/workflows/build.yml` for more details.

## License

CAN Finder is released under the GPL-3.0 license.

---

Contributions are warmly welcomed! Feel free to submit code, raise issues, or create pull requests to help us continuously improve CAN Finder.
