# Medicare Framework Service

This is the backend service for the Medicare Framework. Follow the steps below to set up, run, and build the project.

## Prerequisites

- **Go**: Version 1.25.5 ([Download Here](https://go.dev/dl/))
- **Google Wire**: Required for dependency injection.

If you haven't installed `wire` yet, install it via:

```bash
go install github.com/google/wire/cmd/wire@latest
```

## Getting Started

Follow these steps to get the application up and running on your local machine.

### 1\. Clone the Repository

```bash
git clone https://github.com/onnytra/first-go.git <your-project-name>
cd <your-project-name>
```

### 2\. Environment Setup

Copy the example environment file to create your local `.env` file. This file contains your database configuration and other secrets.

```bash
# For Linux/macOS
cp .env.example .env

# For Windows
copy .env.example .env
```

> **Note:** Open the `.env` file and adjust the configuration values (database credentials, ports, etc.) according to your local environment.

### 3\. Rename Project (Important)

If you are using this repository as a template, you must change the module name from the default `github.com/onnytra/first-go` to your own project name.

**Step A: Change module name in `go.mod`**

```bash
# Replace <your-new-module-path> with your actual project name (e.g., [github.com/username/project](https://github.com/username/project))
go mod edit -module <your-new-module-path>
```

**Step B: Update imports in all files**
You need to replace the old import path with your new one in all `.go` files.

_Option 1: Using VS Code (Recommended)_

1.  Press `Ctrl + Shift + H` (Search & Replace).
2.  Find: `github.com/onnytra/first-go`
3.  Replace with: `<your-new-module-path>`
4.  Click **Replace All**.

_Option 2: Using Terminal (Linux/macOS)_

```bash
grep -rRl "[github.com/onnytra/first-go](https://github.com/onnytra/first-go)" . | xargs sed -i '' 's|[github.com/onnytra/first-go](https://github.com/onnytra/first-go)|<your-new-module-path>|g'
```

### 4\. Install Dependencies

After renaming the project, download and verify all required modules:

```bash
go mod tidy
```

### 5\. Dependency Injection (Wire)

This project uses **Google Wire** for dependency injection. You must generate the wire code inside the `cmd` directory before running the app.

```bash
cd cmd
wire
cd ..
```

### 6\. Run the Application

You can run the application from the root directory:

```bash
go run ./cmd
```

Or, if you are strictly inside the `cmd` folder:

```bash
go run .
```

## Building the Application

For production deployment or creating an executable binary (Docker or standard build), please refer to the specific documentation located in the build directory.

- **Build Guide**: [`build/README.md`]
