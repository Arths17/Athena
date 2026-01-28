package lang

// Node is the base interface for all AST nodes.
type Node interface{}

// TaskNode represents a task definition.
type TaskNode struct {
    Name   string
    Params []string
    Body   []Node
}

// GreetNode prints a message.
type GreetNode struct {
    Message string
}

// BackupNode copies a file or folder to a destination.
type BackupNode struct {
    Source string
    Dest   string
}

// CheckNode evaluates a condition then runs an inline action.
type CheckNode struct {
    Condition string
    Action    string
}

// RepeatNNode repeats a block a fixed number of times.
type RepeatNNode struct {
    Count int
    Body  []Node
}

// RepeatEachNode iterates over items in a list.
type RepeatEachNode struct {
    Var      string
    ListExpr string
    Body     []Node
}

// SetNode assigns the result of an expression to a variable.
type SetNode struct {
    Var   string
    Value string
}

// RunNode calls a named task with optional arguments.
type RunNode struct {
    Target string
}

// UseNode imports a module.
type UseNode struct {
    Module string
}

// ProtectNode represents a protect/handle block.
type ProtectNode struct {
    Protect []Node
    Handle  []Node
}

// HandleInlineNode runs when the preceding protect block fails.
type HandleInlineNode struct {
    ErrorType string
    Action    string
}

// RunParallelNode executes multiple tasks concurrently.
type RunParallelNode struct {
    Tasks []string
}

// ReturnNode exits a task with a value.
type ReturnNode struct {
    Expr string
}
