## [VIM](./vim)

The file contained herein provides syntax highlighting for vim.
The file should be saved in `$HOME/.vim/syntax/vint.vim`.
You should add the following line to your `.vimrc` or the appropriate location:

```vim
au BufRead,BufNewFile *.vint set filetype=vint
```

This provides highlighting for all VintLang keywords, builtins, standard library modules, operators, strings, numbers, and comments.
