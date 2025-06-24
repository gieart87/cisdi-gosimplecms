package services

import (
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"gosimplecms/configs"
	"gosimplecms/models"
	"gosimplecms/repositories"
)

type PostService interface {
	GetPosts(limit, offset int) ([]models.Post, int64, error)
	FindBySlug(string) (*models.Post, error)
	Create(models.CreatePostRequest) (*models.Post, error)
	FindCategoriesByIDs([]uint) ([]models.Category, error)
	FindTagsByIDs([]uint) ([]models.Tag, error)
}

type postService struct {
	postRepository     repositories.PostRepository
	tagRepository      repositories.TagRepository
	categoryRepository repositories.CategoryRepository
}

func NewPostService(postRepository repositories.PostRepository, tagRepository repositories.TagRepository, categoryRepository repositories.CategoryRepository) PostService {
	return &postService{
		postRepository:     postRepository,
		tagRepository:      tagRepository,
		categoryRepository: categoryRepository,
	}
}

func (p postService) GetPosts(limit, offset int) ([]models.Post, int64, error) {
	return p.postRepository.GetPosts(limit, offset)
}

func (p postService) Create(req models.CreatePostRequest) (*models.Post, error) {

	//tx := configs.DB.Begin()
	tx := configs.ConnectDatabase()

	bm := bluemonday.UGCPolicy()
	cleanContent := bm.Sanitize(req.Content)

	var post models.Post
	post = models.Post{
		Title:   req.Title,
		Content: cleanContent,
	}
	if _, err := p.postRepository.CreateTx(tx, &post); err != nil {
		//tx.Rollback()
		return nil, err
	}

	categories, err := p.categoryRepository.FindByIDs(req.CategoryIDs)
	if err != nil {
		//tx.Rollback()
		return nil, err
	}

	if err = p.postRepository.DB().Model(&post).Association("Categories").Replace(categories); err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to associate categories")
	}

	tags, err := p.tagRepository.FindByIDs(req.CategoryIDs)
	if err != nil {
		//tx.Rollback()
		return nil, err
	}
	if err = p.postRepository.DB().Model(&post).Association("Tags").Append(tags); err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to associate tags")
	}

	// 3. Create PostVersion
	var postVersion models.PostVersion
	postVersion = models.PostVersion{
		Title:   req.Title,
		Content: cleanContent,
		PostID:  post.ID,
	}

	_, err = p.postRepository.CreateVersionTx(tx, &postVersion)
	if err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to create post version")
	}

	post.CurrentVersionID = postVersion.ID

	if err := p.postRepository.DB().Model(&postVersion).Association("Categories").Replace(categories); err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to associate categories")
	}

	if err := p.postRepository.DB().Model(&postVersion).Association("Tags").Append(tags); err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to associate tags")
	}

	//err = tx.Commit().Error
	//if err != nil {
	//	return nil, errors.New("error committing transaction")
	//}

	return &post, nil
}

func (p postService) FindBySlug(slug string) (*models.Post, error) {
	return p.postRepository.FindBySlug(slug)
}

func (p postService) FindCategoriesByIDs(categoryIDs []uint) ([]models.Category, error) {
	return p.categoryRepository.FindByIDs(categoryIDs)
}

func (p postService) FindTagsByIDs(tagIDs []uint) ([]models.Tag, error) {
	return p.tagRepository.FindByIDs(tagIDs)
}
