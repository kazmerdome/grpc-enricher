package post

type postService struct {
}

func NewPostService() *postService {
	return &postService{}
}

// Mock Data
func (r *postService) ListPosts() ([]Post, error) {
	return PostData, nil
}
