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
	resp_body, ok := resp_body_gen.(map[string]interface{})
	if ok == false {
		// no block number ?
		return -1
	}

	block_number := int(resp_body["result"].(float64))

	return block_number
}

func RequestBlockHashByNumber(block_height int) string {

	// loop through all shards
	// lower ID shards have prio
	for shard_id := 0; shard_id < 4; shard_id++ {
		if shard_id == 2 {
			continue
		}

		// Prepare request
		req_body_string := fmt.Sprintf(`
			{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "hmyv2_getBlockByNumber",
				"params": [
					%d,
					{
							"fullTx": false,
							"inclTx": false,
							"InclStaking": false
					}
				]
			}
		`, block_height)
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
		resp_body, ok := resp_body_gen.(map[string]interface{})
		if ok == false {
			// no block
			continue
		}

		block_data, ok := resp_body["result"].(map[string]interface{})
		if ok == false {
			// no block
			continue
		}

		block_hash, ok := block_data["hash"].(string)
		if ok == false {
			// no block
			continue
		}

		return block_hash
	}

	return ""
}
