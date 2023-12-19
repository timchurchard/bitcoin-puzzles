package internal

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func BTC32Worker(ch chan bool, num, countnum, worksize int, miniMode bool) {
	var privKeyBytes []byte

	pd := NewRealPuzzleAddress(false, miniMode, false)

	start := time.Now()
	count := 0

	oneBigInt := new(big.Int)
	oneBigInt.SetString("1", 10)

	startBigInt := hexToNum(hex.EncodeToString(randomKey(pd.IsMiniMode())))
	privKeyBytes, _ = hex.DecodeString(numToHex(startBigInt))

	for i := 0; i < countnum; i++ {
		if checkPuzzlePriv(privKeyBytes, pd) {
			fmt.Printf("FOUND! checkPuzzlePriv: hex=%s\n", hex.EncodeToString(privKeyBytes))

			ch <- true
		}

		// Stats!
		count++
		if count%worksize == 0 {
			duration := time.Since(start)
			start = time.Now()

			fmt.Printf("BTC32Worker (num %d) count %d/%d in %f secs (%f/s)\t\t - %s\n",
				num, count, countnum, duration.Seconds(), float64(worksize)/duration.Seconds(), numToHex(startBigInt))
		}

		// Next!
		startBigInt = new(big.Int).Add(startBigInt, oneBigInt)
		privKeyBytes, _ = hex.DecodeString(numToHex(startBigInt))
	}

	// Unblock the channel to indicate worker finished (not found)
	ch <- false
}

// randomKey returns a random key
// If miniMode then between 0x00 and 0xfffffffffffffffff for the 66, 67 and 68-bit ranges
// Else random for 99-bits
func randomKey(miniMode bool) []byte {
	if miniMode {
		// fffffffffffffffff
		buf := make([]byte, 9)
		_, _ = rand.Read(buf)

		return []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, buf[0],
			buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7], buf[8],
		}
	} else {
		// 7ffffffffffffffffffffffff
		buf := make([]byte, 13)
		_, _ = rand.Read(buf)

		return []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, buf[0], buf[1], buf[2], buf[3], buf[4],
			buf[5], buf[6], buf[7], buf[8], buf[9], buf[10], buf[11], buf[12],
		}
	}
}

func checkPuzzlePriv(privKeyBytes []byte, pd PuzzleAddress) bool {
	var (
		firstByte int
		testBytes []byte
	)

	if pd.IsMiniMode() {
		// 66, 67 & 68-bit range 20000000000000000..fffffffffffffffff
		firstByte = 23
		testBytes = []byte{0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	} else {
		// 1ffffffffffffffffffffffff..7ffffffffffffffffffffffff
		firstByte = 19
		testBytes = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	}

	for idx := range testBytes {
		privKeyBytes[firstByte] = testBytes[idx]

		if checkSinglePriv(pd, privKeyBytes) {
			return true
		}
	}

	return false
}

// checkSinglePriv checks a single private key
func checkSinglePriv(pd PuzzleAddress, privKeyBytes []byte) bool {
	_, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
	pubKeyCmp := pubKey.SerializeCompressed()

	addrCmp, _ := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKeyCmp), &chaincfg.MainNetParams)
	addrCmpStr := addrCmp.String()

	return pd.HasAddress(addrCmpStr)
}

func numToHex(current *big.Int) string {
	result := fmt.Sprintf("%064x", current)
	if resultLen := len(result); resultLen > 64 {
		result = result[resultLen-64:]
	}
	// fmt.Printf("DEBUG Private Exponent Hex: %s\n", result)
	return result
}

func hexToNum(hexstr string) *big.Int {
	i := new(big.Int)
	i.SetString(hexstr, 16)
	// fmt.Printf("DEBUG hexToNum returning %s -> %v\n", hexstr, i)
	return i
}
