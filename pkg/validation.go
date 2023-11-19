package pkg

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
)

func CheckParameters(videoURL string) (string, error) {
	var id string
	var err error

	isMatch, err := regexp.MatchString(`https://www\.youtube\.com/watch\?v=[\w-]+`, videoURL) // TODO need better regex pattern
	if err != nil {
		return id, err
	}

	if !isMatch {
		return id, fmt.Errorf("GoTube: Invalid YouTube URL")
	}

	var reprURL *url.URL
	reprURL, err = url.Parse(videoURL)
	if err != nil {
		return id, err
	}
	id = reprURL.Query()["v"][0]

	return id, nil
}

func CheckValidDir(outputDirectory string) (bool, error) {
	var res bool
	var err error
	// Check if the given file/directory exists
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat(outputDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			return res, fmt.Errorf("The output directory '%v' doesn't exist", outputDirectory)
		}
		return false, err
	}
	if !fileInfo.Mode().IsDir() {
		return res, fmt.Errorf("The directory '%v' is a file", outputDirectory)
	}
	return true, err
}
