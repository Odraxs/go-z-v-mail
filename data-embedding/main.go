package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Odraxs/go-z-v-mail/data-embedding/utils"
	"github.com/jhillyerd/enmime"
)

const (
	jsonIndexerPath     = "./index.json"
	indexName           = "emails"
	dataToIndexRootPath = "./maildir/"
	zincsearchBaseUrl   = "http://localhost:4080/api"
	dateFormatLayout    = "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"
)

func main() {
	log.Println("Starting indexer!")
	utils.CpuProfiling()
	indexerData, err := createIndexerFromJsonFile(jsonIndexerPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleting index if exists...")
	deleted := deleteIndexOnZincSearch(indexName)
	if deleted != nil {
		log.Println("Index doesn't exist. Creating...")
	}

	sent := createIndexOnZincSearch(indexerData)
	if sent != nil {
		log.Fatal(sent)
	}

	log.Println("Index created successfully.")
	log.Println("Start indexing, this might take a few minutes...")
	startTime := time.Now()

	var records []utils.EmailData
	var locker sync.Mutex
	var wg sync.WaitGroup

	// Process all the folders contained in the path `dataToIndexRootPath` to obtain all the emails records
	err = filepath.Walk(dataToIndexRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				emailData, err := processFile(p)
				if err != nil {
					log.Println(err)
					return
				}
				locker.Lock()
				records = append(records, emailData)
				locker.Unlock()
			}(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	sendBulkToZincSearch(records)
	utils.MemoryProfiling()
	duration := time.Since(startTime)
	log.Printf("Finished indexing. Time taken: %.2f seconds", duration.Seconds())
}

func processFile(path string) (utils.EmailData, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return utils.EmailData{}, err
	}
	mime, err := enmime.ReadEnvelope(bytes.NewReader(content))
	if err != nil {
		return utils.EmailData{}, err
	}

	date, err := time.Parse(dateFormatLayout, mime.GetHeader("Date"))
	if err != nil {
		return utils.EmailData{}, err
	}

	// Format the subject because many of them have string escapes that affect the `subject` sorting.
	subject := strings.ReplaceAll(mime.GetHeader("Subject"), "\"", "")

	return utils.EmailData{
		From:            mime.GetHeader("From"),
		To:              mime.GetHeader("To"),
		Subject:         subject,
		Content:         mime.Text,
		MessageID:       mime.GetHeader("Message-ID"),
		Date:            date.Format(time.RFC3339),
		ContentType:     mime.GetHeader("Content-Type"),
		MimeVersion:     mime.GetHeader("Mime-Version"),
		ContentEncoding: mime.GetHeader("Content-Transfer-Encoding"),
		XFrom:           mime.GetHeader("X-From"),
		XTo:             mime.GetHeader("X-To"),
		Xcc:             mime.GetHeader("X-cc"),
		Xbcc:            mime.GetHeader("X-bcc"),
		XFolder:         mime.GetHeader("X-Folder"),
		XOrigin:         mime.GetHeader("X-Origin"),
		XFilename:       mime.GetHeader("X-Filename"),
	}, nil
}

func sendBulkToZincSearch(records []utils.EmailData) {
	bulkData := utils.BulkData{
		Index:   "emails",
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", zincsearchBaseUrl+"/_bulkv2", bytes.NewReader(jsonData))
	if err != nil {
		log.Println(err)
		return
	}
	//TODO: change to .evn latter
	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func createIndexerFromJsonFile(filepath string) (utils.IndexerData, error) {
	var indexerData utils.IndexerData

	file, err := os.Open(filepath)
	if err != nil {
		return indexerData, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexerData)
	if err != nil {
		return indexerData, err
	}

	return indexerData, nil
}

func createIndexOnZincSearch(indexerData utils.IndexerData) error {
	jsonData, err := json.Marshal(indexerData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", zincsearchBaseUrl+"/index", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "password")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to create indexer, status code: %d", resp.StatusCode)
	}

	return nil
}

func deleteIndexOnZincSearch(indexName string) error {
	req, err := http.NewRequest("DELETE", zincsearchBaseUrl+"/index/"+indexName, nil)
	if err != nil {
		return err
	}

	//TODO: Change to .env file latter
	req.SetBasicAuth("admin", "password")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete indexer, status code: %d", resp.StatusCode)
	}

	log.Println("Index deleted successfully")
	return nil
}
