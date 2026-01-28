# Athera Programming Language

![Version](https://img.shields.io/badge/version-0.1.0--phase1-blue)
![Language](https://img.shields.io/badge/runtime-Go-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

**Athera** is a unique programming language designed to be **clear, intentional, and expressive**. Write natural, readable code without unnecessary syntax.

**Phase 1 Status**: ✅ Compiled runtime in Go, working interpreter, 8 stdlib modules, interactive REPL

---

## Features

✨ **Natural Syntax** – Code reads like clear instructions  
🚀 **Compiled Runtime** – Single binary, runs anywhere  
📦 **Task-Based** – Organize code into reusable tasks  
🔄 **Simple Loops** – `repeat N times` and `repeat each` syntax  
✅ **Conditions** – Clear conditional execution with `check`  
📁 **File I/O** – Built-in file operations via `io` module  
🔢 **Variables** – Support for strings, numbers, lists, booleans  
📚 **8 Stdlib Modules** – io, text, math, list, dict, time, json, path  
⚡ **REPL** – Interactive shell for exploration  
🔀 **Parallel Execution** – Run tasks concurrently with `run parallel`  

---

## Quick Start (30 seconds)

### 1. Install

**macOS Binary** (recommended):
```bash
# Download from releases or build:
go build -o athera ./cmd/athera
sudo mv athera /usr/local/bin/
```

Verify:
```bash
athera --help
```

### 2. Run Your First Program

Create `hello.ath`:
```athera
greet "Hello, World!"
```

Run it:
```bash
athera run hello.ath
```

That's it! 🎉

### 3. Try the REPL

```bash
athera repl
```

Type commands interactively:
```
athera> greet "Hello"
Hello
athera> set x = 5
athera> greet x
5
athera> exit
```

---

## Examples

### Example 1: Tasks & Variables
```athera
task greet_user with name:
    greet "Hello, "
    greet name

run greet_user "Alice"
```

### Example 2: Loops & Lists
```athera
set colors = ["red", "green", "blue"]
repeat each color in colors:
    greet color
```

### Example 3: File Operations
```athera
use io

io.write "output.txt" "Hello from Athera!"
set content = io.read "output.txt"
greet content
```

### Example 4: Text Processing
```athera
use text

set message = "hello world"
set upper = text.upper message
greet upper
```

**More examples** in [examples/](examples/) folder.

---

## Standard Library

8 modules embedded in the binary:

| Module | Purpose | Key Functions |
|--------|---------|---|
| `io` | File operations | read, write, append, exists, size |
| `text` | String operations | upper, lower, length, split, contains, trim |
| `math` | Arithmetic | add, sub, mul, div, sqrt, abs |
| `list` | List operations | length, append, at, contains |
| `dict` | Key-value storage | get, set, keys, values |
| `time` | Dates & timers | now, timestamp, sleep, format |
| `json` | JSON parsing | parse, stringify |
| `path` | File paths | join, dir, base, ext, exists |

**Usage:**
```athera
use io
use text
use math

set data = io.read "file.txt"
set upper = text.upper data
set len = math.add 5 10
```

---

## Language Syntax

### Commands

| Command | Purpose | Example |
|---------|---------|---------|
| `greet` | Print output | `greet "Hello"` |
| `set` | Assign variable | `set x = 42` |
| `task` | Define function | `task hello:` |
| `run` | Call task | `run hello` |
| `repeat N times` | Loop N times | `repeat 5 times:` |
| `repeat each` | Loop over list | `repeat each x in list:` |
| `check` | Conditional | `check x -> greet "yes"` |
| `use` | Import module | `use io` |
| `protect/handle` | Error handling | `protect: ... handle:` |
| `run parallel` | Concurrent execution | `run parallel task1, task2` |

### Data Types

```athera
set string = "hello"
set number = 42
set float = 3.14
set boolean = true
set list = [1, 2, 3]
set empty_dict = {}
```

### Accessing Properties

```athera
set message = "hello"
greet message.length    # 5

set items = [1, 2, 3]
greet items.length      # 3
greet items.is_empty    # false
```

---

## Learning Path

1. **Read**: [GETTING_STARTED.md](GETTING_STARTED.md) – Installation & basics (10 min)
2. **Try**: [examples/](examples/) – 8 ready-to-run programs
3. **Reference**: [QUICKREF.md](QUICKREF.md) – Syntax cheat sheet
4. **Learn**: [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md) – Language philosophy
5. **Build**: Write your own programs!

---

## Architecture

**Compiled Runtime** – Single-file Go binary, ~3MB  
**Lexer** → **Parser** → **Tree-Walking Interpreter**  
**Embedded Stdlib** – 8 modules compiled in  
**Cross-Platform** – Runs on macOS, Linux, Windows

---

## Development

### Build Binary

```bash
cd /path/to/athera
go build -o athera ./cmd/athera
```

### For macOS ARM64

```bash
GOOS=darwin GOARCH=arm64 go build -o athera ./cmd/athera
```

### Run Tests

```bash
go run ./cmd/athera run examples/hello_world.ath
go run ./cmd/athera run examples/file_ops.ath
go run ./cmd/athera repl
```

---

## File Structure

```
.
├── cmd/athera/           # CLI entrypoint
├── internal/lang/        # Language implementation
│   ├── ast.go           # AST node definitions
│   ├── lexer.go         # Tokenizer
│   ├── parser.go        # Parser
│   ├── interpreter.go   # Execution engine
│   └── stdlib.go        # Standard library modules
├── examples/            # Sample programs (8 examples)
├── GETTING_STARTED.md   # Installation & quick start
├── QUICKREF.md          # Syntax reference
├── DESIGN_BRIEF_V2.1.md # Language design philosophy
├── go.mod              # Go module definition
└── athera              # Compiled binary
```

---

## What's Different About Athera?

### 1. Intent Over Mechanics
```athera
# Clear intent
greet "Processing data"
repeat each item in items:
    process item

# vs. typical:
for item in items:
    print("Processing data")
    process(item)
```

### 2. The Dot Rule
Properties (read-only) vs. Actions (side effects):
```athera
set len = string.length    # Property – no side effects
set upper = text.upper     # Action – transforms value
```

### 3. One Clear Way Per Concept
```athera
# One way to loop over lists:
repeat each item in list:
    greet item

# Not multiple ways like in other languages
```

---

## Roadmap

**Phase 1 (Current)** ✅
- [x] Compiled Go runtime
- [x] Core interpreter (lexer, parser, executor)
- [x] 8 stdlib modules
- [x] REPL
- [x] Basic examples

**Phase 2**
- [ ] Compilation to standalone executables
- [ ] Package manager (Athera Hub)
- [ ] Code formatter (`athera fmt`)
- [ ] Linter (`athera lint`)
- [ ] Testing framework

**Phase 3**
- [ ] VS Code extension
- [ ] More IDE plugins
- [ ] Advanced stdlib
- [ ] Web-based playground

**Phase 4**
- [ ] Docker support
- [ ] Community hub
- [ ] Ecosystem growth

---

## Contributing

Contributions welcome! Areas:
- [ ] More stdlib modules
- [ ] Better error messages
- [ ] Performance improvements
- [ ] Documentation
- [ ] Examples

---

## License

MIT – See LICENSE file

---

## Community

- 📖 [Documentation](GETTING_STARTED.md)
- 🐛 [Issues](https://github.com/Arths17/Athena/issues)
- 💬 [Discussions](https://github.com/Arths17/Athena/discussions)

---

## Next Steps

→ **Start**: [GETTING_STARTED.md](GETTING_STARTED.md)  
→ **Explore**: [examples/](examples/)  
→ **Learn**: [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md)

**Happy coding!** 🚀

## Language Reference

### Task Declaration

Define reusable tasks (similar to functions):

```athera
task morning_routine:
    greet "Good morning!"
    greet "Starting your day..."
```

### Commands

#### `greet` - Print Messages
```athera
greet "Hello, World!"
greet "Welcome to Athera"
```

#### `set` - Variable Assignment
```athera
set name = "Athera"
set count = 42
set enabled = true
set items = ["apple", "banana", "cherry"]
```

#### `repeat N times` - Counted Loop
```athera
repeat 5 times:
    greet "Repeating..."
```

#### `repeat each` - Iteration Loop
```athera
set fruits = ["apple", "banana", "orange"]
repeat each fruit in fruits:
    greet "Fruit: item"
```

#### `check` - Conditional Execution
```athera
check "Backup/file.txt" -> greet "File exists!"
```

#### `backup` - File Operations
```athera
backup "document.txt" to "Backup"
```

#### `run` - Call Another Task
```athera
task helper:
    greet "Helper task running"

task main:
    run helper

run main
```

#### `use` - Import Module (placeholder)
```athera
use filesystem
use networking
```

## Example Programs

### Morning Routine
```athera
task morning_routine:
    greet "Good morning!"
    set files = ["Documents/report.docx", "Pictures/photo.jpg"]
    repeat each file in files:
        backup file to "Backup"
    check "Backup/report.docx" -> greet "Backup complete!"
    greet "Have a nice day!"

run morning_routine
```

### Counter
```athera
task count_to_ten:
    greet "Counting to 10:"
    repeat 10 times:
        greet "Count!"

run count_to_ten
```

### File Manager
```athera
task organize_files:
    set documents = ["report.txt", "notes.txt", "presentation.ppt"]
    
    greet "Organizing files..."
    repeat each doc in documents:
        backup doc to "Archive"
    
    check "Archive/report.txt" -> greet "Organization complete!"

run organize_files
```

## Language Philosophy

Athera is designed with these principles:

1. **Readability First** - Code should read like instructions
2. **Minimal Syntax** - No unnecessary symbols or keywords
3. **Natural Flow** - Commands follow natural language patterns
4. **Expressiveness** - Complex operations with simple syntax
5. **No Confusion** - Unique syntax prevents mixing with other languages

## What Makes Athera Different?

### Traditional Languages:
```python
if os.path.exists("file.txt"):
    print("File exists")
    
for item in items:
    shutil.copy(item, "backup")
```

### Athera:
```athera
check "file.txt" -> greet "File exists"

repeat each item in items:
    backup item to "backup"
```

## Technical Details

### Implementation

- **Parser**: Recursive descent parser
- **Execution**: Tree-walking interpreter
- **Type System**: Dynamic typing
- **Scope**: Task-level scope for variables

### Data Types

- **Strings**: `"text"` or `'text'`
- **Numbers**: `42`, `3.14`
- **Booleans**: `true`, `false`
- **Lists**: `["item1", "item2"]`
- **Dictionaries**: `{"key": "value"}`

### Comments

```athera
# This is a comment
task example:  # Inline comment supported
    greet "Hello"
```

## Examples Included

- `morning_routine.ath` - Basic task and file operations
- `demo.ath` - Complete language showcase
- `fibonacci.ath` - Loop and arithmetic demo

## Running the Examples

```bash
# Run morning routine
python athera_interpreter.py morning_routine.ath

# Run full demo
python athera_interpreter.py demo.ath

# Run Fibonacci
python athera_interpreter.py fibonacci.ath
```

## Future Enhancements

- [ ] Error handling with `handle <error> -> <action>`
- [ ] Function parameters: `task greet_user with name:`
- [ ] Return values from tasks
- [ ] Module system implementation
- [ ] Standard library (file, network, math modules)
- [ ] Interactive REPL mode
- [ ] Debugging support

## Contributing

Athera is an experimental language. Contributions, ideas, and feedback are welcome!

## License

MIT License - Feel free to use, modify, and distribute

---

**Created with ❤️ for simplicity and clarity in programming**

## Contact

For questions or suggestions about Athera, please open an issue or discussion.

---

*"Programming should be as natural as speaking"* - Athera Philosophy
