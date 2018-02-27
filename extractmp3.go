package main

import (
	"os/exec"
	"log"
)

func main(){
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "--download-archive", "archive.txt", "https://www.youtube.com/playlist?list=PLTP0mc_el22QmWmNs1QUdlyhSdStJtIId")
	cmd.Dir = "/your/music/directory"
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
				
}