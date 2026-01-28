package lang

import (
    "regexp"
    "strconv"
    "strings"
)

// Token represents a lexical token.
type Token struct {
    Type  string
    Value string
    Line  int
}

// Lexer converts source code into tokens.
type Lexer struct {
    source string
    lines  []string
    tokens []Token
}

// NewLexer creates a lexer for the provided source.
func NewLexer(src string) *Lexer {
    return &Lexer{source: src, lines: strings.Split(src, "\n"), tokens: []Token{}}
}

// Tokenize lexes the source into a slice of tokens.
func (l *Lexer) Tokenize() []Token {
    for idx, raw := range l.lines {
        lineNum := idx + 1
        trimmed := strings.TrimSpace(raw)

        if trimmed == "" || strings.HasPrefix(trimmed, "#") {
            continue
        }

        indent := len(raw) - len(strings.TrimLeft(raw, " \t"))
        if indent > 0 {
            l.tokens = append(l.tokens, Token{Type: "INDENT", Value: strconv.Itoa(indent), Line: lineNum})
        }

        l.tokenizeLine(trimmed, lineNum)
        l.tokens = append(l.tokens, Token{Type: "NEWLINE", Line: lineNum})
    }

    l.tokens = append(l.tokens, Token{Type: "EOF", Line: len(l.lines)})
    return l.tokens
}

func (l *Lexer) tokenizeLine(line string, lineNum int) {
    switch {
    case strings.HasPrefix(line, "task ") && strings.HasSuffix(line, ":"):
        taskDef := strings.TrimSuffix(strings.TrimSpace(line[5:]), ":")
        if strings.Contains(taskDef, " with ") {
            parts := strings.SplitN(taskDef, " with ", 2)
            name := strings.TrimSpace(parts[0])
            params := splitCSV(parts[1])
            l.tokens = append(l.tokens, Token{Type: "TASK", Value: name + "|" + strings.Join(params, ","), Line: lineNum})
        } else {
            l.tokens = append(l.tokens, Token{Type: "TASK", Value: taskDef, Line: lineNum})
        }
        return
    case strings.HasPrefix(line, "greet "):
        msg := strings.TrimSpace(line[len("greet "):])
        l.tokens = append(l.tokens, Token{Type: "GREET", Value: msg, Line: lineNum})
        return
    case strings.HasPrefix(line, "backup ") && strings.Contains(line, " to "):
        parts := strings.SplitN(strings.TrimSpace(line[len("backup "):]), " to ", 2)
        l.tokens = append(l.tokens, Token{Type: "BACKUP", Value: parts[0]+"|"+parts[1], Line: lineNum})
        return
    case strings.HasPrefix(line, "check ") && strings.Contains(line, " -> "):
        parts := strings.SplitN(strings.TrimSpace(line[len("check "):]), " -> ", 2)
        l.tokens = append(l.tokens, Token{Type: "CHECK", Value: parts[0]+"|"+parts[1], Line: lineNum})
        return
    case strings.HasPrefix(line, "repeat ") && strings.HasSuffix(line, ":"):
        if strings.Contains(line, " each ") && strings.Contains(line, " in ") {
            re := regexp.MustCompile(`repeat each (\w+) in (.+):`)
            if m := re.FindStringSubmatch(line); len(m) == 3 {
                l.tokens = append(l.tokens, Token{Type: "REPEAT_EACH", Value: m[1]+"|"+strings.TrimSpace(m[2]), Line: lineNum})
                return
            }
        }
        re := regexp.MustCompile(`repeat (\d+) times:`)
        if m := re.FindStringSubmatch(line); len(m) == 2 {
            l.tokens = append(l.tokens, Token{Type: "REPEAT_N", Value: m[1], Line: lineNum})
            return
        }
    case strings.HasPrefix(line, "set ") && strings.Contains(line, " = "):
        parts := strings.SplitN(strings.TrimSpace(line[len("set "):]), " = ", 2)
        l.tokens = append(l.tokens, Token{Type: "SET", Value: parts[0]+"|"+parts[1], Line: lineNum})
        return
    case strings.HasPrefix(line, "run "):
        rest := strings.TrimSpace(line[len("run "):])
        if strings.HasPrefix(rest, "parallel ") {
            tasks := strings.Split(strings.TrimSpace(rest[len("parallel "):]), ",")
            l.tokens = append(l.tokens, Token{Type: "RUN_PARALLEL", Value: strings.Join(trimAll(tasks), ","), Line: lineNum})
        } else {
            l.tokens = append(l.tokens, Token{Type: "RUN", Value: rest, Line: lineNum})
        }
        return
    case strings.HasPrefix(line, "use "):
        l.tokens = append(l.tokens, Token{Type: "USE", Value: strings.TrimSpace(line[len("use "):]), Line: lineNum})
        return
    case line == "protect:":
        l.tokens = append(l.tokens, Token{Type: "PROTECT", Line: lineNum})
        return
    case line == "handle:":
        l.tokens = append(l.tokens, Token{Type: "HANDLE", Line: lineNum})
        return
    case strings.HasPrefix(line, "handle ") && strings.Contains(line, " -> "):
        parts := strings.SplitN(strings.TrimSpace(line[len("handle "):]), " -> ", 2)
        l.tokens = append(l.tokens, Token{Type: "HANDLE_INLINE", Value: parts[0]+"|"+parts[1], Line: lineNum})
        return
    case strings.HasPrefix(line, "return "):
        l.tokens = append(l.tokens, Token{Type: "RETURN", Value: strings.TrimSpace(line[len("return "):]), Line: lineNum})
        return
    default:
        l.tokens = append(l.tokens, Token{Type: "UNKNOWN", Value: line, Line: lineNum})
    }
}

// splitCSV splits a comma-separated list respecting simple spacing.
func splitCSV(raw string) []string {
    parts := strings.Split(raw, ",")
    out := make([]string, 0, len(parts))
    for _, p := range parts {
        p = strings.TrimSpace(p)
        if p != "" {
            out = append(out, p)
        }
    }
    return out
}

// trimAll trims whitespace from all items.
func trimAll(in []string) []string {
    out := make([]string, 0, len(in))
    for _, v := range in {
        if strings.TrimSpace(v) != "" {
            out = append(out, strings.TrimSpace(v))
        }
    }
    return out
}
