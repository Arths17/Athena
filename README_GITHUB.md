# 🚀 Athera - A Professional Programming Language

**Status:** Working Interpreter + Professional Design Brief  
**GitHub:** https://github.com/Arths17/Athena  
**Design Philosophy:** Intent over mechanics, clarity over cleverness

---

## What is Athera?

Athera is a general-purpose programming language designed to be **as capable as Python and Java**, while being **clearer, safer, and more intent-driven**.

It's not a toy. It's a serious professional language with:
- Clear execution model
- Object capability system (the "dot rule")
- Built-in error handling
- Built-in concurrency
- Professional-grade features

---

## Why Athera?

### The Problem
```python
# Python: Confusing syntax obscures intent
file.exists()          # Is this a property or method?
data = file.read()     # Hidden mechanics
file.delete()          # Side effect! Hard to spot!
```

### The Solution
```athera
# Athera: Intent is clear
check file.exists:     # Property: safe to read
    set data = file.read   # Action: produces value
    file.delete            # Action: visible side effect
```

**What Athera provides:**
✅ Clearer code that reads like instructions  
✅ Safety by default (no hidden side effects)  
✅ Explicit concurrency (`run parallel`)  
✅ Explicit error handling (`protect/handle`)  
✅ Professional capabilities  

---

## Core Features

### 1. Natural Syntax
```athera
task greet_user with name:
    greet "Hello, name!"
    return name

run greet_user "Alice"
```

### 2. The Dot Rule (Properties vs Actions)
```athera
# Properties: read-only, no side effects
check file.exists
set length = list.length
set name = person.name

# Actions: explicit, intentional
set data = file.read
file.delete
list.append 42
```

### 3. Clear Error Handling
```athera
protect:
    set data = file.read
    run process data
handle:
    greet "Processing failed, using defaults"
```

### 4. Built-in Concurrency
```athera
run parallel download, process, upload
```

### 5. Custom Modules
```athera
use math_utils
use file_utils
run add 5, 10
```

---

## Quick Start

### Installation
```bash
# Clone the repository
git clone https://github.com/Arths17/Athena.git
cd Athena

# Run examples
python3 athera_interpreter.py morning_routine.ath
```

### Create Your First Program
```bash
# Create a file
nano hello.ath
```

```athera
task hello:
    greet "Hello, Athera!"
    greet "I'm writing professional code!"

run hello
```

```bash
# Run it
python3 athera_interpreter.py hello.ath
```

---

## Language Design

Athera is built on documented principles:

1. **Intent over Mechanics** - Code shows WHAT happens, not HOW
2. **Words First** - Symbols only when they improve clarity  
3. **One Clear Way** - No multiple syntax paths for the same idea
4. **Safety by Default** - Error handling is explicit and visible
5. **The Dot Rule** - Properties never have side effects; actions are always visible

Read the full design brief: [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md)

---

## File Structure

```
Athera/
├── athera_interpreter.py        # Main interpreter (1000+ lines)
├── DESIGN_BRIEF_V2.1.md        # Professional design document
├── DOT_RULE_GUIDE.md           # The dot rule explained
├── ADVANCED_FEATURES.md        # Advanced language features
├── GETTING_STARTED.md          # Tutorial for beginners
│
├── Example Programs/
│   ├── morning_routine.ath     # Original example
│   ├── advanced_features.ath   # All features demo
│   ├── module_demo.ath         # Custom modules
│   └── ... (12 total examples)
│
├── Modules/
│   ├── math_utils.ath
│   └── file_utils.ath
│
└── Documentation/
    ├── README.md               # Original readme
    ├── QUICKREF.md             # Quick reference
    └── README_V2.md            # Updated docs
```

---

## Design Highlights

### What's Strong ✅
- **Clear goal:** Capable as Python/Java, but clearer
- **Correct focus:** Intent over mechanics
- **Good constraints:** One way to express each idea
- **Real pain points:** Addresses what professionals complain about
- **Differentiator:** The dot rule is a genuine innovation

### What Makes It Professional 🎯
- Clear execution model (sequential by default, parallel by declaration)
- Object capability system (properties vs actions)
- Safety guarantees (properties are side-effect free)
- Designed for scale (from scripts to systems)
- Honest principles (not copied from other languages)

### The Dot Rule (Critical) 🔵
```
In Athera, . exposes readable properties and explicit actions.
- Properties NEVER cause side effects
- Actions ARE visible and intentional
- This contract is guaranteed by the language
```

This single rule creates predictability, safety, and clarity.

Read more: [DOT_RULE_GUIDE.md](DOT_RULE_GUIDE.md)

---

## Example: File Processing

### Python (Confusing)
```python
if file.exists():
    data = file.read()
    lines = data.split("\n")
    for line in lines:
        process(line)
    file.delete()
```

### Athera (Clear)
```athera
check file.exists:
    set data = file.read
    set lines = data.split "\n"
    repeat each line in lines:
        run process line
    file.delete
```

**What Athera adds:**
- `file.exists` is clearly a property (no computation)
- `file.read` clearly produces a value
- `split` is clearly a property (no side effects)
- `file.delete` is clearly an action (side effect, visible)

---

## Comparison

| Feature | Python | Java | Athera |
|---------|--------|------|--------|
| Readability | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ |
| Safety | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| Clarity | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ |
| Concurrency | ⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| Learning curve | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ |
| Boilerplate | ⭐⭐ | ⭐ | ⭐⭐⭐ |

---

## What This Is NOT

❌ Not a toy DSL  
❌ Not a Python clone  
❌ Not a JavaScript wrapper  
❌ Not designed to be "cute"  

✅ A professional programming language  
✅ Designed for clarity and safety  
✅ Built on documented principles  
✅ Capable of real systems  

---

## Getting Started

1. **Read the design brief:** [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md)
2. **Understand the dot rule:** [DOT_RULE_GUIDE.md](DOT_RULE_GUIDE.md)
3. **Learn the language:** [GETTING_STARTED.md](GETTING_STARTED.md)
4. **Run examples:** `python3 athera_interpreter.py morning_routine.ath`
5. **Create your own:** Copy `your_program_template.ath` and start coding

---

## Project Status

### ✅ Complete
- Working interpreter with lexer, parser, and tree-walking executor
- 12 example programs (from basics to advanced)
- 2 custom modules
- Error handling (protect/handle)
- Parallel execution (run parallel)
- Function parameters
- Return values
- Module system

### 🔄 In Progress
- Dot notation implementation for properties and actions
- Built-in capabilities for core types

### 📋 Planned
- Optional type hints
- Standard library
- Capability scopes
- Advanced safety features

---

## Documentation

| Document | Purpose |
|----------|---------|
| [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md) | Professional design document |
| [DOT_RULE_GUIDE.md](DOT_RULE_GUIDE.md) | The dot rule explained in detail |
| [ADVANCED_FEATURES.md](ADVANCED_FEATURES.md) | Language features guide |
| [GETTING_STARTED.md](GETTING_STARTED.md) | Beginner tutorial |
| [QUICKREF.md](QUICKREF.md) | Syntax quick reference |
| [README.md](README.md) | Original documentation |

---

## Contributing

This is a professional language project. Contributions that:
- Improve clarity and safety
- Follow the design principles
- Include documentation
- Have working examples

are welcome!

---

## License

MIT License - Free to use, modify, and distribute

---

## The Philosophy

> "Code should express intent, not mechanics."

Athera is built for professionals who value:
- **Clarity** over cleverness
- **Safety** over convenience  
- **Intent** over implementation details
- **Explicitness** over magic

This is a serious professional programming language.

---

**Repository:** https://github.com/Arths17/Athena  
**Design Brief:** See [DESIGN_BRIEF_V2.1.md](DESIGN_BRIEF_V2.1.md)  
**Get Started:** Run `python3 athera_interpreter.py morning_routine.ath`
