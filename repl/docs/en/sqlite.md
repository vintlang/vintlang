# SQLite Module in VintLang

In **VintLang**, the `sqlite` module allows interaction with SQLite databases. You can open a database, execute queries, fetch data, and manage tables. This guide covers basic database operations.

## Open a Database

Use `sqlite.open()` to open a connection to an SQLite database.

```vint
db = sqlite.open("example.db")
```

## Close a Database

To close the database connection, use `sqlite.close()`.

```vint
sqlite.close(db)
```

## Execute a Query

You can execute `INSERT`, `UPDATE`, `DELETE`, and other queries using `sqlite.execute()`.

### Insert Data:

```vint
sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
```

### Update Data:

```vint
sqlite.execute(db, "UPDATE users SET age = ? WHERE name = ?", 26, "Alice")
```

## Fetch Data

Use `sqlite.fetchAll()` to retrieve all rows from a query.

```vint
users = sqlite.fetchAll(db, "SELECT * FROM users")
print(users)
```

You can also fetch a single row with `sqlite.fetchOne()`.

```vint
first_user = sqlite.fetchOne(db, "SELECT * FROM users LIMIT 1")
print(first_user)
```

## Create a Table

To create a new table, use `sqlite.createTable()`.

```vint
sqlite.createTable(db, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
```

## Drop a Table

Use `sqlite.dropTable()` to delete a table from the database.

```vint
sqlite.dropTable(db, "users")
```

## Example Usage

```vint
import sqlite

// Open a database
db = sqlite.open("example.db")

// Create a table
sqlite.createTable(db, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")

// Insert data
sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)

// Fetch all rows
users = sqlite.fetchAll(db, "SELECT * FROM users")
print(users)

// Fetch a single row
first_user = sqlite.fetchOne(db, "SELECT * FROM users LIMIT 1")
print(first_user)

// Close the database connection
sqlite.close(db)
```

By using the **VintLang** `sqlite` module, you can easily manage SQLite databases in your programs.
