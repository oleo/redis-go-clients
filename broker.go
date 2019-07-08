package main
import (
	"fmt"
	"strconv"
	"os"
	"github.com/go-redis/redis"
	"regexp"
)
var client *redis.Client;
func main() {
	client := redis.NewClient(&redis.Options{	Addr:	"192.168.0.201:6379",	Password: "",	DB:		0})

  defer client.Close()

	fromchannel := os.Args[1]
  tochannel := os.Args[2]
  myregexp := os.Args[3]
	pubsub := client.Subscribe(fromchannel)

  defer pubsub.Close()

	i := 0
	for i < 100000 {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		match, _ := regexp.MatchString(myregexp,msg.Payload)
		if(match) {
  	fmt.Print("\rHandled ("+strconv.Itoa(i)+") - Now Forwarding " + msg.Channel + " message from " + fromchannel + " to " + tochannel)

    merr := client.Publish(tochannel,msg.Payload).Err()
		i+=1
    if merr != nil {
      panic(err)
    }
		}

	}

	
}
