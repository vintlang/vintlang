# VintLang Async Operations Guide

VintLang now supports native async operations and a simple concurrency model inspired by Go's goroutines and channels.

## Async Functions

Create async functions using the `async func` syntax. These functions return promises that can be awaited:

```javascript
// Define an async function
let fetchData = async func(url) {
    // Simulate async work
    return "Data from " + url
}

// Call the function (returns a promise)
let promise = fetchData("https://api.example.com")

// Wait for the result
let result = await promise
print("Result:", result)
```

## Async/Await

Use `await` to wait for promise resolution:

```javascript
let processData = async func(data) {
    return "Processed: " + data
}

let result = await processData("my data")
print(result)  // Output: Processed: my data
```

## Concurrent Execution with `go`

Use the `go` keyword to execute code concurrently:

```javascript
// Execute concurrently
go print("This runs in a goroutine")
go print("This also runs concurrently")

print("This runs in the main thread")
```

## Channels

Channels provide communication between concurrent operations.

### Creating Channels

```javascript
// Unbuffered channel
let ch = chan

// Buffered channel with size 5
let bufferedCh = chan(5)
```

### Channel Operations

```javascript
// Send to channel
send(ch, "Hello")

// Receive from channel
let message = receive(ch)

// Close channel
close(ch)
```

### Producer-Consumer Pattern

```javascript
let dataChan = chan(3)

// Producer goroutine
go func() {
    send(dataChan, "Item 1")
    send(dataChan, "Item 2")
    send(dataChan, "Item 3")
    close(dataChan)
}()

// Consumer
let item1 = receive(dataChan)
let item2 = receive(dataChan)
let item3 = receive(dataChan)

print("Received:", item1, item2, item3)
```

## Complex Example: Async with Channels

Combine async functions with channels for powerful patterns:

```javascript
let processInBackground = async func(input) {
    let resultChan = chan
    
    // Process in background
    go func() {
        let processed = "Processed: " + input
        send(resultChan, processed)
    }()
    
    // Wait for result
    let result = receive(resultChan)
    return result
}

let promise = processInBackground("data")
let result = await promise
print("Final result:", result)
```

## Error Handling

Async functions that encounter errors will reject their promises:

```javascript
let riskyFunction = async func() {
    // If an error occurs, the promise will be rejected
    return "Success!"
}

let result = await riskyFunction()
print("Result:", result)
```

## Multiple Concurrent Operations

Execute multiple async operations concurrently:

```javascript
let task1 = async func() { return "Task 1 complete" }
let task2 = async func() { return "Task 2 complete" }
let task3 = async func() { return "Task 3 complete" }

// Start all tasks
let p1 = task1()
let p2 = task2()
let p3 = task3()

// Wait for all to complete
let r1 = await p1
let r2 = await p2
let r3 = await p3

print("All tasks done:", r1, r2, r3)
```

## Best Practices

1. Use async functions for operations that might take time
2. Use channels for communication between goroutines
3. Always close channels when done sending
4. Use buffered channels to avoid blocking
5. Combine async/await with goroutines for powerful concurrent patterns

The async operations in VintLang provide a simple yet powerful way to handle concurrency and asynchronous operations in your programs.