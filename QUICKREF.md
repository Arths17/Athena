# Athera Language Quick Reference

## Getting Started

Run any `.ath` file:
```bash
python3 athera_interpreter.py your_program.ath
```

## Syntax Reference

### 1. Task Declaration (Functions)
```athera
task task_name:
    greet "Task body"
    # More commands...
```

### 2. Print/Output
```athera
greet "Hello, World!"
greet "Any message you want"
```

### 3. Variables
```athera
set name = "Athera"           # String
set count = 42                 # Integer
set price = 19.99              # Float
set active = true              # Boolean
set items = ["a", "b", "c"]   # List
```

### 4. Loops

**Counted Loop:**
```athera
repeat 5 times:
    greet "Repeating..."
```

**Iteration Loop:**
```athera
set colors = ["red", "green", "blue"]
repeat each color in colors:
    greet "Color found"
```

### 5. Conditions
```athera
check "file.txt" -> greet "File exists!"
check variable -> greet "Variable is truthy!"
```

### 6. File Operations
```athera
backup "source.txt" to "Destination"
backup file_variable to "Backup"
```

### 7. Call Other Tasks
```athera
task helper:
    greet "Helper task"

task main:
    run helper

run main
```

### 8. Import Modules (Placeholder)
```athera
use filesystem
use networking
```

### 9. Comments
```athera
# This is a comment
task example:    # Inline comment
    greet "Hello"
```

## Data Types

| Type | Example | Description |
|------|---------|-------------|
| String | `"text"` or `'text'` | Text data |
| Integer | `42` | Whole numbers |
| Float | `3.14` | Decimal numbers |
| Boolean | `true`, `false` | Logical values |
| List | `["a", "b", "c"]` | Ordered collection |
| Dict | `{"key": "value"}` | Key-value pairs |

## Language Rules

1. ✅ **Indentation matters** - Use consistent indentation for blocks
2. ✅ **Task names** - Use underscores: `my_task` not `myTask`
3. ✅ **No semicolons** - Each statement on new line
4. ✅ **No braces** - Use indentation instead
5. ✅ **Variables** - Must be set before use
6. ✅ **Tasks** - Define before calling with `run`

## Common Patterns

### Pattern 1: Simple Script
```athera
task main:
    greet "Starting..."
    set x = 10
    greet "Done!"

run main
```

### Pattern 2: File Backup
```athera
task backup_files:
    set files = ["doc1.txt", "doc2.txt"]
    repeat each file in files:
        backup file to "Backup"
    check "Backup/doc1.txt" -> greet "Success!"

run backup_files
```

### Pattern 3: Nested Tasks
```athera
task process_item:
    greet "Processing..."

task process_all:
    repeat 3 times:
        run process_item

run process_all
```

### Pattern 4: Conditional Logic
```athera
task check_and_act:
    set file = "data.txt"
    check file -> greet "File exists!"
    check "missing.txt" -> greet "Won't show"

run check_and_act
```

## Complete Example

```athera
# File organizer program

task organize:
    greet "Starting file organization..."
    
    set documents = ["report.txt", "notes.txt"]
    set backup_dir = "Archive"
    
    repeat each doc in documents:
        backup doc to backup_dir
    
    check "Archive/report.txt" -> greet "Backup verified!"
    greet "Organization complete!"

run organize
```

## Tips & Best Practices

1. **Use descriptive task names**: `backup_photos` not `task1`
2. **Group related operations**: Create tasks for logical units
3. **Check conditions**: Use `check` for file verification
4. **Comment your code**: Explain complex logic
5. **Test incrementally**: Run tasks individually first

## Error Handling

Current behavior:
- Missing files: Warning message shown, continues
- Unknown tasks: Error message, stops execution
- Syntax errors: Parser will report issue

## What's Next?

Athera is under active development. Future features:
- Function parameters
- Return values
- Exception handling
- More built-in functions
- Standard library modules

---

**Happy coding in Athera!**
