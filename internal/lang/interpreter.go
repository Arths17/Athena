package lang

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "sync"
)

// Interpreter executes Athera programs.
type Interpreter struct {
    tasks         map[string]TaskDef
    variables     map[string]any
    modules       map[string]bool
    returnValue   any
    errorOccurred bool
    lastError     error
    stdlib        map[string]map[string]BuiltinFunc
}

// TaskDef stores a task body and parameter list.
type TaskDef struct {
    Body   []Node
    Params []string
}

// NewInterpreter creates a fresh interpreter instance.
func NewInterpreter() *Interpreter {
    return &Interpreter{
        tasks:     make(map[string]TaskDef),
        variables: make(map[string]any),
        modules:   make(map[string]bool),
        stdlib:    StdlibModules,
    }
}

// Execute runs a list of AST nodes.
func (i *Interpreter) Execute(nodes []Node) {
    for _, n := range nodes {
        i.executeNode(n)
    }
}

func (i *Interpreter) executeNode(n Node) {
    switch node := n.(type) {
    case *TaskNode:
        i.tasks[node.Name] = TaskDef{Body: node.Body, Params: node.Params}
    case *GreetNode:
        msg := i.evaluateExpression(node.Message)
        fmt.Println(toString(msg))
    case *BackupNode:
        i.executeBackup(node)
    case *CheckNode:
        if i.evaluateCondition(node.Condition) {
            i.executeInlineAction(node.Action)
        }
    case *RepeatNNode:
        for idx := 0; idx < node.Count; idx++ {
            for _, stmt := range node.Body {
                i.executeNode(stmt)
            }
        }
    case *RepeatEachNode:
        listVal := i.evaluateExpression(node.ListExpr)
        arr, ok := listVal.([]any)
        if !ok {
            // attempt []string fallback
            if sArr, okStr := listVal.([]string); okStr {
                for _, item := range sArr {
                    i.variables[node.Var] = item
                    for _, stmt := range node.Body {
                        i.executeNode(stmt)
                    }
                }
                return
            }
            fmt.Printf("[Warning: expected list, got %T]\n", listVal)
            return
        }
        for _, item := range arr {
            i.variables[node.Var] = item
            for _, stmt := range node.Body {
                i.executeNode(stmt)
            }
        }
    case *SetNode:
        i.variables[node.Var] = i.evaluateExpression(node.Value)
    case *RunNode:
        i.executeRun(node)
    case *UseNode:
        i.executeUse(node)
    case *ProtectNode:
        i.executeProtect(node)
    case *HandleInlineNode:
        if i.errorOccurred {
            i.executeInlineAction(node.Action)
        }
    case *RunParallelNode:
        i.executeRunParallel(node)
    case *ReturnNode:
        i.returnValue = i.evaluateExpression(node.Expr)
    }
}

func (i *Interpreter) executeBackup(node *BackupNode) {
    src := toString(i.evaluateExpression(node.Source))
    dst := toString(i.evaluateExpression(node.Dest))

    src = strings.Trim(src, "\"'")
    dst = strings.Trim(dst, "\"'")

    if src == "" || dst == "" {
        fmt.Println("[Backup error: source or destination missing]")
        return
    }

    if err := os.MkdirAll(dst, 0o755); err != nil {
        fmt.Printf("[Backup error: %v]\n", err)
        return
    }

    info, err := os.Stat(src)
    if err != nil {
        fmt.Printf("[Backup error: %v]\n", err)
        return
    }

    destPath := filepath.Join(dst, filepath.Base(src))

    if info.IsDir() {
        fmt.Printf("[Backup warning: directory backup not yet supported: %s]\n", src)
        return
    }

    in, err := os.Open(src)
    if err != nil {
        fmt.Printf("[Backup error: %v]\n", err)
        return
    }
    defer in.Close()

    out, err := os.Create(destPath)
    if err != nil {
        fmt.Printf("[Backup error: %v]\n", err)
        return
    }
    defer out.Close()

    if _, err := io.Copy(out, in); err != nil {
        fmt.Printf("[Backup error: %v]\n", err)
        return
    }

    fmt.Printf("[Backed up: %s -> %s]\n", src, destPath)
}

func (i *Interpreter) executeConditionAction(action string) {
    i.executeInlineAction(action)
}

func (i *Interpreter) evaluateCondition(expr string) bool {
    expr = strings.TrimSpace(expr)

    // quoted path check
    if strings.HasPrefix(expr, "\"") || strings.HasPrefix(expr, "'") {
        path := strings.Trim(expr, "\"'")
        path = toString(i.evaluateExpression(path))
        _, err := os.Stat(path)
        return err == nil
    }

    if val, ok := i.variables[expr]; ok {
        switch v := val.(type) {
        case bool:
            return v
        case int:
            return v != 0
        case float64:
            return v != 0
        case string:
            return v != ""
        case []any:
            return len(v) > 0
        case []string:
            return len(v) > 0
        }
    }

    // fall back to arithmetic-like evaluation
    if res := i.evaluateExpression(expr); res != nil {
        switch v := res.(type) {
        case bool:
            return v
        case int:
            return v != 0
        case float64:
            return v != 0
        case string:
            return v != ""
        }
    }

    return false
}

func (i *Interpreter) executeInlineAction(action string) {
    action = strings.TrimSpace(action)
    if strings.HasPrefix(action, "greet ") {
        msg := strings.TrimSpace(action[len("greet "):])
        msgVal := i.evaluateExpression(msg)
        fmt.Println(toString(msgVal))
    }
}

func (i *Interpreter) executeRun(node *RunNode) {
    target := strings.TrimSpace(node.Target)
    name := target
    var args []string

    if strings.Contains(target, " ") {
        parts := strings.SplitN(target, " ", 2)
        name = parts[0]
        args = splitArgsRespectingQuotes(parts[1])
    }

    def, ok := i.tasks[name]
    if !ok {
        fmt.Printf("[Error: task '%s' not found]\n", name)
        return
    }

    saved := copyMap(i.variables)

    for idx, param := range def.Params {
        if idx < len(args) {
            i.variables[param] = i.evaluateExpression(args[idx])
        } else {
            i.variables[param] = nil
        }
    }

    i.returnValue = nil
    for _, stmt := range def.Body {
        i.executeNode(stmt)
        if i.returnValue != nil {
            break
        }
    }

    for _, param := range def.Params {
        if val, exists := saved[param]; exists {
            i.variables[param] = val
        } else {
            delete(i.variables, param)
        }
    }
}

func (i *Interpreter) executeUse(node *UseNode) {
    name := strings.TrimSpace(node.Module)
    if _, ok := StdlibModules[name]; ok {
        i.modules[name] = true
        fmt.Printf("[Imported built-in module: %s]\n", name)
        return
    }

    // Future: load external .ath modules if present.
    fmt.Printf("[Warning: module %s not found]\n", name)
}

func (i *Interpreter) executeProtect(node *ProtectNode) {
    i.errorOccurred = false
    i.lastError = nil

    defer func() {
        if r := recover(); r != nil {
            i.errorOccurred = true
            i.lastError = fmt.Errorf("%v", r)
            fmt.Printf("[Error caught: %v]\n", r)
            for _, stmt := range node.Handle {
                i.executeNode(stmt)
            }
        }
    }()

    for _, stmt := range node.Protect {
        i.executeNode(stmt)
        if i.errorOccurred {
            break
        }
    }
}

func (i *Interpreter) executeRunParallel(node *RunParallelNode) {
    var wg sync.WaitGroup
    for _, taskName := range node.Tasks {
        tn := strings.TrimSpace(taskName)
        def, ok := i.tasks[tn]
        if !ok {
            fmt.Printf("[Warning: task %s not found for parallel run]\n", tn)
            continue
        }

        wg.Add(1)
        go func(td TaskDef) {
            defer wg.Done()
            // copy variables for isolation
            local := NewInterpreter()
            local.tasks = i.tasks
            local.variables = copyMap(i.variables)
            for _, stmt := range td.Body {
                local.executeNode(stmt)
            }
        }(def)
    }

    wg.Wait()
    fmt.Printf("[Parallel execution complete: %d tasks]\n", len(node.Tasks))
}

// evaluateExpression resolves literals, variables, and stdlib calls.
func (i *Interpreter) evaluateExpression(expr string) any {
    expr = strings.TrimSpace(expr)
    if expr == "" {
        return ""
    }

    if (strings.HasPrefix(expr, "\"") && strings.HasSuffix(expr, "\"")) || (strings.HasPrefix(expr, "'") && strings.HasSuffix(expr, "'")) {
        return strings.Trim(expr, "\"'")
    }

    if strings.HasPrefix(expr, "[") && strings.HasSuffix(expr, "]") {
        return i.parseList(expr)
    }

    if strings.EqualFold(expr, "true") {
        return true
    }
    if strings.EqualFold(expr, "false") {
        return false
    }

    if v, err := strconv.Atoi(expr); err == nil {
        return v
    }
    if f, err := strconv.ParseFloat(expr, 64); err == nil {
        return f
    }

    if val, ok := i.variables[expr]; ok {
        return val
    }

    // property access like myList.length
    if strings.Contains(expr, ".") && !strings.Contains(expr, " ") {
        parts := strings.Split(expr, ".")
        if len(parts) == 2 {
            base := parts[0]
            prop := parts[1]
            if val, ok := i.variables[base]; ok {
                switch prop {
                case "length":
                    switch v := val.(type) {
                    case string:
                        return len([]rune(v))
                    case []any:
                        return len(v)
                    case []string:
                        return len(v)
                    }
                case "is_empty":
                    switch v := val.(type) {
                    case string:
                        return v == ""
                    case []any:
                        return len(v) == 0
                    case []string:
                        return len(v) == 0
                    }
                }
            }
        }
    }

    if strings.Contains(expr, ".") {
        modCall := strings.SplitN(expr, ".", 2)
        modName := strings.TrimSpace(modCall[0])
        if len(modCall) == 2 {
            rest := strings.TrimSpace(modCall[1])
            fnParts := strings.SplitN(rest, " ", 2)
            fnName := strings.TrimSpace(fnParts[0])
            argStr := ""
            if len(fnParts) > 1 {
                argStr = fnParts[1]
            }

            if modFuncs, ok := i.stdlib[modName]; ok {
                if fn, okFn := modFuncs[fnName]; okFn {
                    rawArgs := splitArgsRespectingQuotes(argStr)
                    args := make([]any, 0, len(rawArgs))
                    for _, a := range rawArgs {
                        args = append(args, i.evaluateExpression(a))
                    }
                    res, err := fn(args)
                    if err != nil {
                        fmt.Printf("[Error: %v]\n", err)
                        return nil
                    }
                    return res
                }
            }
        }
    }

    // simple addition support: a + b
    if strings.Contains(expr, "+") {
        parts := strings.Split(expr, "+")
        if len(parts) == 2 {
            left := i.evaluateExpression(parts[0])
            right := i.evaluateExpression(parts[1])
            switch l := left.(type) {
            case int:
                if r, ok := right.(int); ok {
                    return l + r
                }
                if r, ok := right.(float64); ok {
                    return float64(l) + r
                }
            case float64:
                if r, ok := right.(int); ok {
                    return l + float64(r)
                }
                if r, ok := right.(float64); ok {
                    return l + r
                }
            case string:
                return l + toString(right)
            }
        }
    }

    return expr
}

func (i *Interpreter) parseList(expr string) []any {
    inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(expr, "["), "]"))
    if inner == "" {
        return []any{}
    }

    var items []any
    current := ""
    inQuotes := false
    quoteChar := byte(0)
    bracketDepth := 0

    for idx := 0; idx < len(inner); idx++ {
        ch := inner[idx]
        switch ch {
        case '\'', '"':
            if !inQuotes {
                inQuotes = true
                quoteChar = ch
                current += string(ch)
            } else if quoteChar == ch {
                inQuotes = false
                current += string(ch)
            } else {
                current += string(ch)
            }
        case '[':
            if !inQuotes {
                bracketDepth++
            }
            current += string(ch)
        case ']':
            if !inQuotes {
                if bracketDepth > 0 {
                    bracketDepth--
                }
            }
            current += string(ch)
        case ',':
            if inQuotes || bracketDepth > 0 {
                current += string(ch)
                continue
            }
            items = append(items, i.evaluateExpression(current))
            current = ""
        default:
            current += string(ch)
        }
    }
    if strings.TrimSpace(current) != "" {
        items = append(items, i.evaluateExpression(current))
    }
    return items
}

// Helpers
func splitArgsRespectingQuotes(input string) []string {
    var args []string
    current := strings.Builder{}
    inQuotes := false
    quoteChar := byte(0)
    bracketDepth := 0

    flush := func() {
        val := strings.TrimSpace(current.String())
        if val != "" {
            args = append(args, val)
        }
        current.Reset()
    }

    for idx := 0; idx < len(input); idx++ {
        ch := input[idx]
        switch ch {
        case '\'', '"':
            if !inQuotes {
                inQuotes = true
                quoteChar = ch
            } else if quoteChar == ch {
                inQuotes = false
            }
            current.WriteByte(ch)
        case '[':
            if !inQuotes {
                bracketDepth++
            }
            current.WriteByte(ch)
        case ']':
            if !inQuotes && bracketDepth > 0 {
                bracketDepth--
            }
            current.WriteByte(ch)
        case ',':
            if inQuotes || bracketDepth > 0 {
                current.WriteByte(ch)
                continue
            }
            flush()
        default:
            current.WriteByte(ch)
        }
    }
    flush()
    return args
}

func copyMap(in map[string]any) map[string]any {
    out := make(map[string]any, len(in))
    for k, v := range in {
        out[k] = v
    }
    return out
}

// RunFile loads and runs an Athera program from disk.
func RunFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return RunSource(string(data))
}

// RunSource runs Athera code from a string using a fresh interpreter.
func RunSource(src string) error {
    lexer := NewLexer(src)
    tokens := lexer.Tokenize()
    parser := NewParser(tokens)
    ast := parser.Parse()

    interpreter := NewInterpreter()
    interpreter.Execute(ast)
    return nil
}
