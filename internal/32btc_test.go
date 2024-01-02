package internal

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRandomKey(t *testing.T) {
	res1 := randomKey(true)
	zeroPrefix1 := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	res2 := randomKey(false)
	zeroPrefix2 := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00,
	}

	assert.NotEqual(t, res1, res2)
	assert.True(t, bytes.HasPrefix(res1, zeroPrefix1))
	assert.True(t, bytes.HasPrefix(res2, zeroPrefix2))
}

func Test_numToHex(t *testing.T) {
	type args struct {
		current *big.Int
	}

	num1, _ := new(big.Int).SetString("80085", 10)
	num2, _ := new(big.Int).SetString("12345", 10)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sanity 1",
			args: args{
				current: num1,
			},
			want: "00000000000000000000000000000000000000000000000000000000000138d5",
		},
		{
			name: "sanity 2",
			args: args{
				current: num2,
			},
			want: "0000000000000000000000000000000000000000000000000000000000003039",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numToHex(tt.args.current); got != tt.want {
				t.Errorf("numToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToNum(t *testing.T) {
	type args struct {
		hexstr string
	}

	num1, _ := new(big.Int).SetString("80085", 10)
	num2, _ := new(big.Int).SetString("12345", 10)

	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "80085",
			args: args{
				hexstr: "00000000000000000000000000000000000000000000000000000000000138d5",
			},
			want: num1,
		},
		{
			name: "12345",
			args: args{
				hexstr: "0000000000000000000000000000000000000000000000000000000000003039",
			},
			want: num2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexToNum(tt.args.hexstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hexToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPuzzlePriv(t *testing.T) {
	t.Run("mock true", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPD := NewMockPuzzleAddress(ctrl)
		mockPD.EXPECT().IsTestMode().Return(false)
		mockPD.EXPECT().IsMiniMode().Return(true)
		mockPD.EXPECT().HasAddress("161nAtMQjUrJNKG31mZvigsTRLUFbawseB").Return(false)
		mockPD.EXPECT().HasAddress("1HqajQQ1LgmmbSg54DKaYntvVrjoTR9dLf").Return(false)
		mockPD.EXPECT().HasAddress("13rwmSfUHN2jJ9LqXcKSrKFhPSecw8T1NH").Return(true)

		testKey, _ := hex.DecodeString("000000000000000000000000000000000000000000000000deadbeefdeadbeef")

		assert.True(t, checkPuzzlePriv(testKey, mockPD))
	})

	t.Run("mock true", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPD := NewMockPuzzleAddress(ctrl)
		mockPD.EXPECT().IsTestMode().Return(false)
		mockPD.EXPECT().IsMiniMode().Return(true)
		mockPD.EXPECT().HasAddress("161nAtMQjUrJNKG31mZvigsTRLUFbawseB").Return(false)
		mockPD.EXPECT().HasAddress("1HqajQQ1LgmmbSg54DKaYntvVrjoTR9dLf").Return(false)
		mockPD.EXPECT().HasAddress("13rwmSfUHN2jJ9LqXcKSrKFhPSecw8T1NH").Return(true)

		testKey, _ := hex.DecodeString("000000000000000000000000000000000000000000000000deadbeefdeadbeef")

		assert.True(t, checkPuzzlePriv(testKey, mockPD))
	})

	t.Run("real (minmode) false", func(t *testing.T) {
		pd := NewRealPuzzleAddress(false, true, false)

		testKey, _ := hex.DecodeString("000000000000000000000000000000000000000000000000deadbeefdeadbeef")

		assert.False(t, checkPuzzlePriv(testKey, pd)) // Note: result false as < 66-bit not supported
	})

	t.Run("real testmode", func(t *testing.T) {
		pd := NewRealPuzzleAddress(true, false, false)

		testKey65, _ := hex.DecodeString("000000000000000000000000000000000000000000000001a838b13505b26867")

		assert.True(t, checkPuzzlePriv(testKey65, pd))
	})
}
