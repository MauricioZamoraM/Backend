package main

import (
	"fmt"
	"net/http"
	"time"
)

// func main() {
// 	start := time.Now()

// 	apis := []string{
// 		"https://management.azure.com",
// 		"https://dev.azure.com",
// 		"https://api.github.com",
// 		"https://outlook.office.com/",
// 		"https://api.somewhereintheinternet.com/",
// 		"https://graph.microsoft.com",
// 	}

// 	for _, api := range apis {
// 		_, err := http.Get(api)
// 		if err != nil {
// 			fmt.Printf("ERROR: %s is down!\n", api)
// 			continue
// 		}

// 		fmt.Printf("SUCCESS: %s is up and running!\n", api)
// 	}

// 	elapsed := time.Since(start)
// 	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

// }

// Un canal (channel) es una poderosa herramienta de comunicación que permite a las goroutines
// (hilos ligeros de ejecución) comunicarse entre sí y sincronizar sus ejecuciones. Los canales facilitan la
// transmisión de datos entre goroutines de manera segura y eficiente.


func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

