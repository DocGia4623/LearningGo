package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	// t := Title("string")
	// p := NewPublisher(&t)
	// m := NewMainService(p)
	// m.Run()
	// fx.New(
	// 	fx.Provide(NewMainService),
	// 	fx.Provide(
	// 		fx.Annotate(
	// 			NewPublisher,
	// 			fx.As(new(IPublisher)),
	// 			fx.ParamTags(`group:"titles"`),
	// 		),
	// 	),
	// 	fx.Provide(
	// 		titleComponent("goobye"),
	// 	),

	// 	fx.Provide(
	// 		titleComponent("hello"),
	// 	),
	// 	fx.Provide(
	// 		titleComponent("world"),
	// 	),
	// 	fx.Invoke(func(service *MainService) {
	// 		service.Run()
	// 	}),
	// ).Run()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// Log với các level khác nhau
	logrus.Info("This is an info message")
	logrus.Warn("This is a warning")
	logrus.Error("This is an error message")

}

func titleComponent(title string) any {
	return fx.Annotate(
		func() *Title {
			t := Title(title)
			return &t
		},
		fx.ResultTags(`group:"titles"`),
	)
}

// Main service
type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (service *MainService) Run() {
	service.publisher.Publish()
	log.Println("main program")
}

// Dependency
type IPublisher interface {
	Publish()
}

type Publisher struct {
	title []*Title
}

func NewPublisher(titles ...*Title) *Publisher {
	return &Publisher{title: titles}
}

func (publisher *Publisher) Publish() {
	for _, title := range publisher.title {
		log.Println("publishing", *title)
	}
}

// Dependency of publisher

type Title string
