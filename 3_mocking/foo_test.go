package module3

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestBar_sumNew(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockFoo := NewMockFooInterface(ctrl)

	type fields struct {
		foo FooInterface
	}
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Success", fields{&FooImpl{
			message1: "Message 1",
			message2: "Message 2",
		}}, args{a: 5, b: 7}, "Message 1"},

		{"SuccessMock1", fields{foo: mockFoo}, args{a: 5, b: 7}, "Message 2"},
		{"SuccessMock2", fields{foo: mockFoo}, args{a: 5, b: 7}, "Message 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bar := &Bar{
				foo: tt.fields.foo,
			}

			switch tt.name {
			case "SuccessMock1":
				mockFoo.EXPECT().sum(int64(5), int64(7)).Return(int64(12))
				mockFoo.EXPECT().getMessage(int64(1)).Return("Message 2")

			case "SuccessMock2":
				mockFoo.EXPECT().sum(int64(5), int64(7)).Return(int64(-3))
				mockFoo.EXPECT().getMessage(int64(-1)).Return("Message 2")

			}

			if got := bar.sumNew(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Bar.sumNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
