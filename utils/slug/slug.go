package slug

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

func GenerateSlug(input string) string {
	// lowercase
	slug := strings.ToLower(input)

	// ganti spasi dan underscore jadi tanda hubung
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	// hapus karakter non-alfanumerik dan strip
	re := regexp.MustCompile(`[^a-z0-9\-]`)
	slug = re.ReplaceAllString(slug, "")

	// hapus tanda hubung berurutan
	re2 := regexp.MustCompile(`-+`)
	slug = re2.ReplaceAllString(slug, "-")

	// trim tanda hubung di awal/akhir
	slug = strings.Trim(slug, "-")

	return slug
}

func GenerateUniqueSlug(input string) string {
	return fmt.Sprintf("%s-%s", GenerateSlug(input), uuid.NewString()[:8])
}
