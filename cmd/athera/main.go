package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"

    "athera/internal/lang"
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Athera (Go) - Phase 1 minimal runtime\n")
        fmt.Fprintf(os.Stderr, "Usage:\n")
        fmt.Fprintf(os.Stderr, "  athera run <file.ath>\n")
        fmt.Fprintf(os.Stderr, "  athera repl\n")
    }

    flag.Parse()
    args := flag.Args()

    if len(args) < 1 {
        flag.Usage()
        os.Exit(1)
    }

    switch args[0] {
    case "run":
        if len(args) < 2 {
            fmt.Println("Error: athera run requires a file path")
            os.Exit(1)
        }
        if err := lang.RunFile(args[1]); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }
    case "repl":
        runRepl()
    default:
        fmt.Printf("Unknown command: %s\n", args[0])
        flag.Usage()
        os.Exit(1)
    }
}

func runRepl() {
    fmt.Println("Athera REPL (Go runtime)")
    fmt.Println("Type 'exit' or 'quit' to leave. Enter blank line to execute a multi-line block.")

    interp := lang.NewInterpreter()
    scanner := bufio.NewScanner(os.Stdin)
    var buffer []string

    prompt := func(cont bool) {
        if cont {
            fmt.Print("... ")
        } else {
            fmt.Print("athera> ")
        }
    }

    for {
        prompt(len(buffer) > 0)
        if !scanner.Scan() {
            break
        }
        line := scanner.Text()
        trimmed := strings.TrimSpace(line)

        if trimmed == "exit" || trimmed == "quit" {
            break
        }

        // Blank line triggers execution of accumulated buffer
        if trimmed == "" {
            if len(buffer) == 0 {
                continue
            }
            src := strings.Join(buffer, "\n")
            lexer := lang.NewLexer(src)
            tokens := lexer.Tokenize()
            parser := lang.NewParser(tokens)
            nodes := parser.Parse()
            interp.Execute(nodes)
            buffer = buffer[:0]
            continue
        }

        buffer = append(buffer, line)

        // For single-line commands without trailing ':' execute immediately
        if len(buffer) == 1 && !strings.HasSuffix(trimmed, ":") {
            src := buffer[0]
            buffer = buffer[:0]
            lexer := lang.NewLexer(src)
            tokens := lexer.Tokenize()
            parser := lang.NewParser(tokens)
            nodes := parser.Parse()
            interp.Execute(nodes)
        }
    }

    fmt.Println("Goodbye.")
}
