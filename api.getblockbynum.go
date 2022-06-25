package trongrid

type TGetBlockByNumReq struct {
	Num int `json:"num"` // 数字
}

type TGetBlockByNumRsp struct {
}

func GetBlockByNum(req *TGetBlockByNumReq) (rsp TGetBlockByNumRsp, err error) {
	url := "https://api.shasta.trongrid.io/wallet/getblockbynum"
	err = MyPost(url, req, &rsp)
	return
}

/*

{
  "blockID": "00000000017bfad2869641e02e68563771b038e91fbe35df3fe4333df11d36b0",
  "block_header": {
    "raw_data": {
      "number": 24902354,
      "txTrieRoot": "610a475ae2c9cb27a4ea5f123c0310eb0f2c73370028934291cff8ed20698a33",
      "witness_address": "41f16412b9a17ee9408646e2a21e16478f72ed1e95",
      "parentHash": "00000000017bfad18ca0dd12811737d3585af5fe9af908ebf26df6d6568ffb51",
      "version": 23,
      "timestamp": 1654152273000
    },
    "witness_signature": "8f8b5cc9ece7b823d24c0360623b746a2c556b19134f44e0db643e6dd032e7a32834657dd5dc520224067fc31923cc40e0a2d287f1f04326c06b8feeea16344001"
  },
  "transactions": [
    {
      "ret": [
        {
          "contractRet": "SUCCESS"
        }
      ],
      "signature": [
        "db40809beb78e8dbc43a54e260db04a726b129063d7c0c0501f585cee51fa4159f349916d4bec3aa796b812f386a07cac1cc4017145b5759c37f8fd8159919da01"
      ],
      "txID": "c9a43a821eb5ac32d97eca073030e19bef866734e67152008c319be4e70c8c35",
      "raw_data": {
        "contract": [
          {
            "parameter": {
              "value": {
                "amount": 1000000000,
                "owner_address": "4130ba69b6b9746c2a15d6a7f9b9bea1f97ba3b023",
                "to_address": "4108a8954b7862d07cf4ce6ac78c3a83d72bd917e4"
              },
              "type_url": "type.googleapis.com/protocol.TransferContract"
            },
            "type": "TransferContract"
          }
        ],
        "ref_block_bytes": "fabe",
        "ref_block_hash": "6c357b3c30ce6d78",
        "expiration": 1654152327000,
        "timestamp": 1654152269917
      },
      "raw_data_hex": "0a02fabe22086c357b3c30ce6d7840d8dea59992305a69080112650a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412340a154130ba69b6b9746c2a15d6a7f9b9bea1f97ba3b02312154108a8954b7862d07cf4ce6ac78c3a83d72bd917e4188094ebdc0370dda0a2999230"
    }
  ]
}
*/
