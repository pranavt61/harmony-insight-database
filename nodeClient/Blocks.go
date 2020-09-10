package nodeClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RequestBlockNumber(shard_id int) int {

	// Prepare request
	req_body_string := `
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "hmyv2_blockNumber",
			"params": []
		}
	`
	req_body := strings.NewReader(req_body_string)
	req, err := http.NewRequest("POST", mapShardURLs[shard_id], req_body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response
	resp_body_buffer := new(bytes.Buffer)
	resp_body_buffer.ReadFrom(resp.Body)

	var resp_body_gen interface{}
	json.Unmarshal(resp_body_buffer.Bytes(), &resp_body_gen)
	resp_body := resp_body_gen.(map[string]interface{})

	block_number := int(resp_body["result"].(float64))

	fmt.Printf("SHARD %d - BLOCK NUM %d\n", shard_id, block_number)

	return block_number
}
