# Athera Standard Library Design

**Version:** 1.0  
**Format:** Standard library specification  
**Purpose:** Built-in modules for essential tasks

---

## Overview

The Athera Standard Library (stdlib) provides essential functionality for all programs. It's:
- **Built-in** - Comes with Athera, no installation needed
- **Minimal** - Only what beginners need to be productive
- **Professional** - Used by enterprise applications
- **Extensible** - Advanced users add custom modules
- **Consistent** - Follows Athera's design principles (dot rule)

---

## Design Principles

### 1. Small Core
- Only essential functionality in stdlib
- Keep binary size reasonable (~30MB)
- Advanced features as optional modules

### 2. Consistent API
- All functions use consistent naming
- Dot notation for properties vs actions
- Predictable behavior

### 3. The Dot Rule Applied
```athera
# Properties (read-only, no side effects)
set length = list.length
set exists = file.exists
set name = dict.keys

# Actions (explicit, visible side effects)
set data = file.read
file.write data
list.append item
```

### 4. Beginner-Friendly
- Simple functions for common tasks
- Clear naming
- Good error messages

---

## Module Structure

### Import Pattern

```athera
# Import module
use io
use text
use math

# Use functions
set data = io.read "file.txt"
set lines = text.split data "\n"
set total = math.add 5 10

# Or without import (built-in functions)
greet "Hello"                  # Built-in
set x = 5                      # Built-in
repeat 3 times: greet "Hi"     # Built-in
```

---

## Module Reference

### 1. IO Module - File Operations

**Purpose:** Read/write files, manage directories

```athera
use io

# Properties
set exists = file.exists               # → boolean
set is_file = file.is_file             # → boolean
set is_dir = file.is_directory         # → boolean
set size = file.size                   # → number (bytes)

# Actions - Read
set data = file.read                   # → string
set lines = file.read_lines            # → list of strings
set binary = file.read_binary          # → bytes

# Actions - Write
file.write data
file.append data                       # Append to end
file.write_lines list                  # Write list of strings

# File Management
file.delete
file.copy destination_path
file.move destination_path
file.chmod 0755                        # Change permissions

# Directory Operations
set files = dir.list                   # → list of filenames
set entries = dir.list_all             # → list with types
dir.create
dir.delete
dir.exists
dir.is_empty

# Path Operations
set parent = file.parent               # → path
set name = file.name                   # → filename only
set extension = file.extension         # → ".txt"
```

**Error Handling:**
```athera
use io

protect:
    set data = file.read "missing.txt"
handle:
    greet "File not found"
```

**Common Tasks:**
```athera
# Read entire file
set content = file.read "poem.txt"
greet content

# Read line by line
set lines = file.read_lines "data.txt"
repeat each line in lines:
    greet "Line: " + line

# Write text
file.write "output.txt" "Hello, World!"

# Append to file
file.append "log.txt" "New entry"

# Copy file
file.copy "original.txt" "backup.txt"

# List files in directory
set files = dir.list "."
repeat each file in files:
    greet file
```

**Module Signature:**
```
io.read(path: string) → string
io.write(path: string, data: string) → void
io.append(path: string, data: string) → void
io.read_lines(path: string) → list
io.write_lines(path: string, lines: list) → void
io.read_binary(path: string) → bytes
io.delete(path: string) → void
io.copy(source: string, dest: string) → void
io.move(source: string, dest: string) → void
io.chmod(path: string, mode: number) → void
io.exists(path: string) → boolean
io.is_file(path: string) → boolean
io.is_directory(path: string) → boolean
io.size(path: string) → number

io.dir.list(path: string) → list
io.dir.list_all(path: string) → list
io.dir.create(path: string) → void
io.dir.delete(path: string) → void
io.dir.exists(path: string) → boolean
io.dir.is_empty(path: string) → boolean
```

---

### 2. TEXT Module - String Operations

**Purpose:** Manipulate strings and text

```athera
use text

# Properties (no side effects)
set length = string.length             # → number
set is_empty = string.is_empty         # → boolean
set is_alpha = string.is_alpha         # → boolean
set is_numeric = string.is_numeric     # → boolean

# Query Properties
set contains = string.contains "hello"  # → boolean
set starts_with = string.starts_with "The"  # → boolean
set ends_with = string.ends_with "!"   # → boolean
set index = string.index_of "world"    # → number (-1 if not found)

# Actions - Transform
set upper = string.uppercase           # → "HELLO"
set lower = string.lowercase           # → "hello"
set trimmed = string.trim              # → "hello" (remove whitespace)
set replaced = string.replace "a" "b"  # → "hello" with replacements

# Actions - Extract
set substr = string.substring 0 5      # → first 5 chars
set split = string.split ","           # → list of strings
set chars = string.chars               # → list of characters

# Formatting
set formatted = text.format "Hello, {}"  [name]
set interpolated = text.interpolate "x={x}, y={y}"  [x, y]
set joined = text.join ", "  ["a", "b", "c"]  # → "a, b, c"

# Case Conversion
set title = string.title_case          # → "Hello World"
set snake = string.to_snake_case       # → "hello_world"
set camel = string.to_camel_case       # → "helloWorld"
```

**Common Tasks:**
```athera
use text

# Convert to uppercase
set greeting = "hello world"
greet greeting.uppercase              # Output: HELLO WORLD

# Split and process
set csv = "apple,banana,orange"
set fruits = csv.split ","
repeat each fruit in fruits:
    greet "Fruit: " + fruit.trim

# Check if contains
check "hello world".contains "world":
    greet "Found 'world'"

# Replace text
set original = "The quick brown fox"
set replaced = original.replace "brown" "red"
greet replaced                        # Output: The quick red fox

# Format string
set name = "Alice"
set age = 30
greet text.format "Name: {}, Age: {}" [name, age]
```

**Module Signature:**
```
text.uppercase(string: string) → string
text.lowercase(string: string) → string
text.trim(string: string) → string
text.replace(string: string, old: string, new: string) → string
text.split(string: string, sep: string) → list
text.join(list: list, sep: string) → string
text.substring(string: string, start: number, length: number) → string
text.contains(string: string, sub: string) → boolean
text.starts_with(string: string, prefix: string) → boolean
text.ends_with(string: string, suffix: string) → boolean
text.index_of(string: string, sub: string) → number
text.length(string: string) → number
text.is_empty(string: string) → boolean
text.is_alpha(string: string) → boolean
text.is_numeric(string: string) → boolean
text.title_case(string: string) → string
text.to_snake_case(string: string) → string
text.to_camel_case(string: string) → string
text.format(template: string, args: list) → string
text.interpolate(template: string, vars: dict) → string
text.chars(string: string) → list
```

---

### 3. MATH Module - Arithmetic & Math Functions

**Purpose:** Mathematical operations and calculations

```athera
use math

# Basic Arithmetic (also built-in +, -, *, /)
set result = math.add 5 3              # → 8
set result = math.subtract 10 3        # → 7
set result = math.multiply 4 5         # → 20
set result = math.divide 20 4          # → 5
set result = math.modulo 17 5          # → 2
set result = math.power 2 8            # → 256

# Rounding
set rounded = math.round 3.7           # → 4
set floor = math.floor 3.7             # → 3
set ceil = math.ceil 3.2               # → 4

# Advanced
set root = math.sqrt 16                # → 4
set absolute = math.abs -5             # → 5
set max = math.max 5 10 3              # → 10
set min = math.min 5 10 3              # → 3

# Trigonometry (radians)
set sine = math.sin 0                  # → 0
set cosine = math.cos 0                # → 1
set tangent = math.tan 0               # → 0

# Logarithms
set log = math.log 100                 # Natural log
set log10 = math.log10 100             # Base 10 log
set log2 = math.log2 8                 # Base 2 log

# Random Numbers
set random = math.random               # 0 to 1
set random_int = math.random_int 1 100 # 1 to 100
set random_choice = math.random_choice list  # Random from list

# Constants
set pi = math.pi                       # 3.14159...
set e = math.e                         # 2.71828...
```

**Common Tasks:**
```athera
use math

# Calculate average
set total = 10 + 20 + 30
set average = total / 3
greet "Average: " + average

# Random selection
set options = ["rock", "paper", "scissors"]
set choice = math.random_choice options

# Round currency
set price = 19.9999
set rounded = math.round price         # 20

# Check if in range
set value = 50
check math.min(0, value) >= 0:
    check math.max(100, value) <= 100:
        greet "Value in range"
```

**Module Signature:**
```
math.add(a: number, b: number) → number
math.subtract(a: number, b: number) → number
math.multiply(a: number, b: number) → number
math.divide(a: number, b: number) → number
math.modulo(a: number, b: number) → number
math.power(base: number, exp: number) → number
math.sqrt(n: number) → number
math.abs(n: number) → number
math.round(n: number) → number
math.floor(n: number) → number
math.ceil(n: number) → number
math.max(...numbers: number) → number
math.min(...numbers: number) → number
math.sin(radians: number) → number
math.cos(radians: number) → number
math.tan(radians: number) → number
math.log(n: number) → number
math.log10(n: number) → number
math.log2(n: number) → number
math.random() → number (0-1)
math.random_int(min: number, max: number) → number
math.random_choice(list: list) → any

math.pi → number
math.e → number
```

---

### 4. LIST Module - Array Operations

**Purpose:** Work with lists/arrays

```athera
use list

# Properties
set length = my_list.length            # → number
set is_empty = my_list.is_empty        # → boolean
set first = my_list.first              # → first element
set last = my_list.last                # → last element

# Query
set contains = my_list.contains 5      # → boolean
set index = my_list.index_of 5         # → position (-1 if not found)
set count = my_list.count 5            # → how many times appears

# Actions - Modify
my_list.append item                    # Add to end
my_list.prepend item                   # Add to beginning
my_list.insert index item              # Insert at position
my_list.remove index                   # Remove by index
my_list.remove_value value             # Remove by value
my_list.clear                          # Remove all

# Actions - Reorder
set sorted = my_list.sort              # → new sorted list
set reversed = my_list.reverse         # → reversed list
set shuffled = my_list.shuffle         # → randomized list

# Actions - Transform
set doubled = my_list.map double       # Apply task to each
set filtered = my_list.filter is_even  # Keep matching items
set flattened = my_list.flatten        # Flatten nested lists
set unique = my_list.unique            # Remove duplicates

# Actions - Combine
set joined = my_list.join ", "         # → string
set combined = list.concat list1 list2 # Combine lists
set slice = my_list.slice 0 5          # Get elements 0-5

# Convert
set string = my_list.to_string         # → "[1, 2, 3]"
```

**Common Tasks:**
```athera
use list

# Create and add items
set numbers = [1, 2, 3]
numbers.append 4
numbers.append 5
greet numbers                          # [1, 2, 3, 4, 5]

# Remove items
set items = ["a", "b", "c", "b"]
items.remove_value "b"
greet items                            # ["a", "c"]

# Filter list
set numbers = [1, 2, 3, 4, 5]
set evens = numbers.filter is_even
greet evens                            # [2, 4]

# Sort list
set names = ["charlie", "alice", "bob"]
set sorted = names.sort
greet sorted                           # ["alice", "bob", "charlie"]

# Map transformation
set numbers = [1, 2, 3]
set doubled = numbers.map double
greet doubled                          # [2, 4, 6]
```

**Module Signature:**
```
list.length(list: list) → number
list.is_empty(list: list) → boolean
list.first(list: list) → any
list.last(list: list) → any
list.append(list: list, item: any) → void
list.prepend(list: list, item: any) → void
list.insert(list: list, index: number, item: any) → void
list.remove(list: list, index: number) → void
list.remove_value(list: list, value: any) → void
list.clear(list: list) → void
list.sort(list: list) → list
list.reverse(list: list) → list
list.shuffle(list: list) → list
list.contains(list: list, item: any) → boolean
list.index_of(list: list, item: any) → number
list.count(list: list, item: any) → number
list.join(list: list, sep: string) → string
list.concat(list1: list, list2: list) → list
list.slice(list: list, start: number, end: number) → list
list.map(list: list, task: function) → list
list.filter(list: list, task: function) → list
list.flatten(list: list) → list
list.unique(list: list) → list
list.to_string(list: list) → string
```

---

### 5. DICT Module - Dictionary/Map Operations

**Purpose:** Work with dictionaries/maps

```athera
use dict

# Properties
set length = my_dict.length            # → number of keys
set keys = my_dict.keys                # → list of keys
set values = my_dict.values            # → list of values
set is_empty = my_dict.is_empty        # → boolean

# Query
set value = my_dict.get key            # → value (or null)
set exists = my_dict.has key           # → boolean
set has_value = my_dict.contains_value value  # → boolean

# Actions - Modify
my_dict.set key value                  # Add/update
my_dict.delete key                     # Remove key
my_dict.clear                          # Remove all

# Actions - Transform
set merged = dict.merge dict1 dict2    # Combine dicts
set filtered = my_dict.filter predicate # Keep matching
set mapped = my_dict.map transform     # Transform values

# Convert
set json = my_dict.to_json             # → JSON string
set pairs = my_dict.to_pairs            # → list of [key, value]
```

**Common Tasks:**
```athera
use dict

# Create dictionary
set user = {}
user.set "name" "Alice"
user.set "age" 30
user.set "city" "NYC"

# Access values
greet user.get "name"                  # Output: Alice

# Check if key exists
check user.has "email":
    greet user.get "email"
handle:
    greet "Email not set"

# Get all keys
set keys = user.keys
repeat each key in keys:
    greet key + ": " + user.get key

# Merge dictionaries
set defaults = {}
defaults.set "theme" "light"
defaults.set "language" "en"

set config = dict.merge defaults user  # user values override defaults
```

**Module Signature:**
```
dict.length(dict: dict) → number
dict.keys(dict: dict) → list
dict.values(dict: dict) → list
dict.is_empty(dict: dict) → boolean
dict.get(dict: dict, key: string) → any
dict.set(dict: dict, key: string, value: any) → void
dict.delete(dict: dict, key: string) → void
dict.has(dict: dict, key: string) → boolean
dict.contains_value(dict: dict, value: any) → boolean
dict.clear(dict: dict) → void
dict.merge(dict1: dict, dict2: dict) → dict
dict.filter(dict: dict, predicate: function) → dict
dict.map(dict: dict, transform: function) → dict
dict.to_json(dict: dict) → string
dict.to_pairs(dict: dict) → list
```

---

### 6. TIME Module - Date/Time Operations

**Purpose:** Work with dates and times

```athera
use time

# Current Time
set now = time.now                     # → current timestamp
set current_year = time.year            # → 2026
set current_month = time.month          # → 1-12
set current_day = time.day              # → 1-31
set current_hour = time.hour            # → 0-23

# Create Time
set ts = time.from_date 2026 1 27      # → timestamp for that date
set ts = time.from_datetime 2026 1 27 14 30 45  # Date + time

# Format
set formatted = time.format now "%Y-%m-%d"  # → "2026-01-27"
set formatted = time.format now "%H:%M:%S"  # → "14:30:45"

# Parse
set ts = time.parse "2026-01-27" "%Y-%m-%d"

# Difference
set delta = time.difference ts1 ts2    # → milliseconds between

# Delay
time.sleep 2                           # Sleep 2 seconds

# Timeouts
protect:
    set result = run_with_timeout task 5  # 5 second timeout
handle:
    greet "Task timed out"
```

**Common Tasks:**
```athera
use time

# Get current time
set now = time.now
greet "Current time: " + time.format now "%H:%M:%S"

# Wait between actions
greet "Starting..."
time.sleep 1
greet "Done!"

# Calculate time difference
set start = time.now
run expensive_task
set end = time.now
set elapsed = time.difference start end
greet "Took " + (elapsed / 1000) + " seconds"
```

**Module Signature:**
```
time.now() → timestamp
time.year() → number
time.month() → number
time.day() → number
time.hour() → number
time.minute() → number
time.second() → number
time.from_date(year: number, month: number, day: number) → timestamp
time.from_datetime(year: number, month: number, day: number, 
                   hour: number, minute: number, second: number) → timestamp
time.format(timestamp: number, format: string) → string
time.parse(date_string: string, format: string) → timestamp
time.difference(ts1: number, ts2: number) → number (milliseconds)
time.sleep(seconds: number) → void
```

---

### 7. JSON Module - JSON Parsing

**Purpose:** Parse and generate JSON

```athera
use json

# Parse JSON
set data = json.parse '{"name": "Alice", "age": 30}'
greet data.get "name"                  # Output: Alice

# Check validity
check json.is_valid json_string:
    set obj = json.parse json_string
handle:
    greet "Invalid JSON"

# Generate JSON
set person = {}
person.set "name" "Bob"
person.set "age" 25
set json_string = json.stringify person

# Pretty print
set pretty = json.pretty_string person
greet pretty
```

**Module Signature:**
```
json.parse(json_string: string) → dict/list
json.stringify(object: any) → string
json.pretty_string(object: any) → string
json.is_valid(json_string: string) → boolean
json.pretty_print(object: any) → void
```

---

### 8. PATH Module - File Path Operations

**Purpose:** Work with file paths

```athera
use path

# Construct paths
set full_path = path.join "folder" "subfolder" "file.txt"
set path_str = path.to_string full_path

# Deconstruct paths
set directory = path.dirname "/home/alice/file.txt"  # "/home/alice"
set filename = path.basename "/home/alice/file.txt"  # "file.txt"
set extension = path.extension "/home/alice/file.txt"  # ".txt"
set stem = path.stem "/home/alice/file.txt"  # "file"

# Normalize
set absolute = path.absolute "."       # Full path
set normalized = path.normalize path   # Clean up path
set home = path.home                   # Home directory

# Check
set exists = path.exists path_str      # → boolean
set is_file = path.is_file path_str    # → boolean
set is_dir = path.is_directory path_str # → boolean

# Relative paths
set relative = path.relative "/home/alice" "/home/alice/files/file.txt"
```

**Common Tasks:**
```athera
use path

# Build file path
set base_dir = path.home
set my_file = path.join base_dir "Documents" "notes.txt"

# Check if file exists
check path.exists my_file:
    set content = io.read my_file
handle:
    greet "File doesn't exist"

# Get file extension
set file_path = "document.pdf"
set ext = path.extension file_path
check ext == ".pdf":
    greet "Is a PDF file"
```

**Module Signature:**
```
path.join(...components: string) → path
path.dirname(path: string) → string
path.basename(path: string) → string
path.extension(path: string) → string
path.stem(path: string) → string
path.absolute(path: string) → string
path.normalize(path: string) → string
path.home() → string
path.exists(path: string) → boolean
path.is_file(path: string) → boolean
path.is_directory(path: string) → boolean
path.to_string(path: path) → string
path.relative(base: string, target: string) → string
```

---

## Built-in Functions (No Import Needed)

These are always available without `use`:

```athera
# Output
greet message                  # Print to console
print message                  # Same as greet

# Variables
set name = value              # Assign variable
return value                  # Return from task

# Control Flow
check condition:              # If
repeat N times:               # Loop N times
repeat each item in list:     # Iterate list
run parallel task1, task2     # Parallel execution

# Task Definition
task name:                    # Define task
task name with param1, param2:  # With parameters

# Error Handling
protect:                      # Try block
handle:                       # Catch block

# Type Checking
type_of value                 # Get type
is_string value               # Check type
is_number value
is_boolean value
is_list value
is_dict value

# Comments
# This is a comment
```

---

## Standard Library in Numbers

| Module | Functions | Lines of Code | Size |
|--------|-----------|---------------|------|
| io | 15 | 400 | 15 KB |
| text | 18 | 450 | 18 KB |
| math | 24 | 350 | 12 KB |
| list | 20 | 380 | 14 KB |
| dict | 15 | 320 | 12 KB |
| time | 10 | 250 | 10 KB |
| json | 5 | 200 | 8 KB |
| path | 12 | 280 | 11 KB |
| **Total** | **119** | **2,830** | **100 KB** |

**Embedded in Runtime:**
- Compiles into athera binary
- No external files needed
- Always available
- Automatic on import

---

## Performance Characteristics

### Speed
- All stdlib functions are optimized
- Typically 10-100x faster than naive implementations
- Compiled stdlib (not interpreted)

### Memory
- Lazy loading (modules only loaded when imported)
- Minimal overhead
- Efficient data structures

---

## Future Modules (Phase 2+)

### Phase 2
- `network` - HTTP requests, sockets
- `regex` - Regular expressions
- `crypto` - Hashing, encryption basics

### Phase 3
- `database` - SQL queries
- `csv` - CSV parsing
- `xml` - XML parsing
- `html` - HTML manipulation

### Phase 4
- `async` - Async/await patterns
- `graphics` - Image processing
- `audio` - Audio processing
- `ml` - Machine learning basics

---

## Standard Library Philosophy

The Athera Standard Library embodies the design principles:

1. **Intent over Mechanics**
   - `file.read` clearly expresses intent
   - Not `open().read().close()` boilerplate

2. **Words First**
   - `text.to_uppercase` not `str.upper()`
   - Clear naming over abbreviations

3. **One Clear Way**
   - `list.append item` is the only way to add
   - No multiple syntaxes for same operation

4. **Safety Visible by Default**
   - Properties (`file.exists`) never have side effects
   - Actions (`file.delete`) are always visible
   - Error handling is explicit

5. **Familiar Concepts, Unfamiliar Clarity**
   - Functions developers know (read, write, sort)
   - But clearer API design
   - No hidden side effects

This makes Athera's stdlib better for professionals and beginners alike.
