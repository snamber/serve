# serve

the little, simple command line file server

Well know from node.js, this little tool simply serves static content. The following details can be configured:

- the path at which serve will serve the files, e.g. `/assets/`
- the directory to serve, e.g. `~` or `.` or `dir` or `/app/static/`
- the port to listen on
- whether to use basic auth or not, and if so, with which password
- whether to fall back to serving a file e.g. `index.html`, instead of serving a 404 (for Single Page Applications)

## Examples

Serve the current directory on `localhost:8080`:

```sh
serve
```

Serve the home directory on `localhost:8080`:

```sh
serve --dir $HOME
```

Serve the current directory on `localhost:3000`:

```sh
serve --port 3000
```

Serve the home directory on `localhost:8080`, protected with username "foo" and password "bar":

```sh
serve --dir $HOME --user:pass foo:bar
```

Serve an asset directory on `localhost:8080/static/`:

```sh
serve --dir assets --path static
```

Fall back to `index.html` instead of serving 404s:

```sh
serve --fallback ./index.html
```
