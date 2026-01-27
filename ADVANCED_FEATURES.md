# Athera Advanced Features Guide

## 🚀 New Advanced Features

Athera now includes professional-grade features found in modern programming languages!

### 1. Functions with Parameters

Define tasks that accept parameters:

```athera
task greet_person with name:
    greet "Hello, name!"

task calculate_area with width, height:
    set area = 100
    return area

task show_user with name, age, city:
    greet "User info displayed"

# Call with arguments
run greet_person "Alice"
run calculate_area 10, 20
run show_user "Bob", 25, "NYC"
```

**Syntax**: `task <name> with <param1>, <param2>, ...:`

### 2. Error Handling (Try/Catch Style)

Protect code blocks and handle errors gracefully:

```athera
task safe_operation:
    protect:
        backup "file.txt" to "Backup"
        greet "This might fail"
    handle:
        greet "Error caught! Recovering..."
        greet "Program continues"
```

**How it works**:
- `protect:` block executes normally
- If error occurs, execution jumps to `handle:` block
- Program continues after error is handled

**Nested error handling**:
```athera
protect:
    greet "Outer protection"
    protect:
        greet "Inner protection"
    handle:
        greet "Inner error handler"
handle:
    greet "Outer error handler"
```

### 3. Parallel Execution (Concurrency)

Run multiple tasks simultaneously:

```athera
task download:
    greet "Downloading..."

task process:
    greet "Processing..."

task notify:
    greet "Notifying..."

# Run all three at once!
run parallel download, process, notify
```

**Features**:
- True concurrent execution using threads
- All tasks run simultaneously
- Waits for all to complete before continuing

**Use cases**:
- File processing pipelines
- Network operations
- Data transformation
- Background tasks

### 4. Custom Modules / Libraries

Create reusable Athera modules:

**math_utils.ath**:
```athera
task add with a, b:
    greet "Adding..."
    return result

task multiply with a, b:
    greet "Multiplying..."
    return result
```

**main.ath**:
```athera
# Import the module
use math_utils

# Use module functions
run add 5, 10
run multiply 3, 7
```

**Module Organization**:
- Put modules in `modules/` folder
- Each module is a `.ath` file
- Modules can define multiple tasks
- Tasks become available after import

### 5. Return Values

Tasks can return values:

```athera
task calculate with x:
    set result = 42
    return result

task main:
    run calculate 10
    greet "Calculation complete"
```

## 📊 Complete Feature Comparison

| Feature | Basic Athera | Advanced Athera |
|---------|-------------|-----------------|
| Functions | ✓ | ✓ |
| Parameters | ✗ | ✓ |
| Return values | ✗ | ✓ |
| Error handling | ✗ | ✓ |
| Parallel execution | ✗ | ✓ |
| Modules | Limited | ✓ Full support |

## 🎯 Real-World Example

```athera
# Import modules
use file_utils
use data_processor

# Define tasks with parameters
task process_user with username, age, email:
    protect:
        greet "Processing user..."
        run validate_data username
        run save_to_database username, age, email
    handle:
        greet "Error! Using fallback..."
        run log_error username

# Parallel processing
task batch_process:
    run parallel process_user1, process_user2, process_user3

# Main application
task main:
    run setup_database "users.db"
    run batch_process
    greet "All users processed!"

run main
```

## 🛠️ Advanced Patterns

### Pattern 1: Error Recovery Pipeline
```athera
task safe_pipeline:
    protect:
        run step_one
        run step_two
        run step_three
    handle:
        greet "Pipeline failed, using backup"
        run backup_pipeline
```

### Pattern 2: Parallel Data Processing
```athera
task process_files:
    set files = ["file1.txt", "file2.txt", "file3.txt"]
    run parallel process_file1, process_file2, process_file3
```

### Pattern 3: Modular Architecture
```athera
# main.ath
use database
use api
use logger

task application:
    run database.connect "mydb"
    run api.start_server 8080
    run logger.init "app.log"
```

### Pattern 4: Parameter Passing Chain
```athera
task calculate with x:
    set result = 100
    return result

task process with data:
    run calculate data
    greet "Processing complete"

task main:
    run process 42
```

## 📝 Syntax Reference

### Parameters
```athera
task name with param1, param2, param3:
    # task body
```

### Error Handling
```athera
protect:
    # risky code
handle:
    # error recovery
```

### Parallel Execution
```athera
run parallel task1, task2, task3
```

### Module Import
```athera
use module_name
```

### Return Statement
```athera
return value
```

## 🔥 Best Practices

1. **Use parameters for reusable tasks**
   ```athera
   task backup_file with filename:
       backup filename to "Backup"
   ```

2. **Always protect risky operations**
   ```athera
   protect:
       backup "important.txt" to "Backup"
   handle:
       greet "Backup failed!"
   ```

3. **Parallelize independent tasks**
   ```athera
   # These can run simultaneously
   run parallel download, process, upload
   ```

4. **Organize code into modules**
   ```
   modules/
   ├── file_utils.ath
   ├── data_processor.ath
   └── helpers.ath
   ```

5. **Return values for chainable operations**
   ```athera
   task get_data:
       set data = "result"
       return data
   ```

## 🎓 Learning Path

1. **Start with parameters**: Add parameters to existing tasks
2. **Add error handling**: Wrap risky operations in protect/handle
3. **Experiment with parallel**: Run independent tasks concurrently
4. **Create modules**: Organize code into reusable modules
5. **Build applications**: Combine all features

## 💡 Tips

- Parameters make tasks reusable
- Error handling prevents crashes
- Parallel execution speeds up I/O-bound tasks
- Modules keep code organized
- Return values enable data flow

## 🚀 What's Next?

You now have access to all advanced features! Build real applications:
- File processing pipelines
- Data transformation tools
- Automation scripts
- API clients
- System utilities

---

**Athera is now a full-featured programming language!** 🎉
