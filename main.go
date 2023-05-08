package main

import (
	"context"
	"fmt"
	"net/http"
	api "test-api-blockchain/api/nft"
	apiverify "test-api-blockchain/api/verifly"

	_ "context"
	_ "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Canddidate struct {
	Count string `json:"count"`
}

type CanddidateTotal struct {
	Votes []string `json:"canddidate"`
}

var conn *api.Api
var auth *bind.TransactOpts
var connVerify *apiverify.Api

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	connectBlockchain()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/mint", mintNft)
	e.GET("/mint", getNft)
	// e.GET("/totalVoteFor/:name", getTotalVoteFor)
	// e.GET("/votCaddindate/:name", vote)
	// e.GET("/totalVoteAll", totalVoteAll)
	e.Logger.Fatal(e.Start(":9000"))
}

func connectBlockchain() {
	//runnig blockchain client
	client, err := ethclient.Dial("http://dgs-rpc.digitsoul.co.th")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("efa96e96c8fb1851115aec1b5d39b2499bbe8cea7c21c19d00923cd5ba2a6988")
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	conn, err = api.NewApi(common.HexToAddress("0x63F4c2b4845F380ab823EA1962B7bB57113EbCd5"), client)
	if err != nil {
		panic(err)
	}
	connVerify, err = apiverify.NewApi(common.HexToAddress("0xb10E6Aa2Ef0cB3F1e5C25b4658FEd043EDeb089a"), client)
	if err != nil {
		fmt.Println("ererer")
		panic(err)
	}
}

type Mints struct {
	Address string `json:"address"`
	Uri     string `json:"uri"`
	Items   Itmes  `json:"items"`
}

type Itmes struct {
	Key        string `json:"key"`
	FullName   string `json:"full_name"`
	CourseName string `json:"course_name"`
	CreatedAt  string `json:"created_at"`
}

func mintNft(c echo.Context) error {
	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	// if err != nil {
	// 	return err
	// }
	// metadata := fmt.Sprintf("%v", json_map["metadata"])
	// jsonEncodeUrl := url.PathEscape(metadata)
	// return c.JSON(http.StatusOK, jsonEncodeUrl)
	data := Mints{}

	if err := c.Bind(&data); err != nil {
		fmt.Println("rerer2")
		return c.NoContent(http.StatusBadRequest)
	}
	addressBig := common.HexToAddress(data.Address)
	transection, err := conn.SafeMint(auth, addressBig, data.Uri)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Status Error")
	}
	transection, err = connVerify.Add(auth, data.Items.Key, data.Items.FullName, data.Items.CourseName, data.Items.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Status Error")
	}
	return c.JSON(http.StatusOK, transection)
}

func getNft(c echo.Context) error {
	key := c.QueryParam("key")
	if key == "" {
		return c.JSON(http.StatusInternalServerError, "Status Error")
	}
	fmt.Println(key)
	data, err := connVerify.Ipfs(&bind.CallOpts{}, key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Status Error")
	}
	return c.JSON(http.StatusOK, data)
}
