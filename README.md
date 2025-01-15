# VintLang Installation Guide (Linux & macOS)

Follow the steps below to easily install **VintLang** on your Linux or macOS system.

---

### For Linux:

1. **Download the Binary:**

   First, download the **VintLang** binary for Linux. You can do this using the `curl` command. This will download the `tar.gz` file containing the binary to your current directory.

   ```bash
   curl -O -L https://github.com/vintlang/vintlang/releases/download/latest/vintLang_linux_amd64.tar.gz
   ```

2. **Extract the Binary to a Global Location:**

   After downloading the binary, you need to extract it into a directory that is globally accessible. `/usr/local/bin` is a commonly used directory for this purpose. The `tar` command will extract the contents of the `tar.gz` file and place them in `/usr/local/bin`.

   ```bash
   sudo tar -C /usr/local/bin -xzvf vintLang_linux_amd64.tar.gz
   ```

   This step ensures that the **VintLang** command can be used from anywhere on your system.

3. **Verify the Installation:**

   Once the extraction is complete, confirm that **VintLang** was installed successfully by checking its version. If the installation was successful, it will display the installed version of **VintLang**.

   ```bash
   vint -v
   ```

4. **Initialize a vint project:**

   Create a simple boilerplate vint project

   ```bash
   vint <optional:project-name>
   ```

---

5. **Install the vintlang extension from vscode**

   Install the official vint language support extension int vscode called **`vintlang`**

---

### For macOS:

1. **Download the Binary:**

   Begin by downloading the **VintLang** binary for macOS using the following `curl` command. This will download the `tar.gz` file for macOS to your current directory.

   ```bash
   curl -O -L https://github.com/vintlang/vintlang/releases/download/latest/vintLang_mac_amd64.tar.gz
   ```

2. **Extract the Binary to a Global Location:**

   Next, extract the downloaded binary to a globally accessible location. As with Linux, the standard directory for this on macOS is `/usr/local/bin`. Use the following command to extract the binary:

   ```bash
   sudo tar -C /usr/local/bin -xzvf vintLang_mac_amd64.tar.gz
   ```

   This allows you to run **VintLang** from any terminal window.

3. **Verify the Installation:**

   To check that the installation was successful, run the following command. It will output the version of **VintLang** that was installed:

   ```bash
   vint -v
   ```

4. **Initialize a vint project:**

   Create a simple boilerplate vint project

   ```bash
   vint init <optional:project-name>
   ```

---

5. **Install the vintlang extension from vscode**

   Install the official vint language support extension int vscode called **`vintlang`**

---


### Summary of Installation Steps:

1. **Download the Binary** using `curl` for your system (Linux or macOS).
2. **Extract the Binary** to `/usr/local/bin` (or another globally accessible directory).
3. **Verify the Installation** by checking the version with `vint -v`.
4. **Initialize a vintlang project** by running `vint init <projectname>`.
5. **Install the vintlang extension from vscode** install vintlang extension in vscode

## Sample Code

### Example 1: String Splitting and Printing

```js
// Importing modules
import net       // Importing networking module for HTTP operations
import time      // Importing time module to work with date and time

// Main logic to split and print characters of a string
let name = "VintLang"
s = name.split("")
for i in s {
    print(i)
}
```

### Example 2: Type Conversion and Conditional Statements

```js
// Demonstrating type conversion and conditional statements
age = "10"
convert(age, "INTEGER")  // Convert age string to integer
print(type(age))          // Uncomment to check the type of ageInInt

// Conditional statements to compare the age variable
if (age == 20) {
    print(age)
} else if (age == 10) {
    print("Age is " + age)
} else {
    print((age == "20"))
}
```

### Example 3: Working with Height Variable

```ja
// Working with height variable
height = "6.0" // Height in feet
print("My name is " + name)

// Define a function to print details
let printDetails = func(name, age, height) {
    print("My name is " + name + ", I am " + age + " years old, and my height is " + height + " feet.")
}

// Calling the printDetails function with initial values
printDetails(name, age, height)

// Update height and call the function again
height = "7"
printDetails(name, age, height)
```

### Example 4: Time-Related Operations

```vint
// Print the current timestamp
print(time.now())

// Function to greet a user based on the time of the day
let greet = func(nameParam) {
    let currentTime = time.now()  // Get the current time
    print(currentTime)            // Print the current time
    if (true) {                   // Placeholder condition, modify for actual logic
        print("Good morning, " + nameParam + "!")
    } else {
        print("Good evening, " + nameParam + "!")
    }
}

// Time-related operations
year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))
print(time.subtract(time.now(), "2h30m45s"))

// Call the greet function with a sample name
greet("John")
```

### Example 5: Networking with HTTP GET Request

```vint
// Example of a GET request using the net module
let res = net.get("https://tachera.com")
print(res)
```

### Example 6: Built-in Functions and Output

```vint
// Built-in functions
print(type(123))             // Print the type of an integer
let a = "123"                // Initialize a string variable
convert(a, "INTEGER")        // Convert the string to an integer
type(a)
print(a)                     // Check the type of the variable
print("Hello", "World")      // Print multiple values
write("Hello World")         // Write a string (useful in returning output)
```

### Step 4: Run the Sample Code

Once you have the sample code files saved, you can run them using the following command:

```bash
vint <filename>.vint
```

Replace `<filename>` with the actual name of the file you want to run (e.g., `hello.vint`, `fibonacci.vint`, etc.).

---
