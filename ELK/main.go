package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
)

func main() {
	logPath := "C:\\Users\\Admin\\Desktop\\ELK\\logs\\go.log"
	_, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", "C:\\Users\\Admin\\Desktop\\ELK\\logs\\go.log"}
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	i := 0
	for {
		i++
		time.Sleep(time.Second * 3)
		if rand.Intn(10) == 1 {
			l.Error("test error", zap.Error(fmt.Errorf("error because test: %d", i)))
		} else {
			l.Info(fmt.Sprintf("test log: %d", i))
		}
	}
}
