package lang

import (
    "encoding/json"
    "errors"
    "fmt"
    "math"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

// BuiltinFunc represents a standard library function.
type BuiltinFunc func(args []any) (any, error)

// StdlibModules holds the core modules bundled in the binary.
var StdlibModules = map[string]map[string]BuiltinFunc{
    "io":   ioModule(),
    "text": textModule(),
    "math": mathModule(),
    "list": listModule(),
    "dict": dictModule(),
    "time": timeModule(),
    "json": jsonModule(),
    "path": pathModule(),
}

func ioModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "read": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("io.read expects path")
            }
            path := toString(args[0])
            data, err := os.ReadFile(path)
            if err != nil {
                return "", err
            }
            return string(data), nil
        },
        "write": func(args []any) (any, error) {
            if len(args) < 2 {
                return nil, errors.New("io.write expects path and data")
            }
            path := toString(args[0])
            data := toString(args[1])
            return nil, os.WriteFile(path, []byte(data), 0o644)
        },
        "append": func(args []any) (any, error) {
            if len(args) < 2 {
                return nil, errors.New("io.append expects path and data")
            }
            path := toString(args[0])
            data := toString(args[1])
            f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
            if err != nil {
                return nil, err
            }
            defer f.Close()
            _, err = f.WriteString(data)
            return nil, err
        },
        "exists": func(args []any) (any, error) {
            if len(args) < 1 {
                return false, errors.New("io.exists expects path")
            }
            path := toString(args[0])
            _, err := os.Stat(path)
            return err == nil, nil
        },
        "read_lines": func(args []any) (any, error) {
            if len(args) < 1 {
                return []string{}, errors.New("io.read_lines expects path")
            }
            path := toString(args[0])
            data, err := os.ReadFile(path)
            if err != nil {
                return []string{}, err
            }
            lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
            if len(lines) > 0 && lines[len(lines)-1] == "" {
                lines = lines[:len(lines)-1]
            }
            return lines, nil
        },
        "size": func(args []any) (any, error) {
            if len(args) < 1 {
                return 0, errors.New("io.size expects path")
            }
            info, err := os.Stat(toString(args[0]))
            if err != nil {
                return 0, err
            }
            return info.Size(), nil
        },
        "dirname": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("io.dirname expects path")
            }
            return filepath.Dir(toString(args[0])), nil
        },
        "basename": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("io.basename expects path")
            }
            return filepath.Base(toString(args[0])), nil
        },
    }
}

func textModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "length": func(args []any) (any, error) {
            if len(args) < 1 {
                return 0, errors.New("text.length expects string")
            }
            return len([]rune(toString(args[0]))), nil
        },
        "upper": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("text.upper expects string")
            }
            return strings.ToUpper(toString(args[0])), nil
        },
        "lower": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("text.lower expects string")
            }
            return strings.ToLower(toString(args[0])), nil
        },
        "trim": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("text.trim expects string")
            }
            cutset := " "
            if len(args) > 1 {
                cutset = toString(args[1])
            }
            return strings.Trim(toString(args[0]), cutset), nil
        },
        "split": func(args []any) (any, error) {
            if len(args) < 2 {
                return []string{}, errors.New("text.split expects string and delimiter")
            }
            return strings.Split(toString(args[0]), toString(args[1])), nil
        },
        "contains": func(args []any) (any, error) {
            if len(args) < 2 {
                return false, errors.New("text.contains expects haystack and needle")
            }
            return strings.Contains(toString(args[0]), toString(args[1])), nil
        },
        "starts_with": func(args []any) (any, error) {
            if len(args) < 2 {
                return false, errors.New("text.starts_with expects haystack and prefix")
            }
            return strings.HasPrefix(toString(args[0]), toString(args[1])), nil
        },
        "ends_with": func(args []any) (any, error) {
            if len(args) < 2 {
                return false, errors.New("text.ends_with expects haystack and suffix")
            }
            return strings.HasSuffix(toString(args[0]), toString(args[1])), nil
        },
    }
}

func mathModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "add": func(args []any) (any, error) { return numericFold(args, 0, func(a, b float64) float64 { return a + b }) },
        "sub": func(args []any) (any, error) {
            if len(args) < 2 {
                return 0, errors.New("math.sub expects at least 2 numbers")
            }
            start := toFloat(args[0])
            for _, v := range args[1:] {
                start -= toFloat(v)
            }
            return start, nil
        },
        "mul": func(args []any) (any, error) { return numericFold(args, 1, func(a, b float64) float64 { return a * b }) },
        "div": func(args []any) (any, error) {
            if len(args) < 2 {
                return 0, errors.New("math.div expects at least 2 numbers")
            }
            start := toFloat(args[0])
            for _, v := range args[1:] {
                denom := toFloat(v)
                if denom == 0 {
                    return 0, errors.New("division by zero")
                }
                start /= denom
            }
            return start, nil
        },
        "sqrt": func(args []any) (any, error) {
            if len(args) < 1 {
                return 0, errors.New("math.sqrt expects number")
            }
            return math.Sqrt(toFloat(args[0])), nil
        },
        "abs": func(args []any) (any, error) {
            if len(args) < 1 {
                return 0, errors.New("math.abs expects number")
            }
            return math.Abs(toFloat(args[0])), nil
        },
    }
}

func numericFold(args []any, seed float64, op func(a, b float64) float64) (any, error) {
    if len(args) == 0 {
        return seed, nil
    }
    total := seed
    for _, v := range args {
        total = op(total, toFloat(v))
    }
    return total, nil
}

// LIST MODULE
func listModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "length": func(args []any) (any, error) {
            if len(args) < 1 {
                return 0, errors.New("list.length expects list")
            }
            switch v := args[0].(type) {
            case []any:
                return len(v), nil
            case []string:
                return len(v), nil
            default:
                return 0, fmt.Errorf("list.length: got %T, expected list", args[0])
            }
        },
        "append": func(args []any) (any, error) {
            if len(args) < 2 {
                return nil, errors.New("list.append expects list and item")
            }
            lst, ok := args[0].([]any)
            if !ok {
                return nil, errors.New("list.append: first arg must be list")
            }
            return append(lst, args[1]), nil
        },
        "at": func(args []any) (any, error) {
            if len(args) < 2 {
                return nil, errors.New("list.at expects list and index")
            }
            idx := int(toFloat(args[1]))
            switch v := args[0].(type) {
            case []any:
                if idx < 0 || idx >= len(v) {
                    return nil, errors.New("list.at: index out of bounds")
                }
                return v[idx], nil
            case []string:
                if idx < 0 || idx >= len(v) {
                    return nil, errors.New("list.at: index out of bounds")
                }
                return v[idx], nil
            default:
                return nil, fmt.Errorf("list.at: got %T, expected list", args[0])
            }
        },
        "contains": func(args []any) (any, error) {
            if len(args) < 2 {
                return false, errors.New("list.contains expects list and item")
            }
            item := args[1]
            switch v := args[0].(type) {
            case []any:
                for _, x := range v {
                    if fmt.Sprint(x) == fmt.Sprint(item) {
                        return true, nil
                    }
                }
            case []string:
                itemStr := toString(item)
                for _, x := range v {
                    if x == itemStr {
                        return true, nil
                    }
                }
            }
            return false, nil
        },
    }
}

// DICT MODULE
func dictModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "get": func(args []any) (any, error) {
            if len(args) < 2 {
                return nil, errors.New("dict.get expects dict and key")
            }
            d, ok := args[0].(map[string]any)
            if !ok {
                return nil, errors.New("dict.get: first arg must be dict")
            }
            key := toString(args[1])
            if val, exists := d[key]; exists {
                return val, nil
            }
            return nil, nil
        },
        "set": func(args []any) (any, error) {
            if len(args) < 3 {
                return nil, errors.New("dict.set expects dict, key, and value")
            }
            d, ok := args[0].(map[string]any)
            if !ok {
                d = make(map[string]any)
            }
            key := toString(args[1])
            d[key] = args[2]
            return d, nil
        },
        "keys": func(args []any) (any, error) {
            if len(args) < 1 {
                return []string{}, errors.New("dict.keys expects dict")
            }
            d, ok := args[0].(map[string]any)
            if !ok {
                return []string{}, errors.New("dict.keys: arg must be dict")
            }
            var keys []string
            for k := range d {
                keys = append(keys, k)
            }
            return keys, nil
        },
        "values": func(args []any) (any, error) {
            if len(args) < 1 {
                return []any{}, errors.New("dict.values expects dict")
            }
            d, ok := args[0].(map[string]any)
            if !ok {
                return []any{}, errors.New("dict.values: arg must be dict")
            }
            var vals []any
            for _, v := range d {
                vals = append(vals, v)
            }
            return vals, nil
        },
    }
}

// TIME MODULE
func timeModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "now": func(args []any) (any, error) {
            return time.Now().Format(time.RFC3339), nil
        },
        "timestamp": func(args []any) (any, error) {
            return time.Now().Unix(), nil
        },
        "sleep": func(args []any) (any, error) {
            if len(args) < 1 {
                return nil, errors.New("time.sleep expects milliseconds")
            }
            ms := int(toFloat(args[0]))
            time.Sleep(time.Duration(ms) * time.Millisecond)
            return nil, nil
        },
        "format": func(args []any) (any, error) {
            if len(args) < 2 {
                return "", errors.New("time.format expects timestamp and layout")
            }
            ts := int64(toFloat(args[0]))
            layout := toString(args[1])
            t := time.Unix(ts, 0)
            return t.Format(layout), nil
        },
    }
}

// JSON MODULE
func jsonModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "parse": func(args []any) (any, error) {
            if len(args) < 1 {
                return nil, errors.New("json.parse expects string")
            }
            jsonStr := toString(args[0])
            var result any
            if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
                return nil, err
            }
            return result, nil
        },
        "stringify": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("json.stringify expects value")
            }
            data, err := json.Marshal(args[0])
            if err != nil {
                return "", err
            }
            return string(data), nil
        },
    }
}

// PATH MODULE
func pathModule() map[string]BuiltinFunc {
    return map[string]BuiltinFunc{
        "join": func(args []any) (any, error) {
            if len(args) == 0 {
                return "", errors.New("path.join expects at least one path")
            }
            parts := make([]string, 0, len(args))
            for _, arg := range args {
                parts = append(parts, toString(arg))
            }
            return filepath.Join(parts...), nil
        },
        "dir": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("path.dir expects path")
            }
            return filepath.Dir(toString(args[0])), nil
        },
        "base": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("path.base expects path")
            }
            return filepath.Base(toString(args[0])), nil
        },
        "ext": func(args []any) (any, error) {
            if len(args) < 1 {
                return "", errors.New("path.ext expects path")
            }
            return filepath.Ext(toString(args[0])), nil
        },
        "exists": func(args []any) (any, error) {
            if len(args) < 1 {
                return false, errors.New("path.exists expects path")
            }
            _, err := os.Stat(toString(args[0]))
            return err == nil, nil
        },
    }
}

func toString(v any) string {
    switch val := v.(type) {
    case string:
        return val
    case []byte:
        return string(val)
    default:
        return strings.TrimSpace(strings.Trim(fmt.Sprint(val), "\n"))
    }
}

func toFloat(v any) float64 {
    switch val := v.(type) {
    case int:
        return float64(val)
    case int64:
        return float64(val)
    case float64:
        return val
    case float32:
        return float64(val)
    case string:
        f, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
        return f
    default:
        return 0
    }
}
