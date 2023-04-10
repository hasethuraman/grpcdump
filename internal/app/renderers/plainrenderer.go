package renderers

import (
	"fmt"
	"time"

	"github.com/rmedvedev/grpcdump/internal/app/models"
)

// PlainRenderer ...
type PlainRenderer struct{}

// Render renders model
func (PlainRenderer) Render(model models.RenderModel) string {
	return fmt.Sprintf(
		"%s %s:%s -> %s:%s %s %+v %s",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		model.GetSrcHost(),
		model.GetSrcPort(),
		model.GetDstHost(),
		model.GetDstPort(),
		model.GetPath(),
		model.GetBody(),
		model.GetHeaders(),
	)
}
