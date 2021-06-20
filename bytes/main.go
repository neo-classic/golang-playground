package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("{\"id\":\"26271735-3fc4-4c55-7e28-9c175946fbc1\",\"type\":\"notify_input\",\"name\":\"Demo Notify by Input\",\"goal\":null,\"template\":\"\",\"notify_input\":{\"play_audio\":{\"mode\":\"tts\",\"tts_properties\":{\"text\":\"Hi, {{recipient_name}}. Do you want to buy an elephant?\",\"language\":\"en-US\",\"voice\":\"en-US-Wavenet-C\",\"engine\":\"google\"}},\"input_passed\":{\"description\":\"passed\",\"input_value\":\"1\"},\"input_failed\":{\"description\":\"failed\",\"input_value\":\"2\"},\"sms_passed\":null,\"sms_failed\":null,\"farewell_play_audio_passed\":null,\"farewell_play_audio_failed\":null},\"account_sid\":\"accc2c661a0-3a41-360e-8417-3b4fdec4f980\"}")

	b = bytes.ReplaceAll(b, []byte("farewell_qualified_message"), []byte(""))

	fmt.Println(b)
}
