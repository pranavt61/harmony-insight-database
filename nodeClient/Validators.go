package nodeClient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pranavt61/harmony-insight-database/sql"
)

func RequestAndStoreValidators() {

	// Prepare request
	req_body_string := `
		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "hmyv2_getAllValidatorInformation",
			"params": [
				-1
			]
		}
	`

	req_body := strings.NewReader(req_body_string)
	req, err := http.NewRequest("POST", mapShardURLs[0], req_body)
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

	validators_list := resp_body["result"].([]interface{})

	max_len_n := 0
	max_len_w := 0
	for v_i := 0; v_i < len(validators_list); v_i++ {
		validator := validators_list[v_i].(map[string]interface{})["validator"].(map[string]interface{})

		v_name := validator["name"].(string)
		v_website := validator["website"].(string)
		v_address := validator["address"].(string)

		if max_len_n < len(v_name) {
			max_len_n = len(v_name)
		}
		if max_len_w < len(v_website) {
			max_len_w = len(v_website)
		}

		sql.InsertValidators(v_name, v_website, v_address)
	}
}
