package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ether/v1/utils"

	"github.com/gofiber/fiber/v2"
)

type EthResponse struct {
	GasPrice    string `json:"gas_price"`
	BlockNumber string `json:"block_number"`
	Balance     string `json:"balance"`
}

func GetEthInfo(c *fiber.Ctx) error {
	address := c.Params("address")
	ctx := context.Background()

	gasPrice := utils.GetCached(ctx, "gas_price")
	blockNumber := utils.GetCached(ctx, "block_number")

	if gasPrice == "" || blockNumber == "" {
		gasResp, _ := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_gasPrice&apikey=%s", os.Getenv("ETHERSCAN_API_KEY")))
		blockResp, _ := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s", os.Getenv("ETHERSCAN_API_KEY")))

		gasBody, _ := io.ReadAll(gasResp.Body)
		blockBody, _ := io.ReadAll(blockResp.Body)

		var gasResult map[string]any
		var blockResult map[string]any
		json.Unmarshal(gasBody, &gasResult)
		json.Unmarshal(blockBody, &blockResult)

		gasPrice = gasResult["result"].(string)
		blockNumber = blockResult["result"].(string)

		utils.SetCached(ctx, "gas_price", gasPrice, 30*time.Second)
		utils.SetCached(ctx, "block_number", blockNumber, 30*time.Second)
	}

	balanceResp, _ := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", address, os.Getenv("ETHERSCAN_API_KEY")))
	balanceBody, _ := io.ReadAll(balanceResp.Body)
	var balResult map[string]any
	json.Unmarshal(balanceBody, &balResult)
	balance := balResult["result"].(string)

	utils.SaveBalance(ctx, address, balance)

	return c.JSON(EthResponse{
		GasPrice:    gasPrice,
		BlockNumber: blockNumber,
		Balance:     balance,
	})
}
