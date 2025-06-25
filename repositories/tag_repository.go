package repositories

import (
	"errors"
	"gorm.io/gorm"
	"gosimplecms/configs"
	"gosimplecms/models"
)

type TagRepository interface {
	FirstOrCreate(tag models.Tag) (*models.Tag, error)
	Create(tag models.Tag) (*models.Tag, error)
	FindByIDs(ids []uint) ([]models.Tag, error)
	GetTagScores() ([]models.TagRelationship, error)
	GetTags() ([]models.Tag, error)
	CalculateTagRelationshipScore(tagIDs []uint) ([]models.TagScorePost, float64, error)
}

type tagRepository struct{}

func (t tagRepository) GetTags() ([]models.Tag, error) {
	var tags []models.Tag

	err := configs.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func NewTagRepository() TagRepository {
	return &tagRepository{}
}

func (t tagRepository) Create(tag models.Tag) (*models.Tag, error) {
	err := configs.DB.Create(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t tagRepository) FindByIDs(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag

	err := configs.DB.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t tagRepository) FirstOrCreate(req models.Tag) (*models.Tag, error) {
	var tag models.Tag

	err := configs.DB.FirstOrCreate(&tag, req).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t tagRepository) GetTagScores() ([]models.TagRelationship, error) {
	var scores []models.TagRelationship

	err := configs.DB.Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

// CalculateTagRelationshipScore menghitung skor total dari pasangan tag yang ada pada satu post
func (t tagRepository) CalculateTagRelationshipScore(tagIDs []uint) ([]models.TagScorePost, float64, error) {
	if len(tagIDs) < 2 {
		return nil, 0, nil
	}

	var totalScore float64
	var scores []models.TagScorePost

	for i := 0; i < len(tagIDs); i++ {
		for j := i + 1; j < len(tagIDs); j++ {
			tagA := tagIDs[i]
			tagB := tagIDs[j]

			var rel models.TagRelationship
			err := configs.DB.
				Where("(tag1_id = ? AND tag2_id = ?) OR (tag1_id = ? AND tag2_id = ?)",
					tagA, tagB, tagB, tagA).
				First(&rel).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0, err
			}

			score := rel.Score
			totalScore += score

			scores = append(scores, models.TagScorePost{
				Tag1ID: tagA,
				Tag2ID: tagB,
				Score:  score,
			})
		}
	}

	return scores, totalScore, nil
}
