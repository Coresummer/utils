package main

import (
	"fmt"
	"testing"

	"github.com/Coresummer/utils"
)

func TestUintToLittleEndianByteArray(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "uint16",
			args: args{v: uint16(111)},
			want: "[111 0]",
		},
		{
			name: "uint32",
			args: args{v: uint16(1111)},
			want: "[87 4]",
		},
		{
			name: "uint64",
			args: args{v: uint16(111)},
			want: "[111 0]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmt.Sprint(utils.UintConvertToLittleEndianByteArray(tt.args.v)); got != tt.want {
				t.Errorf("UintToByte() = %v, want %v", got, tt.want)
			} else {
				println("Passed")
			}
		})
	}
}

func TestLittleEndianByteArrayToUint(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args []byte
		want int
	}{
		{
			name: "uint16",
			args: []byte("111"),
			want: 111,
		},
		{
			name: "uint32",
			args: []byte("874"),
			want: 1111,
		},
		{
			name: "uint64",
			args: []byte("1110"),
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ByteArrayConvertToUint(tt.args); got != tt.want {
				t.Errorf("UintToByte() = %v, want %v", got, tt.want)
			} else {
				println("Passed")
			}
		})
	}
}
