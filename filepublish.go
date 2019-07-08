package main
import (
	"fmt"
//	"time"
  "strconv"
  "os"
	"github.com/go-redis/redis"
  "bufio"
  "regexp"
//	"io"
	//"io/ioutil"
)
var client *redis.Client;
func main() {
	client := redis.NewClient(&redis.Options{	Addr:	"192.168.0.201:6379",	Password: "",	DB:		0})

  filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
  defer file.Close()

	var myxml = ""


  channel := os.Args[1]

  var sentbytes=0
	i := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
		var data = scanner.Text()
		myxml = myxml + data
		
		match, _ := regexp.MatchString("</IcwEvent>",data)
		if(match) {
//			fmt.Println("Got XML: "+myxml)

			err := client.Publish(channel,myxml).Err()
			if err != nil {
				panic(err)
			}
			i+=1
			sentbytes+=len(myxml)
      fmt.Print("\rSendt messages : " + strconv.Itoa(i) + " " + strconv.Itoa(sentbytes) + " bytes ")
			myxml = ""
			
		}
  }

fmt.Println("\nDone")

}
