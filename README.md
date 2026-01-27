# Athera Programming Language

![Version](https://img.shields.io/badge/version-1.0-blue)
![License](https://img.shields.io/badge/license-MIT-green)

**Athera** is a unique programming language with natural language-like syntax designed for clarity and expressiveness. Unlike traditional languages, Athera uses intuitive commands that read like English instructions.

## Features

✨ **Natural Syntax** - Commands read like natural language  
📦 **Task-Based** - Organize code into reusable tasks  
🔄 **Simple Loops** - Intuitive repeat constructs  
✅ **Clear Conditions** - Check-based conditional execution  
📁 **File Operations** - Built-in backup and file handling  
🔢 **Variables** - Support for strings, numbers, lists, and booleans  

## Quick Start

### Installation

No installation required! Just Python 3.6+

### Running Your First Program

1. Create a file `hello.ath`:
```athera
task greet:
    greet "Hello, Athera!"

run greet
```

2. Run it:
```bash
python athera_interpreter.py hello.ath
```

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
