package main

import (
	"bilibili/pkg/bpi"
	"bilibili/pkg/client"
	"bilibili/tests"
	"fmt"
	"os"
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
