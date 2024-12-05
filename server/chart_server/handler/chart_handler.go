package handler

import (
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/chart"
	"net/http"
)

// HTTP 핸들러
func ChartHandler(w http.ResponseWriter, r *http.Request) {
	chart := chart.GenerateKLineChart()
	chart.Render(w)
}
