package lang

import (
    "strconv"
    "strings"
)

// Parser builds an AST from tokens.
type Parser struct {
    tokens []Token
    pos    int
}

// NewParser constructs a parser for tokens.
func NewParser(tokens []Token) *Parser {
    return &Parser{tokens: tokens}
}

// Parse converts tokens into a slice of AST nodes.
func (p *Parser) Parse() []Node {
    var nodes []Node
    for !p.isAtEnd() {
        if node := p.parseStatement(0); node != nil {
            nodes = append(nodes, node)
        }
    }
    return nodes
}

func (p *Parser) parseStatement(parentIndent int) Node {
    tok := p.peek()

    if tok.Type == "NEWLINE" {
        p.advance()
        return nil
    }
    if tok.Type == "EOF" {
        return nil
    }

    if tok.Type == "INDENT" {
        indent, _ := strconv.Atoi(tok.Value)
        if indent <= parentIndent {
            return nil
        }
        p.advance()
        tok = p.peek()
    }

    switch tok.Type {
    case "TASK":
        return p.parseTask()
    case "GREET":
        return p.parseGreet()
    case "BACKUP":
        return p.parseBackup()
    case "CHECK":
        return p.parseCheck()
    case "REPEAT_N":
        return p.parseRepeatN()
    case "REPEAT_EACH":
        return p.parseRepeatEach()
    case "SET":
        return p.parseSet()
    case "RUN":
        return p.parseRun()
    case "USE":
        return p.parseUse()
    case "PROTECT":
        return p.parseProtect()
    case "HANDLE_INLINE":
        return p.parseHandleInline()
    case "RUN_PARALLEL":
        return p.parseRunParallel()
    case "RETURN":
        return p.parseReturn()
    default:
        p.advance()
        return nil
    }
}

func (p *Parser) parseTask() Node {
    tok := p.advance()
    val := tok.Value
    name := val
    var params []string

    if strings.Contains(val, "|") {
        parts := strings.SplitN(val, "|", 2)
        name = parts[0]
        if len(parts) > 1 && parts[1] != "" {
            params = splitCSV(parts[1])
        }
    }

    p.consume("NEWLINE")
    body := p.parseBlock()
    return &TaskNode{Name: name, Params: params, Body: body}
}

func (p *Parser) parseBlock() []Node {
    var body []Node
    var baseIndent *int

    for !p.isAtEnd() {
        tok := p.peek()

        if tok.Type == "INDENT" {
            indent, _ := strconv.Atoi(tok.Value)
            if baseIndent == nil {
                baseIndent = &indent
            } else if indent < *baseIndent {
                break
            }
            p.advance()
            tok = p.peek()
        } else if tok.Type == "NEWLINE" {
            p.advance()
            continue
        } else {
            if baseIndent != nil {
                break
            }
        }

        node := p.parseBlockStatement()
        if node != nil {
            body = append(body, node)
        }
    }

    return body
}

func (p *Parser) parseBlockStatement() Node {
    tok := p.peek()
    var node Node

    switch tok.Type {
    case "GREET":
        node = p.parseGreet()
    case "BACKUP":
        node = p.parseBackup()
    case "CHECK":
        node = p.parseCheck()
    case "REPEAT_N":
        node = p.parseRepeatN()
    case "REPEAT_EACH":
        node = p.parseRepeatEach()
    case "SET":
        node = p.parseSet()
    case "RUN":
        node = p.parseRun()
    case "RUN_PARALLEL":
        node = p.parseRunParallel()
    case "USE":
        node = p.parseUse()
    case "PROTECT":
        node = p.parseProtect()
    case "HANDLE_INLINE":
        node = p.parseHandleInline()
    case "RETURN":
        node = p.parseReturn()
    default:
        p.advance()
        return nil
    }

    p.consume("NEWLINE")
    return node
}

func (p *Parser) parseGreet() Node {
    tok := p.advance()
    return &GreetNode{Message: tok.Value}
}

func (p *Parser) parseBackup() Node {
    tok := p.advance()
    parts := strings.SplitN(tok.Value, "|", 2)
    src, dest := parts[0], ""
    if len(parts) > 1 {
        dest = parts[1]
    }
    return &BackupNode{Source: src, Dest: dest}
}

func (p *Parser) parseCheck() Node {
    tok := p.advance()
    parts := strings.SplitN(tok.Value, "|", 2)
    cond, action := parts[0], ""
    if len(parts) > 1 {
        action = parts[1]
    }
    return &CheckNode{Condition: cond, Action: action}
}

func (p *Parser) parseRepeatN() Node {
    tok := p.advance()
    count, _ := strconv.Atoi(tok.Value)
    p.consume("NEWLINE")
    body := p.parseBlock()
    return &RepeatNNode{Count: count, Body: body}
}

func (p *Parser) parseRepeatEach() Node {
    tok := p.advance()
    parts := strings.SplitN(tok.Value, "|", 2)
    varName, listExpr := parts[0], ""
    if len(parts) > 1 {
        listExpr = parts[1]
    }
    p.consume("NEWLINE")
    body := p.parseBlock()
    return &RepeatEachNode{Var: varName, ListExpr: listExpr, Body: body}
}

func (p *Parser) parseSet() Node {
    tok := p.advance()
    parts := strings.SplitN(tok.Value, "|", 2)
    name, value := parts[0], ""
    if len(parts) > 1 {
        value = parts[1]
    }
    return &SetNode{Var: name, Value: value}
}

func (p *Parser) parseRun() Node {
    tok := p.advance()
    return &RunNode{Target: tok.Value}
}

func (p *Parser) parseUse() Node {
    tok := p.advance()
    return &UseNode{Module: tok.Value}
}

func (p *Parser) parseProtect() Node {
    p.advance() // consume protect
    p.consume("NEWLINE")

    var protectBody []Node
    for !p.isAtEnd() {
        tok := p.peek()
        if tok.Type == "HANDLE" || (tok.Type == "INDENT" && p.peekAhead(1).Type == "HANDLE") {
            break
        }
        if tok.Type == "INDENT" {
            p.advance()
        }
        if node := p.parseBlockStatement(); node != nil {
            protectBody = append(protectBody, node)
        }
    }

    var handleBody []Node
    if p.peek().Type == "HANDLE" || (p.peek().Type == "INDENT" && p.peekAhead(1).Type == "HANDLE") {
        if p.peek().Type == "INDENT" {
            p.advance()
        }
        p.consume("HANDLE")
        p.consume("NEWLINE")
        handleBody = p.parseBlock()
    }

    return &ProtectNode{Protect: protectBody, Handle: handleBody}
}

func (p *Parser) parseHandleInline() Node {
    tok := p.advance()
    parts := strings.SplitN(tok.Value, "|", 2)
    errType, action := parts[0], ""
    if len(parts) > 1 {
        action = parts[1]
    }
    return &HandleInlineNode{ErrorType: errType, Action: action}
}

func (p *Parser) parseRunParallel() Node {
    tok := p.advance()
    tasks := splitCSV(tok.Value)
    return &RunParallelNode{Tasks: tasks}
}

func (p *Parser) parseReturn() Node {
    tok := p.advance()
    return &ReturnNode{Expr: tok.Value}
}

func (p *Parser) peekAhead(offset int) Token {
    pos := p.pos + offset
    if pos < len(p.tokens) {
        return p.tokens[pos]
    }
    return Token{Type: "EOF"}
}

func (p *Parser) peek() Token {
    if p.pos < len(p.tokens) {
        return p.tokens[p.pos]
    }
    return Token{Type: "EOF"}
}

func (p *Parser) advance() Token {
    tok := p.peek()
    p.pos++
    return tok
}

func (p *Parser) consume(expected string) {
    if p.peek().Type == expected {
        p.advance()
    }
}

func (p *Parser) isAtEnd() bool {
    return p.peek().Type == "EOF"
}
