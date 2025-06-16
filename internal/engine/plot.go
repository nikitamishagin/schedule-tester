package engine

import (
	"fmt"
	"image/color"
	"os"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// Color palette for algorithm lines
var algorithmColors = []color.RGBA{
	{255, 0, 0, 255},    // Red
	{0, 128, 0, 255},    // Green
	{255, 140, 0, 255},  // Orange
	{128, 0, 128, 255},  // Purple
	{64, 224, 208, 255}, // Turquoise
	{139, 69, 19, 255},  // Brown
	{255, 20, 147, 255}, // Pink
}

// PlotMultiRowLoads draws a grid of plots: for each test, the first row is the naive load, following rows are each algorithm version
func PlotMultiRowLoads(tests []TestData, filename string) error {
	if len(tests) == 0 {
		return fmt.Errorf("no test data provided")
	}

	// Get all version names by scanning the test data (assuming same set for all tests)
	versionSet := make(map[string]struct{})
	for i := range tests {
		for version := range tests[i].Loads {
			versionSet[version] = struct{}{}
		}
	}
	var versions []string
	for ver := range versionSet {
		versions = append(versions, ver)
	}
	sort.Strings(versions)

	// First row is always naive, rest are for algorithm versions
	numRows := 1 + len(versions)
	numCols := len(tests)

	// Define plot sizes
	heightPerPlot := 4 * vg.Inch
	widthPerPlot := 10 * vg.Inch // Wide format for easy reading
	totalWidth := widthPerPlot * vg.Length(numCols)
	totalHeight := heightPerPlot * vg.Length(numRows)

	// Create a grid of plots
	plots := make([][]*plot.Plot, numRows)
	for r := range plots {
		plots[r] = make([]*plot.Plot, numCols)
		for c := range plots[r] {
			plots[r][c] = plot.New()
		}
	}

	// Fill in the plots
	for c := range tests {
		// Top row: Naive load
		p := plots[0][c]
		naivePoints := make(plotter.XYs, len(tests[c].NaiveLoad))
		for j := range tests[c].NaiveLoad {
			naivePoints[j].X = float64(j)
			naivePoints[j].Y = float64(tests[c].NaiveLoad[j])
		}

		p.Title.Text = fmt.Sprintf("%s\nNaive Load (No Balancing)", tests[c].Title)
		p.X.Label.Text = "Time"
		p.Y.Label.Text = "Load"

		naiveLine, err := plotter.NewLine(naivePoints)
		if err != nil {
			return err
		}
		naiveLine.Color = color.RGBA{0, 0, 255, 255} // Blue
		p.Add(naiveLine)
		p.Legend.Add("Naive", naiveLine)

		// Each following row: one for each algorithm version
		for vIdx, version := range versions {
			lp := plots[vIdx+1][c]
			loadPoints := tests[c].Loads[version]
			pts := make(plotter.XYs, len(loadPoints))
			for j := range loadPoints {
				pts[j].X = float64(j)
				pts[j].Y = float64(loadPoints[j])
			}

			lp.Title.Text = fmt.Sprintf("%s\n%s Load", tests[c].Title, version)
			if duration, ok := tests[c].Durations[version]; ok {
				lp.Title.Text += fmt.Sprintf(" (Duration: %s)", duration.String())
			}

			lp.X.Label.Text = "Time"
			lp.Y.Label.Text = "Load"
			algoLine, err := plotter.NewLine(pts)
			if err != nil {
				return err
			}

			algoLine.Color = algorithmColors[vIdx%len(algorithmColors)]
			lp.Add(algoLine)
			lp.Legend.Add(version, algoLine)
		}
	}

	// Create the image with higher resolution (300 DPI)
	img := vgimg.NewWith(
		vgimg.UseWH(totalWidth, totalHeight),
		vgimg.UseDPI(300), // Higher DPI for better resolution
	)
	dc := draw.New(img)

	// Draw each plot in its position in the grid
	for r := range plots {
		for c := range plots[r] {
			canvas := draw.Canvas{
				Canvas: dc,
				Rectangle: vg.Rectangle{
					Min: vg.Point{X: vg.Length(c) * widthPerPlot, Y: vg.Length(numRows-r-1) * heightPerPlot},
					Max: vg.Point{X: vg.Length(c+1) * widthPerPlot, Y: vg.Length(numRows-r) * heightPerPlot},
				},
			}
			plots[r][c].Draw(canvas)
		}
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = vgimg.PngCanvas{Canvas: img}.WriteTo(w)
	return err
}
