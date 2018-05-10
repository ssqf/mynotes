package main

import (
	"bytes"
	"fmt"
	"os"

	chart "github.com/wcharczuk/go-chart"
)

func main() {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.2, 3.0, 5.5},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Println("图片失败:", err)
	}
	f, err := os.OpenFile("./chart.png", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open ./chart.png error: % v ", err)
		return
	}
	buffer.WriteTo(f)
}
