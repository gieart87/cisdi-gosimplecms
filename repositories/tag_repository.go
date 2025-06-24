package repositories

import (
	"gosimplecms/configs"
	"gosimplecms/models"
)

type TagRepository interface {
	FirstOrCreate(tag models.Tag) (*models.Tag, error)
	FindByIDs(ids []uint) ([]models.Tag, error)
	GetTagScores() ([]models.TagRelationship, error)
}

type tagRepository struct{}

func (t tagRepository) FindByIDs(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag

	err := configs.DB.Where("id IN (?)", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func NewTagRepository() TagRepository {
	return &tagRepository{}
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

//
//func (t tagRepository) GetTagScores([]models.TagScore, error) {
//	var tagScores []models.TagScore
//
//	sqlStr := `SELECT
//	LEAST(pt1.tag_id, pt2.tag_id) AS tag1_id,
//		GREATEST(pt1.tag_id, pt2.tag_id) AS tag2_id,
//		COUNT(DISTINCT pt1.post_id) AS co_occurrence_count
//	FROM post_tags pt1
//	JOIN post_tags pt2
//	ON pt1.post_id = pt2.post_id
//	AND pt1.tag_id < pt2.tag_id
//	GROUP BY tag1_id, tag2_id;`
//
//	err := configs.DB.Exec(sqlStr)
//	if err != nil {
//		return nil, err
//	}
//}
