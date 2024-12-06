package module

import (
	"bytes"
	// "image/png"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/ekilie/vint-lang/object"
)

var VintChartFunctions = map[string]object.ModuleFunction{}

func init() {
	VintChartFunctions["barChart"] = barChart
	VintChartFunctions["pieChart"] = pieChart
	VintChartFunctions["lineGraph"] = lineGraph
}

// Create a bar chart
func barChart(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "barChart requires three arguments: labels, values, and output file"}
	}

	labels := objectArrayToStringSlice(args[0].(*object.Array).Elements)
	values := objectArrayToFloatSlice(args[1].(*object.Array).Elements)
	output := args[2].Inspect()

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Bar Chart"}))
	bar.SetXAxis(labels).AddSeries("Values", generateBarData(values))

	buf := new(bytes.Buffer)
	bar.Render(buf)
	return saveChartToFile(output, buf.Bytes())
}

// Create a pie chart
func pieChart(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "pieChart requires three arguments: labels, values, and output file"}
	}

	labels := objectArrayToStringSlice(args[0].(*object.Array).Elements)
	values := objectArrayToFloatSlice(args[1].(*object.Array).Elements)
	output := args[2].Inspect()

	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Pie Chart"}))
	pie.AddSeries("Values", generatePieData(labels, values))

	buf := new(bytes.Buffer)
	pie.Render(buf)
	return saveChartToFile(output, buf.Bytes())
}

// Create a line graph
func lineGraph(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "lineGraph requires three arguments: labels, values, and output file"}
	}

	labels := objectArrayToStringSlice(args[0].(*object.Array).Elements)
	values := objectArrayToFloatSlice(args[1].(*object.Array).Elements)
	output := args[2].Inspect()

	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Line Graph"}))
	line.SetXAxis(labels).AddSeries("Values", generateLineData(values))

	buf := new(bytes.Buffer)
	line.Render(buf)
	return saveChartToFile(output, buf.Bytes())
}

// Helper functions
func generateBarData(values []float64) []opts.BarData {
	data := []opts.BarData{}
	for _, v := range values {
		data = append(data, opts.BarData{Value: v})
	}
	return data
}

func generatePieData(labels []string, values []float64) []opts.PieData {
	data := []opts.PieData{}
	for i, v := range values {
		data = append(data, opts.PieData{Name: labels[i], Value: v})
	}
	return data
}

func generateLineData(values []float64) []opts.LineData {
	data := []opts.LineData{}
	for _, v := range values {
		data = append(data, opts.LineData{Value: v})
	}
	return data
}

func saveChartToFile(filename string, data []byte) object.Object {
	file, err := os.Create(filename)
	if err != nil {
		return &object.Error{Message: "Failed to create file: " + err.Error()}
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return &object.Error{Message: "Failed to write file: " + err.Error()}
	}
	return &object.Boolean{Value: true}
}

func objectArrayToStringSlice(objects []object.Object) []string {
	result := []string{}
	for _, obj := range objects {
		result = append(result, obj.Inspect())
	}
	return result
}

func objectArrayToFloatSlice(objects []object.Object) []float64 {
	result := []float64{}
	for _, obj := range objects {
		if num, ok := obj.(*object.Integer); ok {
			result = append(result, float64(num.Value))
		}
	}
	return result
}
