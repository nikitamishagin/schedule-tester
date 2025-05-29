package engine

import (
	"fmt"
	"os"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type LoadPlot struct {
	Load     []int
	Duration time.Duration
}

func PlotDoubleLoad(naiveLoad, scheduledLoad LoadPlot, filename string) error {
	heightPerPlot := 4 * vg.Inch
	width := 12 * vg.Inch
	totalHeight := 2 * heightPerPlot

	naivePoints := make(plotter.XYs, len(naiveLoad.Load))
	for i := range naiveLoad.Load {
		naivePoints[i].X = float64(i)
		naivePoints[i].Y = float64(naiveLoad.Load[i])
	}

	naivePlot := plot.New()
	naivePlot.Title.Text = fmt.Sprintf("Naive Load (without balancing)\nDuration of computing: %s", naiveLoad.Duration.String())
	naivePlot.X.Label.Text = "Time"
	naivePlot.Y.Label.Text = "Load"
	naiveLine, err := plotter.NewLine(naivePoints)
	if err != nil {
		return err
	}
	naivePlot.Add(naiveLine)

	balancePoints := make(plotter.XYs, len(scheduledLoad.Load))
	for i := range scheduledLoad.Load {
		balancePoints[i].X = float64(i)
		balancePoints[i].Y = float64(scheduledLoad.Load[i])
	}
	balancePlot := plot.New()
	balancePlot.Title.Text = fmt.Sprintf("Load with scheduling\nDuration of computing: %s", scheduledLoad.Duration.String())
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
