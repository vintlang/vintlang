
# VintChart Module (Experimental)

The `vintChart` module provides functions for creating various types of charts and saving them as HTML files. This module is experimental and its API may change in the future.

## Functions

### `barChart(labels, values, outputFile)`

Creates a bar chart.

- `labels` (array): An array of strings for the x-axis labels.
- `values` (array): An array of numbers for the y-axis values.
- `outputFile` (string): The path to save the HTML file (e.g., `"bar_chart.html"`).

**Usage:**

```vint
import vintChart

let labels = ["A", "B", "C"]
let values = [10, 20, 15]
vintChart.barChart(labels, values, "my_bar_chart.html")
```

### `pieChart(labels, values, outputFile)`

Creates a pie chart.

- `labels` (array): An array of strings for the pie slices.
- `values` (array): An array of numbers for the values of the slices.
- `outputFile` (string): The path to save the HTML file (e.g., `"pie_chart.html"`).

**Usage:**

```vint
import vintChart

let labels = ["Work", "Sleep", "Play"]
let values = [8, 8, 8]
vintChart.pieChart(labels, values, "my_pie_chart.html")
```

### `lineGraph(labels, values, outputFile)`

Creates a line graph.

- `labels` (array): An array of strings for the x-axis labels.
- `values` (array): An array of numbers for the y-axis values.
- `outputFile` (string): The path to save the HTML file (e.g., `"line_graph.html"`).

**Usage:**

```vint
import vintChart

let labels = ["Jan", "Feb", "Mar"]
let values = [100, 120, 110]
vintChart.lineGraph(labels, values, "my_line_graph.html")
```
