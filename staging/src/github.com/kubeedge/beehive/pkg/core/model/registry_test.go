package model

import (
	"reflect"
	"testing"
)

type T1 struct {
	Arg1 string
}

type T2 struct {
	Arg1 int
	Arg2 string
}

type T3 struct {
	Arg1 string
}

type Tnotregistrer struct {
	Arg1 int
	Arg2 string
}

var testRegistry = newRegistry()

func init() {
	testRegistry.Register("T1", T1{})
	testRegistry.Register("T2", T2{})
	testRegistry.Register("T3", &T3{})
}

func Test_registry_GetName(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		args   args
		want   string
		wantOk bool
	}{
		{
			args:   args{},
			wantOk: false,
		},
		{
			args: args{
				T1{},
			},
			want:   "T1",
			wantOk: true,
		},
		{
			args: args{
				T2{},
			},
			want:   "T2",
			wantOk: true,
		},
		{
			args: args{
				T3{},
			},
			want:   "T3",
			wantOk: true,
		},
		{
			args: args{
				&T1{},
			},
			want:   "T1",
			wantOk: true,
		},
		{
			args: args{
				&T2{},
			},
			want:   "T2",
			wantOk: true,
		},
		{
			args: args{
				&T3{},
			},
			want:   "T3",
			wantOk: true,
		},
		{
			args: args{
				Tnotregistrer{},
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := testRegistry.GetName(tt.args.val)
			if got != tt.want {
				t.Errorf("GetName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("GetName() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func Test_registry_New(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			args: args{
				name: "T1",
			},
			want: &T1{},
		},
		{
			args: args{
				name: "T2",
			},
			want: &T2{},
		},
		{
			args: args{
				name: "T3",
			},
			want: &T3{},
		},
		{
			args: args{
				name: "Tnotregistrer",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testRegistry.New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
