# WinPortKill 🚀

![GitHub release](https://img.shields.io/github/release/BiruRyuu/winportkill.svg) ![GitHub issues](https://img.shields.io/github/issues/BiruRyuu/winportkill.svg) ![GitHub forks](https://img.shields.io/github/forks/BiruRyuu/winportkill.svg) ![GitHub stars](https://img.shields.io/github/stars/BiruRyuu/winportkill.svg)

## Overview

**WinPortKill** is a lightweight, powerful tool designed for Windows. Built with Go, this tool helps you identify and terminate processes that occupy specific ports. It features a simple command-line interface that makes it easy to free up ports. This utility is ideal for developers, system administrators, and anyone troubleshooting port conflicts.

You can download the latest release [here](https://github.com/BiruRyuu/winportkill/releases). Simply download the file and execute it to get started.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Features

- **Lightweight**: Minimal resource usage ensures that the tool runs smoothly on any Windows machine.
- **Powerful**: Quickly identifies and terminates processes using specified ports.
- **Command-Line Interface**: Easy to use for both beginners and advanced users.
- **Cross-Platform**: Built with Go, making it easy to adapt for other platforms if needed.
- **Open Source**: Contribute to the project or customize it to fit your needs.

## Installation

To install WinPortKill, follow these steps:

1. **Download the latest release** from the [Releases page](https://github.com/BiruRyuu/winportkill/releases).
2. **Extract the downloaded file** to your preferred directory.
3. **Open Command Prompt** and navigate to the directory where you extracted the files.
4. **Run the executable**.

### Prerequisites

- Windows Operating System
- Go (optional, for development)

## Usage

Using WinPortKill is straightforward. Open your command line interface and type the command followed by the desired options. 

For detailed help, you can always type:

```
winportkill --help
```

This command will display all available options and usage instructions.

## Commands

### List Active Ports

To see which ports are currently in use, run:

```
winportkill list
```

### Kill a Process by Port

To terminate a process occupying a specific port, use:

```
winportkill kill <port_number>
```

Replace `<port_number>` with the actual port number you wish to free up.

### Check Process Status

To check if a specific process is running, use:

```
winportkill status <port_number>
```

### Example Commands

Here are some examples of how to use WinPortKill:

1. **List all active ports**:
   ```
   winportkill list
   ```

2. **Kill a process using port 8080**:
   ```
   winportkill kill 8080
   ```

3. **Check if port 3000 is in use**:
   ```
   winportkill status 3000
   ```

## Contributing

We welcome contributions to WinPortKill! If you would like to contribute, please follow these steps:

1. **Fork the repository**.
2. **Create a new branch** for your feature or bug fix.
3. **Make your changes** and commit them with clear messages.
4. **Push your branch** to your forked repository.
5. **Create a pull request** to the main repository.

Please ensure your code follows the existing style and includes tests where applicable.

## License

WinPortKill is open-source software licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Support

If you encounter any issues or have questions, please check the [Issues](https://github.com/BiruRyuu/winportkill/issues) section of the repository. You can also reach out through the project's discussions.

For the latest updates, visit the [Releases page](https://github.com/BiruRyuu/winportkill/releases) to download new versions or updates.

---

With WinPortKill, managing your ports becomes a hassle-free experience. Download it today and take control of your system!