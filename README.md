
# GoShorty URL Shortener

This is a simple URL shortener written in Go. It uses JSON and YAML files to map short paths to long URLs.

The project is based on the [Build a URL Shortener with Go](https://gophercises.com/exercises/urlshort) exercise from [Gophercises](https://gophercises.com/).

## Packages

The main packages used in this project are:

- `net/http`: This is a built-in Go package used for building HTTP servers and clients.
- `flag`: This is a built-in Go package used for command-line option parsing.
- `fmt`: This is a built-in Go package used for formatted I/O.
- `github.com/mattn/go-sqlite3`: This is a third-party package used for SQLite3 database support.
- `github.com/bladev/goshorty/Database`: This is a custom package that contains the database functions. It is used to create a database connection and to query the database.
- `github.com/bladev/goshorty/Handler`: This is a custom package that contains the handlers for the HTTP server.

### Handlers

- defaultHandler: This handler is used to redirect the user to the default URL when they go to the root path.
- mapHandler: This handler is uses a map to map short paths to long URLs.
- dbHandler: This handler uses tokens stored in a database to map short paths to long URLs.
- yamlHandler and jsonHandler: These handlers are used to parse the YAML and JSON files respectively and return a mapHandler.
  
# Usage

You can run the server with either a JSON or YAML file by passing the file path as a command-line argument. Here's how you can do it:

```bash
go run main.go -json <path-to-file>
```

or

```bash
go run main.go -yaml <path-to-file>
```

The JSON or YAML file should contain a list of objects with the following structure:

```json
{
  {"path":"/google", "url":"https://www.google.com"},
  {"path":"/github", "url":"https://www.github.com"}
}
```

or

```yaml
- path: /<short-path>
  url: <long-url>
```

### Database

You can also use a database to store the short paths and long URLs. To do this, you need to create a database file and pass the file path as a command-line argument. Here's how you can do it:

```bash
go run main.go -db <path-to-file>
```

The database file should contain a table called `url_shortner` with the following structure:

```sql
CREATE TABLE url_shortner (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  time DATETIME NOT NULL,
  Url TEXT NOT NULL,
  Token TEXT NOT NULL
);
```

Example:

```sql
INSERT INTO url_shortner (time, Url, Token) VALUES ('2021-01-01 00:00:00', 'https://www.google.com', 'google');
INSERT INTO url_shortner (time, Url, Token) VALUES ('2021-01-01 00:00:00', 'https://www.github.com', 'github');
```


When the server is running, you can access the short URLs by going to `localhost:8080/<short-path>`. For example, if you have a short path `/google`, you can access it by going to `localhost:8080/google`.
