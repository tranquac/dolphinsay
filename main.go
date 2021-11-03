package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"embed"
)

// Hey, I want to embed "dolphins" folder in the executable binary
// Use embed go 1.16 new feature (for embed dolphins static files)
//go:embed dolphins
var embedDolphinFiles embed.FS

func main() {

	// Display usage/help message
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "-h") || (len(os.Args) == 2 && os.Args[1] == "--help") {
		usage := "DolphinSay is inspired by Cowsay program.\nDolphinSay allow you to display a message said by a cute random Dolphin.\n\nUsage:\n   dolphinsay MESSAGE\n\nExample: dolphinsay hello Gopher lovers"

		fmt.Println(usage)
		return
	} else if len(os.Args) > 1 {

		message := strings.Join(os.Args[1:], " ")
		nbChar := len(message)

		line := " "
		for i := 0; i <= nbChar; i++ {
			line += "-"
		}

		fmt.Println(line)
		fmt.Println("< " + message + " >")
		fmt.Println(line)
		fmt.Println("        \\")
		fmt.Println("         \\")

		// Generate a random integer depending on get the number of ascii files
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(getNbOfDolphinFiles() - 1)

		// Display random gopher asii embed files
		fileData, err := embedDolphinFiles.ReadFile("dolphins/dolphin" + strconv.Itoa(randInt) + ".txt")
		if err != nil {
			log.Fatal("Error during read dolphin ascii file", err)
		}
		fmt.Println(string(fileData))
	}
}

func getNbOfDolphinFiles() int {

	files, err := embedDolphinFiles.ReadDir("dolphins")
	if err != nil {
		log.Fatal("Error during reading dolphins folder", err)
	}

	nbOfFiles := 0
	for _, _ = range files {
		nbOfFiles++
	}

	return nbOfFiles
}
