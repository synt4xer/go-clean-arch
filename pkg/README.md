# `/pkg`

The `/pkg` directory in a Go project is a conventional location for storing packages that can be imported by other projects or applications. It contains reusable code that can be shared across multiple applications or services. The contents of this directory are intended to be imported and used by other codebases, promoting code reusability and modularity.

## Structure

Typically, the `/pkg` directory is organized into subdirectories, each representing a package. Here's an example structure:
````
├── /pkg/
│ ├── /logger/
│ │ ├── logger.go
│ └── /utils/
│ ├── utils.go
````

### `/logger`
The `/logger` package could be responsible for logging activities across the application.

- `logger.go`: Defines a logging utility that can be used throughout the application for consistent and structured logging.

### `/utils`
The `/utils` package generally contains utility functions that are commonly used throughout the application but don't necessarily belong to a specific domain or service.

- `utils.go`: Provides various helper functions, like string manipulation, date formatting, etc.

## Usage

To use the packages in the `/pkg` directory, you simply import them in your Go files like so:

```go
import (
    "your_project/pkg/logger"
    "your_project/pkg/utils"
)
```
By keeping reusable code in `/pkg`, you maintain a clean and organized project structure, making it easier to manage and share code across different projects or teams.