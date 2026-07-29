" VintLang syntax highlighting for Vim
" Language: VintLang
" Maintainer: VintLang Team
" Latest Revision: 2026

if exists("b:current_syntax")
  finish
endif

" Keywords and Control Flow
syntax keyword vintKeyword let const return break continue
syntax keyword vintKeyword import package include
syntax keyword vintKeyword defer repeat throw

" Conditionals and Loops
syntax keyword vintConditional if else switch case default
syntax keyword vintLoop for while in

" Concurrency Keywords
syntax keyword vintConcurrency async await go chan

" Pattern Matching
syntax keyword vintMatch match as is

" Boolean and Null Literals
syntax keyword vintBool true false
syntax keyword vintNull null

" Declarative Statements (comment-like: lowercase)
syntax keyword vintDeclarative todo warn error info debug note success
syntax keyword vintDeclarative trace fatal critical log

" Declarative Statements (active output: capitalized)
syntax keyword vintDeclarative Todo Warn Error Info Debug Note Success
syntax keyword vintDeclarative Trace Fatal Critical Log

" Type Keywords
syntax keyword vintType func struct enum chan

" Builtin Type Names (used in type annotations)
syntax keyword vintTypeName int float string bool byte any never

" Builtin Functions
syntax keyword vintBuiltin print println printErr printlnErr input
syntax keyword vintBuiltin len type format startsWith endsWith chr ord debounce
syntax keyword vintBuiltin range append pop indexOf unique
syntax keyword vintBuiltin keys values has_key copy clone pow
syntax keyword vintBuiltin convert string int parseInt parseFloat
syntax keyword vintBuiltin and or not xor nand nor eq
syntax keyword vintBuiltin send receive close
syntax keyword vintBuiltin exit sleep args
syntax keyword vintBuiltin open write
syntax keyword vintBuiltin import

" Type Predicate Builtins
syntax keyword vintBuiltin is_null is_int is_float is_string is_bool
syntax keyword vintBuiltin is_array is_dict is_function is_error is_number

" Standard Library Modules
syntax keyword vintModule os time datetime net http json math cli term uuid
syntax keyword vintModule string styled crypto regex shell dotenv sysinfo
syntax keyword vintModule sqlite mysql postgres path random csv encoding colors
syntax keyword vintModule vintSocket vintChart llm openai schedule logger hash
syntax keyword vintModule xml url email reflect yaml clipboard redis kv jwt
syntax keyword vintModule excel fmt make errors

" Arithmetic Operators
syntax match vintOperator "+"
syntax match vintOperator "-"
syntax match vintOperator "\*"
syntax match vintOperator "/"
syntax match vintOperator "%"
syntax match vintOperator "\*\*"

" Assignment Operators
syntax match vintOperator "="
syntax match vintOperator "+="
syntax match vintOperator "-="
syntax match vintOperator "\*="
syntax match vintOperator "/="
syntax match vintOperator "%="

" Increment/Decrement
syntax match vintOperator "++"
syntax match vintOperator "--"

" Comparison Operators
syntax match vintOperator "=="
syntax match vintOperator "!="
syntax match vintOperator "<="
syntax match vintOperator ">="
syntax match vintOperator "<"
syntax match vintOperator ">"

" Logical Operators
syntax match vintOperator "&&"
syntax match vintOperator "||"
syntax match vintOperator "!"

" Bitwise Operators
syntax match vintOperator "&"
syntax match vintOperator "|"

" Type System Operators
syntax match vintOperator "?"

" Other Operators
syntax match vintOperator "=>"
syntax match vintOperator "\.\.\."
syntax match vintOperator "\.\."
syntax match vintOperator "??"
" Special symbols
syntax match vintSpecial "@"
syntax match vintSpecial "::"

" Numbers (integer, float, scientific notation, hex, octal, binary)
syntax match vintNumber '\<0[xX][0-9a-fA-F]\+\>'
syntax match vintNumber '\<0[oO][0-7]\+\>'
syntax match vintNumber '\<0[bB][01]\+\>'
syntax match vintNumber '\<\d\+\>'
syntax match vintNumber '\<\d\+\.\d*\>'
syntax match vintNumber '\<\d*\.\d\+\>'
syntax match vintNumber '\<\d\+[eE][+-]\=\d\+\>'
syntax match vintNumber '\<\d\+\.\d*[eE][+-]\=\d\+\>'
syntax match vintNumber '\<\d*\.\d\+[eE][+-]\=\d\+\>'
syntax match vintNumber '\<\d\+_\=\d\+\>'

" Strings (double-quoted with interpolation, single-quoted raw)
syntax region vintString start=/"/ skip=/\\"/ end=/"/ contains=vintStringEscape,vintStringInterp
syntax region vintString start=/'/ skip=/\\'/ end=/'/

" String Escape Sequences
syntax match vintStringEscape contained "\\[nrt\\\"'0]"
syntax match vintStringEscape contained "\\x[0-9a-fA-F]\{2}"
syntax match vintStringEscape contained "\\u[0-9a-fA-F]\{4}"

" String Interpolation
syntax region vintStringInterp contained start="\${" end="}" contains=TOP

" Backtick strings (raw)
syntax region vintString start=/`/ end=/`/

" Function definitions: func name(...)
syntax match vintFuncDef '\<func\>\s\+[a-zA-Z_][a-zA-Z0-9_]*'hs=s+5 contains=vintKeyword

" Function calls: name(
syntax match vintFunctionCall '[a-zA-Z_][a-zA-Z0-9_]*\s*('he=e-1

" Builtin function calls via :: prefix
syntax match vintBuiltinCall '::[a-zA-Z_][a-zA-Z0-9_]*\s*('he=e-1

" Method calls: .name(
syntax match vintMethodCall '\.[a-zA-Z_][a-zA-Z0-9_]*\s*('hs=s+1,he=e-1

" Type annotations: :type after variable/parameter names
syntax match vintTypeAnnotation ':\s*\(int\|float\|string\|bool\|byte\|any\|never\|\[\]\(int\|float\|string\|bool\|byte\|any\)\)'hs=s+1

" Struct instantiation: TypeName{...}
syntax match vintStructLiteral '\<[A-Z][a-zA-Z0-9_]*\s*{'he=e-1

" Comments
syntax match vintComment "//.*$" contains=vintCommentTodo
syntax region vintComment start="/\*" end="\*/" contains=vintCommentTodo
syntax keyword vintCommentTodo contained TODO FIXME XXX NOTE HACK BUG WORKAROUND OPTIMIZE

" Delimiters
syntax match vintDelimiter "[\[\]{}(),;:]"
syntax match vintDot "\."

" Shebang
syntax match vintShebang "^#!.*"

" Define the default highlighting
let b:current_syntax = "vint"

" Highlight groups
highlight def link vintKeyword Keyword
highlight def link vintConditional Conditional
highlight def link vintLoop Repeat
highlight def link vintConcurrency Keyword
highlight def link vintMatch Keyword
highlight def link vintBool Boolean
highlight def link vintNull Constant
highlight def link vintDeclarative PreProc
highlight def link vintType Type
highlight def link vintTypeName Type
highlight def link vintBuiltin Function
highlight def link vintModule Include
highlight def link vintOperator Operator
highlight def link vintSpecial Special
highlight def link vintNumber Number
highlight def link vintString String
highlight def link vintStringEscape SpecialChar
highlight def link vintStringInterp Special
highlight def link vintFuncDef Function
highlight def link vintFunctionCall Function
highlight def link vintBuiltinCall Function
highlight def link vintMethodCall Function
highlight def link vintTypeAnnotation Type
highlight def link vintStructLiteral Structure
highlight def link vintComment Comment
highlight def link vintCommentTodo Todo
highlight def link vintDelimiter Delimiter
highlight def link vintDot Delimiter
highlight def link vintShebang PreProc
