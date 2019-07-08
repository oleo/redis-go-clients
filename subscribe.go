package main
import (
	"fmt"
	"strconv"
	"os"
//	"bytes"
//	"strings"
	"github.com/go-redis/redis"
)
var client *redis.Client;
var recievedbytes=0
func main() {
	client := redis.NewClient(&redis.Options{	Addr:	"192.168.0.201:6379",	Password: "",	DB:		0})

  defer client.Close()

	channel := os.Args[1]
	pubsub := client.Subscribe(channel)

  defer pubsub.Close()

	i := 0
	for i < 10000 {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		i+=1
		recievedbytes+=len(msg.Payload)
  	fmt.Print("\rReceived (" + strconv.Itoa(i) +"msgs, "+ strconv.Itoa(recievedbytes)+" bytes) : Now Received via " + msg.Channel + " ===> ")
	}

	
}
