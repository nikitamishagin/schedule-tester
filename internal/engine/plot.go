package engine

import (
	"fmt"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func PlotDoubleLoad(tests []TestData, filename string) error {
	heightPerPlot := 4 * vg.Inch
	width := 20 * vg.Inch
	totalHeight := heightPerPlot * vg.Length(len(tests))

	plots := make([]plot.Plot, len(tests))
	for i := range tests {
		p := plot.New()
		points := make(plotter.XYs, len(tests[i].Load))

		for j := range tests[i].Load {
			points[j].X = float64(j)
			points[j].Y = float64(tests[i].Load[j])
		}

		p.Title.Text = fmt.Sprintf("%s\nDuration of computing: %s", tests[i].Title, tests[i].Duration.String())
		p.X.Label.Text = "Time"
		p.Y.Label.Text = "Load"

		line, err := plotter.NewLine(points)
		if err != nil {
			return err
		}

		p.Add(line)
		plots[i] = *p
	}

	img := vgimg.New(width, totalHeight)
	dc := draw.New(img)

	canvas := make([]draw.Canvas, len(tests))
	for i := range tests {
		canvas[i] = draw.Canvas{
			Canvas: dc,
			Rectangle: vg.Rectangle{
				Min: vg.Point{X: 0, Y: vg.Length(i) * heightPerPlot},
				Max: vg.Point{X: width, Y: vg.Length(i+1) * heightPerPlot},
			},
		}
		plots[i].Draw(canvas[i])
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = vgimg.PngCanvas{Canvas: img}.WriteTo(w)
	return err
}
