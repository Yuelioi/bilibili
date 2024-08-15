package main

import (
	"fmt"
	"os"

	"github.com/Yuelioi/bilibili/pkg/bpi"
	"github.com/Yuelioi/bilibili/tests"
)

func main() {

	tests.LoadEnv()
	bpi := bpi.New()
	bpi.Client.SESSDATA = os.Getenv("SESSDATA")
	bpi.Client.CSRF = os.Getenv("CSRF")
	resp, _ := bpi.Video().Info(0, "BV18x4y1s7jP")

	fmt.Printf("resp: %v\n", resp)

}
