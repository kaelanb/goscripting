package main

import (
	"os/exec"
	"log"
)

func main(){
	
	name := "youtube-dl"
	args := []string{
		"-x",
		"--audio-format",
		"mp3",
		"--download-archive",
		"archive.txt",
		"-o",
		"%(title)s.%(ext)s",
		"https://www.youtube.com/playlist?list=<your playlist id>"}
	
	cmd := exec.Command(name, args...)
	cmd.Dir = "/your/music/directory"
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
				
}