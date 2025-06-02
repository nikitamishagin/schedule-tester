package engine

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func PlotDoubleLoad(tests []TestData, filename string) error {
	// Define grid dimensions
	numCols := len(tests)
	numRows := 2 // One row for naive load, one for algorithm load

	// Define plot sizes
	heightPerPlot := 4 * vg.Inch
	widthPerPlot := 10 * vg.Inch // Increased from 6 to 10 inches to stretch horizontally
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
		// Naive load plot (top row)
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

		// Algorithm load plot (bottom row)
		p = plots[1][c]
		algoPoints := make(plotter.XYs, len(tests[c].Load))
		for j := range tests[c].Load {
			algoPoints[j].X = float64(j)
			algoPoints[j].Y = float64(tests[c].Load[j])
		}

		p.Title.Text = fmt.Sprintf("%s\nAlgorithm v1 Load\nDuration: %s", tests[c].Title, tests[c].Duration.String())
		p.X.Label.Text = "Time"
		p.Y.Label.Text = "Load"

		algoLine, err := plotter.NewLine(algoPoints)
		if err != nil {
			return err
		}
		algoLine.Color = color.RGBA{255, 0, 0, 255} // Red
		p.Add(algoLine)
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
