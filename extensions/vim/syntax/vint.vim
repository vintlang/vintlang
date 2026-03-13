" VintLang syntax highlighting for Vim
" Language: VintLang
" Maintainer: VintLang Team
" Latest Revision: 2026

if exists("b:current_syntax")
  finish
endif

" Keywords and Control Flow
syntax keyword vintKeyword let const func return break continue null
syntax keyword vintKeyword import package include async await go chan throw
syntax keyword vintKeyword defer repeat match as is enum struct

" Conditional and Loop
syntax keyword vintConditional if else switch case default
syntax keyword vintLoop for while in

" Boolean and Null Literals
syntax keyword vintBool true false
syntax keyword vintNull null

" Declarative Statements (both lowercase and capitalized)
syntax keyword vintDeclarative todo warn error info debug note success trace fatal critical log
syntax keyword vintDeclarative Todo Warn Error Info Debug Note Success Trace Fatal Critical Log

" Types
syntax keyword vintType func async struct enum

" Builtin Functions
syntax keyword vintBuiltin print println printErr printlnErr input
syntax keyword vintBuiltin len type format startsWith endsWith chr ord debounce
syntax keyword vintBuiltin range append pop indexOf unique
syntax keyword vintBuiltin keys values has_key
syntax keyword vintBuiltin convert string int parseInt parseFloat
syntax keyword vintBuiltin and or not xor nand nor eq
syntax keyword vintBuiltin send receive close
syntax keyword vintBuiltin exit sleep args
syntax keyword vintBuiltin open write

" Standard Library Modules
syntax keyword vintModule os time datetime net http json math cli term uuid
syntax keyword vintModule string styled crypto regex shell dotenv sysinfo
syntax keyword vintModule sqlite mysql postgres path random csv encoding colors
syntax keyword vintModule vintSocket vintChart llm openai schedule logger hash
syntax keyword vintModule xml url email reflect yaml clipboard redis kv jwt
syntax keyword vintModule excel fmt make errors

" Operators
syntax match vintOperator "=>"
syntax match vintOperator "\.\.\."
syntax match vintOperator "\.\."
syntax match vintOperator "??"
syntax match vintOperator "++"
syntax match vintOperator "--"
syntax match vintOperator "+="
syntax match vintOperator "-="
syntax match vintOperator "\*="
syntax match vintOperator "/="
syntax match vintOperator "%="
syntax match vintOperator "\*\*"
syntax match vintOperator "&&"
syntax match vintOperator "||"
syntax match vintOperator "=="
syntax match vintOperator "!="
syntax match vintOperator "<="
syntax match vintOperator ">="
syntax match vintOperator "="
syntax match vintOperator "+"
syntax match vintOperator "-"
syntax match vintOperator "\*"
syntax match vintOperator "/"
syntax match vintOperator "%"
syntax match vintOperator "<"
syntax match vintOperator ">"
syntax match vintOperator "!"
syntax match vintOperator "&"
syntax match vintOperator "|"

" Special symbols
syntax match vintSpecial "@"

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

" Strings
syntax region vintString start=/"/ skip=/\\"/ end=/"/ contains=vintStringEscape,vintStringInterp
syntax region vintString start=/'/ skip=/\\'/ end=/'/ contains=vintStringEscape
syntax match vintStringEscape contained "\\[nrt\\\"'0]"
syntax region vintStringInterp contained start="\${" end="}" contains=TOP

" Function definitions
syntax match vintFuncDef '\<func\>\s\+[a-zA-Z_][a-zA-Z0-9_]*'hs=s+5 contains=vintKeyword

" Function calls
syntax match vintFunctionCall '[a-zA-Z_][a-zA-Z0-9_]*\s*('he=e-1

" Method calls (dot notation)
syntax match vintMethodCall '\.[a-zA-Z_][a-zA-Z0-9_]*\s*('hs=s+1,he=e-1

" Comments
syntax match vintComment "//.*$" contains=vintCommentTodo
syntax region vintComment start="/\*" end="\*/" contains=vintCommentTodo
syntax keyword vintCommentTodo contained TODO FIXME XXX NOTE HACK BUG

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
highlight def link vintBool Boolean
highlight def link vintNull Constant
highlight def link vintDeclarative PreProc
highlight def link vintType Type
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
highlight def link vintMethodCall Function
highlight def link vintComment Comment
highlight def link vintCommentTodo Todo
highlight def link vintDelimiter Delimiter
highlight def link vintDot Delimiter
highlight def link vintShebang PreProc
