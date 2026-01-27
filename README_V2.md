# 🎉 Athera v2.0 - Complete Guide

## 🚀 What's New - Advanced Features Added!

Athera now has **professional-grade features**:

### ✨ New in v2.0

1. **Functions with Parameters** - Pass arguments to tasks
2. **Error Handling** - Protect/handle blocks (like try/catch)
3. **Parallel Execution** - Run multiple tasks concurrently
4. **Custom Modules** - Create and import reusable libraries
5. **Return Values** - Tasks can return data

---

## 📦 Installation & Setup

**Requirements**: Python 3.6+
**Dependencies**: None (uses standard library only)

```bash
# Clone or download the Athera folder
cd Athera

# Make runner executable (optional)
chmod +x run_athera.sh

# Test installation
python3 athera_interpreter.py morning_routine.ath
```

---

## 🎯 Quick Start

### Method 1: Use the Runner
```bash
./run_athera.sh          # Show all examples
./run_athera.sh 1        # Run example 1
./run_athera.sh 4        # Run advanced features
```

### Method 2: Direct Execution
```bash
python3 athera_interpreter.py your_program.ath
```

### Method 3: Create Your Own
```bash
# Copy the template
cp your_program_template.ath my_first_program.ath

# Edit it with your code
nano my_first_program.ath

# Run it
python3 athera_interpreter.py my_first_program.ath
```

---

## 📚 Language Reference

### Basic Syntax

#### 1. Tasks (Functions)
```athera
task hello:
    greet "Hello, World!"

run hello
```

#### 2. Tasks with Parameters (NEW!)
```athera
task greet_person with name, age:
    greet "Hello there!"
    greet "Welcome!"

run greet_person "Alice", 25
```

#### 3. Variables
```athera
set name = "Athera"
set count = 42
set active = true
set items = ["a", "b", "c"]
```

#### 4. Loops
```athera
# Counted loop
repeat 5 times:
    greet "Repeating..."

# List iteration
set colors = ["red", "green", "blue"]
repeat each color in colors:
    greet "Found color"
```

#### 5. Conditions
```athera
check "file.txt" -> greet "File exists!"
check variable -> greet "Variable is truthy"
```

#### 6. File Operations
```athera
backup "source.txt" to "Destination"
backup file_variable to "Backup"
```

#### 7. Error Handling (NEW!)
```athera
protect:
    backup "important.txt" to "Backup"
    greet "Success!"
handle:
    greet "Error occurred, but handled!"
```

#### 8. Parallel Execution (NEW!)
```athera
task worker1:
    greet "Worker 1"

task worker2:
    greet "Worker 2"

run parallel worker1, worker2
```

#### 9. Modules (NEW!)
```athera
use math_utils
use file_utils

run add 5, 10
```

#### 10. Return Values (NEW!)
```athera
task calculate with x:
    set result = 42
    return result
```

---

## 📖 Complete Examples

### Example 1: Hello World
```athera
task hello:
    greet "Hello, World!"

run hello
```

### Example 2: With Parameters
```athera
task greet_user with name:
    greet "Hello, friend!"

run greet_user "Alice"
```

### Example 3: Error Handling
```athera
task safe_backup:
    protect:
        backup "data.txt" to "Backup"
    handle:
        greet "Backup failed, using alternative"

run safe_backup
```

### Example 4: Parallel Processing
```athera
task download:
    greet "Downloading..."

task process:
    greet "Processing..."

task main:
    run parallel download, process

run main
```

### Example 5: Complete Application
```athera
use file_utils

task process_data with filename:
    protect:
        greet "Processing file..."
        backup filename to "Processed"
        greet "Success!"
    handle:
        greet "Failed! Retrying..."

task main:
    set files = ["data1.txt", "data2.txt"]
    repeat each file in files:
        run process_data file
    greet "All files processed!"

run main
```

---

## 🗂️ Project Structure

```
Athera/
├── athera_interpreter.py          # Main interpreter
├── run_athera.sh                  # Easy launcher
│
├── PROGRAMS (12 examples):
│   ├── morning_routine.ath        # Original example
│   ├── complete_example.ath       # All features
│   ├── demo.ath                   # Visual demo
│   ├── advanced_features.ath     # NEW: Advanced features
│   ├── module_demo.ath           # NEW: Module system
│   ├── error_handling_demo.ath   # NEW: Error handling
│   ├── parallel_demo.ath         # NEW: Parallel tasks
│   ├── real_world_app.ath        # NEW: Complete app
│   ├── real_backup_test.ath      # File operations
│   ├── file_test.ath             # File handling
│   ├── fibonacci.ath             # Algorithm
│   └── your_program_template.ath # NEW: Your template
│
├── MODULES (2 custom modules):
│   └── modules/
│       ├── math_utils.ath        # NEW: Math operations
│       └── file_utils.ath        # NEW: File utilities
│
└── DOCUMENTATION (7 guides):
    ├── README_V2.md              # This file
    ├── GETTING_STARTED.md        # Tutorial
    ├── QUICKREF.md               # Quick reference
    ├── ADVANCED_FEATURES.md      # NEW: Advanced guide
    ├── WHATS_NEW.txt             # NEW: What's new
    ├── PROJECT_SUMMARY.md        # Project overview
    └── INDEX.txt                 # Complete index
```

---

## 🎓 Learning Path

### Level 1: Beginner
1. Read [GETTING_STARTED.md](GETTING_STARTED.md)
2. Run: `./run_athera.sh 1`
3. Study: [morning_routine.ath](morning_routine.ath)
4. Create simple programs using the template

### Level 2: Intermediate
1. Run: `./run_athera.sh 2`
2. Study: [complete_example.ath](complete_example.ath)
3. Practice with variables, loops, conditions

### Level 3: Advanced
1. Read [ADVANCED_FEATURES.md](ADVANCED_FEATURES.md)
2. Run: `./run_athera.sh 4`
3. Study: [advanced_features.ath](advanced_features.ath)
4. Create modules and use advanced features

---

## 💡 Real-World Use Cases

### 1. File Backup Automation
```athera
task daily_backup:
    set files = ["documents/", "photos/", "code/"]
    repeat each folder in files:
        backup folder to "DailyBackup"
    greet "Daily backup complete!"
```

### 2. Data Processing Pipeline
```athera
use data_processor

task pipeline with input_file:
    protect:
        run validate input_file
        run transform input_file
        run save input_file
    handle:
        greet "Pipeline failed, logging error"
```

### 3. Parallel Batch Processing
```athera
task batch_process:
    run parallel process_batch1, process_batch2, process_batch3
    greet "All batches processed!"
```

---

## 🔧 Creating Custom Modules

### Step 1: Create Module File
Create `modules/my_utils.ath`:
```athera
task helper_function with param:
    greet "Helper is working"
    return result

task another_function:
    greet "Another function"
```

### Step 2: Use Your Module
Create `main.ath`:
```athera
use my_utils

task main:
    run helper_function "data"
    run another_function

run main
```

### Step 3: Run
```bash
python3 athera_interpreter.py main.ath
```

---

## 📊 Feature Comparison

| Feature | Python | Athera |
|---------|--------|--------|
| Function def | `def func(x):` | `task func with x:` |
| Print | `print("msg")` | `greet "msg"` |
| Variables | `x = 10` | `set x = 10` |
| For loop | `for i in list:` | `repeat each i in list:` |
| Try/catch | `try: ... except:` | `protect: ... handle:` |
| Threading | `threading.Thread()` | `run parallel` |
| Import | `import module` | `use module` |

---

## 🎯 Command Reference

| Command | Syntax | Example |
|---------|--------|---------|
| Task | `task name:` | `task hello:` |
| Parameters | `task name with a, b:` | `task add with x, y:` |
| Print | `greet "text"` | `greet "Hello!"` |
| Variable | `set x = value` | `set count = 5` |
| Repeat | `repeat N times:` | `repeat 10 times:` |
| For each | `repeat each x in list:` | `repeat each item in items:` |
| Condition | `check x -> action` | `check "file" -> greet "OK"` |
| Backup | `backup src to dst` | `backup "x.txt" to "Backup"` |
| Call | `run task` | `run process` |
| Call with args | `run task a, b` | `run add 5, 10` |
| Parallel | `run parallel t1, t2` | `run parallel job1, job2` |
| Import | `use module` | `use math_utils` |
| Protect | `protect: ... handle:` | See error handling |
| Return | `return value` | `return result` |
| Comment | `# comment` | `# This is a comment` |

---

## 🚀 Running Examples

```bash
# Show all examples
./run_athera.sh

# Basic examples
./run_athera.sh 1    # Morning routine
./run_athera.sh 2    # Complete showcase
./run_athera.sh 3    # Visual demo

# Advanced examples (NEW!)
./run_athera.sh 4    # Advanced features
./run_athera.sh 5    # Module system
./run_athera.sh 6    # Error handling
./run_athera.sh 7    # Parallel execution
./run_athera.sh 8    # Real-world app

# File operations
./run_athera.sh 9    # Real file backup
./run_athera.sh 10   # File handling

# Algorithms
./run_athera.sh 11   # Fibonacci
```

---

## 💻 Creating Your First Program

```bash
# Step 1: Copy the template
cp your_program_template.ath my_program.ath

# Step 2: Edit it
# Open my_program.ath in your editor
# Modify the tasks and logic

# Step 3: Run it
python3 athera_interpreter.py my_program.ath

# Step 4: Iterate
# Make changes, run again, repeat!
```

---

## 📝 Tips & Best Practices

1. **Start Simple**: Begin with basic tasks, add complexity gradually
2. **Use Comments**: Document your code with `#`
3. **Test Often**: Run after each change
4. **Use Parameters**: Make tasks reusable with parameters
5. **Handle Errors**: Wrap risky operations in protect/handle
6. **Parallelize**: Use `run parallel` for independent tasks
7. **Modularize**: Create modules for reusable functionality
8. **Follow Examples**: Study the provided examples

---

## 🐛 Troubleshooting

### "Task not found"
- Make sure the task is defined before calling it
- Check spelling of task name

### "Module not found"
- Ensure module file exists in `modules/` folder
- Check file extension is `.ath`

### Indentation errors
- Use consistent indentation (spaces or tabs)
- Each block level should be indented the same

### File not found warnings
- Check file paths are correct
- Use relative paths from where you run the interpreter

---

## 📚 Documentation Files

- **README_V2.md** (this file) - Complete guide
- **GETTING_STARTED.md** - Beginner tutorial
- **QUICKREF.md** - Quick syntax reference
- **ADVANCED_FEATURES.md** - Advanced features guide
- **WHATS_NEW.txt** - Version 2.0 changes
- **PROJECT_SUMMARY.md** - Project overview

---

## 🎉 What You Can Build

- ✅ File automation scripts
- ✅ Backup systems
- ✅ Data processing pipelines
- ✅ Batch processing tools
- ✅ System administration tasks
- ✅ Build & deployment scripts
- ✅ API automation
- ✅ Task orchestration
- ✅ Educational programs
- ✅ Rapid prototypes

---

## 🌟 Contributing

Athera is an experimental language. Ideas and feedback welcome!

---

## 📄 License

MIT License - Free to use, modify, and distribute

---

## 🎊 Summary

**Athera v2.0 is a complete, modern programming language with:**

✅ Natural, readable syntax
✅ Functions with parameters
✅ Error handling
✅ Parallel execution
✅ Custom modules
✅ Return values
✅ 12 example programs
✅ 2 custom modules
✅ 7 documentation guides
✅ Zero dependencies
✅ Ready for real-world use

**Start coding in Athera today!** 🚀

```bash
./run_athera.sh
```

---

*"Programming should be as natural as speaking"* - Athera Philosophy
