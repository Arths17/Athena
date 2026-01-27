# Athera Distribution & Installation Architecture

**Version:** 1.0  
**Status:** Design Specification  
**Target:** Professional distribution like Python, Node.js, or Go

---

## Vision

Athera should be installable and usable exactly like Python or Node.js:

```bash
# After installation, users type:
athera my_program.ath

# Or run with explicit command:
athera run my_program.ath
```

**Key Design Principles:**
- ✅ Single command execution (no boilerplate)
- ✅ Global availability after installation
- ✅ .ath files are portable and shareable
- ✅ Built-in standard library for immediate productivity
- ✅ Optional advanced modules for professionals
- ✅ Native installers for all major OSes
- ✅ No Python dependency for end users

---

## Architecture Overview

### 1. Distribution Strategy

```
                    ┌─────────────────────┐
                    │   athera.dev        │
                    │  (Official Website) │
                    └──────────┬──────────┘
                               │
                ┌──────────────┼──────────────┐
                │              │              │
         ┌──────▼────┐  ┌──────▼────┐  ┌────▼──────┐
         │   macOS   │  │ Windows   │  │   Linux   │
         │ Installer │  │ Installer │  │ Installer │
         └──────┬────┘  └──────┬────┘  └────┬──────┘
                │              │              │
         ┌──────▼─────────────────────────────▼────┐
         │    Athera Runtime + Stdlib (Compiled)   │
         │         /usr/local/bin/athera           │
         │         (Mac/Linux) or C:\Program...    │
         └──────────────────────────────────────────┘
                           │
            ┌──────────────┴───────────────┐
            │                              │
    ┌───────▼────────┐          ┌─────────▼────────┐
    │  Standard Lib  │          │ Athera CLI Tool  │
    │  (Built-in)    │          │  (Run, Compile)  │
    └────────────────┘          └──────────────────┘
```

### 2. Components

#### A. Athera Runtime (`athera` binary)
- **What it is:** Compiled executable (Go, Rust, or optimized Python)
- **Where it lives:**
  - Mac: `/usr/local/bin/athera`
  - Linux: `/usr/bin/athera` or `/usr/local/bin/athera`
  - Windows: `C:\Program Files\Athera\athera.exe` (in PATH)
- **Size:** ~20-50MB (compiled + embedded stdlib)
- **Dependencies:** Zero (fully self-contained)

#### B. Standard Library
- **Embedded in runtime** (no external files needed)
- **Organized modules:**
  - `io` - file operations, streams, input/output
  - `text` - string manipulation, formatting, parsing
  - `math` - arithmetic, trigonometry, statistics
  - `list` - array operations, sorting, filtering
  - `dict` - map operations, key manipulation
  - `time` - dates, times, scheduling
  - `network` - HTTP requests, sockets (advanced)
  - `json` - JSON parsing and generation
  - `path` - file path operations, directory handling

#### C. CLI Tool (`athera` command)
```
athera [COMMAND] [OPTIONS] [FILE]

Commands:
  run         Execute an .ath file (default)
  compile     Compile .ath to standalone executable
  repl        Interactive REPL session
  install     Install third-party modules
  build       Package application for distribution
  version     Show version info
  help        Show help
```

#### D. Package Manager (Athera Hub)
- **Registry:** https://hub.athera.dev
- **Optional modules:** For advanced users
- **Similar to:** npm, pip, but lightweight
- **Used with:** `athera install math_lib`

---

## Phase 1: MVP Distribution (Months 1-3)

### Goal
Basic installation on macOS and Linux. Users can run `.ath` files globally.

### Deliverables

#### 1. macOS Installer (.pkg)
```
athera-1.0.pkg (25MB)
├── Installs to: /usr/local/bin/athera
├── Installs stdlib (embedded)
├── Creates uninstaller
└── Signs with developer certificate
```

**Installation:** Download from athera.dev → Double-click → Follow prompts

#### 2. Linux Installer (.deb for Ubuntu/Debian)
```
athera_1.0_amd64.deb (25MB)
├── sudo apt install ./athera_1.0_amd64.deb
├── Installs to: /usr/bin/athera
└── Includes stdlib
```

**Also provide:** Arch package, Fedora RPM, generic tarball

#### 3. Windows Installer (.msi or .exe)
```
athera-1.0-installer.msi (25MB)
├── GUI installer wizard
├── Installs to: C:\Program Files\Athera\
├── Adds to PATH
└── Includes uninstaller
```

#### 4. CLI Implementation
```bash
# Phase 1 - Basic execution
athera program.ath              # Run with default settings
athera run program.ath          # Explicit run command
athera --version                # Show version
athera --help                   # Show help

# Example:
$ athera hello.ath
Hello, Athera!
```

#### 5. Standard Library (Phase 1)
Minimal, essential modules:
- `io.read file_path` - Read files
- `io.write file_path data` - Write files
- `io.print text` - Output (same as greet)
- `text.split text separator` - String operations
- `text.join list separator` - String join
- `list.length list` - Array size
- `list.append list item` - Add to array
- `dict.keys dict` - Get dictionary keys
- `math.add a b` - Basic arithmetic (already built-in)

**Access pattern:**
```athera
# Using standard library
use io
use text

set data = io.read "file.txt"
set lines = text.split data "\n"
repeat each line in lines:
    greet line
```

---

## Phase 2: Advanced Installation (Months 4-6)

### Goal
Full-featured distribution, package manager, compilation support.

### Deliverables

#### 1. Package Manager (Athera Hub)
```bash
athera install uuid_gen         # Install module
athera install express --dev    # Install for development
athera search database          # Search registry
athera list                      # List installed modules
```

**Registry structure:**
```
hub.athera.dev/
├── packages/
│   ├── uuid_gen/
│   │   ├── 1.0/
│   │   │   ├── module.ath
│   │   │   ├── doc.md
│   │   │   └── examples/
│   ├── express/
│   ├── database/
```

#### 2. Compilation to Executable
```bash
# Compile .ath to standalone executable
athera compile hello.ath -o hello

# Now users can run:
./hello
# Without needing Athera installed!
```

**How it works:**
- Athera compiler bundles interpreter + stdlib + code into single binary
- Standalone executable ~30MB but fully self-contained
- No .ath source code visible in compiled version

#### 3. Extended Standard Library
Add Phase 1 modules:
- `time` - Dates, times, scheduling
- `json` - JSON parsing/generation
- `path` - File path operations
- `math` - Full math library (sin, cos, log, etc.)
- `random` - Random number generation
- `regex` - Regular expressions (optional, advanced)

#### 4. REPL (Interactive Mode)
```bash
$ athera repl
Athera 1.0 REPL
Type 'help' for assistance, 'exit' to quit

athera> set x = 5
athera> set y = 10
athera> greet "Sum: " + (x + y)
Sum: 15
athera> exit
```

#### 5. Build System
```bash
# For complex projects with multiple files
athera build -o my_app

# Creates:
# - Compiled executable
# - my_app.ath (if source distribution desired)
# - Metadata file for sharing
```

---

## Phase 3: Enterprise & Advanced (Months 7-12)

### Goal
Full-featured professional distribution ecosystem.

### Deliverables

#### 1. Docker Support
```dockerfile
FROM athera:latest
COPY app.ath /app/
CMD ["athera", "run", "/app/app.ath"]
```

#### 2. Version Management
```bash
athera versions                 # List installed versions
athera use 1.1                  # Switch version
athera use latest               # Use latest version
```

#### 3. Project Templates
```bash
athera init my_project          # Create new project
# Creates:
# ├── main.ath
# ├── modules/
# ├── examples/
# ├── tests/
# └── athera.lock (dependency lock file)
```

#### 4. Dependency Management
```
athera.yaml format:
───────────────────
name: my_app
version: 1.0
description: My awesome Athera app

stdlib: 1.0         # Requires stdlib version 1.0+

dependencies:
  uuid_gen: ^1.0    # Semantic versioning
  express: ~2.5
  
dev-dependencies:
  test_framework: 1.0
```

#### 5. Testing Framework
```bash
athera test                     # Run all tests
athera test tests/unit/         # Run specific tests
```

#### 6. Code Formatter & Linter
```bash
athera fmt my_program.ath       # Auto-format code
athera lint my_program.ath      # Check style
```

---

## Platform-Specific Details

### macOS Distribution

**Installation Methods:**
1. **Direct download:** athera-1.0.pkg from athera.dev
2. **Homebrew:** `brew install athera` (requires Homebrew formula)
3. **Direct binary:** Download athera binary and `sudo mv` to `/usr/local/bin/`

**Installer Contents:**
```
athera-1.0.pkg
├── athera binary (compiled, signed, notarized)
├── Stdlib (embedded)
├── man pages
├── launchd script (if needed for services)
└── Uninstaller
```

**Signing & Notarization:**
- Sign with Apple Developer ID (prevents "Unknown developer" warning)
- Notarize with Apple (required for macOS 10.15+)
- Code: `codesign -s "Developer ID Application" athera`

**Verification:**
```bash
$ which athera
/usr/local/bin/athera

$ athera --version
Athera 1.0.0 (macOS amd64)
```

---

### Linux Distribution

**Installation Methods:**

**1. Ubuntu/Debian (.deb):**
```bash
# Download
wget https://athera.dev/releases/athera_1.0_amd64.deb

# Install
sudo apt install ./athera_1.0_amd64.deb

# Verify
athera --version
```

**2. Fedora/RHEL (.rpm):**
```bash
# Download
wget https://athera.dev/releases/athera-1.0-1.x86_64.rpm

# Install
sudo dnf install ./athera-1.0-1.x86_64.rpm
```

**3. Arch Linux (AUR):**
```bash
yay -S athera
```

**4. Generic Linux (tarball):**
```bash
# Download
wget https://athera.dev/releases/athera-1.0-linux-amd64.tar.gz

# Install
tar -xzf athera-1.0-linux-amd64.tar.gz
sudo mv athera /usr/local/bin/

# Verify
athera --version
```

**Man Page:**
```bash
man athera
```

---

### Windows Distribution

**Installation Methods:**

**1. MSI Installer (Recommended):**
- Download: athera-1.0-installer.msi
- Double-click
- Follow GUI wizard
- Adds athera.exe to PATH
- Includes "Add/Remove Programs" entry

**2. Standalone Executable:**
- Download: athera-1.0.exe
- `athera-1.0.exe install` to install
- Or run directly from any location (temporary)

**3. Chocolatey Package:**
```powershell
choco install athera
```

**4. Windows Package Manager:**
```powershell
winget install Athera
```

**Verification:**
```powershell
# PowerShell or CMD
athera --version
```

**PATH Configuration:**
- Installer automatically adds `C:\Program Files\Athera\` to PATH
- Works in PowerShell, CMD, Git Bash, WSL

---

## Standard Library Design

### Module Organization

```
stdlib/
├── io/
│   ├── read(path)           → string
│   ├── write(path, data)    → void
│   ├── exists(path)         → boolean
│   ├── delete(path)         → void
│   ├── list_files(dir)      → list
│   └── create_dir(path)     → void
│
├── text/
│   ├── split(text, sep)     → list
│   ├── join(list, sep)      → string
│   ├── uppercase(text)      → string
│   ├── lowercase(text)      → string
│   ├── trim(text)           → string
│   ├── replace(text, old, new) → string
│   ├── contains(text, sub)  → boolean
│   ├── length(text)         → number
│   └── format(template, ...) → string
│
├── math/
│   ├── add(a, b)            → number
│   ├── subtract(a, b)       → number
│   ├── multiply(a, b)       → number
│   ├── divide(a, b)         → number
│   ├── power(base, exp)     → number
│   ├── sqrt(n)              → number
│   ├── sin(n)               → number
│   ├── cos(n)               → number
│   ├── floor(n)             → number
│   ├── ceil(n)              → number
│   ├── round(n)             → number
│   └── random(max)          → number
│
├── list/
│   ├── length(list)         → number
│   ├── append(list, item)   → void
│   ├── remove(list, index)  → void
│   ├── get(list, index)     → any
│   ├── sort(list)           → list
│   ├── reverse(list)        → list
│   ├── contains(list, item) → boolean
│   ├── join(list, sep)      → string
│   ├── map(list, task)      → list
│   └── filter(list, task)   → list
│
├── dict/
│   ├── keys(dict)           → list
│   ├── values(dict)         → list
│   ├── get(dict, key)       → any
│   ├── set(dict, key, val)  → void
│   ├── delete(dict, key)    → void
│   ├── exists(dict, key)    → boolean
│   ├── length(dict)         → number
│   └── merge(dict1, dict2)  → dict
│
├── time/
│   ├── now()                → timestamp
│   ├── sleep(seconds)       → void
│   ├── format_date(ts, fmt) → string
│   ├── parse_date(str, fmt) → timestamp
│   └── get_timestamp()      → number
│
├── json/
│   ├── parse(text)          → dict/list
│   ├── stringify(obj)       → string
│   ├── pretty_print(obj)    → string
│   └── validate(text)       → boolean
│
└── path/
    ├── join(...)            → path
    ├── dirname(path)        → path
    ├── basename(path)       → string
    ├── extension(path)      → string
    ├── absolute(path)       → path
    ├── exists(path)         → boolean
    ├── is_file(path)        → boolean
    ├── is_dir(path)         → boolean
    └── home()               → path
```

### Usage Pattern

```athera
# Import modules
use io
use text
use json

# Use functions with dot notation
set data = io.read "config.json"
set config = json.parse data

set name = config.name
greet "Hello, " + text.uppercase name

# Or with property access
check config.enabled:
    greet "System is enabled"
```

---

## CLI Specification Summary

### Basic Commands

```bash
# Run program
athera program.ath
athera run program.ath

# Check version
athera --version
athera -v

# Interactive shell
athera repl
athera interactive

# Help
athera help
athera help run
athera --help
```

### Advanced Commands

```bash
# Compile to executable
athera compile program.ath -o program
athera compile program.ath --output program

# Package management
athera install module_name
athera install module@1.0
athera install --dev test_framework
athera search database
athera uninstall module_name
athera list                    # List installed modules
athera update                  # Update modules

# Project management
athera init my_project
athera build -o my_app

# Code quality
athera fmt program.ath         # Format code
athera lint program.ath        # Lint code
athera test                    # Run tests
athera test tests/unit/

# Version management
athera versions
athera use 1.1
athera use latest

# Information
athera info
athera config
athera clear-cache
```

---

## File Distribution & Sharing

### Sharing a Single .ath File

**Before Athera Installation:**
- User downloads `athera-installer.pkg` (one-time)
- Installs to PATH

**Sharing Program:**
- Send single `.ath` file
- Recipient runs: `athera my_program.ath`
- No additional setup needed!

**Example:**
```bash
# Developer packages app
athera compile my_app.ath -o my_app_macos

# Send executable to user who doesn't have Athera
./my_app_macos
# Works without Athera installed!
```

### Dependency Management

**If .ath uses external modules:**
```athera
# program.ath requires uuid_gen module
use uuid_gen

task create_id:
    return uuid_gen.v4
```

**Sharing with dependencies:**
1. **Option A:** Compile to executable (bundles everything)
   ```bash
   athera compile program.ath -o program
   # Send executable - includes all dependencies
   ```

2. **Option B:** Send with athera.yaml
   ```
   send:
   - program.ath
   - athera.yaml
   
   # Recipient runs:
   athera install
   athera run program.ath
   ```

3. **Option C:** Minimal dependency file
   ```yaml
   # .athedep file (simple)
   uuid_gen@1.0
   
   # athera installs before running
   ```

---

## Development Roadmap

### Q1: MVP Phase (Months 1-3)
- [ ] macOS installer (.pkg)
- [ ] Linux installer (.deb, generic tarball)
- [ ] Windows installer (.msi)
- [ ] Basic CLI (run command)
- [ ] Essential stdlib modules (io, text, math, list)
- [ ] `athera --version` and `athera --help`

### Q2: Advanced Phase (Months 4-6)
- [ ] Compilation to standalone executables
- [ ] Package manager (Athera Hub)
- [ ] Extended stdlib
- [ ] REPL (interactive mode)
- [ ] Build system
- [ ] Code formatter

### Q3: Professional Phase (Months 7-9)
- [ ] Testing framework
- [ ] Linter
- [ ] Version manager
- [ ] Docker support
- [ ] Project templates

### Q4: Ecosystem Phase (Months 10-12)
- [ ] IDE plugins (VS Code, etc.)
- [ ] Web-based playground
- [ ] Community package hub
- [ ] Analytics & telemetry (opt-in)
- [ ] Enterprise features

---

## Success Metrics

### Adoption
- 10,000 downloads in first 6 months
- 1,000 GitHub stars
- 100+ packages in Hub

### Ease of Use
- Installation takes <2 minutes
- First program runs in <5 minutes
- 95% of features discoverable without docs

### Quality
- 99.9% uptime for distribution servers
- Zero security vulnerabilities
- <2% crash rate

---

## Comparison with Existing Tools

| Feature | Python | Node.js | Go | Athera (Proposed) |
|---------|--------|---------|----|----|
| Installation | Easy | Easy | Easy | Easy ✅ |
| Global Command | `python` | `node` | `go` | `athera` ✅ |
| Stdlib | Large (bloat) | Medium | Medium | Small, Essential ✅ |
| Compilation | No (default) | No (default) | Yes | Yes (optional) ✅ |
| Shareability | Requires environment | Requires environment | Single binary | Single .ath file ✅ |
| Learning Curve | Medium | Easy | Hard | Very Easy ✅ |
| Clarity | Good | Good | Good | Very Good (dot rule) ✅ |
| Package Manager | pip | npm | None | Athera Hub ✅ |
| REPL | Yes | Yes | No | Yes ✅ |

---

## Next Steps

1. **Review this architecture** with core team
2. **Prioritize features** for Phase 1
3. **Start implementation** with basic CLI
4. **Build MVP installer** for macOS first
5. **Gather user feedback** from early adopters
6. **Iterate** on installer UX

This design positions Athera as a **professional, standalone language** that users install and use like Python or Node.js — but with Athera's unique clarity and the dot rule at its core.
