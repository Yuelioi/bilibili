package main

import (
	"bilibili/pkg/client"
	"bilibili/pkg/endpoints/article"
	"fmt"
	"log"
)

func main() {
	cli := client.New()
	cli.SESSDATA = ""
	cli.CSRF = ""

	// 添加自定义请求头

	articleID := 1 // 替换为实际的文章 ID

	// 发送 POST 请求
	br, err := article.New(cli).UnFavorite(articleID)
	if err != nil {
		log.Fatal("Error sending request:")
	}

	fmt.Printf("br: %v\n", br)

}
