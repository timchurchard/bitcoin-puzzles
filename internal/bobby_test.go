package internal

import "testing"

func TestTryBip38Password(t *testing.T) {
	type args struct {
		addr     string
		bip38wif string
		password string
	}

	const (
		realAddress = "1JxWyNrkgYvgsHu8hVQZqTXEB9RftRGP5m"
		realPriv    = "6PnQmAyBky9ZXJyZBv9QSGRUXkKh9HfnVsZWPn4YtcwoKy5vufUgfA3Ld7"

		testAddress = "1AKX5vvjvF5m8UdyPTGzwDSJ4DMX2eb6h9"
		testPriv    = "6PnSatcySnJjpJDLWkM2nXMB4PZoRVaTSzeg3Y2mBNQs2FTPCrsDXBv35X"
		testPass    = "240W-KNTR-LGBD-7LPY-A7C0"
		testResult  = "L3DSS6LSxCe6EFMKc3jNTkVjcB5fu6cVD95oto8fZeyrANKjagAZ" // b2e279ef8e47f0e35c94bb319c311b5874032d69d5859b5e184f09ec2dc34890"

		// From twitter https://twitter.com/bobbyclee/status/1289011807160397824
		anotherAddress = "1LL6Xy92LwGDRfQP9fBU7f1477cEKctr7c"
		anotherPriv    = "6PnWfKaBfDW6mFFhhFsbNRHnVgojUhdf2b5NXP3FfwXiQ69MxEzVK2J4cH"
		anotherPass    = "335Y-K745-C8WT-4D2W-80WP"
		anotherResult  = "L2mrYyo5a6rpyQdC88UitNeH5n1rAqPcq8Qv5gwtQE8KTvW3ZTeH"
	)

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				addr:     testAddress,
				bip38wif: testPriv,
				password: testPass,
			},
			want:    testResult,
			wantErr: false,
		},
		{
			name: "another",
			args: args{
				addr:     anotherAddress,
				bip38wif: anotherPriv,
				password: anotherPass,
			},
			want:    anotherResult,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				addr:     realAddress,
				bip38wif: realPriv,
				password: anotherPriv,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tryBip38Password(tt.args.addr, tt.args.bip38wif, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("tryBip38Password() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tryBip38Password() got = %v, want %v", got, tt.want)
			}
		})
	}
}
