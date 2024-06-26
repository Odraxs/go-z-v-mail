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
	"runtime/pprof"
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

var (
	zincUser     = os.Getenv("ZINC_USER")
	zincPassword = os.Getenv("ZINC_PASSWORD")
)

func main() {
	log.Println("Starting indexer!")

	//utils.CpuProfiling() doesn't generates a proper cpu profiling
	f, err := os.Create("./profs/cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

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
	// This channel was created with the intention to limit the cpu usage from the goroutines but after running it 
	// it seems like it doesn't work as expected, anyways I'm gonna let it there since it could be helpful to someone trying to optimize this code. 
	routines := make(chan int, 1000)
	jobCounter := 0

	// Process all the folders contained in the path `dataToIndexRootPath` to obtain all the emails records
	err = filepath.Walk(dataToIndexRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			jobCounter++
			go func(p string, routines <-chan int) {
				defer wg.Done()
				for range routines {
					emailData, err := processFile(p)
					if err != nil {
						log.Println(err)
						return
					}
					locker.Lock()
					records = append(records, emailData)
					locker.Unlock()
				}
			}(path, routines)
			routines <- jobCounter
		}
		return nil
	})
	close(routines)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	err = sendBulkToZincSearch(records)
	if err != nil {
		log.Println("something happened while storing the records: ", err)
	}
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

func sendBulkToZincSearch(records []utils.EmailData) error {
	bulkData := utils.BulkData{
		Index:   indexName,
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("failed to decode data into json: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, zincsearchBaseUrl+"/_bulkv2", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create the http request: %w", err)
	}

	req.SetBasicAuth(zincUser, zincPassword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("something happened while doing the request to zincsearch:  %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status response: %d", resp.StatusCode)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
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
		return fmt.Errorf("failed to encode the index data: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, zincsearchBaseUrl+"/index", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create the index request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(zincUser, zincPassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("zincsearch request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create indexer, status code: %d", resp.StatusCode)
	}

	return nil
}

func deleteIndexOnZincSearch(indexName string) error {
	req, err := http.NewRequest(http.MethodDelete, zincsearchBaseUrl+"/index/"+indexName, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(zincUser, zincPassword)

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
