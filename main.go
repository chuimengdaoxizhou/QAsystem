package main

import (
	"qasystem/text-generate-client"
)

func main() {
	prompt := "给我讲一个笑话"
	text_generate_client.GenerateToken(prompt)
}
