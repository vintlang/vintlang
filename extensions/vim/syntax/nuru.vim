" Sintaksia ya vint kwenye programu ya "vim"
" Lugha: vint

" Maneno tengwa
syntax keyword vintKeyword unda pakeji rudisha vunja endelea tupu
syntax keyword vintType fanya
syntax keyword vintBool kweli sikweli
syntax keyword vintConditional kama sivyo au
syntax match vintComparision /[!\|<>]/
syntax keyword vintLoop ktk while badili
syntax keyword vintLabel ikiwa kawaida

" Nambari
syntax match vintInt '[+-]\d\+' contained display
syntax match vintFloat '[+-]\d+\.\d*' contained display

" Viendeshaji
syntax match vintAssignment '='
syntax match vintLogicalOP /[\&!|]/

" Vitendakazi 
syntax keyword vintFunction andika aina jaza fungua

" Tungo
syntax region vintString start=/"/ skip=/\\"/ end=/"/
syntax region vintString start=/'/ skip=/\\'/ end=/'/

" Maoni
syntax match vintComment "//.*"
syntax region vintComment start="/\*" end="\*/"

" Fafanua sintaksia
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

