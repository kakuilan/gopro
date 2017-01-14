//m3u 读入后缀名为.m3u的音乐播放列表文件并输出一个等同的.pls播列表文件
//使用 
//go build t737.go
// ./t737 David-Bowie-Singles.m3u >> mu3.pls
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

type Song struct {
    Title string 
    Filename string 
    Seconds int
}

func main () {
    //检查命令行是否指定了包含.m3u后缀名的文件
    if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") { 
        //filepath.Base()返回给定路径的基本名
        fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0])) 
        //os.Exit() 函数会在退出前清理所有资源
        os.Exit(1)
    }

    if rawBytes,err := ioutil.ReadFile(os.Args[1]); err != nil {
        log.Fatal(err)
    }else{
        songs :=  readM3uPlaylist(string(rawBytes)) 
        writePlsPlaylist(songs)
    }

}

func readM3uPlaylist(data string) (songs []Song) {
    var song Song
    for _,line := range strings.Split(data, "\n") {
        line = strings.TrimSpace(line) 
        if line =="" || strings.HasPrefix(line, "#EXTM3U") {
            continue
        } 

        if strings.HasPrefix(line, "#EXTINF:") {
            song.Title, song.Seconds =  parseExtinfLine(line)
        }else{
            song.Filename = strings.Map(mapPlatformDirSeparator, line)
        }

        if song.Filename !="" && song.Title!="" && song.Seconds!=0 {
            songs = append(songs, song)
            song = Song{}
        }
    }

    return songs
}


func parseExtinfLine(line string) (title string, seconds int) {
    if i:= strings.IndexAny(line, "-0123456789"); i >-1 {
        const  separator = ","  
        line = line[i:] 
        if j := strings.Index(line, separator); j > -1 {
            title = line[j+len(separator):]  
            var err error 
            if  seconds,err = strconv.Atoi(line[:j]); err != nil {
                log.Printf("faled to read the duration for '%s' : %v\n", title, err) 
                seconds = -1
            }
        }
    }

    return title,seconds
}

func mapPlatformDirSeparator(char rune) rune {
    if char == '/' || char == '\\' {
        return filepath.Separator
    }

    return char
}

func writePlsPlaylist(songs []Song) {
    fmt.Println("[playlist]")
    for i,song := range songs {
        i++ 
        fmt.Printf("File%d=%s\n" , i, song.Filename) 
        fmt.Printf("Title%d=%s\n", i, song.Title) 
        fmt.Printf("Length%d=%d\n", i, song.Seconds)
    }
    fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}
