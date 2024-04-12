# Language Server Protocol in Golang

## A Basic LSP for neovim editor and for now its built for markdown file

### add below code to your autocmd.vim

````lua
local client = vim.lsp.start_client({
  name = "go_lsp",
  cmd = { "/Users/bhatraj/pgm files/my projects/lsp_go/main" },
})

if not client then
  print("hey, you didnt do clint thing good")
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```
