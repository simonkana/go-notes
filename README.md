# go-notes

A minimal and friendly notes manager

To run the project:

```
# Clone the project and move to it
git clone https://github.com/simonkana/go-notes
cd go-notes

# Create the SQLite database
touch notes.db

# Init the db schema from SQLC generated files
sqlite3 notes.db < internal/db/schema.sql

# Start the server
go run ./cmd/server
```

Dependencies:

- runtime:

  - Go (you can check if installed with the `go version` command): `brew install go`

<br/>

- development:

  - Air for hot-reloading: `go install github.com/air-verse/air@latest`

  - Tailwind CLI: `npm install tailwindcss @tailwindcss/cli`
