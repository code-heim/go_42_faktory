// Job Producer
package main

import (
	"fmt"
	"strconv"

	faktory "github.com/contribsys/faktory/client"
)

func main() {
	client, err := faktory.Open()
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		job := faktory.NewJob("report", "test"+strconv.Itoa(i)+"@codeheim.io")
		job.Queue = "critical"
		err = client.Push(job)
		if err != nil {
			fmt.Println("Error pushing job")
		}
	}

}
