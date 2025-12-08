package profile

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const (
	placeholderStart   = "<!-- XP_CALC_START -->"
	placeholderEnd     = "<!-- XP_CALC_END -->"
	profileDescription = "%s\nI'm a seasoned Software Engineer with over %d years of experience in the IT world, passionate about technology and creating efficient, scalable solutions.\n%s"
)

// UpdateReadme reads the readme to update the time of experience
func UpdateReadme(file string, startAt int) error {
	experience := calculateExperience(startAt)

	if experience <= 0 {
		return fmt.Errorf("year calculation should not be 0 or less: %d years (career start: %d)", experience, startAt)
	}

	content, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", file, err)
	}

	fileContent := string(content)

	pattern := regexp.MustCompile(fmt.Sprintf("(?s)%s.*?%s", regexp.QuoteMeta(placeholderStart), regexp.QuoteMeta(placeholderEnd)))

	if !pattern.MatchString(fileContent) {
		return fmt.Errorf("error matching string for file %s", file)
	}

	newContent := fmt.Sprintf(profileDescription, placeholderStart, experience, placeholderEnd)

	updatedContent := pattern.ReplaceAllString(fileContent, newContent)

	if fileContent == updatedContent {
		fmt.Printf("no update needed. Current: %d years experience\n", experience)
		return nil
	}

	err = os.WriteFile(file, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("could not write the file %s: %v", file, err)
	}

	fmt.Printf("Successfully updated: %d years experience (since %d)\n", experience, startAt)

	return nil
}

func calculateExperience(startAt int) int {
	currentYear := time.Now().Year()
	return currentYear - startAt
}
