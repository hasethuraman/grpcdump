package renderers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rmedvedev/grpcdump/internal/app/models"
	"github.com/sirupsen/logrus"
)

// JSONRenderer ...
type JSONRenderer struct{}

type jsonView struct {
	Time    string            `json:"time"`
	Src     string            `json:"src"`
	Dst     string            `json:"dst"`
	Path    string            `json:"path"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

// Render renders model
func (JSONRenderer) Render(model models.RenderModel) string {
	bytes, err := json.Marshal(jsonView{
		Time:    time.Now().Format("2006-01-02 15:04:05.000000"),
		Src:     model.GetSrcHost() + ":" + model.GetSrcPort(),
		Dst:     model.GetDstHost() + ":" + model.GetDstPort(),
		Path:    model.GetPath(),
		Body:    fmt.Sprintf("%v", model.GetBody()),
		Headers: model.GetHeaders(),
	})

	if err != nil {
		logrus.Errorf("Error to marshal model: %s", err.Error())
	}

	return string(bytes)
}
