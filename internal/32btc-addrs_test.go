package internal

import (
	"reflect"
	"testing"
)

func Test_realPuzzleAddress_HasAddress(t *testing.T) {
	type fields struct {
		testMode    bool
		minMode     bool
		overHundred bool
	}

	type args struct {
		addr string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test mode",
			fields: fields{
				testMode:    true,
				minMode:     false,
				overHundred: false,
			},
			args: args{
				addr: "15JhYXn6Mx3oF4Y7PcTAv2wVVAuCFFQNiP",
			},
			want: true,
		},
		{
			name: "normal mode",
			fields: fields{
				testMode:    false,
				minMode:     false,
				overHundred: false,
			},
			args: args{
				addr: "13zb1hQbWVsc2S7ZTZnP2G4undNNpdh5so",
			},
			want: true,
		},
		{
			name: "normal mode false",
			fields: fields{
				testMode:    false,
				minMode:     false,
				overHundred: false,
			},
			args: args{
				addr: "1AAA",
			},
			want: false,
		},
		{
			name: "hundred",
			fields: fields{
				testMode:    false,
				minMode:     false,
				overHundred: true,
			},
			args: args{
				addr: "1PXv28YxmYMaB8zxrKeZBW8dt2HK7RkRPX",
			},
			want: true,
		},
		{
			name: "min mode",
			fields: fields{
				testMode:    false,
				minMode:     true,
				overHundred: false,
			},
			args: args{
				addr: "not found",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := realPuzzleAddress{
				testMode:    tt.fields.testMode,
				overHundred: tt.fields.overHundred,
			}
			if got := r.HasAddress(tt.args.addr); got != tt.want {
				t.Errorf("HasAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRealPuzzleAddress(t *testing.T) {
	type args struct {
		testMode    bool
		minMode     bool
		overHundred bool
	}

	tests := []struct {
		name string
		args args
		want realPuzzleAddress
	}{
		{
			name: "",
			args: args{
				testMode:    true,
				minMode:     false,
				overHundred: true,
			},
			want: realPuzzleAddress{
				testMode:    true,
				minMode:     false,
				overHundred: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRealPuzzleAddress(tt.args.testMode, tt.args.minMode, tt.args.overHundred); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRealPuzzleAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
