package internal

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/piotrnar/gocoin/lib/btc"
	"github.com/sour-is/bitcoin/address"
	opendime "github.com/timchurchard/opendime-utils/pkg"
	"go.step.sm/crypto/randutil"
	"golang.org/x/crypto/scrypt"
)

// letters used by Ballet wallet passwords
const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Bobby(ch chan bool, num, worksize int) {
	// Note: https://www.takebobbysbitcoin.com/
	const (
		realAddress = "1JxWyNrkgYvgsHu8hVQZqTXEB9RftRGP5m"
		realPriv    = "6PnQmAyBky9ZXJyZBv9QSGRUXkKh9HfnVsZWPn4YtcwoKy5vufUgfA3Ld7"
	)

	var (
		result string
		err    error
	)

	fmt.Printf("Bobby %d started\n", num)

	start := time.Now()
	count := 0
	loops := 0
	lenLetters := len(letters)

	for {
		randomBase := fmt.Sprintf("%s-%s-%s-%s", randStringBobby(4), randStringBobby(4), randStringBobby(4), randStringBobby(4))

		randomStart := randStringBobby(4)
		a := strings.Index(letters, string(randomStart[0]))
		b := strings.Index(letters, string(randomStart[1]))
		c := strings.Index(letters, string(randomStart[2]))
		d := strings.Index(letters, string(randomStart[3]))

		for {
			c++
			if c >= lenLetters {
				break
			}

			for {
				d++
				if d >= lenLetters {
					d = 0
					break
				}

				randomPassword := fmt.Sprintf("%s-%s", randomBase, numStringBobby(a, b, c, d))

				result, err = tryBip38Password(realAddress, realPriv, randomPassword)
				if err == nil {
					fmt.Printf("-----\nFound! %s = %s\n", randomPassword, result)

					ch <- true
					return
				}

				count++
				if count%worksize == 0 {
					loops++
					duration := time.Since(start)

					fmt.Printf("Bobby %d count %d in %f secs (%f/s)\t- %s\n", num, count, duration.Seconds(), float64(worksize)/duration.Seconds(), randomPassword)

					count = 0
					start = time.Now()
				}
			}
		}
	}
}

func numStringBobby(a, b, c, d int) string {
	return fmt.Sprintf("%s%s%s%s", string(letters[a]), string(letters[b]), string(letters[c]), string(letters[d]))
}

func randStringBobby(i int) string {
	result, err := randutil.String(i, letters)
	if err != nil {
		panic(err) // TODO!
	}

	return result
}

func tryBip38Password(addr, bip38wif, password string) (string, error) {
	const prefixBitcoinHex = "80"

	secretExponentBytes, err := bip38Decrypt(bip38wif, password)
	if err != nil {
		fmt.Printf("Error on bip38.Decrypt: %v", err)
		return "", err
	}

	if len(secretExponentBytes) == 0 {
		return "", errors.New("wrong password empty secret")
	}

	// fmt.Printf("Got here with secretExponentBytes = %v\n", hex.EncodeToString(secretExponentBytes))

	priv := new(address.PrivateKey)
	priv.SetBytes(secretExponentBytes)

	// fmt.Printf("Got here with priv.PublicKeyHex = %v\n", hex.EncodeToString(priv.PublicKey.Bytes()))

	addresses, _ := opendime.GetAddresses(opendime.VerifiedMessage{
		PublicKeyHex: priv.PublicKey.String(), // todo: Bit hacky. GetAddresses only uses this field currently.
	})
	if addr != addresses.BitcoinP2PKH &&
		addr != addresses.BitcoinP2PKHCompressed {
		return "", errors.New("probably wrong password")
	}

	p2pkh := opendime.ToWif(prefixBitcoinHex, hex.EncodeToString(secretExponentBytes), false)
	p2pkhComp := opendime.ToWif(prefixBitcoinHex, hex.EncodeToString(secretExponentBytes), true)

	if addr == addresses.BitcoinP2PKH {
		return p2pkh, nil
	}

	return p2pkhComp, nil
}

func bip38Decrypt(bip38Wif string, password string) ([]byte, error) {
	dec, err := address.FromBase58(bip38Wif)
	if err != nil {
		return nil, err
	}

	if dec[0] == 0x01 && dec[1] == 0x42 {
		return nil, errors.New("TODO: implement decryption when EC multiply mode not used")
	} else if dec[0] == 0x01 && dec[1] == 0x43 {
		compress := dec[2]&0x20 == 0x20
		hasLotSequence := dec[2]&0x04 == 0x04

		var ownerSalt, ownerEntropy []byte
		if hasLotSequence {
			ownerSalt = dec[7:11]
			ownerEntropy = dec[7:15]
		} else {
			ownerSalt = dec[7:15]
			ownerEntropy = ownerSalt
		}

		prefactorA, err := scrypt.Key([]byte(password), ownerSalt, 16384, 8, 8, 32)
		if prefactorA == nil {
			return nil, err
		}

		var passFactor []byte
		if hasLotSequence {
			prefactorB := bytes.Join([][]byte{prefactorA, ownerEntropy}, nil)
			passFactor = sha256Twice(prefactorB)
		} else {
			passFactor = prefactorA
		}

		passpoint := btc.PublicFromPrivate(passFactor, true)
		if passpoint == nil {
			return nil, err
		}

		encryptedpart1 := dec[15:23]
		encryptedpart2 := dec[23:39]

		derived, err := scrypt.Key(passpoint, bytes.Join([][]byte{dec[3:7], ownerEntropy}, nil), 1024, 1, 1, 64)
		if derived == nil {
			return nil, err
		}

		h, err := aes.NewCipher(derived[32:])
		if h == nil {
			return nil, err
		}

		unencryptedpart2 := make([]byte, 16)
		h.Decrypt(unencryptedpart2, encryptedpart2)
		for i := range unencryptedpart2 {
			unencryptedpart2[i] ^= derived[i+16]
		}

		encryptedpart1 = bytes.Join([][]byte{encryptedpart1, unencryptedpart2[:8]}, nil)

		unencryptedpart1 := make([]byte, 16)
		h.Decrypt(unencryptedpart1, encryptedpart1)
		for i := range unencryptedpart1 {
			unencryptedpart1[i] ^= derived[i]
		}

		seeddb := bytes.Join([][]byte{unencryptedpart1[:16], unencryptedpart2[8:]}, nil)
		factorb := sha256Twice(seeddb)

		bigN, success := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
		if !success {
			return nil, errors.New("Failed to create Int for N")
		}

		passFactorBig := new(big.Int).SetBytes(passFactor)
		factorbBig := new(big.Int).SetBytes(factorb)

		privKey := new(big.Int)
		privKey.Mul(passFactorBig, factorbBig)
		privKey.Mod(privKey, bigN)

		pubKey := btc.PublicFromPrivate(privKey.Bytes(), compress)
		if pubKey == nil {
			return nil, err
		}

		addr := btc.NewAddrFromPubkey(pubKey, 0).String()

		addrHashed := sha256Twice([]byte(addr))

		if addrHashed[0] != dec[3] || addrHashed[1] != dec[4] || addrHashed[2] != dec[5] || addrHashed[3] != dec[6] {
			return nil, err
		}

		return privKey.Bytes(), nil
	}

	return nil, errors.New("Unreachable")
}

func sha256Twice(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	hashedOnce := h.Sum(nil)
	h.Reset()
	h.Write(hashedOnce)
	return h.Sum(nil)
}
