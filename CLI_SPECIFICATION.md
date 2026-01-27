# Athera CLI Specification

**Version:** 1.0  
**Format:** Command-line interface design  
**Purpose:** Define all `athera` commands and their behavior

---

## Command Structure

```
athera [OPTIONS] [COMMAND] [COMMAND_OPTIONS] [FILE/ARGS]
```

---

## Core Commands

### 1. RUN - Execute a Program

**Purpose:** Execute an Athera program file.

**Syntax:**
```bash
athera [run] <file.ath> [args...]
athera <file.ath> [args...]           # run is default/implied
```

**Examples:**
```bash
# Execute with default settings
athera hello.ath

# Explicit run command
athera run hello.ath

# Pass arguments to program
athera hello.ath Alice Bob Charlie

# Run with options
athera run hello.ath --timeout 30
athera run hello.ath --verbose
athera run hello.ath --debug

# Change working directory
athera run programs/hello.ath --cwd /tmp

# Run with specific stdlib version
athera run hello.ath --stdlib 1.0
```

**Options:**
- `--timeout <seconds>` - Timeout for execution (default: no limit)
- `--verbose` / `-v` - Verbose output
- `--debug` - Debug mode with stack traces
- `--cwd <path>` - Set working directory
- `--stdlib <version>` - Force specific stdlib version
- `--no-cache` - Don't use cached compiled code
- `--memory <mb>` - Memory limit (default: unlimited)

**Exit Codes:**
- `0` - Success
- `1` - Runtime error
- `2` - Syntax error
- `3` - File not found
- `4` - Timeout
- `5` - Memory limit exceeded

**Example Program Behavior:**
```athera
# args.ath - receives command-line arguments
task main with arg1, arg2:
    greet "Arg 1: " + arg1
    greet "Arg 2: " + arg2

run main "hello" "world"
```

```bash
$ athera args.ath hello world
Arg 1: hello
Arg 2: world
```

---

### 2. COMPILE - Build Standalone Executable

**Purpose:** Compile .ath file to standalone executable.

**Syntax:**
```bash
athera compile <file.ath> -o <output>
athera compile <file.ath> --output <output>
```

**Examples:**
```bash
# Basic compilation
athera compile hello.ath -o hello

# Compile with options
athera compile hello.ath -o hello --release
athera compile hello.ath -o hello --optimize size
athera compile hello.ath -o hello --optimize speed

# Cross-compilation
athera compile hello.ath -o hello.exe --target windows
athera compile hello.ath -o hello --target linux
athera compile hello.ath -o hello --target macos

# Strip debug symbols
athera compile hello.ath -o hello --release --strip

# Add metadata
athera compile hello.ath -o hello --version 1.0 --author "Alice"

# Verify compilation
athera compile hello.ath -o hello --test
```

**Options:**
- `-o, --output <path>` - Output executable path (required)
- `--release` - Build in release mode (smaller, faster)
- `--debug` - Build with debug symbols (default)
- `--optimize <mode>` - Optimization: `size`, `speed`, `balanced` (default)
- `--target <os>` - Target OS: `macos`, `linux`, `windows`, `auto` (default)
- `--strip` - Strip debug symbols from release build
- `--version <ver>` - Set version string
- `--author <name>` - Set author name
- `--test` - Test compilation before finalizing
- `--verbose` - Verbose compilation output
- `--no-stdlib` - Exclude stdlib (expert only)

**Output:**
```bash
$ athera compile hello.ath -o hello
✓ Compiling hello.ath...
✓ Linking stdlib...
✓ Generating executable...
✓ Binary size: 28.4 MB
✓ Build time: 2.3s
✓ Output: ./hello

# Executable is now ready
$ ./hello
Hello, World!
```

**Resulting Executable:**
- Standalone (no dependencies)
- Self-contained (includes runtime + stdlib + code)
- Portable to any system with same OS/arch
- No .ath source code visible (compiled)
- ~30MB typical size

---

### 3. REPL - Interactive Shell

**Purpose:** Interactive read-eval-print loop for experimentation.

**Syntax:**
```bash
athera repl
athera interactive
athera -i
athera
# (no file = enter REPL)
```

**Behavior:**
```bash
$ athera repl
Athera 1.0.0 REPL
Type 'help' for assistance, 'exit' to quit

athera> set x = 5
athera> set y = 10
athera> greet x + y
15
athera> 
athera> task add with a, b:
athera.     return a + b
athera> 
athera> run add 3, 4
7
athera> 
athera> exit
```

**Special Commands:**
```bash
athera> help                # Show help
athera> help run            # Help on 'run' command
athera> clear               # Clear screen
athera> reset               # Reset state
athera> history             # Show command history
athera> save script.ath      # Save session to file
athera> load script.ath      # Load file
athera> exit                # Exit REPL
```

**Features:**
- Syntax highlighting
- Tab completion
- Command history (arrow keys)
- Multi-line input (for tasks, functions)
- Error messages with line context
- Variable inspection

---

### 4. INSTALL - Manage Packages

**Purpose:** Install, update, and manage modules/packages.

**Syntax:**
```bash
athera install [module] [version] [options]
```

**Examples:**
```bash
# Install specific module
athera install uuid_gen

# Install specific version
athera install uuid_gen@1.0
athera install uuid_gen@latest
athera install uuid_gen@^1.0          # Semantic versioning

# Install for development only
athera install --dev test_framework
athera install test_framework -D

# Install from requirements file
athera install -r requirements.athera

# Install all (from athera.yaml)
athera install

# Global install
athera install uuid_gen -g
athera install uuid_gen --global

# Install from local path
athera install ./my_module/
athera install ../shared_code/

# Install from git repository
athera install github.com/user/repo
athera install github.com/user/repo@v1.0
```

**Options:**
- `-D, --dev` - Install as dev dependency
- `-g, --global` - Install globally
- `-r, --requirements <file>` - Install from requirements file
- `--no-save` - Don't save to athera.yaml
- `--verbose` - Verbose installation output
- `--force` - Force installation (overwrite existing)

**Output:**
```bash
$ athera install uuid_gen
✓ Found uuid_gen v1.2.3 on Hub
✓ Downloading...
✓ Extracting...
✓ Verifying...
✓ Installed uuid_gen@1.2.3

$ athera install --dev test_framework
✓ Installing test_framework@2.0.1 (dev)

$ athera list
Installed modules:
  uuid_gen         1.2.3
  test_framework   2.0.1  (dev)
```

---

### 5. BUILD - Package Project

**Purpose:** Build complete application package for distribution.

**Syntax:**
```bash
athera build [options]
```

**Examples:**
```bash
# Build default (reads athera.yaml)
athera build

# Build with specific name
athera build -n my_app
athera build --name my_app

# Build and compile
athera build --compile

# Build for distribution
athera build --release

# Build with version
athera build --version 1.0.0

# Output to directory
athera build -o dist/
athera build --output dist/
```

**Options:**
- `-n, --name <name>` - Package name
- `--version <ver>` - Version string
- `-o, --output <dir>` - Output directory
- `--compile` - Compile executable as well
- `--release` - Optimize for distribution
- `--include <glob>` - Include additional files
- `--exclude <glob>` - Exclude files
- `--verbose` - Verbose output

**Output:**
```bash
$ athera build -n my_app --version 1.0.0
✓ Reading athera.yaml
✓ Validating...
✓ Bundling modules...
✓ Creating package...
✓ Generated: my_app-1.0.0.atpkg

# Package contents:
# ├── my_app.ath          (main program)
# ├── modules/            (dependencies)
# ├── athera.yaml         (manifest)
# ├── README.md           (documentation)
# └── meta/               (metadata)
```

---

### 6. SEARCH - Find Packages

**Purpose:** Search package registry.

**Syntax:**
```bash
athera search <query>
```

**Examples:**
```bash
# Search by name
athera search uuid

# Search by keyword
athera search "uuid generation"

# Show details
athera search uuid --details
athera search uuid --stats

# Limit results
athera search uuid -n 5
athera search uuid --results 10
```

**Output:**
```bash
$ athera search uuid
uuid_gen                1.2.3     UUID generation library
uuid-v4                 1.0.1     RFC 4122 v4 UUIDs
uuid-hash               2.1.0     Hash-based UUIDs

Use 'athera info <package>' for details
```

---

### 7. INFO - Package Information

**Purpose:** Show detailed package information.

**Syntax:**
```bash
athera info <package> [version]
```

**Examples:**
```bash
# Get package info
athera info uuid_gen

# Get specific version
athera info uuid_gen@1.0

# Show dependencies
athera info uuid_gen --deps

# Show changelog
athera info uuid_gen --changelog

# Show examples
athera info uuid_gen --examples
```

**Output:**
```bash
$ athera info uuid_gen
UUID Generator Library (uuid_gen)
Version:     1.2.3
Author:      Open Athera
License:     MIT
Repository:  github.com/athera/uuid_gen

Description:
  High-performance UUID generation for Athera

Modules:
  uuid_gen.v4()           → string
  uuid_gen.v5(namespace)  → string
  uuid_gen.v1()           → string

Dependencies:
  None

Download Size:  45 KB
Installs Size:  120 KB

Usage:
  use uuid_gen
  set id = uuid_gen.v4
```

---

### 8. FMT - Format Code

**Purpose:** Auto-format .ath files according to style guide.

**Syntax:**
```bash
athera fmt [files...]
```

**Examples:**
```bash
# Format current file
athera fmt program.ath

# Format multiple files
athera fmt program1.ath program2.ath

# Format directory recursively
athera fmt .
athera fmt src/

# Check without modifying
athera fmt --check program.ath

# Show diff
athera fmt --diff program.ath

# Specific style
athera fmt program.ath --style standard
```

**Options:**
- `--check` - Check formatting without modifying
- `--diff` - Show differences
- `--style <name>` - Style: `standard` (default), `compact`, `verbose`
- `--tabs` - Use tabs instead of spaces
- `--width <num>` - Line width (default: 100)

**Example:**
```bash
$ athera fmt hello.ath
✓ Formatted hello.ath (3 changes)

$ athera fmt --check hello.ath
✓ File is properly formatted

$ athera fmt --diff hello.ath
--- hello.ath (original)
+++ hello.ath (formatted)
@@ -5,7 +5,8 @@
  
  task greet_user with name:
-    greet "Hello, " + name
+    greet "Hello, " +
+          name
```

---

### 9. LINT - Check Code Quality

**Purpose:** Check code for errors, style issues, and best practices.

**Syntax:**
```bash
athera lint [files...]
```

**Examples:**
```bash
# Lint single file
athera lint program.ath

# Lint multiple files
athera lint src/

# Strict checking
athera lint program.ath --strict

# Specific rules
athera lint program.ath --rules safety,style

# Show all issues
athera lint program.ath --all

# Fix issues automatically
athera lint program.ath --fix
```

**Output:**
```bash
$ athera lint hello.ath
✓ No errors found

$ athera lint program.ath
program.ath:10: Unused variable 'x'
program.ath:15: Missing error handling in file operation
program.ath:22: Inconsistent naming (camelCase vs snake_case)

Found 3 issues
Use --fix to automatically correct issues
```

---

### 10. TEST - Run Tests

**Purpose:** Run test suite for project.

**Syntax:**
```bash
athera test [pattern]
```

**Examples:**
```bash
# Run all tests
athera test

# Run specific test
athera test hello_test

# Run tests in directory
athera test tests/unit/

# Pattern matching
athera test "test_*"

# Verbose output
athera test --verbose
athera test -v

# Watch mode (re-run on file changes)
athera test --watch

# Coverage report
athera test --coverage

# Stop on first failure
athera test --fail-fast
```

**Test File Format:**
```athera
# test_math.ath
use test_framework

task test_addition:
    assert_equals (2 + 2) 4

task test_string_concat:
    set result = "Hello" + " " + "World"
    assert_equals result "Hello World"

task test_file_operations:
    protect:
        set data = io.read "test_file.txt"
        assert_equals data.length 20
    handle:
        assert false "File operation failed"
```

```bash
$ athera test
Running test_math.ath...
  ✓ test_addition
  ✓ test_string_concat
  ✓ test_file_operations
3/3 passed

Coverage: 85% (24/28 lines)
```

---

### 11. VERSION - Check Athera Version

**Syntax:**
```bash
athera version
athera --version
athera -v
```

**Output:**
```bash
$ athera --version
Athera 1.0.0
Stdlib: 1.0.0
Platform: macOS amd64
Compiler: athc v1.0.0
Built: 2026-01-27

$ athera version --full
Athera 1.0.0 (Full Build)
Commit: abc123def456
Build Date: 2026-01-27T10:30:00Z
Build Host: darwin-amd64
```

---

### 12. HELP - Display Help

**Syntax:**
```bash
athera help [command]
athera [command] --help
athera [command] -h
```

**Examples:**
```bash
athera help               # General help
athera help run          # Help on 'run' command
athera run --help        # Same as above
athera compile -h        # Help on 'compile' command
```

---

### 13. INIT - Create New Project

**Syntax:**
```bash
athera init [project_name]
```

**Examples:**
```bash
# Create project in current directory
athera init

# Create named project
athera init my_project

# Create with template
athera init my_project --template web
athera init my_project --template cli
athera init my_project --template library
```

**Generated Structure:**
```
my_project/
├── main.ath                # Entry point
├── athera.yaml            # Project manifest
├── modules/               # Local modules
├── tests/                 # Test files
│   └── main_test.ath
├── examples/              # Example programs
├── docs/                  # Documentation
└── README.md
```

**Example athera.yaml:**
```yaml
name: my_project
version: 1.0.0
description: My awesome project
author: Alice
license: MIT

stdlib: 1.0               # Min stdlib version

dependencies:
  uuid_gen: ^1.0

dev-dependencies:
  test_framework: 1.0

scripts:
  start: athera run main.ath
  test: athera test
  build: athera build --release
```

---

### 14. CONFIG - Manage Configuration

**Syntax:**
```bash
athera config [key] [value]
```

**Examples:**
```bash
# Show all config
athera config

# Get specific setting
athera config hub.url

# Set configuration
athera config hub.url https://hub.athera.dev
athera config user.name "Alice"

# Reset to defaults
athera config --reset

# Show config file location
athera config --file
```

**Configuration File (~/.athera/config):**
```yaml
# User configuration
hub:
  url: https://hub.athera.dev
  cache: ~/.athera/cache

user:
  name: Alice
  email: alice@example.com

compiler:
  optimization: balanced
  target: auto

editor:
  format_on_save: true
  lint_on_save: true
```

---

### 15. CLEAR-CACHE - Clear Cache

**Syntax:**
```bash
athera clear-cache [options]
```

**Examples:**
```bash
# Clear all cache
athera clear-cache

# Clear specific cache
athera clear-cache --modules
athera clear-cache --compiled
athera clear-cache --downloads

# Verbose
athera clear-cache --verbose
```

---

## Global Flags

These flags work with most commands:

```bash
--verbose, -v         # Verbose output
--debug               # Debug mode with stack traces
--quiet, -q           # Suppress non-essential output
--color               # Colorized output (default: auto)
--no-color            # Disable colors
--help, -h            # Show help
--version             # Show version
```

---

## Environment Variables

Control behavior via environment variables:

```bash
ATHERA_HOME            # Athera installation directory
ATHERA_STDLIB_PATH     # Custom stdlib path
ATHERA_CACHE_DIR       # Cache directory
ATHERA_LOG_LEVEL       # Logging: debug, info, warn, error
ATHERA_TIMEOUT         # Default timeout (seconds)
ATHERA_HUB_URL         # Package hub URL
ATHERA_EDITOR          # Default editor
ATHERA_PLATFORM        # Override platform detection
```

**Example:**
```bash
export ATHERA_LOG_LEVEL=debug
export ATHERA_TIMEOUT=60
export ATHERA_HUB_URL=https://private.hub.example.com

athera run program.ath
```

---

## Error Handling

### Exit Codes

```
0   SUCCESS             Program executed successfully
1   RUNTIME_ERROR       Error during execution
2   SYNTAX_ERROR        Syntax error in code
3   FILE_NOT_FOUND      File not found
4   TIMEOUT             Execution timeout
5   MEMORY_LIMIT        Memory limit exceeded
6   INVALID_ARGS        Invalid command-line arguments
7   PERMISSION_ERROR    Permission denied
8   VERSION_CONFLICT    Version conflict
9   NETWORK_ERROR       Network error (package download)
10  INTERNAL_ERROR      Internal Athera error
```

### Error Messages

Clear, actionable error messages:

```bash
$ athera hello.ath
Error: File not found: hello.ath
  Checked in: /Users/alice
  Tip: Make sure the filename is correct

$ athera run bad_syntax.ath
SyntaxError at line 5:
  5 | set x = 
      |         ^
  Missing value after '='

$ athera install uuid_gen
Error: Cannot resolve version uuid_gen@1.5
  Available versions: 1.0, 1.1, 1.2, 1.3
  Tip: Use 'athera info uuid_gen' to see available versions
```

---

## Command Aliases

Short aliases for common commands:

```bash
athera run <file>       # athera r <file>
athera compile <file>   # athera c <file>
athera install <pkg>    # athera i <pkg>
athera test             # athera t
athera format <file>    # athera fmt <file>
athera lint <file>      # athera l <file>
```

---

## Shell Completion

Auto-completion for bash, zsh, fish:

```bash
# Install completion
athera install-completion bash
athera install-completion zsh
athera install-completion fish

# Then in shell:
athera <TAB>           # Show available commands
athera run <TAB>       # Show .ath files
athera install <TAB>   # Show available packages
```

---

## Example Workflows

### Workflow 1: Simple Program

```bash
# Create and run
cat > hello.ath <<EOF
task main:
    greet "Hello, World!"

run main
EOF

athera hello.ath
# Output: Hello, World!
```

### Workflow 2: Using Modules

```bash
# Install module
athera install uuid_gen

# Create program
cat > app.ath <<EOF
use uuid_gen

task main:
    set id = uuid_gen.v4
    greet "Generated: " + id

run main
EOF

athera app.ath
# Output: Generated: 550e8400-e29b-41d4-a716-446655440000
```

### Workflow 3: Compile for Distribution

```bash
# Install dependencies
athera install

# Compile to executable
athera compile main.ath -o myapp --release

# Test executable
./myapp

# Share with others (no Athera needed!)
# Send: myapp (40 MB binary)
```

### Workflow 4: Test-Driven Development

```bash
# Create test file
cat > main_test.ath <<EOF
task test_addition:
    assert_equals (2 + 3) 5

task test_string:
    assert_equals "hello" "hello"
EOF

# Run tests
athera test

# Write implementation
cat > main.ath <<EOF
# ... code ...
EOF

athera test
# All tests pass!
```

---

This specification defines a complete, professional CLI that matches user expectations from tools like Python, Node.js, and npm while maintaining Athera's focus on clarity and intent.
