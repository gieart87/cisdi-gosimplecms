package seeds

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gosimplecms/models"
	"log"
	"math"
)

func SeedTagRelationships(db *gorm.DB) error {
	// 1. Total post
	var totalPosts int64
	if err := db.Table("post_tags").Distinct("post_id").Count(&totalPosts).Error; err != nil {
		return err
	}

	if totalPosts == 0 {
		log.Println("No posts found. Skipping tag relationship scoring.")
		return nil
	}

	// 2. Ambil count tag -> map[tag_id]count
	var tagCounts []struct {
		TagID uint
		Count int64
	}
	if err := db.
		Table("post_tags").
		Select("tag_id, COUNT(DISTINCT post_id) AS count").
		Group("tag_id").
		Scan(&tagCounts).Error; err != nil {
		return err
	}

	tagCountMap := make(map[uint]int64)
	for _, t := range tagCounts {
		tagCountMap[t.TagID] = t.Count
	}

	// 3. Hitung co-occurrence antar pasangan tag
	var coOccurrences []struct {
		Tag1ID uint
		Tag2ID uint
		Count  int64
	}
	query := `
		SELECT
			LEAST(pt1.tag_id, pt2.tag_id) AS tag1_id,
			GREATEST(pt1.tag_id, pt2.tag_id) AS tag2_id,
			COUNT(DISTINCT pt1.post_id) AS count
		FROM post_tags pt1
		JOIN post_tags pt2
			ON pt1.post_id = pt2.post_id
			AND pt1.tag_id < pt2.tag_id
		GROUP BY tag1_id, tag2_id;
	`
	if err := db.Raw(query).Scan(&coOccurrences).Error; err != nil {
		return err
	}

	// 4. Insert/update score
	for _, co := range coOccurrences {
		countA := tagCountMap[co.Tag1ID]
		countB := tagCountMap[co.Tag2ID]
		if countA == 0 || countB == 0 || co.Count == 0 {
			continue
		}

		numerator := float64(co.Count) * float64(totalPosts)
		denominator := float64(countA) * float64(countB)
		score := math.Log2(numerator / denominator)

		relationship := models.TagRelationship{
			Tag1ID:            co.Tag1ID,
			Tag2ID:            co.Tag2ID,
			Score:             score,
			CoOccurrenceCount: int(co.Count),
		}

		if err := db.Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).Create(&relationship).Error; err != nil {
			return fmt.Errorf("failed inserting relationship for tags %d & %d: %w", co.Tag1ID, co.Tag2ID, err)
		}
	}

	return nil
}
