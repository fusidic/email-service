// email-service/main.go
package main

import (
	"context"
	"log"

	pb "github.com/fusidic/user-service/proto/user"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/googlepubsub"
)

const topic = "user.created"

// Subscriber handler for subscribe message
// 响应订阅信息
type Subscriber struct{}

// Process 订阅消息的响应函数
func (s *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to: ", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("email"),
		micro.Version("latest"),
	)

	srv.Init()

	// 不再需要第三方代理了
	// // 通过环境变量获取代理信息
	// pubsub := srv.Server().Options().Broker
	// if err := pubsub.Connect(); err != nil {
	// 	log.Fatal(err)
	// }

	// // 订阅消息，定义消息的回调函数，对消息进行反序列化
	// _, err := pubsub.Subscribe(topic, func(e broker.Event) error {
	// 	var user *pb.User
	// 	if err := json.Unmarshal(e.Message().Body, &user); err != nil {
	// 		return err
	// 	}
	// 	log.Println(user)
	// 	go sendEmail(user)
	// 	return nil
	// })

	// if err != nil {
	// 	log.Println(err)
	// }

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	// 运行服务
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

// sendEmail 未实现，仅作测试
func sendEmail(user *pb.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
