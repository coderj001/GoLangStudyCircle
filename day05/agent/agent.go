package agent

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Job struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func execute(job Job) Result {
	log.Println("executing ", job.Name)
	return Result{ID: job.ID, Status: "SUCCESS"}
}

func process(jobs *[]Job) ([]Result, error) {
	resultCh := make(chan Result)

	for _, job := range *jobs {
		go func(job Job) {
			resultCh <- execute(job)
		}(job)
	}

	var results []Result
	for range *jobs {
		results = append(results, <-resultCh)
	}

	return results, nil
}

func processFile() {
	jobsData, err := ioutil.ReadFile("jobs.json")
	if err != nil {
		log.Fatal(err)
	}

	var jobs []Job
	if err := json.Unmarshal(jobsData, &jobs); err != nil {
		log.Fatal(err)
	}

	results, _ := process(&jobs)

	resultsData, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("status.json", resultsData, 0644); err != nil {
		log.Fatal(err)
	}
}

func main() {
	processFile()
}
