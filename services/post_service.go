package services

import (
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"gosimplecms/configs"
	"gosimplecms/models"
	"gosimplecms/repositories"
	"gosimplecms/utils/helper"
)

type PostService interface {
	GetPosts(limit, offset int) ([]models.Post, int64, error)
	GetActivePosts(limit, offset int, orderClause string) ([]models.Post, int64, error)
	FindBySlug(string) (*models.Post, error)
	FindByID(uint) (*models.Post, error)
	Create(models.CreatePostRequest) (*models.Post, error)
	Update(id uint, request models.UpdatePostRequest) (*models.Post, error)
	FindCategoriesByIDs([]uint) ([]models.Category, error)
	FindTagsByIDs([]uint) ([]models.Tag, error)
	GetTagRelationshipScores() ([]models.TagRelationship, error)
	CalculateTagRelationshipScore(tagIDs []uint) ([]models.TagScorePost, float64, error)
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

func (p postService) GetActivePosts(limit, offset int, orderClause string) ([]models.Post, int64, error) {
	return p.postRepository.GetActivePosts(limit, offset, orderClause)
}

func (p postService) GetTagRelationshipScores() ([]models.TagRelationship, error) {
	return p.tagRepository.GetTagScores()
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
		UserID:  req.UserID,
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
		Title:         req.Title,
		Content:       cleanContent,
		PostID:        post.ID,
		VersionNumber: p.postRepository.GenerateSequentialNumber(post.ID),
	}

	_, err = p.postRepository.CreateVersionTx(tx, &postVersion)
	if err != nil {
		//tx.Rollback()
		return nil, errors.New("failed to create post version")
	}

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

func (p postService) Update(id uint, req models.UpdatePostRequest) (*models.Post, error) {
	//tx := configs.DB.Begin()
	tx := configs.ConnectDatabase()

	bm := bluemonday.UGCPolicy()
	cleanContent := bm.Sanitize(req.Content)

	post, err := p.postRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	needCreateVersion := p.needCreateVersion(post, req)
	needUpdateCurrentVersionNumber := post.Status != models.PostStatusPublished && req.Status == models.PostStatusPublished

	post.Title = req.Title
	post.Content = cleanContent
	if req.Status != "" {
		post.Status = req.Status
	}

	if _, err := p.postRepository.UpdateTx(tx, post); err != nil {
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

	// get updated post
	post, err = p.postRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if needCreateVersion {
		// 3. Create PostVersion
		var postVersion models.PostVersion
		postVersion = models.PostVersion{
			Title:         post.Title,
			Content:       post.Content,
			PostID:        post.ID,
			VersionNumber: p.postRepository.GenerateSequentialNumber(post.ID),
			Status:        post.Status,
		}

		_, err = p.postRepository.CreateVersionTx(tx, &postVersion)
		if err != nil {
			//tx.Rollback()
			return nil, errors.New("failed to create post version")
		}

		if needUpdateCurrentVersionNumber {
			err = p.postRepository.UpdateVersion(post.ID, postVersion.VersionNumber)
			if err != nil {
				return nil, errors.New("failed to update version")
			}
		}

		if err := p.postRepository.DB().Model(&postVersion).Association("Categories").Replace(categories); err != nil {
			//tx.Rollback()
			return nil, errors.New("failed to associate categories")
		}

		if err := p.postRepository.DB().Model(&postVersion).Association("Tags").Append(tags); err != nil {
			//tx.Rollback()
			return nil, errors.New("failed to associate tags")
		}
	}

	//err = tx.Commit().Error
	//if err != nil {
	//	return nil, errors.New("error committing transaction")
	//}

	return post, nil
}

func (p postService) needCreateVersion(post *models.Post, req models.UpdatePostRequest) bool {
	var catIDs, tagIDs []uint
	for _, c := range post.Categories {
		catIDs = append(catIDs, c.ID)
	}
	for _, t := range post.Tags {
		tagIDs = append(tagIDs, t.ID)
	}

	return !(post.Title == req.Title &&
		post.Content == req.Content &&
		helper.EqualUintSliceIgnoreOrder(req.CategoryIDs, catIDs) &&
		helper.EqualUintSliceIgnoreOrder(req.TagIDs, tagIDs))
}

func (p postService) FindBySlug(slug string) (*models.Post, error) {
	return p.postRepository.FindBySlug(slug)
}

func (p postService) FindByID(userID uint) (*models.Post, error) {
	return p.postRepository.FindByID(userID)
}

func (p postService) FindCategoriesByIDs(categoryIDs []uint) ([]models.Category, error) {
	return p.categoryRepository.FindByIDs(categoryIDs)
}

func (p postService) FindTagsByIDs(tagIDs []uint) ([]models.Tag, error) {
	return p.tagRepository.FindByIDs(tagIDs)
}

func (p postService) CalculateTagRelationshipScore(tagIDs []uint) ([]models.TagScorePost, float64, error) {
	return p.tagRepository.CalculateTagRelationshipScore(tagIDs)
}
