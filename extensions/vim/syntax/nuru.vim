" Vint language syntax for Vim
" Language: vint

" Keywords
syntax keyword vintKeyword create package return break continue empty
syntax keyword vintType func
syntax keyword vintBool true false
syntax keyword vintConditional if else or
syntax match vintComparision /[!\|<>]/
syntax keyword vintLoop for while change
syntax keyword vintLabel if default

" Numbers
syntax match vintInt '[+-]\d\+' contained display
syntax match vintFloat '[+-]\d+\.\d*' contained display

" Operators
syntax match vintAssignment '='
syntax match vintLogicalOP /[\&!|]/

" Functions
syntax keyword vintFunction write type fill open

" Strings
syntax region vintString start=/"/ skip=/\\"/ end=/"/
syntax region vintString start=/'/ skip=/\\'/ end=/'/

" Comments
syntax match vintComment "//.*"
syntax region vintComment start="/\*" end="\*/"

" Define syntax
let b:current_syntax = "vint"

highlight def link vintComment Comment
highlight def link vintBool Boolean
highlight def link vintFunction Function
highlight def link vintComparision Conditional
highlight def link vintConditional Conditional
highlight def link vintKeyword Keyword
highlight def link vintString String
highlight def link vintVariable Identifier
highlight def link vintLoop Repeat
highlight def link vintInt Number
highlight def link vintFloat Float
highlight def link vintAssignment Operator
highlight def link vintLogicalOP Operator
highlight def link vintAriOP Operator
highlight def link vintType Type
highlight def link vintLabel Label
