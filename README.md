# jekyll.nvim
## Usage
1. plz write graphql query in the current buffer.
``` graphql
query { viewer { login } }
```
2. Execute this command.
``` vim
:call JekyllCurl("https://api.github.com/graphql?access_token=<GITHUB_TOKEN>")
```

- 1st arg's URL. Next args's Header-Value pair.
``` vim
:call JekyllCurl(URL, HEADER1, VALUE1, HEADER2, VALUE2, ...)
```

- e.g, setting Authorization Header for GitHub API v4
``` vim
:call JekyllCurl("https://api.github.com/graphql", "Authorization", "Bearer <GITHUB_TOKEN>")
```
