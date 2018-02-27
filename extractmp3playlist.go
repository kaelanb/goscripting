package main

import (
	"os/exec"
	"log"
)

func main(){
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "--download-archive", "archive.txt", "<your playlist link>")
	cmd.Dir = "/your/music/directory"
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
				
}