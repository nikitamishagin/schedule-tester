package engine

import (
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func PlotDoubleLoad(naiveLoad, balancingLoad []int, filename string) error {
	heightPerPlot := 4 * vg.Inch
	width := 12 * vg.Inch
	totalHeight := 2 * heightPerPlot

	naivePoints := make(plotter.XYs, len(naiveLoad))
	for i := range naiveLoad {
		naivePoints[i].X = float64(i)
		naivePoints[i].Y = float64(naiveLoad[i])
	}

	naivePlot := plot.New()
	naivePlot.Title.Text = "Naive Load (without balancing)"
	naivePlot.X.Label.Text = "Time"
	naivePlot.Y.Label.Text = "Load"
	naiveLine, err := plotter.NewLine(naivePoints)
	if err != nil {
		return err
	}
	naivePlot.Add(naiveLine)

	balancePoints := make(plotter.XYs, len(balancingLoad))
	for i := range balancingLoad {
		balancePoints[i].X = float64(i)
		balancePoints[i].Y = float64(balancingLoad[i])
	}
	balancePlot := plot.New()
	balancePlot.Title.Text = "Balanced Load"
	balancePlot.X.Label.Text = "Time"
	balancePlot.Y.Label.Text = "Load"
	balanceLine, err := plotter.NewLine(balancePoints)
	if err != nil {
		return err
	}
	balancePlot.Add(balanceLine)

	img := vgimg.New(width, totalHeight)
	dc := draw.New(img)

	top := draw.Canvas{
		Canvas: dc,
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: 0, Y: heightPerPlot},
			Max: vg.Point{X: width, Y: totalHeight},
		},
	}
	bottom := draw.Canvas{
		Canvas: dc,
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: 0, Y: 0},
			Max: vg.Point{X: width, Y: heightPerPlot},
		},
	}

	naivePlot.Draw(top)
	balancePlot.Draw(bottom)

	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = vgimg.PngCanvas{Canvas: img}.WriteTo(w)
	return err
}
