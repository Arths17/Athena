# 🚀 Getting Started with Athera

Welcome to **Athera** - a unique programming language with natural, expressive syntax!

## ⚡ Quick Start (30 seconds)

### Method 1: Using the Runner Script
```bash
./run_athera.sh 2
```
This runs the complete example showcase!

### Method 2: Direct Execution
```bash
python3 athera_interpreter.py complete_example.ath
```

## 📚 Step-by-Step Tutorial

### Step 1: Your First Program

Create a file called `hello.ath`:
```athera
task greet:
    greet "Hello, Athera!"

run greet
```

Run it:
```bash
python3 athera_interpreter.py hello.ath
```

Output:
```
Hello, Athera!
```

### Step 2: Working with Variables

Create `variables.ath`:
```athera
task show_variables:
    set name = "Alice"
    set age = 25
    set hobbies = ["coding", "reading", "music"]
    
    greet "Variables set successfully!"

run show_variables
```

### Step 3: Using Loops

Create `loops.ath`:
```athera
task count:
    greet "Counting to 5:"
    repeat 5 times:
        greet "  Count!"

run count
```

### Step 4: Working with Lists

Create `lists.ath`:
```athera
task process_list:
    set colors = ["red", "green", "blue"]
    
    greet "Processing colors:"
    repeat each color in colors:
        greet "  Found a color!"

run process_list
```

### Step 5: Multiple Tasks

Create `tasks.ath`:
```athera
task helper:
    greet "  Helper is working..."

task main:
    greet "Main task starting"
    run helper
    greet "Main task ending"

run main
```

### Step 6: Real File Operations

First, create a test file:
```bash
echo "Test content" > test.txt
```

Then create `backup.ath`:
```athera
task backup_file:
    greet "Backing up test.txt..."
    backup "test.txt" to "MyBackup"
    check "MyBackup/test.txt" -> greet "Backup successful!"

run backup_file
```

## 🎯 Try the Examples

We've included 6 ready-to-run examples:

### 1. Morning Routine (Basic)
```bash
./run_athera.sh 1
```
Shows: tasks, loops, file operations, conditions

### 2. Complete Example (Comprehensive)
```bash
./run_athera.sh 2
```
Shows: ALL language features in action

### 3. Demo (Visual)
```bash
./run_athera.sh 3
```
Shows: Formatted output, multiple tasks

### 4. Real Backup Test (Practical)
```bash
./run_athera.sh 4
```
Shows: Real file operations with verification

### 5. File Test (Operations)
```bash
./run_athera.sh 5
```
Shows: File handling patterns

### 6. Fibonacci (Algorithm)
```bash
./run_athera.sh 6
```
Shows: Loops and arithmetic

## 📖 Language Cheat Sheet

| Feature | Syntax | Example |
|---------|--------|---------|
| Task | `task name:` | `task hello:` |
| Print | `greet "msg"` | `greet "Hi!"` |
| Variable | `set x = value` | `set age = 25` |
| Loop (count) | `repeat N times:` | `repeat 5 times:` |
| Loop (each) | `repeat each x in list:` | `repeat each item in items:` |
| Condition | `check x -> action` | `check "file" -> greet "OK"` |
| Backup | `backup src to dst` | `backup "x.txt" to "Backup"` |
| Call Task | `run task_name` | `run helper` |

## 💡 Common Patterns

### Pattern: Initialize and Process
```athera
task process:
    set items = ["a", "b", "c"]
    repeat each item in items:
        greet "Processing..."
    greet "Done!"

run process
```

### Pattern: Conditional Actions
```athera
task check_files:
    check "important.txt" -> greet "File found!"
    check "missing.txt" -> greet "Won't show if missing"

run check_files
```

### Pattern: Helper Tasks
```athera
task log:
    greet "  [LOG] Action performed"

task main:
    greet "Starting"
    run log
    greet "Finished"

run main
```

## 🔍 Understanding Errors

### "Task not found"
```
[Error: Task 'xyz' not found]
```
**Fix**: Define the task before calling it with `run`

### "Source file not found"
```
[Warning: Source file not found: file.txt]
```
**Fix**: Make sure the file exists before backing it up

### Indentation errors
```
Parsing error...
```
**Fix**: Use consistent indentation (spaces or tabs, not mixed)

## 🎓 Next Steps

1. **Read the Full Docs**: Check [README.md](README.md)
2. **Quick Reference**: See [QUICKREF.md](QUICKREF.md)
3. **Study Examples**: Look at the `.ath` files
4. **Write Your Own**: Create custom programs
5. **Experiment**: Try combining features

## 🛠️ Tips for Success

1. **Start Simple**: Begin with one feature at a time
2. **Test Often**: Run your program after each change
3. **Use Comments**: Document your code with `#`
4. **Name Clearly**: Use descriptive task names
5. **Check Output**: Look at what your program prints

## 🎯 Challenge Yourself

Try these exercises:

**Easy**: Create a task that greets 10 times
**Medium**: Create a file organizer with 3 folders
**Hard**: Create a task that calls multiple helper tasks

## 📞 Help & Resources

- **Full Documentation**: `README.md`
- **Quick Reference**: `QUICKREF.md`  
- **Project Summary**: `PROJECT_SUMMARY.md`
- **Interpreter Code**: `athera_interpreter.py` (well commented!)

## 🎉 You're Ready!

You now know enough to start writing Athera programs. Happy coding!

```athera
task celebrate:
    greet "🎉 You're an Athera programmer now!"
    greet "Start creating amazing things!"

run celebrate
```

---

**Need help?** Study the example files - they're your best teachers!

**Ready to dive deeper?** Read the comprehensive [README.md](README.md)
