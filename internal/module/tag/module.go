package tag

type tagModule struct {
	tagEnricher TagEnricher
}

func NewTagModule() *tagModule {
	tagDataloader := NewTagDataloader()
	tagEnricher := NewTagEnricher(tagDataloader)
	return &tagModule{tagEnricher: tagEnricher}
}

func (m *tagModule) GetTagEnricher() TagEnricher {
	return m.tagEnricher
}
