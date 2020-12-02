#! /usr/bin/env bash
set -eu -o pipefail

# https://github.com/fatih/vim-go
# https://github.com/fatih/vim-go/blob/master/doc/vim-go.txt
# https://github.com/fatih/vim-go/wiki/Tutorial#identifier-highlighting
# https://github.com/fatih/vim-go/wiki/Tutorial#beautify-it
# https://vim.fandom.com/wiki/Copy,_cut_and_paste

mkdir -p ~/.vim/autoload/
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
git clone https://github.com/fatih/molokai.git ~/.vim/pack/plugins/start/molokai

# colorscheme: default, blue, darkblue, delek, desert, elford, evening, industry, koehler, morning,
#              murphy, pablo, peachpuff, ron, shine, slate, torte, zellner

# set softtabstop=4
# set listchars=tab: ,trail:.,eol:¬,extends:>,precedes:<
# set list
# colorscheme industry

# let g:go_highlight_operators = 1

test -s ~/.vimrc && cp ~/.vimrc. ~/.vimrc.bk || true


cat > ~/.vimrc << 'EOF'
set number
highlight LineNr ctermfg=grey ctermbg=black
set tabstop=4
set colorcolumn=101
set textwidth=100
highlight ColorColumn ctermbg=lightgrey guibg=black
syntax on
set autoindent
set list listchars=tab:❘⠀,trail:$,extends:»,precedes:«,nbsp:×,eol:$
set statusline=%f\ %l,%c
nmap <silent> <c-k> :wincmd k<CR>
nmap <silent> <c-j> :wincmd j<CR>
nmap <silent> <c-h> :wincmd h<CR>
nmap <silent> <c-l> :wincmd l<CR>

call plug#begin()
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
Plug 'fatih/molokai'
let g:go_code_completion_enabled = 1
let g:go_code_completion_icase = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_fields = 1
let g:go_highlight_variable_declarations = 1
let g:go_highlight_variable_assignments = 1
let g:go_highlight_types = 1
let g:go_highlight_extra_types = 1
let g:rehash256 = 1
let g:molokai_original = 1
call plug#end()
EOF
