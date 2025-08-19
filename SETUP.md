# klo Development Setup

## Prerequisites

### Install Go

1. **Download Go**:
   - Visit https://golang.org/dl/
   - Download the latest stable version for Windows (`.msi` installer)

2. **Install Go**:
   - Run the downloaded `.msi` file
   - Follow the installation wizard
   - Default installation path: `C:\Program Files\Go`

3. **Verify Installation**:
   ```powershell
   go version
   ```
   This should output something like: `go version go1.21.0 windows/amd64`

4. **Set GOPATH (if needed)**:
   ```powershell
   # Check current GOPATH
   go env GOPATH
   
   # Usually defaults to: C:\Users\<username>\go
   ```

### Setup klo Development Environment

1. **Clone/Navigate to klo directory**:
   ```powershell
   cd "c:\Users\mesut\Documents\stuff\sourceLocal\klolang"
   ```

2. **Initialize Go module**:
   ```powershell
   go mod init klo
   ```

3. **Download dependencies**:
   ```powershell
   go mod tidy
   ```

4. **Build klo**:
   ```powershell
   go build -o klo.exe
   ```

5. **Test with example**:
   ```powershell
   .\klo.exe examples\hello.klo
   ```

### Add klo to PATH (Optional)

**Windows:**
```powershell
# Option 1: Create a local bin directory (Recommended)
mkdir C:\tools\bin
copy klo.exe C:\tools\bin\klo.exe

# Add to PATH permanently
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\tools\bin", "User")

# Restart PowerShell, then test
klo version
```

**Alternative: System directory (requires admin)**
```powershell
# Copy to Windows\System32 (requires administrator)
copy klo.exe C:\Windows\System32\
```

**Linux/macOS:**
```bash
sudo mv klo /usr/local/bin/klo
```

## Development Workflow

### Building
```powershell
go build -o klo.exe
```

### Running Tests
```powershell
go test ./...
```

### Running Examples
```powershell
.\klo.exe examples\hello.klo
.\klo.exe examples\calculator.klo
.\klo.exe examples\fibonacci.klo
```

### Transpile Only
```powershell
.\klo.exe --transpile examples\hello.klo
```

### Verbose Output
```powershell
.\klo.exe --verbose examples\hello.klo
```

## Troubleshooting

### "go is not recognized"
- Go is not installed or not in PATH
- Reinstall Go from https://golang.org/dl/
- Restart PowerShell after installation

### Module errors
```powershell
go mod download
go mod tidy
```

### Build errors
- Check Go version: `go version` (requires 1.21+)
- Clear module cache: `go clean -modcache`

### Runtime errors
- Ensure Go is installed and accessible
- Check that .klo file exists and has correct syntax

## IDE Setup

### VS Code
1. Install Go extension
2. Open workspace folder
3. VS Code will automatically detect Go module

### Recommended Extensions
- Go (Google)
- Go Outline
- Go Test Explorer
