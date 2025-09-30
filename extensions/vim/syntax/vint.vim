" VintLang syntax highlighting for Vim
" Language: VintLang
" Maintainer: VintLang Team
" Latest Revision: 2025

if exists("b:current_syntax")
  finish
endif

" Keywords and Control Flow
syntax keyword vintKeyword let const func return break continue null
syntax keyword vintKeyword import package include async await go chan throw
syntax keyword vintKeyword defer repeat match as is
syntax keyword vintConditional if else switch case default
syntax keyword vintLoop for while in
syntax keyword vintBool true false

" Declarative Statements (both lowercase and capitalized)
syntax keyword vintDeclarative todo warn error info debug note success trace fatal critical log
syntax keyword vintDeclarative Todo Warn Error Info Debug Note Success Trace Fatal Critical Log

" Types and Functions
syntax keyword vintType func async
syntax keyword vintBuiltin println print len type

" Operators
syntax match vintOperator "="
syntax match vintOperator "+="
syntax match vintOperator "-="
syntax match vintOperator "*="
syntax match vintOperator "/="
syntax match vintOperator "%="
syntax match vintOperator "=="
syntax match vintOperator "!="
syntax match vintOperator "<="
syntax match vintOperator ">="
syntax match vintOperator "<"
syntax match vintOperator ">"
syntax match vintOperator "+"
syntax match vintOperator "-"
syntax match vintOperator "*"
syntax match vintOperator "/"
syntax match vintOperator "%"
syntax match vintOperator "**"
syntax match vintOperator "&&"
syntax match vintOperator "||"
syntax match vintOperator "!"
syntax match vintOperator "&"
syntax match vintOperator "|"
syntax match vintOperator "??"
syntax match vintOperator "++"
syntax match vintOperator "--"
syntax match vintOperator "=>"
syntax match vintOperator "\.\."

" Special symbols
syntax match vintSpecial "@"

" Numbers
syntax match vintNumber '\<\d\+\>'
syntax match vintNumber '\<\d\+\.\d*\>'
syntax match vintNumber '\<\d*\.\d\+\>'
syntax match vintNumber '\<\d\+[eE][+-]\=\d\+\>'
syntax match vintNumber '\<\d\+\.\d*[eE][+-]\=\d\+\>'
syntax match vintNumber '\<\d*\.\d\+[eE][+-]\=\d\+\>'

" Strings
syntax region vintString start=/"/ skip=/\\"/ end=/"/ contains=vintStringEscape
syntax region vintString start=/'/ skip=/\\'/ end=/'/ contains=vintStringEscape
syntax match vintStringEscape contained "\\[nrt\\\"']"

" Identifiers
syntax match vintIdentifier '[a-zA-Z_][a-zA-Z0-9_]*'

" Function calls
syntax match vintFunctionCall '[a-zA-Z_][a-zA-Z0-9_]*\s*('he=e-1

" Comments
syntax match vintComment "//.*$"
syntax region vintComment start="/\*" end="\*/" contains=vintCommentTodo
syntax keyword vintCommentTodo contained TODO FIXME XXX NOTE

" Delimiters
syntax match vintDelimiter "[\[\]{}(),;:]"
syntax match vintDelimiter "\."

" Shebang
syntax match vintShebang "^#!.*"

" Define the default highlighting
let b:current_syntax = "vint"

" Highlight groups
highlight def link vintKeyword Keyword
highlight def link vintConditional Conditional
highlight def link vintLoop Repeat
highlight def link vintBool Boolean
highlight def link vintDeclarative PreProc
highlight def link vintType Type
highlight def link vintBuiltin Function
highlight def link vintOperator Operator
highlight def link vintSpecial Special
highlight def link vintNumber Number
highlight def link vintString String
highlight def link vintStringEscape SpecialChar
highlight def link vintIdentifier Identifier
highlight def link vintFunctionCall Function
highlight def link vintComment Comment
highlight def link vintCommentTodo Todo
highlight def link vintDelimiter Delimiter
highlight def link vintShebang PreProc
