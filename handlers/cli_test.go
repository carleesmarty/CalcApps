package handlers

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/carleesmarty/CalcLib"
)

func TestAddition(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zeros",
			args: args{
				a: 0,
				b: 0,
			},
			want: 0,
		},

		{
			name: "negative",
			args: args{
				a: -1,
				b: 0,
			},
			want: -1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a2 := CalcLib.Addition{}
			if got := a2.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_Handle(t *testing.T) {
	type fields struct {
		calculator CalcLib.Calculator
		out        io.Writer
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "double",
			fields: fields{
				calculator: CalcLib.Addition{},
			},
			args: args{
				args: []string{"add", "5", "5"},
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				calculator: tt.fields.calculator,
				out:        tt.fields.out,
			}
			if err := h.Handle(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		calculator CalcLib.Calculator
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		want    *Handler
	}{
		{
			name: "double",
			args: args{
				calculator: CalcLib.Addition{},
			},
			wantOut: "10",
			want:    &Handler{calculator: CalcLib.Addition{}},
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got := NewHandler(tt.args.calculator, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("NewHandler() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
