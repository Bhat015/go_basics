package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://golang.org"}

	c := make(chan string) //channel

	for _, link := range links {
		go checkLink(link, c)
		//fmt.Println(<-c)
	}

	//for i := 0; i < len(links); i++ {
	//fmt.Println(<-c)
	//	checkLink(<-c, c)
	//}

	//for { // infinite loop
	//	go checkLink(<-c, c)
	//}

	// other way

	for l := range c {
		//go checkLink(l, c)

		go func(link string) { // func literal
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// channels is a way to communicate between goroutines

func checkLink(link string, c chan string) {
	//print("Checking link: " + link)
	_, err := http.Get(link) // auto awaits for the response

	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- " Channel msg : Yep, it's up"
	c <- link
}
