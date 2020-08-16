package main 

import (
    // "bufio"
    "fmt"
    // "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	mydir, _ := os.Getwd() 
	
	outputPlaylist := mydir + "/" + "playlist.m3u8"
	// playlist, err := m3u8.ReadFile(outputPlaylist)
	// fmt.Println("err: ", err)
	

	if _, err := os.Stat(outputPlaylist); err == nil {
		message := "Add this content at end"
		f, err := os.OpenFile(outputPlaylist, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
 
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		defer f.Close()
	 
		fmt.Fprintf(f, "%s\n", message)
	  
	  } else if os.IsNotExist(err) {
		fmt.Println("not exist")
		f, _ := os.Create(outputPlaylist)
		defer f.Close()
		l, _ := f.WriteString("#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-MEDIA-SEQUENCE:1\n#EXT-X-TARGETDURATION:180\n\nyooo\n")
		fmt.Println(l, "bytes written successfully")
	  } else {
		fmt.Println("aaaaaa")
	  
	  
	  }
}

