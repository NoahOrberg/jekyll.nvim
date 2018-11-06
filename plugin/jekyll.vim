scriptencoding utf-8

if exists('g:loaded_jekyll')
    finish
endif
let g:loaded_jekyll = 1

let s:save_cpo = &cpo
set cpo&vim

function! s:RequireJekyll(host) abort
  return jobstart(['jekyll.nvim'], { 'rpc': v:true })
endfunction

call remote#host#Register('jekyll.nvim', '0', function('s:RequireJekyll'))
call remote#host#RegisterPlugin('jekyll.nvim', '0', [
\ {'type': 'function', 'name': 'JekyllCurl', 'sync': 1, 'opts': {}},
\ ])

let &cpo = s:save_cpo
unlet s:save_cpo
