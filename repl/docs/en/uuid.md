# Using the UUID Module in Vint

The **UUID** (Universal Unique Identifier) module in Vint allows you to generate unique identifiers that are globally unique. These identifiers can be used for creating unique keys for database records, user sessions, or other purposes where a unique value is needed.

## Generating a UUID

You can generate a new UUID using the `uuid.generate()` function. It returns a unique identifier each time it is called.

### Example:

```vint
import uuid

// Generate and print a new UUID
print(uuid.generate())
```

Each time `uuid.generate()` is called, it generates a new, unique UUID value. This is useful for ensuring that each identifier is distinct across systems.

The generated UUID can be used for various purposes such as tracking sessions, unique IDs for objects, or database keys.