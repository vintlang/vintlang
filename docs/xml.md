# XML Module in Vint

The XML module in Vint provides basic XML processing capabilities including validation, value extraction, and character escaping/unescaping. This module helps you work with XML data safely and efficiently.

---

## Importing the XML Module

To use the XML module, simply import it:
```js
import xml
```

---

## Functions and Examples

### 1. Validate XML (`validate`)
The `validate` function checks if a given XML string is well-formed and valid.

**Syntax**:
```js
validate(xmlString)
```

**Example**:
```js
import xml

print("=== XML Validation Example ===")

// Valid XML
valid_xml = "<root><name>John</name><age>30</age></root>"
is_valid = xml.validate(valid_xml)
print("Valid XML:", is_valid)
// Output: Valid XML: true

// Invalid XML
invalid_xml = "<root><name>John</age></root>"
is_invalid = xml.validate(invalid_xml)
print("Invalid XML:", is_invalid)
// Output: Invalid XML: false
```

---

### 2. Extract Value from XML Tag (`extract`)
The `extract` function extracts the value from a specific XML tag.

**Syntax**:
```js
extract(xmlString, tagName)
```

**Example**:
```js
import xml

print("=== XML Value Extraction Example ===")
xml_data = "<user><name>John Doe</name><email>john@example.com</email></user>"

name = xml.extract(xml_data, "name")
email = xml.extract(xml_data, "email")

print("Name:", name)
print("Email:", email)
// Output: 
// Name: John Doe
// Email: john@example.com
```

---

### 3. Escape XML Characters (`escape`)
The `escape` function escapes special XML characters to make text safe for XML content.

**Syntax**:
```js
escape(text)
```

**Example**:
```js
import xml

print("=== XML Escape Example ===")
unsafe_text = "<script>alert('Hello & Goodbye');</script>"
safe_text = xml.escape(unsafe_text)

print("Original:", unsafe_text)
print("Escaped: ", safe_text)
// Output: Escaped: &lt;script&gt;alert(&#39;Hello &amp; Goodbye&#39;);&lt;/script&gt;
```

---

### 4. Unescape XML Entities (`unescape`)
The `unescape` function converts XML entities back to their original characters.

**Syntax**:
```js
unescape(escapedText)
```

**Example**:
```js
import xml

print("=== XML Unescape Example ===")
escaped_text = "&lt;tag&gt;Hello &amp; Goodbye&lt;/tag&gt;"
unescaped_text = xml.unescape(escaped_text)

print("Escaped:  ", escaped_text)
print("Unescaped:", unescaped_text)
// Output: Unescaped: <tag>Hello & Goodbye</tag>
```

---

## Complete Usage Example

```js
import xml

print("=== XML Module Complete Example ===")

// Create XML data
user_name = "John & Jane"
user_email = "user@example.com"

// Escape data for safe XML insertion
safe_name = xml.escape(user_name)
safe_email = xml.escape(user_email)

// Build XML string
xml_string = "<user><name>" + safe_name + "</name><email>" + safe_email + "</email></user>"
print("Generated XML:", xml_string)

// Validate the XML
if xml.validate(xml_string) {
    print("XML is valid!")
    
    // Extract values
    extracted_name = xml.extract(xml_string, "name")
    extracted_email = xml.extract(xml_string, "email")
    
    // Unescape extracted values
    final_name = xml.unescape(extracted_name)
    final_email = xml.unescape(extracted_email)
    
    print("Final Name:", final_name)
    print("Final Email:", final_email)
} else {
    print("Generated XML is invalid!")
}
```

---

## Use Cases

- **XML Document Processing**: Parse and extract data from XML files
- **Web Scraping**: Extract information from XML responses
- **Configuration Files**: Read XML configuration data
- **Data Exchange**: Safely prepare data for XML transmission
- **Template Processing**: Build XML documents dynamically

---

## Summary of Functions

| Function    | Description                                         | Return Type |
|-------------|-----------------------------------------------------|-------------|
| `validate`  | Validates XML structure and syntax                  | Boolean     |
| `extract`   | Extracts value from a specific XML tag             | String      |
| `escape`    | Escapes special characters for safe XML content    | String      |
| `unescape`  | Converts XML entities back to original characters  | String      |

The XML module provides essential functionality for working with XML data safely and efficiently in VintLang applications.