package renderers

import (
	"fmt"
	"time"

	c "github.com/logrusorgru/aurora"
	"github.com/rmedvedev/grpcdump/internal/app/models"
)

// PrettyRenderer ...
type PrettyRenderer struct{}

// Render renders model
func (PrettyRenderer) Render(model models.RenderModel) string {
	return fmt.Sprintf(
		"%s %s:%s -> %s:%s %s %+v %s",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		c.Green(model.GetSrcHost()),
		c.Green(model.GetSrcPort()),
		c.Yellow(model.GetDstHost()),
		c.Yellow(model.GetDstPort()),
		c.BgBrightBlue(c.White(model.GetPath())),
		model.GetBody(),
		renderHeaders(model.GetHeaders()),
	)
}

func renderHeaders(headers map[string]string) string {
	result := "["

	for name, val := range headers {
		result += c.Cyan(name).String() + ":" + val + ", "
	}

	return result + "]"
}
