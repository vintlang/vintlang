package module

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/vintlang/vintlang/object"
)

var FileWatcherFunctions = map[string]object.ModuleFunction{}

func init() {
	FileWatcherFunctions["watch"] = watchFile
	FileWatcherFunctions["watchDir"] = watchDirectory
	FileWatcherFunctions["stopWatch"] = stopWatch
	FileWatcherFunctions["isWatching"] = isWatching
}

// Watcher structure to track file watchers
type fileWatcher struct {
	path       string
	callback   *object.Function
	interval   time.Duration
	lastMod    time.Time
	stopChan   chan struct{}
	isDir      bool
	recursive  bool
	extensions []string
}

var (
	watchers      = make(map[string]*fileWatcher)
	watchersMutex sync.Mutex
)

// watchFile watches a file for changes and calls a callback function when changes are detected
func watchFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "watch requires at least 2 arguments: file path and callback function"}
	}

	// Get file path
	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "file path must be a string"}
	}

	// Get callback function
	callback, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "callback must be a function"}
	}

	// Check if file exists
	fileInfo, err := os.Stat(path.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("file not found: %s", path.Value)}
	}

	if fileInfo.IsDir() {
		return &object.Error{Message: fmt.Sprintf("%s is a directory, use watchDir instead", path.Value)}
	}

	// Get interval (optional, default 1 second)
	interval := 1 * time.Second
	if intervalObj, ok := defs["interval"]; ok {
		if intervalInt, ok := intervalObj.(*object.Integer); ok {
			interval = time.Duration(intervalInt.Value) * time.Millisecond
		}
	}

	// Create a new watcher
	watcherID := path.Value
	watchersMutex.Lock()
	defer watchersMutex.Unlock()

	// Stop existing watcher if any
	if existingWatcher, exists := watchers[watcherID]; exists {
		close(existingWatcher.stopChan)
	}

	watcher := &fileWatcher{
		path:     path.Value,
		callback: callback,
		interval: interval,
		lastMod:  fileInfo.ModTime(),
		stopChan: make(chan struct{}),
		isDir:    false,
	}

	// Start watching
	go watchFileLoop(watcher)

	// Store the watcher
	watchers[watcherID] = watcher

	return &object.String{Value: watcherID}
}

// watchDirectory watches a directory for changes and calls a callback function when changes are detected
func watchDirectory(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "watchDir requires at least 2 arguments: directory path and callback function"}
	}

	// Get directory path
	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "directory path must be a string"}
	}

	// Get callback function
	callback, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: "callback must be a function"}
	}

	// Check if directory exists
	fileInfo, err := os.Stat(path.Value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("directory not found: %s", path.Value)}
	}

	if !fileInfo.IsDir() {
		return &object.Error{Message: fmt.Sprintf("%s is not a directory", path.Value)}
	}

	// Get interval (optional, default 1 second)
	interval := 1 * time.Second
	if intervalObj, ok := defs["interval"]; ok {
		if intervalInt, ok := intervalObj.(*object.Integer); ok {
			interval = time.Duration(intervalInt.Value) * time.Millisecond
		}
	}

	// Get recursive flag (optional, default false)
	recursive := false
	if recursiveObj, ok := defs["recursive"]; ok {
		if recursiveBool, ok := recursiveObj.(*object.Boolean); ok {
			recursive = recursiveBool.Value
		}
	}

	// Get extensions filter (optional)
	var extensions []string
	if extensionsObj, ok := defs["extensions"]; ok {
		if extensionsArr, ok := extensionsObj.(*object.Array); ok {
			for _, ext := range extensionsArr.Elements {
				if extStr, ok := ext.(*object.String); ok {
					extensions = append(extensions, extStr.Value)
				}
			}
		}
	}

	// Create a new watcher
	watcherID := path.Value
	watchersMutex.Lock()
	defer watchersMutex.Unlock()

	// Stop existing watcher if any
	if existingWatcher, exists := watchers[watcherID]; exists {
		close(existingWatcher.stopChan)
	}

	watcher := &fileWatcher{
		path:       path.Value,
		callback:   callback,
		interval:   interval,
		lastMod:    fileInfo.ModTime(),
		stopChan:   make(chan struct{}),
		isDir:      true,
		recursive:  recursive,
		extensions: extensions,
	}

	// Start watching
	go watchDirLoop(watcher)

	// Store the watcher
	watchers[watcherID] = watcher

	return &object.String{Value: watcherID}
}

// stopWatch stops a file or directory watcher
func stopWatch(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "stopWatch requires exactly 1 argument: watcher ID"}
	}

	// Get watcher ID
	watcherID, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "watcher ID must be a string"}
	}

	watchersMutex.Lock()
	defer watchersMutex.Unlock()

	// Check if watcher exists
	watcher, exists := watchers[watcherID.Value]
	if !exists {
		return &object.Boolean{Value: false}
	}

	// Stop the watcher
	close(watcher.stopChan)
	delete(watchers, watcherID.Value)

	return &object.Boolean{Value: true}
}

// isWatching checks if a file or directory is being watched
func isWatching(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "isWatching requires exactly 1 argument: path"}
	}

	// Get path
	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "path must be a string"}
	}

	watchersMutex.Lock()
	defer watchersMutex.Unlock()

	_, exists := watchers[path.Value]
	return &object.Boolean{Value: exists}
}

// watchFileLoop is the main loop for watching a file
func watchFileLoop(watcher *fileWatcher) {
	ticker := time.NewTicker(watcher.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check if file has been modified
			fileInfo, err := os.Stat(watcher.path)
			if err != nil {
				// File might have been deleted
				continue
			}

			if fileInfo.ModTime() != watcher.lastMod {
				// File has been modified, call the callback
				watcher.lastMod = fileInfo.ModTime()
				
				// Create event object
				event := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
				
				// Add path
				pathKey := object.HashKey{Type: object.STRING_OBJ, Value: 1}
				event.Pairs[pathKey] = object.DictPair{
					Key:   &object.String{Value: "path"},
					Value: &object.String{Value: watcher.path},
				}
				
				// Add event type
				typeKey := object.HashKey{Type: object.STRING_OBJ, Value: 2}
				event.Pairs[typeKey] = object.DictPair{
					Key:   &object.String{Value: "type"},
					Value: &object.String{Value: "modified"},
				}
				
				// Add timestamp
				timeKey := object.HashKey{Type: object.STRING_OBJ, Value: 3}
				event.Pairs[timeKey] = object.DictPair{
					Key:   &object.String{Value: "time"},
					Value: &object.String{Value: watcher.lastMod.String()},
				}
				
				// Call the callback with the event
				watcher.callback.Env.Set("@", event)
				Eval(watcher.callback.Body, watcher.callback.Env)
			}
		case <-watcher.stopChan:
			// Watcher has been stopped
			return
		}
	}
}

// watchDirLoop is the main loop for watching a directory
func watchDirLoop(watcher *fileWatcher) {
	ticker := time.NewTicker(watcher.interval)
	defer ticker.Stop()

	// Map to store last modification times of files
	fileModTimes := make(map[string]time.Time)

	// Initialize with current files
	err := filepath.Walk(watcher.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories if we're only interested in files
		if info.IsDir() && path != watcher.path {
			if !watcher.recursive {
				return filepath.SkipDir
			}
			return nil
		}

		// Check file extension if extensions filter is set
		if len(watcher.extensions) > 0 && !info.IsDir() {
			ext := filepath.Ext(path)
			matched := false
			for _, allowedExt := range watcher.extensions {
				if ext == allowedExt {
					matched = true
					break
				}
			}
			if !matched {
				return nil
			}
		}

		fileModTimes[path] = info.ModTime()
		return nil
	})

	if err != nil {
		// Error initializing, just return
		return
	}

	for {
		select {
		case <-ticker.C:
			// Check for changes in the directory
			newFiles := make(map[string]bool)
			
			err := filepath.Walk(watcher.path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Skip directories if we're only interested in files
				if info.IsDir() && path != watcher.path {
					if !watcher.recursive {
						return filepath.SkipDir
					}
					return nil
				}

				// Check file extension if extensions filter is set
				if len(watcher.extensions) > 0 && !info.IsDir() {
					ext := filepath.Ext(path)
					matched := false
					for _, allowedExt := range watcher.extensions {
						if ext == allowedExt {
							matched = true
							break
						}
					}
					if !matched {
						return nil
					}
				}

				newFiles[path] = true

				// Check if file is new or modified
				lastMod, exists := fileModTimes[path]
				if !exists || lastMod != info.ModTime() {
					// File is new or modified
					fileModTimes[path] = info.ModTime()
					
					// Create event object
					event := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
					
					// Add path
					pathKey := object.HashKey{Type: object.STRING_OBJ, Value: 1}
					event.Pairs[pathKey] = object.DictPair{
						Key:   &object.String{Value: "path"},
						Value: &object.String{Value: path},
					}
					
					// Add event type
					typeKey := object.HashKey{Type: object.STRING_OBJ, Value: 2}
					eventType := "modified"
					if !exists {
						eventType = "created"
					}
					event.Pairs[typeKey] = object.DictPair{
						Key:   &object.String{Value: "type"},
						Value: &object.String{Value: eventType},
					}
					
					// Add timestamp
					timeKey := object.HashKey{Type: object.STRING_OBJ, Value: 3}
					event.Pairs[timeKey] = object.DictPair{
						Key:   &object.String{Value: "time"},
						Value: &object.String{Value: info.ModTime().String()},
					}
					
					// Call the callback with the event
					watcher.callback.Env.Set("@", event)
					Eval(watcher.callback.Body, watcher.callback.Env)
				}
				
				return nil
			})

			if err != nil {
				// Error walking directory, just continue
				continue
			}

			// Check for deleted files
			for path := range fileModTimes {
				if !newFiles[path] {
					// File has been deleted
					
					// Create event object
					event := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
					
					// Add path
					pathKey := object.HashKey{Type: object.STRING_OBJ, Value: 1}
					event.Pairs[pathKey] = object.DictPair{
						Key:   &object.String{Value: "path"},
						Value: &object.String{Value: path},
					}
					
					// Add event type
					typeKey := object.HashKey{Type: object.STRING_OBJ, Value: 2}
					event.Pairs[typeKey] = object.DictPair{
						Key:   &object.String{Value: "type"},
						Value: &object.String{Value: "deleted"},
					}
					
					// Add timestamp
					timeKey := object.HashKey{Type: object.STRING_OBJ, Value: 3}
					event.Pairs[timeKey] = object.DictPair{
						Key:   &object.String{Value: "time"},
						Value: &object.String{Value: time.Now().String()},
					}
					
					// Call the callback with the event
					watcher.callback.Env.Set("@", event)
					Eval(watcher.callback.Body, watcher.callback.Env)
					
					// Remove from map
					delete(fileModTimes, path)
				}
			}
			
		case <-watcher.stopChan:
			// Watcher has been stopped
			return
		}
	}
}