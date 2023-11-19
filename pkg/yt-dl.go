package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func getMetaData(id string) (string, string, error) {
	log.Printf("getMetaData for ID: %v", id)

	metaURL := "https://www.youtube.com/get_video_info?video_id=" + id

	info(fmt.Sprintf("Making a HTTP GET request thru %s...", metaURL))

	resp, err := http.Get(metaURL)
	var fileName string
	var downloadURL string

	if err != nil {
		return fileName, downloadURL, fmt.Errorf("GoTube: Failed to acquire video info: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fileName, downloadURL, fmt.Errorf("GoTube: Bad status: %s (%s)", resp.Status, http.StatusText(resp.StatusCode))
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	log.Printf("received: %v", string(byteArray))

	data := make(map[string]interface{})
	//err = parseStr(string(byteArray[:]), data)
	if err != nil {
		return fileName, downloadURL, fmt.Errorf("GoTube: Failed to parse video info response: %v", err)
	}

	// We only need to retrieve video title, format and download url nothing else

	log.Printf("player_response: %v", data["player_response"])
	var videoData map[string]interface{}
	err = json.Unmarshal([]byte(data["player_response"].(string)), &videoData)
	if err != nil {
		return fileName, downloadURL, fmt.Errorf("GoTube: Failed to unmarshal video info data: %v", err)
	}

	log.Printf("videoData: %v", videoData)
	for key, value := range videoData {
		log.Printf("videoData: %v - %v", key, value)
	}
	log.Printf("videoDetails: %v", videoData["videoDetails"])
	log.Printf("streamingData: %v", videoData["streamingData"])

	if videoData["streamingData"] == nil {
		return fileName, downloadURL, fmt.Errorf("GoTube: streamingData is missing from this video '%v'", id)
	}

	videoDetails := videoData["videoDetails"].(map[string]interface{})
	streamingData := videoData["streamingData"].(map[string]interface{})
	formats := streamingData["formats"].([]interface{})

	// Let's try the first format...
	moreData := formats[0].(map[string]interface{})
	moreData["mime"] = moreData["mimeType"]
	s := moreData["mime"].(string)

	title := strings.Replace(strings.ToLower(videoDetails["title"].(string)), " ", "_", -1)
	format := s[strings.Index(s, "/")+1 : strings.Index(s, ";")]
	downloadURL = moreData["url"].(string)

	// Remove characters like ':' and '?' in the video title
	re := regexp.MustCompile(`[^A-Za-z0-9.\_\-]`)
	fileName = re.ReplaceAllString(title+"."+format, "")

	return fileName, downloadURL, nil
}
