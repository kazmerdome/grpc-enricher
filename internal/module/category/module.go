package category

import "github.com/kazmerdome/grpc-enricher/internal/module/tag"

type categoryModule struct {
	categoryEnricher CategoryEnricher
}

func NewCategoryModule(tagEnricher tag.TagEnricher) *categoryModule {
	categoryDataloader := NewCategoryDataloader()
	categoryEnricher := NewCategoryEnricher(tagEnricher, categoryDataloader)
	return &categoryModule{categoryEnricher: categoryEnricher}
}

func (m *categoryModule) GetCategoryEnricher() CategoryEnricher {
	return m.categoryEnricher
}
