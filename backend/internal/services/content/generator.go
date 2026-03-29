package content

import "opflow-backend/internal/models"

type Generator struct {
	apiKey string
}

func NewGenerator(apiKey string) *Generator {
	return &Generator{apiKey: apiKey}
}

func (g *Generator) Generate(hotspot models.Hotspot, platform string) (string, error) {
	// TODO: Cline implement here
	return "Generated: " + hotspot.Title + " for " + platform, nil
}