package main

import (
	"encoding/csv"
	"fmt"
	"gopdf/mtTesting/api"
	"gopdf/mtTesting/login"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

var (
	jobs    = make(chan string)
	results = make(chan transcriptResult)

	domainAdmin = login.Person{}
)

type transcriptResult struct {
	personId       int
	statusCode     int
	transcriptSize int
}

func allocateJob() {
	filePtr, err := os.Open("/Users/rrode/jmeter/userListCSV/userList_JCP.csv")
	defer func() {
		if err = filePtr.Close(); err != nil {
			log.Fatal("Error when closing file: ", err)
		}
	}()
	if err != nil {
		fmt.Println("Failed to open file, error: ", err)
		return
	}
	r := csv.NewReader(filePtr)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(record[0])
		jobs <- record[0]
	}
	close(jobs)

}

func getTranscriptWorker(wg *sync.WaitGroup) {
	for personId := range jobs {
		pt := api.GetCourseTranscript + personId
		req, _ := http.NewRequest(http.MethodGet, pt, nil)
		req.Close = true
		req.Header.Add("x-session-token", domainAdmin.LoginResponse.Token)
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Expires", "0")
		req.Header.Add("DNT", "1")

		resp, err := domainAdmin.HTTPClient.Do(req)
		if err != nil {
			if v, ok := err.(*url.Error); ok {
				if v.Err.Error() == "EOF" {
					fmt.Println("EOF Error: ", err, "for studentId:", personId)
					return
				} else {
					panic(err)
				}
			}
		}
		//fmt.Println("---------Request Headers---------")
		//fmt.Println(req.Header)
		//fmt.Println("---------Response Headers---------")
		//fmt.Println(resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		intPersonId, _ := strconv.Atoi(personId)
		results <- transcriptResult{intPersonId,
			resp.StatusCode,
			len(body),
		}
		resp.Body.Close()
	}
	wg.Done()
}

func createWorkerPool(nOfW int) {
	var wg sync.WaitGroup
	for i := 0; i < nOfW; i++ {
		wg.Add(1)
		go getTranscriptWorker(&wg)
	}
	wg.Wait()
	close(results)
}

func printResults(done chan bool) {
	f, err := os.Create("/Users/rrode/Desktop/results.csv")
	if err != nil {
		fmt.Println("Cannot create result.csv, ", err)
		panic(err)
	}
	w := csv.NewWriter(f)
	w.Write([]string{"studentId", "status code", "transcript size"})
	for result := range results {
		//fmt.Println(result)
		record := []string{strconv.Itoa(result.personId),
			strconv.Itoa(result.statusCode),
			strconv.Itoa(result.transcriptSize),
		}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	done <- true
}

func main() {
	noOfWorkers := 50
	domainAdmin = login.Login("g_domain", "elearn")
	done := make(chan bool)
	go allocateJob()                 // Populate jobs chan with studentId
	go createWorkerPool(noOfWorkers) // create goroutines, each goroutine will fetch transcript
	go printResults(done)
	<-done
}
