package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gosimplecms/configs"
	"gosimplecms/models"
)

type PostRepository interface {
	FindBySlug(slug string) (*models.Post, error)
	CreateTx(tx *gorm.DB, post *models.Post) (*models.Post, error)
	CreateVersionTx(tx *gorm.DB, post *models.PostVersion) (*models.PostVersion, error)
	UpdateTx(tx *gorm.DB, post *models.Post) (*models.Post, error)
	DeleteTx(tx *gorm.DB, id uuid.UUID) error
	GetPosts(limit, offset int) ([]models.Post, int64, error)
	FindByID(id uuid.UUID) (*models.Post, error)
	FindTagsByIDs(ids []string) ([]models.Tag, error)
	FindCategoriesByIDs(ids []string) ([]models.Category, error)
	DB() *gorm.DB
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (p postRepository) CreateTx(tx *gorm.DB, post *models.Post) (*models.Post, error) {
	return post, tx.Create(post).Error
}

func (p postRepository) CreateVersionTx(tx *gorm.DB, postVersion *models.PostVersion) (*models.PostVersion, error) {
	return postVersion, tx.Create(postVersion).Error
}

func (p postRepository) UpdateTx(tx *gorm.DB, post *models.Post) (*models.Post, error) {
	return post, tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(post).Error
}

func (p postRepository) DeleteTx(tx *gorm.DB, id uuid.UUID) error {
	return tx.Delete(&models.Post{}, "id = ?", id).Error
}

func (p postRepository) GetPosts(limit, offset int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	configs.DB.Model(&models.Post{}).Count(&total)
	configs.DB.Preload("Categories").Preload("Tags").
		Limit(limit).Offset(offset).Find(&posts)

	return posts, total, nil
}

func (p postRepository) FindTagsByIDs(ids []string) ([]models.Tag, error) {
	var tags []models.Tag

	err := configs.DB.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (p postRepository) FindCategoriesByIDs(ids []string) ([]models.Category, error) {
	var categories []models.Category

	err := configs.DB.Where("id IN (?)", ids).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (p postRepository) DB() *gorm.DB {
	return p.db
}

func (p postRepository) GetAll() ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) Create(post models.Post) (*models.Post, error) {
	err := configs.DB.Create(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p postRepository) FindByID(id uuid.UUID) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) FindBySlug(slug string) (*models.Post, error) {
	var post models.Post

	err := configs.DB.Where("slug = ?", slug).Take(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p postRepository) Update(post models.Post) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}
