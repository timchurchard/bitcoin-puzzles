//go:generate mockgen -package internal -destination 32btc-addrs_mock.go -source 32btc-addrs.go
package internal

type PuzzleAddress interface {
	HasAddress(addr string) bool
	IsTestMode() bool
	IsMiniMode() bool
}

// realPuzzleAddress holds real addresses from https://privatekeys.pw/puzzles/bitcoin-puzzle-tx
type realPuzzleAddress struct {
	testMode    bool
	minMode     bool
	overHundred bool
}

func NewRealPuzzleAddress(testMode, minMode, overHundred bool) realPuzzleAddress {
	return realPuzzleAddress{
		testMode:    testMode,
		minMode:     minMode,
		overHundred: overHundred,
	}
}

func (r realPuzzleAddress) IsTestMode() bool {
	return r.testMode
}

func (r realPuzzleAddress) IsMiniMode() bool {
	return r.minMode
}

// HasAddress function holds a list of the bitcoin 000 puzzle with test address and 2^65 onwards
func (r realPuzzleAddress) HasAddress(addr string) bool {
	var exists bool

	if r.minMode {
		// Minimum mode! 66, 67 & 68-bits only.
		minAddrs := map[string]bool{
			"13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so": true,
			"1BY8GQbnueYofwSuFAT3USAhGjPrkxDdW9": true,
			"1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ": true,
		}

		_, exists = minAddrs[addr]
		return exists
	}

	testAddrs := map[string]bool{
		"15JhYXn6Mx3oF4Y7PcTAv2wVVAuCFFQNiP": true, // 2^25 for test     0000000000000000000000000000000000000000000000000000000001fa5ee5
		"16jY7qLJnxb7CHZyqBP8qca9d51gAjyXQN": true, // 2^64 (2022-09-10) 000000000000000000000000000000000000000000000000f7051f27b09112d4
		"18ZMbwUFLMHoZBbfpCjUJQTCMCbktshgpe": true, // 2^65 (2019-06-07) 000000000000000000000000000000000000000000000001a838b13505b26867
	}

	addrs := map[string]bool{
		"13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so": true, // 2^66 from         0000000000000000000000000000000000000000000000020000000000000000
		"1BY8GQbnueYofwSuFAT3USAhGjPrkxDdW9": true,
		"1MVDYgVaSN6iKKEsbzRUAYFrYJadLYZvvZ": true,
		"19vkiEajfhuZ8bs8Zu2jgmC6oqZbWqhxhG": true,
		"1PWo3JeB9jrGwfHDNpdGK54CRas7fsVzXU": true,
		"1JTK7s9YVYywfm5XUH7RNhHJH1LshCaRFR": true,
		"12VVRNPi4SJqUTsp6FmqDqY5sGosDtysn4": true,
		"1FWGcVDK3JGzCC3WtkYetULPszMaK2Jksv": true,
		"1DJh2eHFYQfACPmrvpyWc8MSTYKh7w9eRF": true,
		"1Bxk4CQdqL9p22JEtDfdXMsng1XacifUtE": true,
		"15qF6X51huDjqTmF9BJgxXdt1xcj46Jmhb": true,
		"1ARk8HWJMn8js8tQmGUJeQHjSE7KRkn2t8": true,
		"15qsCm78whspNQFydGJQk5rexzxTQopnHZ": true,
		"13zYrYhhJxp6Ui1VV7pqa5WDhNWM45ARAC": true,
		"14MdEb4eFcT3MVG5sPFG4jGLuHJSnt1Dk2": true,
		"1CMq3SvFcVEcpLMuuH8PUcNiqsK1oicG2D": true,
		"1K3x5L6G57Y494fDqBfrojD28UJv4s5JcK": true,
		"1PxH3K1Shdjb7gSEoTX7UPDZ6SH4qGPrvq": true,
		"16AbnZjZZipwHMkYKBSfswGWKDmXHjEpSf": true,
		"19QciEHbGVNY4hrhfKXmcBBCrJSBZ6TaVt": true,
		"1EzVHtmbN4fs4MiNk3ppEnKKhsmXYJ4s74": true,
		"1AE8NzzgKE7Yhz7BWtAcAAxiFMbPo82NB5": true,
		"17Q7tuG2JwFFU9rXVj3uZqRtioH3mx2Jad": true,
		"1K6xGMUbs6ZTXBnhw1pippqwK6wjBWtNpL": true,
		"15ANYzzCp5BFHcCnVFzXqyibpzgPLWaD8b": true,
		"18ywPwj39nGjqBrQJSzZVq2izR12MDpDr8": true,
		"1CaBVPrwUxbQYYswu32w7Mj4HR4maNoJSX": true,
		"1JWnE6p6UN7ZJBN7TtcbNDoRcjFtuDWoNL": true,
	}

	hundredAddrs := map[string]bool{
		"1CKCVdbDJasYmhswB6HKZHEAnNaDpK7W4n": true, // 2^100
		"1PXv28YxmYMaB8zxrKeZBW8dt2HK7RkRPX": true,
		"1AcAmB6jmtU6AiEcXkmiNE9TNVPsj9DULf": true,
		"1EQJvpsmhazYCcKX5Au6AZmZKRnzarMVZu": true,
		"18KsfuHuzQaBTNLASyj15hy4LuqPUo1FNB": true,
		"15EJFC5ZTs9nhsdvSUeBXjLAuYq3SWaxTc": true,
		"1HB1iKUqeffnVsvQsbpC6dNi1XKbyNuqao": true,
		"1GvgAXVCbA8FBjXfWiAms4ytFeJcKsoyhL": true,
		"1824ZJQ7nKJ9QFTRBqn7z7dHV5EGpzUpH3": true,
		"18A7NA9FTsnJxWgkoFfPAFbQzuQxpRtCos": true,
		"1NeGn21dUDDeqFQ63xb2SpgUuXuBLA4WT4": true,
		"174SNxfqpdMGYy5YQcfLbSTK3MRNZEePoy": true,
		"1MnJ6hdhvK37VLmqcdEwqC3iFxyWH2PHUV": true,
		"1KNRfGWw7Q9Rmwsc6NT5zsdvEb9M2Wkj5Z": true,
		"1PJZPzvGX19a7twf5HyD2VvNiPdHLzm9F6": true,
		"1GuBBhf61rnvRe4K8zu8vdQB3kHzwFqSy7": true,
		"17s2b9ksz5y7abUm92cHwG8jEPCzK3dLnT": true,
		"1GDSuiThEV64c166LUFC9uDcVdGjqkxKyh": true,
		"1Me3ASYt5JCTAK2XaC32RMeH34PdprrfDx": true,
		"1CdufMQL892A69KXgv6UNBD17ywWqYpKut": true,
		"1BkkGsX9ZM6iwL3zbqs7HWBV7SvosR6m8N": true,
		"1PXAyUB8ZoH3WD8n5zoAthYjN15yN5CVq5": true,
		"1AWCLZAjKbV1P7AHvaPNCKiB7ZWVDMxFiz": true,
		"1G6EFyBRU86sThN3SSt3GrHu1sA7w7nzi4": true,
		"1MZ2L1gFrCtkkn6DnTT2e4PFUTHw9gNwaj": true,
		"1Hz3uv3nNZzBVMXLGadCucgjiCs5W9vaGz": true,
		"1Fo65aKq8s8iquMt6weF1rku1moWVEd5Ua": true,
		"16zRPnT8znwq42q7XeMkZUhb1bKqgRogyy": true,
		"1KrU4dHE5WrW8rhWDsTRjR21r8t3dsrS3R": true,
		"17uDfp5r4n441xkgLFmhNoSW1KWp6xVLD":  true,
		"13A3JrvXmvg5w9XGvyyR4JEJqiLz8ZySY3": true,
		"16RGFo6hjq9ym6Pj7N5H7L1NR1rVPJyw2v": true,
		"1UDHPdovvR985NrWSkdWQDEQ1xuRiTALq":  true,
		"15nf31J46iLuK1ZkTnqHo7WgN5cARFK3RA": true,
		"1Ab4vzG6wEQBDNQM1B2bvUz4fqXXdFk2WT": true,
		"1Fz63c775VV9fNyj25d9Xfw3YHE6sKCxbt": true,
		"1QKBaU6WAeycb3DbKbLBkX7vJiaS8r42Xo": true,
		"1CD91Vm97mLQvXhrnoMChhJx4TP9MaQkJo": true,
		"15MnK2jXPqTMURX4xC3h4mAZxyCcaWWEDD": true,
		"13N66gCzWWHEZBxhVxG18P8wyjEWF9Yoi1": true,
		"1NevxKDYuDcCh1ZMMi6ftmWwGrZKC6j7Ux": true,
		"19GpszRNUej5yYqxXoLnbZWKew3KdVLkXg": true,
		"1M7ipcdYHey2Y5RZM34MBbpugghmjaV89P": true,
		"18aNhurEAJsw6BAgtANpexk5ob1aGTwSeL": true,
		"1FwZXt6EpRT7Fkndzv6K4b4DFoT4trbMrV": true,
		"1CXvTzR6qv8wJ7eprzUKeWxyGcHwDYP1i2": true,
		"1MUJSJYtGPVGkBCTqGspnxyHahpt5Te8jy": true,
		"13Q84TNNvgcL3HJiqQPvyBb9m4hxjS3jkV": true,
		"1LuUHyrQr8PKSvbcY1v1PiuGuqFjWpDumN": true,
		"18192XpzzdDi2K11QVHR7td2HcPS6Qs5vg": true,
		"1NgVmsCCJaKLzGyKLFJfVequnFW9ZvnMLN": true,
		"1AoeP37TmHdFh8uN72fu9AqgtLrUwcv2wJ": true,
		"1FTpAbQa4h8trvhQXjXnmNhqdiGBd1oraE": true,
		"14JHoRAdmJg3XR4RjMDh6Wed6ft6hzbQe9": true,
		"19z6waranEf8CcP8FqNgdwUe1QRxvUNKBG": true,
		"14u4nA5sugaswb6SZgn5av2vuChdMnD9E5": true,
		"1NBC8uXJy1GiJ6drkiZa1WuKn51ps7EPTv": true, // 2^160
	}

	if r.testMode {
		_, exists = testAddrs[addr]
		if exists {
			return true
		}
	}

	if r.overHundred {
		_, exists = hundredAddrs[addr]
		if exists {
			return true
		}
	}

	_, exists = addrs[addr]
	return exists
}
