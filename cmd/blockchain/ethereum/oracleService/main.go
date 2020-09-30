package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/oracleService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	/*
	 * Read in Oracle address
	 */
	var deployedContract = flag.String("deployedContract", "", "Address of the deployed oracle contract")
	var topCoins = flag.Int("topCoins", 15, "Number of coins to push with the oracle")
	flag.Parse()

	/*
	 * Read secrets for unlocking the ETH account
	 */
	var lines []string
	file, err := os.Open("/run/secrets/oracle_keys") // Read in key information
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 2 {
		log.Fatal("Secrets file should have exactly two lines")
	}
	key := lines[0]
	key_password := lines[1]

	/*
	 * Setup connection to contract, deploy if necessary
	 */
	// TO DO: Switch from Infura to our node
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/ec6581408f09414b8e4446067cd3ba08")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), key_password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *oracleService.DiaOracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	periodicOracleUpdateHelper(topCoins, auth, contract)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(topCoins, auth, contract)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(topCoins *int, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	// fortube Data
	rawFortube, err := getDefiRatesFromDia("FORTUBE", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve forTube data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawFortube, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Fortube Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)
	// ddex Data
	rawDdex, err := getDefiRatesFromDia("DDEX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve ddex data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawDdex, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update ddex Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)
	// nuo Data
	rawNuo, err := getDefiRatesFromDia("NUO", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Nuo data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawNuo, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Nuo Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// bZx Data
	rawBzx, err := getDefiRatesFromDia("BZX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve bZx data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawBzx, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update bZx Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Compound Data
	rawCompound, err := getDefiRatesFromDia("COMPOUND", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Compound data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawCompound, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Compound Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Compound State Data
	rawCompoundState, err := getDefiStateFromDia("COMPOUND")
	if err != nil {
		log.Fatalf("Failed to retrieve Compound state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawCompoundState, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Compound state Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// DYDX Data
	rawDydx, err := getDefiRatesFromDia("DYDX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve DYDX data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawDydx, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update DYDX Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// DYDX State Data
	rawDydxState, err := getDefiStateFromDia("DYDX")
	if err != nil {
		log.Fatalf("Failed to retrieve DYDX state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawDydxState, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update DYDX state Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Aave Data
	rawAave, err := getDefiRatesFromDia("AAVE", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Aave data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawAave, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Aave Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// ECB Data
	rawECB, err := getECBRatesFromDia("EUR")
	if err != nil {
		log.Fatalf("Failed to retrieve ECB from DIA: %v", err)
		return err
	}
	err = updateECBRate(rawECB, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update ECB Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Maker DEX data
	rawMaker, err := getDEXFromDia("Maker", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Maker from DIA: %v", err)
		return err
	}

	err = updateDEX(rawMaker, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Maker Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Curvefi DEX data
	rawCurvefi, err := getDEXFromDia("Curvefi", "DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve Curvefi from DIA: %v", err)
		return err
	}

	err = updateDEX(rawCurvefi, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Curvefi Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Gnosis DEX data
	rawGnosis, err := getDEXFromDia("Gnosis", "DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve Gnosis from DIA: %v", err)
		return err
	}

	err = updateDEX(rawGnosis, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Gnosis Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Uniswap data
	rawUniswap, err := getDEXFromDia("Uniswap", "DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve Uniswap from DIA: %v", err)
		return err
	}

	err = updateDEX(rawUniswap, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Uniswap Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Loopring  data
	rawLoopring, err := getDEXFromDia("Loopring", "LRC")
	if err != nil {
		log.Fatalf("Failed to retrieve Loopring from DIA: %v", err)
		return err
	}

	err = updateDEX(rawLoopring, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Loopring Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Bancor data
	rawBancor, err := getDEXFromDia("Bancor", "USDT")
	if err != nil {
		log.Fatalf("Failed to retrieve Bancor from DIA: %v", err)
		return err
	}

	err = updateDEX(rawBancor, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Bancor Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Top 15 coins
	rawCoins, err := getToplistFromDia()
	if err != nil {
		log.Fatalf("Failed to retrieve toplist from DIA: %v", err)
		return err
	}

	cleanedCoins := []models.Coin{}

	for key := range rawCoins.Coins {
		if rawCoins.Coins[key].CirculatingSupply != nil {
			cleanedCoins = append(cleanedCoins, rawCoins.Coins[key])
		}
	}
	sort.Slice(cleanedCoins, func(i, j int) bool {
		return cleanedCoins[i].Price**cleanedCoins[i].CirculatingSupply > cleanedCoins[j].Price**cleanedCoins[j].CirculatingSupply
	})
	topCoinSlice := cleanedCoins[:*topCoins]

	err = updateTopCoins(topCoinSlice, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coins Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	// Coingecko Data
	// TO DO: Add quotations
	rawQuotBTC, err := getForeignQuotationFromDia("Coingecko", "BTC")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotBTC, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotETH, err := getForeignQuotationFromDia("Coingecko", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotETH, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotUSDT, err := getForeignQuotationFromDia("Coingecko", "USDT")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotUSDT, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotXRP, err := getForeignQuotationFromDia("Coingecko", "XRP")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotXRP, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotBNB, err := getForeignQuotationFromDia("Coingecko", "BNB")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotBNB, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotBCH, err := getForeignQuotationFromDia("Coingecko", "BCH")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotBCH, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotDOT, err := getForeignQuotationFromDia("Coingecko", "DOT")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotDOT, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotLINK, err := getForeignQuotationFromDia("Coingecko", "LINK")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotLINK, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotBSV, err := getForeignQuotationFromDia("Coingecko", "BSV")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotBSV, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotADA, err := getForeignQuotationFromDia("Coingecko", "ADA")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotADA, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotCRO, err := getForeignQuotationFromDia("Coingecko", "CRO")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotCRO, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotLTC, err := getForeignQuotationFromDia("Coingecko", "LTC")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotLTC, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotUSDC, err := getForeignQuotationFromDia("Coingecko", "USDC")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotUSDC, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotEOS, err := getForeignQuotationFromDia("Coingecko", "EOS")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotEOS, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotXMR, err := getForeignQuotationFromDia("Coingecko", "XMR")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotXMR, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotTRX, err := getForeignQuotationFromDia("Coingecko", "TRX")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotTRX, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotOKB, err := getForeignQuotationFromDia("Coingecko", "OKB")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotOKB, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotXTZ, err := getForeignQuotationFromDia("Coingecko", "XTZ")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotXTZ, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotXLM, err := getForeignQuotationFromDia("Coingecko", "XLM")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotXLM, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	rawQuotNEO, err := getForeignQuotationFromDia("Coingecko", "NEO")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotNEO, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle: %v", err)
		return err
	}
	time.Sleep(5 * time.Minute)

	return nil
}

func updateTopCoins(topCoins []models.Coin, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	for _, element := range topCoins {
		symbol := strings.ToUpper(element.Symbol)
		name := element.Name
		supply := element.CirculatingSupply
		price := element.Price
		// Get 5 digits after the comma by multiplying price with 100000
		err := updateOracle(contract, auth, name, symbol, int64(price*100000), int64(*supply))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
		time.Sleep(5 * time.Minute)
	}
	return nil
}

func updateDEX(dexData *models.Points, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	symbol := strings.ToUpper(dexData.DataPoints[0].Series[0].Values[0][3].(string))
	name := dexData.DataPoints[0].Series[0].Values[0][1].(string)
	supply := 0
	price := dexData.DataPoints[0].Series[0].Values[0][4].(float64)
	// Get 5 digits after the comma by multiplying price with 100000
	// Set supply to 0, as we don't have a supply for one exchange
	err := updateOracle(contract, auth, name, symbol, int64(price*100000), int64(supply))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}
	return nil
}

func updateECBRate(ecbRate *models.CurrencyChange, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	symbol := strings.ToUpper(ecbRate.Symbol)
	name := strings.ToUpper(ecbRate.Symbol)
	price := ecbRate.Rate
	// Get 5 digits after the comma by multiplying price with 100000
	// Set supply to 0, as we don't have a supply for fiat currencies
	err := updateOracle(contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateDefiRate(defiRate *dia.DefiRate, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	symbol := strings.ToUpper(defiRate.Asset)
	name := strings.ToUpper(defiRate.Protocol)
	price := defiRate.LendingRate
	// Get 5 digits after the comma by multiplying price with 100000
	// Set supply to 0, as we don't have a supply for fiat currencies
	err := updateOracle(contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateDefiState(defiState *dia.DefiProtocolState, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	symbol := ""
	name := strings.ToUpper(defiState.Protocol.Name) + "-state"
	price := defiState.TotalUSD
	// Get 5 digits after the comma by multiplying price with 100000
	// Set supply to 0, as we don't have a supply for fiat currencies
	err := updateOracle(contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateForeignQuotation(foreignQuotation *models.ForeignQuotation, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	name := foreignQuotation.Source + "-" + foreignQuotation.Name
	symbol := foreignQuotation.Symbol
	price := foreignQuotation.Price
	err := updateOracle(contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **oracleService.DiaOracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = oracleService.NewDiaOracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = oracleService.DeployDiaOracle(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

func getCoinDetailsFromDia(symbol string) (*models.Coin, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/symbol/" + symbol)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.SymbolDetails
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b.Coin, nil
		}
		return nil, err
	}
}

func getToplistFromDia() (*models.Coins, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/coins")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Coins
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

// Getting EUR vs XXX rate
func getECBRatesFromDia(symbol string) (*models.CurrencyChange, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/coins")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Coins
		err = b.UnmarshalBinary(contents)
		if err != nil {
			return nil, err
		}

		for _, change := range b.Change.USD {
			if strings.ToUpper(change.Symbol) == strings.ToUpper(symbol) {
				return &change, nil
			}
		}
		return nil, nil
	}
}

// Getting defi rate
func getDefiRatesFromDia(protocol string, symbol string) (*dia.DefiRate, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingRate/" + strings.ToUpper(protocol) + "/" + strings.ToUpper(symbol) + "/" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiRate
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

// Getting defi state
func getDefiStateFromDia(protocol string) (*dia.DefiProtocolState, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingState/" + strings.ToUpper(protocol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiProtocolState
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getDEXFromDia(dexname string, symbol string) (*models.Points, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/chartPoints/MAIR120/" + dexname + "/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Points
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getForeignQuotationFromDia(source, symbol string) (*models.ForeignQuotation, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/foreignQuotation/" + strings.Title(strings.ToLower(source)) + "/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quotation models.ForeignQuotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func updateOracle(
	contract *oracleService.DiaOracle,
	auth *bind.TransactOpts,
	name string,
	symbol string,
	price int64,
	supply int64) error {
	// Write values to smart contract
	tx, err := contract.UpdateCoinInfo(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 800725,
		//	Nonce: big.NewInt(time.Now().Unix()),
	}, name, symbol, big.NewInt(price), big.NewInt(supply), big.NewInt(time.Now().Unix()))
	// prices are with 5 digits after the comma
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("Symbol: %s\n", symbol)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}
