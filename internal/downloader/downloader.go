package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// The base URL to the repository that provides the .gitignore files
const BaseURL = "https://raw.githubusercontent.com/github/gitignore/main/"

// capitalize returns a new string with the first alphabet being capitalized and the rest being lower cased.
func capitalize(word string) string {
	if word == "" {
		return word
	}

	return strings.ToUpper(word[:1]) + word[1:]
}

// getGitIgnoreURL constructs the URL to fetch the .gitignore file.
func getGitIgnoreURL(language string) string {
	return BaseURL + capitalize(language) + ".gitignore"
}

// fetchGitIgnoreContents performs the HTTP GET request to retrieve the .gitignore file.
func fetchGitIgnoreContents(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch .gitignore: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("receieved status code %d", resp.StatusCode)
	}
	return resp, nil
}

// saveGitIgnoreFile creates a file and writes the contents from the response body.
func saveGitIgnoreFile(filePath string, body io.Reader) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
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

	fmt.Printf(".gitignore for %s downloaded successfully at %s\n", language, filePath)
	return nil
}
