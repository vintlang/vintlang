export const codeExamples = {
    basics: `// Basic VintLang Operations
  let name = "VintLang"
  s = name.split("") 
  for i in s { 
      print(i)
  }
  
  // Type conversion
  age = "10"
  convert(age, "INTEGER")
  print(type(age))`,
    
    functions: `// Function Definition in VintLang
  let printDetails = func(name, age, height) {
      print("My name is " + name + 
            ", I am " + age + 
            " years old, and my height is " + 
            height + " feet.")
  }
  
  // Function call
  printDetails("VintLang", "10", "6.0")`,
    
    timeAndNet: `// Time and Network Operations
  import time
  import net
  
  // Time operations
  print(time.format(time.now(), "02-01-2006 15:04:05"))
  print(time.add(time.now(), "1h"))
  
  // Network request
  let res = net.get("https://example.com")
  print(res)`,
}