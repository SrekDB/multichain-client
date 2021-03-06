package multichain

func (client *Client) IssueMore(accountAddress, assetName string, value float64) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id": CONST_ID,
		"method": "issuemore",
		"params": []interface{}{
			accountAddress,
			assetName,
			value,
		},
	}

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
