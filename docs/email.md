# Email Module in Vint

The Email module in Vint provides email validation and processing functions. This module helps you validate email addresses, extract components, and normalize email formats for consistent processing.

---

## Importing the Email Module

To use the Email module, simply import it:
```js
import email
```

---

## Functions and Examples

### 1. Validate Email Address (`validate`)
The `validate` function checks if a given string is a valid email address format.

**Syntax**:
```js
validate(emailAddress)
```

**Example**:
```js
import email

print("=== Email Validation Example ===")

// Valid email addresses
valid_emails = [
    "user@example.com",
    "john.doe@company.org",
    "test.email+tag@domain.co.uk"
]

// Invalid email addresses
invalid_emails = [
    "not-an-email",
    "@example.com",
    "user@",
    "user..name@example.com"
]

print("Valid Emails:")
for valid_email in valid_emails {
    result = email.validate(valid_email)
    print("  " + valid_email + " -> " + string(result))
}

print("\nInvalid Emails:")
for invalid_email in invalid_emails {
    result = email.validate(invalid_email)
    print("  " + invalid_email + " -> " + string(result))
}
```

---

### 2. Extract Domain (`extractDomain`)
The `extractDomain` function extracts the domain part from an email address.

**Syntax**:
```js
extractDomain(emailAddress)
```

**Example**:
```js
import email

print("=== Email Domain Extraction Example ===")
email_address = "john.doe@example.com"
domain = email.extractDomain(email_address)
print("Email:  ", email_address)
print("Domain: ", domain)
// Output: Domain: example.com
```

---

### 3. Extract Username (`extractUsername`)
The `extractUsername` function extracts the username part (before @) from an email address.

**Syntax**:
```js
extractUsername(emailAddress)
```

**Example**:
```js
import email

print("=== Email Username Extraction Example ===")
email_address = "john.doe@example.com"
username = email.extractUsername(email_address)
print("Email:    ", email_address)
print("Username: ", username)
// Output: Username: john.doe
```

---

### 4. Normalize Email (`normalize`)
The `normalize` function converts an email address to lowercase and trims whitespace for consistent formatting.

**Syntax**:
```js
normalize(emailAddress)
```

**Example**:
```js
import email

print("=== Email Normalization Example ===")
messy_email = "  John.DOE@EXAMPLE.COM  "
normalized = email.normalize(messy_email)
print("Original:   ", messy_email)
print("Normalized: ", normalized)
// Output: Normalized: john.doe@example.com
```

---

## Complete Usage Example

```js
import email

print("=== Email Module Complete Example ===")

// Process a list of email addresses
email_list = [
    "  Alice@COMPANY.COM  ",
    "bob.smith@example.org",
    "CHARLIE@DOMAIN.CO.UK",
    "invalid-email",
    "diana.jones@website.net"
]

print("Processing email addresses:")
for email_addr in email_list {
    print("\n--- Processing: " + email_addr + " ---")
    
    // Normalize the email first
    normalized = email.normalize(email_addr)
    print("Normalized: " + normalized)
    
    // Validate the email
    is_valid = email.validate(normalized)
    print("Valid: " + string(is_valid))
    
    if is_valid {
        // Extract components
        username = email.extractUsername(normalized)
        domain = email.extractDomain(normalized)
        
        print("Username: " + username)
        print("Domain: " + domain)
        
        // Example: Check for specific domains
        if domain == "company.com" {
            print("✓ Corporate email detected")
        } else {
            print("→ External email")
        }
    } else {
        print("✗ Invalid email format")
    }
}
```

---

## Use Cases

- **User Registration**: Validate email addresses during account creation
- **Email Lists**: Clean and normalize email addresses in mailing lists
- **Domain Analysis**: Extract domains for statistical analysis
- **Email Routing**: Route emails based on domain or username patterns
- **Data Cleaning**: Standardize email formats in databases

---

## Advanced Example: Email Domain Statistics

```js
import email

print("=== Email Domain Statistics ===")

emails = [
    "user1@gmail.com",
    "user2@company.com",
    "user3@gmail.com",
    "user4@yahoo.com",
    "user5@company.com",
    "user6@outlook.com"
]

domain_count = {}

for email_addr in emails {
    if email.validate(email_addr) {
        domain = email.extractDomain(email_addr)
        
        if domain in domain_count {
            domain_count[domain] = domain_count[domain] + 1
        } else {
            domain_count[domain] = 1
        }
    }
}

print("Domain statistics:")
for domain, count in domain_count {
    print("  " + domain + ": " + string(count) + " emails")
}
```

---

## Summary of Functions

| Function           | Description                                    | Return Type |
|--------------------|------------------------------------------------|-------------|
| `validate`         | Validates email address format                 | Boolean     |
| `extractDomain`    | Extracts domain part from email               | String      |
| `extractUsername`  | Extracts username part from email             | String      |
| `normalize`        | Normalizes email to lowercase and trims       | String      |

The Email module provides essential functionality for working with email addresses safely and efficiently in VintLang applications.