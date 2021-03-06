package nodeClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pranavt61/harmony-insight-database/sql"
)

func RequestAndStoreBlockGasUsed(shard_id int, block_height int) {

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
		// no block at height
		return
	}

	result, ok := resp_body["result"].(map[string]interface{})
	if ok == false {
		// no block at height
		return
	}

	gas_used := int(result["gasUsed"].(float64))

	// Store in DB
	if gas_used > 0 {
		fmt.Printf("SHARD %d - BLOCK %d - GAS %d\n", shard_id, block_height, gas_used)
		sql.InsertBlockGasUsed(shard_id, block_height, gas_used)
	}

	return
}
