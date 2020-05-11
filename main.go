// email-service/main.go
package main

import (
	"encoding/json"
	"log"

	pb "github.com/fusidic/user-service/proto/user"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("email"),
		micro.Version("latest"),
	)

	srv.Init()

	// 通过环境变量获取代理示例信息
	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// 将消息发布到代理上
	_, err := pubsub.Subscribe(topic, func(e broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(e.Message().Body, &user); err != nil {
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

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
