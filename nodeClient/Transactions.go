package nodeClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pranavt61/harmony-insight-database/sql"
)

type PendingTransaction struct {
  From        string
  To          string
  Hash        string
  Shard_id    int
  Shard_id_to int
  Value       string
  Gas_used    string
  Gas_price   string
  Nonce       int
  Status      string
}

func RequestAndStoreBlockTransactionCount(shard_id int, block_height int) {

	// Prepare request
	req_body_string := fmt.Sprintf(`
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "hmyv2_getBlockTransactionCountByNumber",
			"params": [
				 %d
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

	tx_count, ok := resp_body["result"].(float64)
  if ok == false {
    // no block at height
    return
  }

	// Store in DB
	if int(tx_count) > 0 {
		fmt.Printf("SHARD %d - BLOCK %d - COUNT %d\n", shard_id, block_height, int(tx_count))
		sql.InsertBlockTransactionCount(shard_id, block_height, int(tx_count))
	}

	return
}

func RequestPendingTransactions(shard_id int) string {

	// Prepare request
	req_body_string := `
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "hmyv2_pendingTransactions",
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
		return ""
	}

  return fmt.Sprintf("%v", resp_body["result"])
}
