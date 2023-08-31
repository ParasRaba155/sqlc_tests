package queries

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
)

func readFile(path string) [300]string {
	file, err := os.Open(path)
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("could not open the file : %s with error : %s", path, err)
		}
	}()
	if err != nil {
		log.Fatalf("could not open the file : %s with error : %s", path, err)
	}
	var result [300]string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		result[i] = scanner.Text()
		i++
	}
	if scanner.Err() != nil {
		log.Fatalf("could not scan the file : %s with error : %s", path, err)
	}
	return result
}

func generateRandomEmployee(q Querier) [300]CreateEmployeesParams {
	var result [300]CreateEmployeesParams
	emails := readFile("./emails.txt")
	names := readFile("./names.txt")
	ids, err := q.GetTentantID(context.TODO())

	if err != nil {
		log.Printf("could not get ids: %s", err)
	}
	for i := range result {
		result[i] = CreateEmployeesParams{
			Email:    emails[i],
			UserName: names[i],
			TenantID: getRandomElelmentFromList(ids),
		}
	}
	return result
}

func getRandomElelmentFromList[k any](list []k) k {
	return list[rand.Intn(len(list))]
}
