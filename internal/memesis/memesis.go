package memesis


type MemesisCoreService struct {
}

func newMemesisCoreService() *MemesisCoreService {
    return &MemesisCoreService{}
}

func (m *MemesisCoreService) CreateMeme(imageUrl string) (filePath string) {
    filePath = "some/path"
    return
}
