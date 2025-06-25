# MySQL Module in VintLang

The `mysql` module in **VintLang** provides a way to interact with MySQL databases. You can connect to a database, execute queries, and fetch data.

## Connecting to a MySQL Database

To connect to a MySQL database, use `mysql.open()`. You need to provide a connection string in the following format: `user:password@tcp(host:port)/dbname`.

```vint
conn = mysql.open("user:password@tcp(127.0.0.1:3306)/testdb")
```

## Closing the Connection

Always close the connection when you're done with `mysql.close()`.

```vint
mysql.close(conn)
```

## Executing Queries

Use `mysql.execute()` for `INSERT`, `UPDATE`, `DELETE`, or any other queries that don't return rows.

```vint
// Inserting data with placeholders
insert_query = "INSERT INTO users (name, age) VALUES (?, ?)"
mysql.execute(conn, insert_query, "Alice", 30)
```

## Fetching Data

### Fetch All Rows

To get all rows from a query result, use `mysql.fetchAll()`.

```vint
users = mysql.fetchAll(conn, "SELECT * FROM users")
print(users)
```

### Fetch a Single Row

To get only the first row from a query result, use `mysql.fetchOne()`.

```vint
user = mysql.fetchOne(conn, "SELECT * FROM users WHERE id = ?", 1)
print(user)
```

## Full Example

Here's a complete example of how to use the `mysql` module:

```vint
import mysql

// Replace with your actual credentials
conn_str = "user:password@tcp(127.0.0.1:3306)/testdb"
conn = mysql.open(conn_str)

if conn.type() == "ERROR" {
    print("Error connecting to MySQL:", conn.message())
} else {
    print("Successfully connected to MySQL")

    // Create a table
    create_query = "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT)"
    mysql.execute(conn, create_query)

    // Insert data
    mysql.execute(conn, "INSERT INTO users (name, age) VALUES (?, ?)", "Bob", 35)

    // Fetch and print data
    users = mysql.fetchAll(conn, "SELECT * FROM users")
    print("All users:", users)

    // Close the connection
    mysql.close(conn)
    print("Connection closed")
} 