package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to Nerijus Gabalis' Bitcoin full node with JSON-RPC -----------------------------------------------------
	connCfg := &rpcclient.ConnConfig{
		Host:         "91.103.255.169:8332",
		User:         "bitcoin",
		Pass:         "VUISI",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Handle routing --------------------------------------------------------------------------------------------------
	router := gin.Default()

	// GET the page
	router.LoadHTMLFiles("./front/index.html")
	router.Static("/css", "./front/css")
	router.Static("/js", "./front/js")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// Get current block count
	router.GET("/getBlockCount", func(ctx *gin.Context) {
		blockCount, err := client.GetBlockCount()
		if err != nil {
			log.Fatal(err)
		}

		ctx.String(http.StatusOK, "%d", blockCount)
	})

	// Get the latest block
	router.GET("/getLatestBlock", func(ctx *gin.Context) {
		blockCount, _ := client.GetBlockCount()
		blockHash, _ := client.GetBlockHash(blockCount)

		ctx.JSON(http.StatusOK, gin.H{
			"Hash":   blockHash.String(),
			"Height": blockCount,
		})
	})

	// Get the header of a block from provided hash
	router.GET("/getBlockHeaderFromHash/:hash", func(ctx *gin.Context) {
		hashString := ctx.Param("hash")

		// Parse hash into a chainhash
		var hash chainhash.Hash
		_ = chainhash.Decode(&hash, hashString)
		header, _ := client.GetBlockHeaderVerbose(&hash)

		ctx.JSON(http.StatusOK, gin.H{
			"Timestamp":     header.Time,
			"Hash":          header.Hash,
			"PrevBlock":     header.PreviousHash,
			"NextBlock":     header.NextHash,
			"MerkleRoot":    header.MerkleRoot,
			"Nonce":         header.Nonce,
			"Version":       header.Version,
			"Bits":          header.Bits,
			"Confirmations": header.Confirmations,
			"Difficulty":    header.Difficulty,
			"Height":        header.Height,
		})
	})

	// Get the header of a block from provided height
	router.GET("/getBlockHeaderFromHeight/:height", func(ctx *gin.Context) {
		height, _ := strconv.ParseInt(ctx.Param("height"), 10, 64)
		hash, _ := client.GetBlockHash(height)

		header, _ := client.GetBlockHeaderVerbose(hash)

		ctx.JSON(http.StatusOK, gin.H{
			"Timestamp":     header.Time,
			"BlockHash":     header.Hash,
			"PrevBlock":     header.PreviousHash,
			"NextBlock":     header.NextHash,
			"MerkleRoot":    header.MerkleRoot,
			"Nonce":         header.Nonce,
			"Version":       header.Version,
			"Bits":          header.Bits,
			"Confirmations": header.Confirmations,
			"Difficulty":    header.Difficulty,
			"Height":        header.Height,
		})
	})

	// Get the transactions of a block
	router.GET("/getBlockTransactions/:hash", func(ctx *gin.Context) {
		hashString := ctx.Param("hash")

		// Parse hash into a chainhash
		var hash chainhash.Hash
		_ = chainhash.Decode(&hash, hashString)

		block, _ := client.GetBlockVerbose(&hash)
		txs := block.Tx

		ctx.JSON(http.StatusOK, gin.H{
			"amount":       len(txs),
			"transactions": txs,
		})
	})

	// Get block hash by block height
	router.GET("/getBlockHashFromHeight/:height", func(ctx *gin.Context) {
		height, _ := strconv.ParseInt(ctx.Param("height"), 10, 64)
		hash, _ := client.GetBlockHash(height)

		ctx.JSON(http.StatusOK, gin.H{
			"Hash": hash.String(),
		})
	})

	// Get block height by block hash
	router.GET("/getBlockHeightFromHash/:hash", func(ctx *gin.Context) {
		hashString := ctx.Param("hash")

		// Parse hash into a chainhash
		var hash chainhash.Hash
		_ = chainhash.Decode(&hash, hashString)
		block, _ := client.GetBlockHeaderVerbose(&hash)
		height := block.Height

		ctx.JSON(http.StatusOK, gin.H{
			"Height": height,
		})
	})

	_ = router.Run()
}
