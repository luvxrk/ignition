package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// The base URL to the repository that provides the .gitignore files
const BaseURL = "https://raw.githubusercontent.com/github/gitignore/main/"

// URL to Gitignore templates provided by GitHub's API
const GitignoreAPI = "https://api.github.com/gitignore/templates"

// Capitalize modifies the provided string such that the first alphabet is upper case and the rest is lower case.
func CapitalizeString(s *string) {
	if s == nil || *s == "" {
		return
	}
	*s = strings.ToUpper((*s)[:1]) + strings.ToLower((*s)[1:])
}

// getGitIgnoreURL constructs the URL to fetch the .gitignore file.
func getGitIgnoreURL(language string) string {
	return BaseURL + language + ".gitignore"
}

// fetchGitIgnoreContents performs the HTTP GET request to retrieve the .gitignore file.
func fetchGitIgnoreContents(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch .gitignore: %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("receieved status code %d\n", resp.StatusCode)
	}
	return resp, nil
}

// saveGitIgnoreFile creates a file and writes the contents from the response body.
func saveGitIgnoreFile(filePath string, body io.Reader) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v\n", err)
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v\n", err)
	}
	return nil
}

// FetchGitIgnore downloads a .gitignore file for a specific language and saves it.
func FetchGitIgnore(language, output string) error {
	url := getGitIgnoreURL(language)
	resp, err := fetchGitIgnoreContents(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filePath := filepath.Join(output, ".gitignore")
	err = saveGitIgnoreFile(filePath, resp.Body)
	if err != nil {
		return err
	}

	// Only include filePath in the message if output flag is not provided
	// Otherwise, the message basically shows:
	// '.gitignore for %s generated successfully at .gitignore'
	if filePath == ".gitignore" {
		fmt.Printf(".gitignore for %s generated successfully!\n", language)
	} else {
		fmt.Printf(".gitignore for %s generated successfully at %s", language, filePath)
	}
	return nil
}

// FetchAvailableLanguages makes an HTTP request to GitHub's API for gitignore templates
// and returns an array of string representing the name of available programming languages.
func FetchAvailableLanguages() ([]string, error) {
	resp, err := http.Get(GitignoreAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Github API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var languages []string
	if err := json.Unmarshal(body, &languages); err != nil {
		return nil, err
	}

	return languages, nil
}
