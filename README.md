# WinPortKill

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go) ![License](https://img.shields.io/badge/License-MIT-green?style=flat-square) ![Platform](https://img.shields.io/badge/Platform-Windows-blue?style=flat-square)

**WinPortKill** is a simple and powerful tool written in Go that helps you identify and free up occupied ports on Windows. When a port you need is in use, this program lists the processes using it and allows you to terminate them easily.

## Features
- **Quick Detection**: Displays a list of all processes using the specified port.
- **Easy Termination**: Terminate related processes with a single command.
- **Lightweight & Fast**: Built with Go, no complex dependencies required.
- **User-Friendly**: Suitable for both advanced users and beginners.

![image](https://github.com/user-attachments/assets/ca5850d1-8f00-447b-8d79-4486b2f4b729)


## Installation

### Prerequisites
- Go version 1.18 or higher
- Windows operating system

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/winportkill.git
   ```
2. Navigate to the project directory:
   ```bash
   cd winportkill
   ```
3. Build the program:
   ```bash
   go build
   ```
4. The executable `winportkill.exe` is ready to use!

### Install via Binary
You can download the pre-built executable from the [Releases](https://github.com/imrostami/winportkill/releases) section.

## Usage
1. Run the program:
   ```bash
   ./winportkill.exe
   ```
2. Enter the port number (e.g., `8080`).
3. A list of processes using the port will be displayed (including PID and process name).
4. To terminate a process, enter its PID or use the interactive options provided.

### Example
```bash
$ ./winportkill.exe
Enter port number: 8080
Processes using port 8080:
- PID: 1234, Name: example.exe
- PID: 5678, Name: anotherapp.exe
Enter PID to kill (or 0 to exit): 1234
Process 1234 terminated successfully.
```

## Important Notes
- **Run as Administrator**: Terminating some processes may require running the program with Administrator privileges.
- **Caution**: Terminating system processes may cause system instability. Proceed with care.

## Contributing
If you have ideas for improvements or encounter issues, please open an [Issue](https://github.com/yourusername/winportkill/issues) or submit a Pull Request.

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add YourFeature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License
This project is licensed under the GNU General Public License v3.0 License - see the [LICENSE](LICENSE) file for details.
