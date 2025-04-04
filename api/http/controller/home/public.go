package home

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stonksdex/externalapi/api/common"
	"github.com/stonksdex/externalapi/codes"
)

func Public(c *gin.Context) {
	res := common.Response{}
	res.Timestamp = time.Now().Unix()

	res.Code = codes.CODE_SUCCESS
	res.Msg = "success"
	res.Data = struct {
		Rpc map[string]string `json:"rpc"`
	}{
		Rpc: map[string]string{
			"Solana":   "https://mainnet.helius-rpc.com/?api-key=4b1030d1-e346-4788-a65d-29c065efa012",
			"Ethereum": "https://eth.llamarpc.com",
			"Base":     "https://base-mainnet.infura.io/v3/15d81a19824c41159daec8327f691720",
			"Arbitrum": "https://arbitrum-mainnet.infura.io/v3/15d81a19824c41159daec8327f691720",
			"Bsc":      "https://binance.llamarpc.com",
		},
	}

	c.JSON(http.StatusOK, res)
}
