# Getting Started with Athera

Welcome! Athera is a programming language designed to be **clear, intentional, and expressive**. This guide gets you writing Athera programs in 10 minutes.

---

## Installation (Phase 1)

### macOS – Fastest Setup

**Step 1: Download Binary**
- Download `athera` from [GitHub Releases](https://github.com/Arths17/Athena/releases/tag/v0.1.0-phase1)

**Step 2: Make it Executable & Install**
```bash
chmod +x athera
sudo mv athera /usr/local/bin/
```

**Step 3: Verify**
```bash
athera --help
which athera
```

You should see: `/usr/local/bin/athera`

### Build from Source
```bash
git clone https://github.com/Arths17/Athena.git
cd Athena
go build -o athera ./cmd/athera
sudo mv athera /usr/local/bin/
```

---

## Your First Program (30 Seconds)

Create `hello.ath`:
```athera
greet "Hello, World!"
```

Run it:
```bash
athera run hello.ath
```

Output:
```
Hello, World!
```

**That's it!** 🎉

---

## Five Essential Concepts

### 1. **Greet** – Print Output
```athera
greet "This prints to the screen"
```

### 2. **Set** – Variables
```athera
set name = "Alice"
set age = 30
set items = ["apple", "banana"]
greet name
```

### 3. **Repeat** – Loops
```athera
# Fixed count
repeat 3 times:
    greet "Hello!"

# Over a list
set fruits = ["apple", "banana"]
repeat each fruit in fruits:
    greet fruit
```

### 4. **Task** – Reusable Code
```athera
task greet_user with name:
    greet "Hello, "
    greet name

run greet_user "Alice"
```

### 5. **Check** – Conditions
```athera
check "file.txt" -> greet "File exists!"
check "missing.txt" -> greet "Won't print"
```

---

## Complete Examples

### Example 1: Variables & Output
```athera
set greeting = "Hello, World!"
greet greeting

set count = 42
greet count
```

### Example 2: Lists & Loops
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

set words = text.split message " "
repeat each word in words:
    greet word
```

### Example 5: Tasks & Functions
```athera
task add with a, b:
    set result = a + b
    greet result

run add 5 10
```

---

## Using Standard Library Modules

### **io** – File Operations
```athera
use io
io.write "file.txt" "content"
set data = io.read "file.txt"
io.append "log.txt" "new line"
check io.exists "file.txt" -> greet "Found!"
```

### **text** – String Functions
```athera
use text
set upper = text.upper "hello"
set len = text.length "hi"
set has_it = text.contains "hello" "ell"
set words = text.split "a,b,c" ","
```

### **math** – Calculations
```athera
use math
set sum = math.add 2 3 4
set product = math.mul 5 6
set root = math.sqrt 16
```

### **list** – List Operations
```athera
use list
set arr = [1, 2, 3]
set len = list.length arr
set item = list.at arr 0
set has_2 = list.contains arr 2
```

### **dict** – Dictionaries
```athera
use dict
set obj = {}
set obj = dict.set obj "name" "Alice"
set name = dict.get obj "name"
set keys = dict.keys obj
```

### **time** – Dates & Timers
```athera
use time
greet time.now
time.sleep 1000
greet time.timestamp
```

### **json** – JSON Parsing
```athera
use json
set text = "{\"name\": \"Alice\"}"
set obj = json.parse text
greet obj
```

### **path** – File Paths
```athera
use path
set file = path.join "home" "user" "file.txt"
set dir = path.dir file
set name = path.base file
```

---

## REPL – Interactive Mode

Try commands live:
```bash
athera repl
```

Type commands:
```
athera> greet "Hello"
Hello
athera> set x = 5
athera> greet x
5
athera> exit
Goodbye.
```

---

## Common Patterns

### Pattern 1: File Processing
```athera
use io
set lines = io.read_lines "data.txt"
repeat each line in lines:
    greet line
```

### Pattern 2: Error Handling
```athera
protect:
    set data = io.read "missing.txt"
handle:
    greet "File not found!"
    set data = "default"
```

### Pattern 3: Parallel Tasks
```athera
task fetch:
    greet "Fetching..."

task process:
    greet "Processing..."

run parallel fetch, process
```

---

---

## Tips & Tricks

1. **Comments**: Use `#` for comments
   ```athera
   # This is a comment
   greet "Code runs here"
   ```

2. **String concatenation**: Use `+`
   ```athera
   set msg = "Hello" + " World"
   greet msg
   ```

3. **List literals**: Use `[...]`
   ```athera
   set items = [1, 2, 3]
   ```

4. **Access properties**: Use `.` for properties
   ```athera
   set msg = "hello"
   greet msg.length  # Returns 5
   ```

---

## Troubleshooting

| Problem | Solution |
|---------|----------|
| "Command not found: athera" | Make sure `/usr/local/bin` is in PATH |
| "File not found" | Check file exists with `ls filename.ath` |
| "Task not found" | Define task before calling with `run` |
| Indentation errors | Use consistent spacing (2-4 spaces) |
| Module not found | Use `use module_name` at top |

---

## Next Steps

1. Read [README.md](README.md) for full documentation
2. Check [QUICKREF.md](QUICKREF.md) for syntax reference
3. Try examples in [examples/](examples/) folder
4. Explore [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md) for language design
5. Write your own program!

---

## Share Your Code

Athera programs are single-file, portable `.ath` scripts. Share freely!

```bash
# Someone with athera installed can run it:
athera run your_program.ath
```

No dependencies. No setup. Just code.

---

**You're ready!** Start with a simple `greet` and build from there. 🚀

For issues or questions, visit [GitHub Issues](https://github.com/Arths17/Athena/issues).
    greet "Start creating amazing things!"

run celebrate
```

---

**Need help?** Study the example files - they're your best teachers!

**Ready to dive deeper?** Read the comprehensive [README.md](README.md)
