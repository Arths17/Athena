#!/usr/bin/env python3
"""
Athera Programming Language Interpreter

A unique programming language with natural language-like syntax.
Features: tasks, greet, backup, check conditions, loops, variables
"""

import re
import os
import shutil
from pathlib import Path
from typing import Any, Dict, List, Optional


# ============================================================================
# LEXER - Tokenizes Athera source code
# ============================================================================

class Token:
    """Represents a single token in the source code"""
    def __init__(self, type_: str, value: Any, line: int = 0):
        self.type = type_
        self.value = value
        self.line = line
    
    def __repr__(self):
        return f"Token({self.type}, {self.value})"


class Lexer:
    """Tokenizes Athera source code into tokens"""
    
    def __init__(self, source: str):
        self.source = source
        self.lines = source.split('\n')
        self.tokens = []
        self.current_line = 0
    
    def tokenize(self) -> List[Token]:
        """Convert source code into list of tokens"""
        for line_num, line in enumerate(self.lines):
            self.current_line = line_num + 1
            stripped = line.strip()
            
            # Skip empty lines and comments
            if not stripped or stripped.startswith('#'):
                continue
            
            # Calculate indentation level
            indent_level = len(line) - len(line.lstrip())
            if indent_level > 0:
                self.tokens.append(Token('INDENT', indent_level, self.current_line))
            
            # Tokenize the line
            self._tokenize_line(stripped)
            self.tokens.append(Token('NEWLINE', None, self.current_line))
        
        self.tokens.append(Token('EOF', None, self.current_line))
        return self.tokens
    
    def _tokenize_line(self, line: str):
        """Tokenize a single line"""
        # Task declaration with parameters: task <name> with <params>:
        if line.startswith('task ') and line.endswith(':'):
            task_def = line[5:-1].strip()
            if ' with ' in task_def:
                parts = task_def.split(' with ', 1)
                task_name = parts[0].strip()
                params = [p.strip() for p in parts[1].split(',')]
                self.tokens.append(Token('TASK', {'name': task_name, 'params': params}, self.current_line))
            else:
                self.tokens.append(Token('TASK', {'name': task_def, 'params': []}, self.current_line))
            return
        
        # Greet command: greet "message"
        if line.startswith('greet '):
            message = self._extract_string(line[6:])
            self.tokens.append(Token('GREET', message, self.current_line))
            return
        
        # Backup command: backup <source> to <dest>
        if line.startswith('backup ') and ' to ' in line:
            parts = line[7:].split(' to ', 1)
            source = parts[0].strip()
            dest = parts[1].strip()
            self.tokens.append(Token('BACKUP', {'source': source, 'dest': dest}, self.current_line))
            return
        
        # Check condition: check <condition> -> <action>
        if line.startswith('check ') and ' -> ' in line:
            parts = line[6:].split(' -> ', 1)
            condition = parts[0].strip()
            action = parts[1].strip()
            self.tokens.append(Token('CHECK', {'condition': condition, 'action': action}, self.current_line))
            return
        
        # Repeat loop: repeat <N> times: or repeat each <item> in <list>:
        if line.startswith('repeat ') and line.endswith(':'):
            if ' each ' in line and ' in ' in line:
                # repeat each <item> in <list>:
                match = re.match(r'repeat each (\w+) in (.+):', line)
                if match:
                    item_var = match.group(1)
                    list_expr = match.group(2).strip()
                    self.tokens.append(Token('REPEAT_EACH', {'var': item_var, 'list': list_expr}, self.current_line))
                    return
            else:
                # repeat <N> times:
                match = re.match(r'repeat (\d+) times:', line)
                if match:
                    count = int(match.group(1))
                    self.tokens.append(Token('REPEAT_N', count, self.current_line))
                    return
        
        # Variable assignment: set <var> = <value>
        if line.startswith('set ') and ' = ' in line:
            parts = line[4:].split(' = ', 1)
            var_name = parts[0].strip()
            value_expr = parts[1].strip()
            self.tokens.append(Token('SET', {'var': var_name, 'value': value_expr}, self.current_line))
            return
        
        # Run task: run <task_name> or run parallel <tasks>
        if line.startswith('run '):
            rest = line[4:].strip()
            if rest.startswith('parallel '):
                # Parallel execution
                tasks = [t.strip() for t in rest[9:].split(',')]
                self.tokens.append(Token('RUN_PARALLEL', tasks, self.current_line))
            else:
                # Regular task call
                self.tokens.append(Token('RUN', rest, self.current_line))
            return
        
        # Use module: use <module>
        if line.startswith('use '):
            module_name = line[4:].strip()
            self.tokens.append(Token('USE', module_name, self.current_line))
            return
        
        # Protect block: protect:
        if line == 'protect:':
            self.tokens.append(Token('PROTECT', None, self.current_line))
            return
        
        # Handle error: handle <error> -> <action> or handle:
        if line.startswith('handle'):
            if line == 'handle:':
                self.tokens.append(Token('HANDLE', None, self.current_line))
            elif ' -> ' in line:
                parts = line[7:].split(' -> ', 1)
                error_type = parts[0].strip()
                action = parts[1].strip()
                self.tokens.append(Token('HANDLE_INLINE', {'error': error_type, 'action': action}, self.current_line))
            return
        
        # Return statement: return <value>
        if line.startswith('return '):
            value = line[7:].strip()
            self.tokens.append(Token('RETURN', value, self.current_line))
            return
        
        # If nothing matched, it's an unknown command
        self.tokens.append(Token('UNKNOWN', line, self.current_line))
    
    def _extract_string(self, s: str) -> str:
        """Extract string from quotes"""
        s = s.strip()
        if (s.startswith('"') and s.endswith('"')) or (s.startswith("'") and s.endswith("'")):
            return s[1:-1]
        return s


# ============================================================================
# PARSER - Builds Abstract Syntax Tree (AST)
# ============================================================================

class ASTNode:
    """Base class for AST nodes"""
    pass


class TaskNode(ASTNode):
    """Represents a task definition"""
    def __init__(self, name: str, body: List[ASTNode]):
        self.name = name
        self.body = body


class GreetNode(ASTNode):
    """Represents a greet command"""
    def __init__(self, message: str):
        self.message = message


class BackupNode(ASTNode):
    """Represents a backup command"""
    def __init__(self, source: str, dest: str):
        self.source = source
        self.dest = dest


class CheckNode(ASTNode):
    """Represents a check condition"""
    def __init__(self, condition: str, action: str):
        self.condition = condition
        self.action = action


class RepeatNNode(ASTNode):
    """Represents a repeat N times loop"""
    def __init__(self, count: int, body: List[ASTNode]):
        self.count = count
        self.body = body


class RepeatEachNode(ASTNode):
    """Represents a repeat each item in list loop"""
    def __init__(self, var: str, list_expr: str, body: List[ASTNode]):
        self.var = var
        self.list_expr = list_expr
        self.body = body


class SetNode(ASTNode):
    """Represents a variable assignment"""
    def __init__(self, var: str, value_expr: str):
        self.var = var
        self.value_expr = value_expr


class RunNode(ASTNode):
    """Represents a task call"""
    def __init__(self, task_name: str):
        self.task_name = task_name


class UseNode(ASTNode):
    """Represents a module import"""
    def __init__(self, module_name: str):
        self.module_name = module_name


class ProtectNode(ASTNode):
    """Represents a protect-handle error handling block"""
    def __init__(self, protect_body: List[ASTNode], handle_body: List[ASTNode]):
        self.protect_body = protect_body
        self.handle_body = handle_body


class HandleInlineNode(ASTNode):
    """Represents inline error handling"""
    def __init__(self, error_type: str, action: str):
        self.error_type = error_type
        self.action = action


class RunParallelNode(ASTNode):
    """Represents parallel task execution"""
    def __init__(self, task_names: List[str]):
        self.task_names = task_names


class ReturnNode(ASTNode):
    """Represents a return statement"""
    def __init__(self, value_expr: str):
        self.value_expr = value_expr


class Parser:
    """Parses tokens into an Abstract Syntax Tree"""
    
    def __init__(self, tokens: List[Token]):
        self.tokens = tokens
        self.pos = 0
        self.current_indent = 0
    
    def parse(self) -> List[ASTNode]:
        """Parse tokens into AST"""
        nodes = []
        while not self._is_at_end():
            node = self._parse_statement()
            if node:
                nodes.append(node)
        return nodes
    
    def _parse_statement(self, parent_indent: int = 0) -> Optional[ASTNode]:
        """Parse a single statement"""
        token = self._peek()
        
        if token.type == 'NEWLINE':
            self._advance()
            return None
        
        if token.type == 'EOF':
            return None
        
        # Check indent for block structure
        if token.type == 'INDENT':
            indent = token.value
            if indent <= parent_indent:
                return None
            self._advance()
            token = self._peek()
        
        if token.type == 'TASK':
            return self._parse_task()
        elif token.type == 'GREET':
            return self._parse_greet()
        elif token.type == 'BACKUP':
            return self._parse_backup()
        elif token.type == 'CHECK':
            return self._parse_check()
        elif token.type == 'REPEAT_N':
            return self._parse_repeat_n()
        elif token.type == 'REPEAT_EACH':
            return self._parse_repeat_each()
        elif token.type == 'SET':
            return self._parse_set()
        elif token.type == 'RUN':
            return self._parse_run()
        elif token.type == 'USE':
            return self._parse_use()
        elif token.type == 'PROTECT':
            return self._parse_protect()
        elif token.type == 'HANDLE_INLINE':
            return self._parse_handle_inline()
        elif token.type == 'RUN_PARALLEL':
            return self._parse_run_parallel()
        elif token.type == 'RETURN':
            return self._parse_return()
        else:
            self._advance()
            return None
    
    def _parse_task(self) -> TaskNode:
        """Parse task definition"""
        token = self._advance()
        if isinstance(token.value, dict):
            task_name = token.value['name']
            params = token.value['params']
        else:
            task_name = token.value
            params = []
        self._consume('NEWLINE')
        
        # Parse task body
        body = self._parse_block()
        node = TaskNode(task_name, body)
        node.params = params
        return node
    
    def _parse_block(self) -> List[ASTNode]:
        """Parse a block of statements (indented body)"""
        body = []
        base_indent = None
        
        while not self._is_at_end():
            token = self._peek()
            
            if token.type == 'INDENT':
                indent = token.value
                if base_indent is None:
                    base_indent = indent
                elif indent < base_indent:
                    break
                
                self._advance()
                token = self._peek()
            elif token.type == 'NEWLINE':
                self._advance()
                continue
            else:
                # No indent means end of block
                if base_indent is not None:
                    break
            
            # Parse statement in block
            node = self._parse_block_statement()
            if node:
                body.append(node)
        
        return body
    
    def _parse_block_statement(self) -> Optional[ASTNode]:
        """Parse a statement within a block"""
        token = self._peek()
        
        if token.type == 'GREET':
            node = self._parse_greet()
        elif token.type == 'BACKUP':
            node = self._parse_backup()
        elif token.type == 'CHECK':
            node = self._parse_check()
        elif token.type == 'REPEAT_N':
            node = self._parse_repeat_n()
        elif token.type == 'REPEAT_EACH':
            node = self._parse_repeat_each()
        elif token.type == 'SET':
            node = self._parse_set()
        elif token.type == 'RUN':
            node = self._parse_run()
        elif token.type == 'RUN_PARALLEL':
            node = self._parse_run_parallel()
        elif token.type == 'USE':
            node = self._parse_use()
        elif token.type == 'PROTECT':
            node = self._parse_protect()
        elif token.type == 'HANDLE_INLINE':
            node = self._parse_handle_inline()
        elif token.type == 'RETURN':
            node = self._parse_return()
        else:
            self._advance()
            return None
        
        self._consume('NEWLINE')
        return node
    
    def _parse_greet(self) -> GreetNode:
        """Parse greet command"""
        token = self._advance()
        return GreetNode(token.value)
    
    def _parse_backup(self) -> BackupNode:
        """Parse backup command"""
        token = self._advance()
        return BackupNode(token.value['source'], token.value['dest'])
    
    def _parse_check(self) -> CheckNode:
        """Parse check condition"""
        token = self._advance()
        return CheckNode(token.value['condition'], token.value['action'])
    
    def _parse_repeat_n(self) -> RepeatNNode:
        """Parse repeat N times loop"""
        token = self._advance()
        count = token.value
        self._consume('NEWLINE')
        body = self._parse_block()
        return RepeatNNode(count, body)
    
    def _parse_repeat_each(self) -> RepeatEachNode:
        """Parse repeat each loop"""
        token = self._advance()
        var = token.value['var']
        list_expr = token.value['list']
        self._consume('NEWLINE')
        body = self._parse_block()
        return RepeatEachNode(var, list_expr, body)
    
    def _parse_set(self) -> SetNode:
        """Parse variable assignment"""
        token = self._advance()
        return SetNode(token.value['var'], token.value['value'])
    
    def _parse_run(self) -> RunNode:
        """Parse task call"""
        token = self._advance()
        return RunNode(token.value)
    
    def _parse_use(self) -> UseNode:
        """Parse module import"""
        token = self._advance()
        return UseNode(token.value)
    
    def _parse_protect(self) -> ProtectNode:
        """Parse protect-handle block"""
        self._advance()  # consume 'protect'
        self._consume('NEWLINE')
        
        # Parse protect body
        protect_body = []
        while not self._is_at_end():
            token = self._peek()
            if token.type == 'HANDLE' or (token.type == 'INDENT' and self._peek_ahead(1).type == 'HANDLE'):
                break
            if token.type == 'INDENT':
                self._advance()
            node = self._parse_block_statement()
            if node:
                protect_body.append(node)
        
        # Parse handle block
        handle_body = []
        if self._peek().type == 'HANDLE' or (self._peek().type == 'INDENT' and self._peek_ahead(1).type == 'HANDLE'):
            if self._peek().type == 'INDENT':
                self._advance()
            self._consume('HANDLE')
            self._consume('NEWLINE')
            handle_body = self._parse_block()
        
        return ProtectNode(protect_body, handle_body)
    
    def _parse_handle_inline(self) -> HandleInlineNode:
        """Parse inline error handler"""
        token = self._advance()
        return HandleInlineNode(token.value['error'], token.value['action'])
    
    def _parse_run_parallel(self) -> RunParallelNode:
        """Parse parallel task execution"""
        token = self._advance()
        return RunParallelNode(token.value)
    
    def _parse_return(self) -> ReturnNode:
        """Parse return statement"""
        token = self._advance()
        return ReturnNode(token.value)
    
    def _peek_ahead(self, offset: int) -> Token:
        """Look ahead at token at offset"""
        pos = self.pos + offset
        if pos < len(self.tokens):
            return self.tokens[pos]
        return Token('EOF', None)
    
    def _peek(self) -> Token:
        """Look at current token without consuming"""
        if self.pos < len(self.tokens):
            return self.tokens[self.pos]
        return Token('EOF', None)
    
    def _advance(self) -> Token:
        """Consume and return current token"""
        token = self._peek()
        self.pos += 1
        return token
    
    def _consume(self, expected_type: str):
        """Consume token of expected type"""
        token = self._peek()
        if token.type == expected_type:
            self._advance()
    
    def _is_at_end(self) -> bool:
        """Check if we're at end of tokens"""
        return self._peek().type == 'EOF'


# ============================================================================
# INTERPRETER - Executes the AST
# ============================================================================

class Interpreter:
    """Executes Athera programs"""
    
    def __init__(self):
        self.tasks = {}  # Store defined tasks
        self.variables = {}  # Store variables
        self.modules = {}  # Store imported modules
        self.return_value = None  # Store return value from tasks
        self.error_occurred = False  # Track if error occurred in protect block
        self.last_error = None  # Store last error message
    
    def execute(self, nodes: List[ASTNode]):
        """Execute list of AST nodes"""
        for node in nodes:
            self._execute_node(node)
    
    def _execute_node(self, node: ASTNode):
        """Execute a single AST node"""
        if isinstance(node, TaskNode):
            self._execute_task_definition(node)
        elif isinstance(node, GreetNode):
            self._execute_greet(node)
        elif isinstance(node, BackupNode):
            self._execute_backup(node)
        elif isinstance(node, CheckNode):
            self._execute_check(node)
        elif isinstance(node, RepeatNNode):
            self._execute_repeat_n(node)
        elif isinstance(node, RepeatEachNode):
            self._execute_repeat_each(node)
        elif isinstance(node, SetNode):
            self._execute_set(node)
        elif isinstance(node, RunNode):
            self._execute_run(node)
        elif isinstance(node, UseNode):
            self._execute_use(node)
        elif isinstance(node, ProtectNode):
            self._execute_protect(node)
        elif isinstance(node, HandleInlineNode):
            self._execute_handle_inline(node)
        elif isinstance(node, RunParallelNode):
            self._execute_run_parallel(node)
        elif isinstance(node, ReturnNode):
            self._execute_return(node)
    
    def _execute_task_definition(self, node: TaskNode):
        """Store task definition"""
        self.tasks[node.name] = {
            'body': node.body,
            'params': getattr(node, 'params', [])
        }
    
    def _execute_greet(self, node: GreetNode):
        """Execute greet command - print message"""
        message = self._resolve_value(node.message)
        print(message)
    
    def _execute_backup(self, node: BackupNode):
        """Execute backup command - copy files"""
        source = self._resolve_value(node.source)
        dest = self._resolve_value(node.dest)
        
        # Remove quotes from dest if present
        dest = dest.strip('"').strip("'")
        source = source.strip('"').strip("'")
        
        try:
            # Create destination directory if it doesn't exist
            dest_path = Path(dest)
            dest_path.mkdir(parents=True, exist_ok=True)
            
            # Get source file name
            source_path = Path(source)
            if source_path.exists():
                dest_file = dest_path / source_path.name
                shutil.copy2(source_path, dest_file)
                print(f"[Backed up: {source} -> {dest_file}]")
            else:
                print(f"[Warning: Source file not found: {source}]")
        except Exception as e:
            print(f"[Backup error: {e}]")
    
    def _execute_check(self, node: CheckNode):
        """Execute check condition"""
        condition = node.condition
        
        # Evaluate condition
        if self._evaluate_condition(condition):
            # Execute the action (parse and execute as a mini-statement)
            self._execute_inline_action(node.action)
    
    def _evaluate_condition(self, condition: str) -> bool:
        """Evaluate a condition expression"""
        condition = condition.strip()
        
        # Check if it's a file/path check (string in quotes)
        if condition.startswith('"') or condition.startswith("'"):
            # File existence check
            path = condition.strip('"').strip("'")
            path = self._resolve_value(path)
            return Path(path).exists()
        
        # Check if it's a variable
        if condition in self.variables:
            value = self.variables[condition]
            # Truthy check
            if isinstance(value, bool):
                return value
            elif isinstance(value, (int, float)):
                return value != 0
            elif isinstance(value, str):
                return len(value) > 0
            elif isinstance(value, (list, dict)):
                return len(value) > 0
            return bool(value)
        
        # Try to evaluate as expression
        try:
            # Replace variables in expression
            for var, val in self.variables.items():
                if var in condition:
                    if isinstance(val, str):
                        condition = condition.replace(var, f'"{val}"')
                    else:
                        condition = condition.replace(var, str(val))
            
            # Safe evaluation
            result = eval(condition, {"__builtins__": {}}, {})
            return bool(result)
        except:
            return False
    
    def _execute_inline_action(self, action: str):
        """Execute an inline action from check statement"""
        # Simple greet action
        if action.startswith('greet '):
            message = self._extract_string(action[6:])
            message = self._resolve_value(message)
            print(message)
    
    def _execute_repeat_n(self, node: RepeatNNode):
        """Execute repeat N times loop"""
        for _ in range(node.count):
            for stmt in node.body:
                self._execute_node(stmt)
    
    def _execute_repeat_each(self, node: RepeatEachNode):
        """Execute repeat each loop"""
        # Evaluate list expression
        list_value = self._evaluate_expression(node.list_expr)
        
        if not isinstance(list_value, list):
            print(f"[Warning: Expected list, got {type(list_value).__name__}]")
            return
        
        # Iterate over list
        for item in list_value:
            # Set loop variable
            self.variables[node.var] = item
            
            # Execute loop body
            for stmt in node.body:
                self._execute_node(stmt)
    
    def _execute_set(self, node: SetNode):
        """Execute variable assignment"""
        value = self._evaluate_expression(node.value_expr)
        self.variables[node.var] = value
    
    def _execute_run(self, node: RunNode):
        """Execute task call"""
        # Parse task name and arguments
        task_call = node.task_name.strip()
        task_name = task_call
        args = []
        
        # Check for arguments: task_name arg1, arg2, arg3
        if ' ' in task_call:
            parts = task_call.split(' ', 1)
            task_name = parts[0]
            args_str = parts[1]
            # Parse comma-separated arguments
            args = [self._evaluate_expression(arg.strip()) for arg in args_str.split(',')]
        
        if task_name in self.tasks:
            task_def = self.tasks[task_name]
            task_body = task_def['body']
            task_params = task_def['params']
            
            # Save current variable state
            saved_vars = self.variables.copy()
            
            # Bind parameters to arguments
            for i, param in enumerate(task_params):
                if i < len(args):
                    self.variables[param] = args[i]
                else:
                    self.variables[param] = None
            
            # Execute task body
            self.return_value = None
            for stmt in task_body:
                self._execute_node(stmt)
                if self.return_value is not None:
                    break
            
            # Restore variables (but keep global changes)
            for param in task_params:
                if param in saved_vars:
                    self.variables[param] = saved_vars[param]
                elif param in self.variables:
                    del self.variables[param]
        else:
            print(f"[Error: Task '{task_name}' not found]")
    
    def _execute_use(self, node: UseNode):
        """Execute module import - load .ath files"""
        module_name = node.module_name
        
        # Try to load .ath file
        module_path = f"{module_name}.ath"
        if not Path(module_path).exists():
            # Try with modules/ prefix
            module_path = f"modules/{module_name}.ath"
        
        if Path(module_path).exists():
            try:
                with open(module_path, 'r') as f:
                    source = f.read()
                
                # Parse and execute the module
                from copy import deepcopy
                saved_tasks = deepcopy(self.tasks)
                
                lexer = Lexer(source)
                tokens = lexer.tokenize()
                parser = Parser(tokens)
                ast = parser.parse()
                
                # Execute module (this will define its tasks)
                for node in ast:
                    self._execute_node(node)
                
                # Store module reference
                self.modules[module_name] = {
                    'path': module_path,
                    'tasks': {k: v for k, v in self.tasks.items() if k not in saved_tasks}
                }
                
                print(f"[Imported module: {module_name} from {module_path}]")
            except Exception as e:
                print(f"[Error importing module {module_name}: {e}]")
        else:
            print(f"[Warning: Module {module_name} not found]")
    
    def _execute_protect(self, node: ProtectNode):
        """Execute protect-handle error handling block"""
        self.error_occurred = False
        self.last_error = None
        
        try:
            # Execute protected block
            for stmt in node.protect_body:
                self._execute_node(stmt)
        except Exception as e:
            self.error_occurred = True
            self.last_error = str(e)
            print(f"[Error caught: {e}]")
            
            # Execute handle block
            if node.handle_body:
                for stmt in node.handle_body:
                    self._execute_node(stmt)
    
    def _execute_handle_inline(self, node: HandleInlineNode):
        """Execute inline error handler"""
        if self.error_occurred:
            # Execute action if error matches
            self._execute_inline_action(node.action)
    
    def _execute_run_parallel(self, node: RunParallelNode):
        """Execute tasks in parallel using threads"""
        import threading
        
        def run_task(task_name):
            """Run a task in a thread"""
            # Create a copy of interpreter state for this thread
            if task_name.strip() in self.tasks:
                task_def = self.tasks[task_name.strip()]
                for stmt in task_def['body']:
                    self._execute_node(stmt)
        
        threads = []
        for task_name in node.task_names:
            thread = threading.Thread(target=run_task, args=(task_name,))
            threads.append(thread)
            thread.start()
        
        # Wait for all threads to complete
        for thread in threads:
            thread.join()
        
        print(f"[Parallel execution complete: {len(threads)} tasks]")
    
    def _execute_return(self, node: ReturnNode):
        """Execute return statement"""
        self.return_value = self._evaluate_expression(node.value_expr)
        return self.return_value
    
    def _evaluate_expression(self, expr: str) -> Any:
        """Evaluate an expression and return its value"""
        expr = expr.strip()
        
        # String literal
        if (expr.startswith('"') and expr.endswith('"')) or \
           (expr.startswith("'") and expr.endswith("'")):
            return expr[1:-1]
        
        # List literal: ["item1", "item2"]
        if expr.startswith('[') and expr.endswith(']'):
            return self._parse_list(expr)
        
        # Dictionary literal: {"key": "value"}
        if expr.startswith('{') and expr.endswith('}'):
            return self._parse_dict(expr)
        
        # Number
        try:
            if '.' in expr:
                return float(expr)
            return int(expr)
        except ValueError:
            pass
        
        # Boolean
        if expr.lower() == 'true':
            return True
        if expr.lower() == 'false':
            return False
        
        # Variable reference
        if expr in self.variables:
            return self.variables[expr]
        
        # Try to evaluate as Python expression (for math, etc.)
        try:
            # Replace variables
            eval_expr = expr
            for var, val in self.variables.items():
                if var in eval_expr:
                    if isinstance(val, str):
                        eval_expr = eval_expr.replace(var, f'"{val}"')
                    else:
                        eval_expr = eval_expr.replace(var, str(val))
            
            result = eval(eval_expr, {"__builtins__": {}}, {})
            return result
        except:
            pass
        
        # Return as string if nothing else works
        return expr
    
    def _parse_list(self, expr: str) -> List[Any]:
        """Parse list literal"""
        expr = expr[1:-1].strip()  # Remove brackets
        if not expr:
            return []
        
        items = []
        current = ""
        in_quotes = False
        quote_char = None
        
        for char in expr:
            if char in ('"', "'") and (not in_quotes or char == quote_char):
                in_quotes = not in_quotes
                quote_char = char if in_quotes else None
                current += char
            elif char == ',' and not in_quotes:
                items.append(self._evaluate_expression(current.strip()))
                current = ""
            else:
                current += char
        
        if current.strip():
            items.append(self._evaluate_expression(current.strip()))
        
        return items
    
    def _parse_dict(self, expr: str) -> Dict[str, Any]:
        """Parse dictionary literal"""
        # Simplified dict parsing
        try:
            return eval(expr)
        except:
            return {}
    
    def _resolve_value(self, value: Any) -> Any:
        """Resolve a value (could be variable or literal)"""
        if isinstance(value, str):
            # Check if it's a variable
            if value in self.variables:
                return self.variables[value]
        return value
    
    def _extract_string(self, s: str) -> str:
        """Extract string from quotes"""
        s = s.strip()
        if (s.startswith('"') and s.endswith('"')) or (s.startswith("'") and s.endswith("'")):
            return s[1:-1]
        return s


# ============================================================================
# MAIN - Run Athera programs
# ============================================================================

def run_athera_file(filename: str):
    """Run an Athera program from a file"""
    try:
        with open(filename, 'r') as f:
            source = f.read()
        
        run_athera_code(source)
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found")
    except Exception as e:
        print(f"Error running Athera program: {e}")


def run_athera_code(source: str):
    """Run Athera code from a string"""
    try:
        # Lexical analysis
        lexer = Lexer(source)
        tokens = lexer.tokenize()
        
        # Parsing
        parser = Parser(tokens)
        ast = parser.parse()
        
        # Interpretation
        interpreter = Interpreter()
        interpreter.execute(ast)
    except Exception as e:
        print(f"Error: {e}")
        import traceback
        traceback.print_exc()


def main():
    """Main entry point"""
    import sys
    
    if len(sys.argv) < 2:
        print("Athera Programming Language Interpreter")
        print("Usage: python athera_interpreter.py <filename.ath>")
        print("\nExample:")
        print("  python athera_interpreter.py morning_routine.ath")
        return
    
    filename = sys.argv[1]
    run_athera_file(filename)


if __name__ == "__main__":
    main()
