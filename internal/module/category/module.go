package category

type categoryModule struct {
	categoryEnricher CategoryEnricher
}

func NewCategoryModule() *categoryModule {
	categoryDataloader := NewCategoryDataloader()
	categoryEnricher := NewCategoryEnricher(categoryDataloader)
	return &categoryModule{categoryEnricher: categoryEnricher}
}

func (m *categoryModule) GetCategoryEnricher() CategoryEnricher {
	return m.categoryEnricher
}
