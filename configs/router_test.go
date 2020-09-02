package configs

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateRouter(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunServer(t *testing.T) {
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunServer(tt.args.router)
		})
	}
}
