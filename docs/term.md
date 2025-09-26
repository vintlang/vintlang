
# Terminal UI (`term`) Module

The `term` module provides a rich set of tools for building beautiful and interactive terminal user interfaces (TUIs). It is built on top of the powerful `lipgloss` library and offers a wide range of components, from simple text styling to complex layouts and widgets.

## Basic I/O

### `print(message, [color])`

Prints a message to the terminal.

- `message` (string): The text to print.
- `color` (string, optional): The color to apply to the text.

### `println(message, [color])`

Prints a message to the terminal, followed by a newline.

- `message` (string): The text to print.
- `color` (string, optional): The color to apply to the text.

### `input(prompt)`

Prompts the user for input and returns the entered text.

- `prompt` (string): The message to display to the user.

### `password(prompt)`

Prompts the user for a password without displaying the entered characters.

- `prompt` (string): The message to display to the user.

### `confirm(prompt)`

Asks a yes/no question and returns `true` or `false`.

- `prompt` (string): The question to ask the user.

## Styling & Formatting

### `style(text, options)`

Applies a set of styles to a string.

- `text` (string): The text to style.
- `options` (dict): A dictionary of style options, such as `color`, `background`, `bold`, `italic`, and `underline`.

### `banner(text)`

Creates a large, stylized banner.

- `text` (string): The text to display in the banner.

### `box(text)`

Draws a box around a piece of text.

- `text` (string): The content of the box.

### `badge(text)`

Creates a small, colored badge with text.

- `text` (string): The text for the badge.

### `avatar(text)`

Creates a circular avatar with the first letter of the given text.

- `text` (string): The text to create the avatar from.

## Interactive Components

### `select(options)`

Displays a list of options and allows the user to select one.

- `options` (array): An array of strings representing the choices.

### `checkbox(options)`

Displays a list of options and allows the user to select multiple.

- `options` (array): An array of strings representing the choices.

### `radio(options)`

Displays a list of options and allows the user to select one (similar to `select`).

- `options` (array): An array of strings representing the choices.

### `menu(items)`

Creates a numbered menu from a list of items.

- `items` (array): An array of strings for the menu.

### `form(config)`

Creates a simple form with labeled fields.

- `config` (dict): A dictionary where keys are the field labels.

### `wizard(steps)`

Creates a multi-step wizard or form.

- `steps` (array): An array of strings representing the steps.

## Notifications & Alerts

### `alert(message)`

Displays a prominent alert message.

- `message` (string): The alert message.

### `notify(message)`

Shows a notification message.

- `message` (string): The notification text.

### `error(message)`

Displays a formatted error message.

- `message` (string): The error text.

### `success(message)`

Displays a formatted success message.

- `message` (string): The success text.

### `info(message)`

Displays a formatted informational message.

- `message` (string): The info text.

### `warning(message)`

Displays a formatted warning message.

- `message` (string): The warning text.

## Layouts & Widgets

### `layout(config)`

Creates a flexible layout with configurable direction, padding, and borders.

- `config` (dict): A dictionary of layout options.

### `grid(items, config)`

Creates a grid layout for a list of items.

- `items` (array): The items to display in the grid.
- `config` (dict): A dictionary with grid options, such as `columns`.

### `tabs(config)`

Creates a tabbed interface.

- `config` (dict): A dictionary where keys are the tab titles.

### `accordion(sections)`

Creates a collapsible accordion view.

- `sections` (dict): A dictionary where keys are the section titles and values are the content.

### `tree(config)`

Creates a tree view from a nested dictionary.

- `config` (dict): The nested dictionary representing the tree structure.

### `table(rows)`

Creates a formatted table from an array of rows.

- `rows` (array): An array of arrays, where each inner array is a row.

### `card(config)`

Creates a card with a title and content.

- `config` (dict): A dictionary with `title` and `content` keys.

### `list(items)`

Creates a bulleted list.

- `items` (array): An array of strings to display in the list.

### `split(left, right)`

Creates a split view with two panels.

- `left` (string): The content for the left panel.
- `right` (string): The content for the right panel.

### `modal(config)`

Creates a modal dialog.

- `config` (dict): A dictionary with `title` and `content` for the modal.

### `tooltip(text, message)`

Adds a tooltip to a piece of text.

- `text` (string): The text to add the tooltip to.
- `message` (string): The tooltip message.

## Visualizations

### `chart(data)`

Creates a simple bar chart from an array of numbers.

- `data` (array): An array of integers or floats.

### `gauge(value)`

Creates a gauge or progress indicator.

- `value` (integer): A value between 0 and 100.

### `heatmap(data)`

Creates a heatmap from an array of numbers.

- `data` (array): An array of integers or floats.

### `calendar(config)`

Displays a calendar for a given month and year.

- `config` (dict): A dictionary with `year` and `month` keys.

### `timeline(events)`

Creates a timeline from a list of events.

- `events` (array): An array of dictionaries, where each dictionary has `title` and `time` keys.

### `kanban(columns)`

Creates a Kanban board.

- `columns` (dict): A dictionary where keys are the column titles and values are arrays of tasks.

## Other Utilities

### `clear()`

Clears the terminal screen.

### `spinner(message)`

Displays a loading spinner with a message.

- `message` (string): The message to display next to the spinner.

### `progress(total)`

Initializes a progress bar.

- `total` (integer): The total value for the progress bar.

### `cursor(visible)`

Shows or hides the terminal cursor.

- `visible` (boolean): `true` to show the cursor, `false` to hide it.

### `beep()`

Plays a terminal beep sound.

### `moveCursor(x, y)`

Moves the cursor to a specific position.

- `x` (integer): The column to move to.
- `y` (integer): The row to move to.

### `getSize()`

Returns the size of the terminal as a dictionary with `width` and `height` keys.

### `countdown(duration)`

Creates a countdown timer.

- `duration` (integer): The duration of the countdown in seconds.

### `loading(message)`

Displays a loading message.

- `message` (string): The message to display.

### `dashboard(widgets)`

Creates a dashboard layout with multiple widgets.

- `widgets` (dict): A dictionary where keys are the widget titles and values are their content.
