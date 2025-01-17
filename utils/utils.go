package utils

import (
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// RemoveFrontmatter removes the front matter header of a markdown file.
func RemoveFrontmatter(content []byte) []byte {
	if frontmatterBoundaries := detectFrontmatter(content); frontmatterBoundaries[0] == 0 {
		return content[frontmatterBoundaries[1]:]
	}
	return content
}

var yamlPattern = regexp.MustCompile(`(?m)^---\r?\n(\s*\r?\n)?`)

func detectFrontmatter(c []byte) []int {
	if matches := yamlPattern.FindAllIndex(c, 2); len(matches) > 1 {
		return []int{matches[0][0], matches[1][1]}
	}
	return []int{-1, -1}
}

// Expands tilde and all environment variables from the given path.
func ExpandPath(path string) string {
	s, err := homedir.Expand(path)
	if err == nil {
		return os.ExpandEnv(s)
	}
	return os.ExpandEnv(path)
}

// Returns a slice containing the pager location/command and its args.
// Returns the pager location from the environment variable if present, else uses `less -r`.
func GetPagerCommand(varName string) []string {
	pagerCmd := os.Getenv(varName)
	pa := []string{pagerCmd}
	if pagerCmd == "" {
		pagerCmd = "less -r"
		pa = strings.Split(pagerCmd, " ")
	}
	return pa
}
