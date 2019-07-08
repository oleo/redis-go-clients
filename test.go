package main
import (
	"fmt"
	"time"
  "strconv"
	"github.com/go-redis/redis"
)
var client *redis.Client;
func main() {
	client := redis.NewClient(&redis.Options{	Addr:	"192.168.0.201:6379",	Password: "",	DB:		0})

	i := 0
	for  i < 1000 {
	  time.Sleep(1 * time.Second)
		i+=1
		err := client.Set("mygolang:key1","value is " + strconv.Itoa(i),0).Err()
		fmt.Println("Sending value is " + strconv.Itoa(i) + "\n")
		if err != nil {
			panic(err)
		}
	}	
	pong,err := client.Ping().Result()
	fmt.Println(pong,err)
}
