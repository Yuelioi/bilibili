package main

import (
	"fmt"
	"os"

	"github.com/Yuelioi/bilibili/pkg/bpi"
	"github.com/Yuelioi/bilibili/pkg/client"
	"github.com/Yuelioi/bilibili/tests"
)

func main() {

	tests.LoadEnv()
	cli := client.New()
	cli.SESSDATA = os.Getenv("SESSDATA")
	cli.CSRF = os.Getenv("CSRF")
	bpi := bpi.New(cli)
	resp, _ := bpi.Video().Info(0, "BV18x4y1s7jP")

	fmt.Printf("resp: %v\n", resp)

}
