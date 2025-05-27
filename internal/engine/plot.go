package engine

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotLoad(load []int, filename string) error {
	points := make(plotter.XYs, len(load))
	for i := range load {
		points[i].X = float64(i)
		points[i].Y = float64(load[i])
	}

	p := plot.New()
	p.Title.Text = "Load over Time"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Load"

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	p.Add(line)

	return p.Save(12*vg.Inch, 4*vg.Inch, filename)
}
