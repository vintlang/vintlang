# PostgreSQL Module in VintLang

The `postgres` module in **VintLang** allows you to interact with PostgreSQL databases. This guide will walk you through connecting, executing queries, and fetching data.

## Connecting to a PostgreSQL Database

To connect to a PostgreSQL database, use `postgres.open()`. The connection string should be in the format: `"user=youruser password=yourpassword dbname=yourdbname sslmode=disable"`.

```js
conn = postgres.open("user=postgres password=password dbname=testdb sslmode=disable")
```

## Closing the Connection

Make sure to close the connection with `postgres.close()` when you are finished.

```js
postgres.close(conn)
```

## Executing Queries

Use `postgres.execute()` for `INSERT`, `UPDATE`, `DELETE`, and other statements that do not return data. PostgreSQL uses `$1`, `$2`, etc., as placeholders.

```js
// Inserting data with placeholders
insert_query = "INSERT INTO users (name, age) VALUES ($1, $2)"
postgres.execute(conn, insert_query, "Alice", 30)
```

## Fetching Data

### Fetch All Rows

To retrieve all rows from a query, use `postgres.fetchAll()`.

```js
users = postgres.fetchAll(conn, "SELECT * FROM users")
print(users)
```

### Fetch a Single Row

To retrieve just one row, use `postgres.fetchOne()`.

```js
user = postgres.fetchOne(conn, "SELECT * FROM users WHERE id = $1", 1)
print(user)
```

## Full Example

Here is a complete example demonstrating the use of the `postgres` module:

```js
import postgres

// Replace with your actual credentials
conn_str = "user=postgres password=password dbname=testdb sslmode=disable"
conn = postgres.open(conn_str)

if conn.type() == "ERROR" {
    print("Error connecting to PostgreSQL:", conn.message())
} else {
    print("Successfully connected to PostgreSQL")

    // Create a table
    create_query = "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), age INT)"
    postgres.execute(conn, create_query)

    // Insert data
    postgres.execute(conn, "INSERT INTO users (name, age) VALUES ($1, $2)", "Bob", 35)

    // Fetch and print data
    users = postgres.fetchAll(conn, "SELECT * FROM users")
    print("All users:", users)

    // Close the connection
    postgres.close(conn)
    print("Connection closed")
} 