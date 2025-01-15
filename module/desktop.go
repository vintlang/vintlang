package module

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"
// 	"fyne.io/fyne/v2/theme"

// 	"github.com/vintlang/vintlang/object"
// )

// var DesktopAppFunctions = map[string]object.ModuleFunction{}

// func init() {
// 	DesktopAppFunctions["createApp"] = createApp
// 	DesktopAppFunctions["setTheme"] = setTheme
// 	DesktopAppFunctions["newLabel"] = newLabel
// 	DesktopAppFunctions["newButton"] = newButton
// 	DesktopAppFunctions["newEntry"] = newEntry
// 	DesktopAppFunctions["newList"] = newList
// 	DesktopAppFunctions["newCheckbox"] = newCheckbox
// 	DesktopAppFunctions["newSlider"] = newSlider
// 	DesktopAppFunctions["newProgressBar"] = newProgressBar
// 	DesktopAppFunctions["newTabs"] = newTabs
// 	DesktopAppFunctions["setWindowSize"] = setWindowSize
// 	DesktopAppFunctions["newImage"] = newImage
// 	DesktopAppFunctions["addToLayout"] = addToLayout
// 	DesktopAppFunctions["setTitle"] = setTitle
// 	DesktopAppFunctions["runApp"] = runApp
// }

// var appInstance fyne.App
// var mainWindow fyne.Window
// var layoutContainer *fyne.Container

// // Themes map
// var themeMap = map[string]fyne.Theme{
// 	"light":  theme.LightTheme(),
// 	"dark":   theme.DarkTheme(),
// 	"default": theme.DefaultTheme(),
// }

// func createApp(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "createApp requires exactly one argument: the app name"}
// 	}

// 	appName := args[0].Inspect()
// 	appInstance = app.NewWithID(appName)
// 	mainWindow = appInstance.NewWindow(appName)
// 	layoutContainer = container.NewVBox() // Default layout
// 	mainWindow.SetContent(layoutContainer)

// 	return &object.Boolean{Value: true}
// }

// func setTheme(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "setTheme requires exactly one argument: 'light', 'dark', or 'default'"}
// 	}

// 	themeName := args[0].Inspect()
// 	selectedTheme, exists := themeMap[themeName]
// 	if !exists {
// 		return &object.Error{Message: "Invalid theme name"}
// 	}

// 	fyne.CurrentApp().Settings().SetTheme(selectedTheme)
// 	return &object.Boolean{Value: true}
// }

// func newSlider(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "newSlider requires two arguments: min and max values"}
// 	}

// 	min, ok1 := args[0].(*object.Integer)
// 	max, ok2 := args[1].(*object.Integer)
// 	if !ok1 || !ok2 {
// 		return &object.Error{Message: "Arguments must be numbers"}
// 	}

// 	slider := widget.NewSlider(float64(min.Value), float64(max.Value))
// 	layoutContainer.Add(slider)
// 	return &object.Boolean{Value: true}
// }

// func newProgressBar(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "newProgressBar requires one argument: initial value (0.0 - 1.0)"}
// 	}

// 	value, ok := args[0].(*object.Integer)
// 	if !ok || value.Value < 0.0 || value.Value > 1.0 {
// 		return &object.Error{Message: "Value must be a number between 0.0 and 1.0"}
// 	}

// 	progressBar := widget.NewProgressBar()
// 	progressBar.SetValue(value.Value)
// 	layoutContainer.Add(progressBar)
// 	return &object.Boolean{Value: true}
// }

// func newTabs(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args)%2 != 0 {
// 		return &object.Error{Message: "newTabs requires an even number of arguments: pairs of tab name and tab content"}
// 	}

// 	var tabs []*container.TabItem
// 	for i := 0; i < len(args); i += 2 {
// 		tabName := args[i].Inspect()
// 		tabContent, ok := args[i+1].(*object.String)
// 		if !ok {
// 			return &object.Error{Message: "Tab content must be a string"}
// 		}
// 		tabs = append(tabs, container.NewTabItem(tabName, widget.NewLabel(tabContent.Value)))
// 	}

// 	tabContainer := container.NewAppTabs(tabs...)
// 	layoutContainer.Add(tabContainer)
// 	return &object.Boolean{Value: true}
// }

// func newLabel(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "newLabel requires exactly one argument: the text"}
// 	}

// 	labelText := args[0].Inspect()
// 	label := widget.NewLabel(labelText)
// 	layoutContainer.Add(label)
// 	return &object.Boolean{Value: true}
// }

// func newButton(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "newButton requires two arguments: the text and a callback"}
// 	}

// 	buttonText := args[0].Inspect()
// 	callbackFn, ok := args[1].(*object.Function)
// 	if !ok {
// 		return &object.Error{Message: "second argument must be a function"}
// 	}

// 	button := widget.NewButton(buttonText, func() {
// 		callbackFn.Call([]object.Object{}, map[string]object.Object{})
// 	})
// 	layoutContainer.Add(button)
// 	return &object.Boolean{Value: true}
// }

// func newEntry(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 0 {
// 		return &object.Error{Message: "newEntry requires no arguments"}
// 	}

// 	entry := widget.NewEntry()
// 	layoutContainer.Add(entry)
// 	return &object.Boolean{Value: true}
// }

// func newList(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "newList requires one argument: a list of strings"}
// 	}

// 	items, ok := args[0].(*object.Array)
// 	if !ok {
// 		return &object.Error{Message: "Argument must be a list"}
// 	}

// 	var data []string
// 	for _, item := range items.Elements {
// 		data = append(data, item.Inspect())
// 	}

// 	list := widget.NewList(
// 		func() int { return len(data) },
// 		func() fyne.CanvasObject { return widget.NewLabel("") },
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 			o.(*widget.Label).SetText(data[i])
// 		},
// 	)
// 	layoutContainer.Add(list)
// 	return &object.Boolean{Value: true}
// }

// func newCheckbox(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "newCheckbox requires two arguments: the text and a callback"}
// 	}

// 	checkboxText := args[0].Inspect()
// 	callbackFn, ok := args[1].(*object.Function)
// 	if !ok {
// 		return &object.Error{Message: "second argument must be a function"}
// 	}

// 	checkbox := widget.NewCheck(checkboxText, func(checked bool) {
// 		callbackFn.Call([]object.Object{&object.Boolean{Value: checked}}, map[string]object.Object{})
// 	})
// 	layoutContainer.Add(checkbox)
// 	return &object.Boolean{Value: true}
// }

// func setWindowSize(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "setWindowSize requires two arguments: width and height"}
// 	}

// 	width, ok1 := args[0].(*object.Integer)
// 	height, ok2 := args[1].(*object.Integer)
// 	if !ok1 || !ok2 {
// 		return &object.Error{Message: "Both arguments must be numbers"}
// 	}

// 	mainWindow.Resize(fyne.NewSize(float32(width.Value), float32(height.Value)))
// 	return &object.Boolean{Value: true}
// }

// func newImage(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "newImage requires one argument: the image file path"}
// 	}

// 	imagePath := args[0].Inspect()
// 	image := widget.NewImageFromFile(imagePath)
// 	layoutContainer.Add(image)
// 	return &object.Boolean{Value: true}
// }

// func addToLayout(args []object.Object, defs map[string]object.Object) object.Object {
// 	// Custom logic to allow more advanced layout setups could be added here.
// 	return &object.Boolean{Value: true}
// }

// func setTitle(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "setTitle requires one argument: the title"}
// 	}

// 	title := args[0].Inspect()
// 	mainWindow.SetTitle(title)
// 	return &object.Boolean{Value: true}
// }
// func runApp(args []object.Object, defs map[string]object.Object) object.Object {
// 	if appInstance == nil || mainWindow == nil {
// 		return &object.Error{Message: "App has not been created. Call createApp first."}
// 	}
// 	mainWindow.ShowAndRun()
// 	return &object.Boolean{Value: true}
// }
