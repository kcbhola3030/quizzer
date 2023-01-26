package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Person struct {
	name  string
	score string
}

func main() {
	url := "http://localhost:5000/questions"
	var name string
	fmt.Print("Enter your name : ")
	nameC := make(chan string)
	go func() {
		fmt.Scanf("%s", &name)
		nameC <- name
	}()

	nam := <-nameC
	m := map[string]string{
		"name": nam,
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Print("conversion to json went wrong ")
		return
	}
	postData, err := http.Post("http://localhost:5000/user", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print("Something went wrong in post req")
	}
	id, _ := ioutil.ReadAll(postData.Body)
	fmt.Print(string(id))

	problems := questionPuller(url)
	tobj := time.NewTimer(20 * time.Second) // Time for all the questions --> 1 question => 10 seconds

	var correctAns int = 0
ProblemLoop:

	for i, problem := range problems {
		var answer string
		fmt.Printf("\nProblem %d: %s", i+1, problem.Question)
		fmt.Printf("\n a. %s \n b. %s \n c. %s \n d. %s \n Select Option 'a','b','c','d' : ",
			problem.Options.A, problem.Options.B, problem.Options.C, problem.Options.D)
		ansC := make(chan string)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()

		select {
		case <-tobj.C:
			fmt.Println("\nTime Over !!! Your Quiz has been Submitted\n")
			break ProblemLoop
		case iAns := <-ansC:
			if iAns == problem.Answer {
				correctAns++
			}
			if i == len(problems)-1 {
				fmt.Print("All Questions Submitted Successfully...:)\n ")
			}

		}

	}

	var scnid string

	fmt.Println("Enter You id")
	fmt.Scan(&scnid)

	m2 := map[string]string{
		"_id":   scnid,
		"score": strconv.Itoa(correctAns),
	}
	jsonData2, err := json.Marshal(m2)
	if err != nil {
		fmt.Print("conversion to json went wrong ")
		return
	}
	//postData2, err := http.put("http://localhost:5000/user", "application/json", bytes.NewBuffer(jsonData2));
	// postData2, err := http.NewRequest("PUT", "http://localhost:5000/updateUser", bytes.NewBuffer(jsonData2))
	// if err != nil {
	// 	fmt.Print("Something went wrong in post req")
	// }
	// id2, _ := ioutil.ReadAll(postData2.Body)
	// fmt.Print(string(id2))


	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/updateUser", bytes.NewBuffer(jsonData2))
    if err != nil {
        panic(err)
    }

    // set the request header Content-Type for json
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	id2, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(id2))
	fmt.Printf("Correct %d out of %d", correctAns, len(problems))


}


