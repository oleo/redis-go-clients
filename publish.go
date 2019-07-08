package main
import (
	"fmt"
//	"time"
  //"strconv"
  "os"
	"github.com/go-redis/redis"
)
var client *redis.Client;
func main() {
	client := redis.NewClient(&redis.Options{	Addr:	"192.168.0.201:6379",	Password: "",	DB:		0})

	i := 0
	for  i < 1{
	  //time.Sleep(1 * time.Millisecond)
		i+=1
    channel := os.Args[1]
    msg := os.Args[2]
		err := client.Publish(channel,msg).Err()
		if err != nil {
			panic(err)
		}
	}	
	pong,err := client.Ping().Result()
	fmt.Println(pong,err)
}
