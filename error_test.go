package caller

import (
	"errors"
	"testing"
)

var testWrapVal = Wrap(errors.New("test"))

var testWrapCb = func() error {
	return Wrap(errors.New("test"))
}()

func testWrapFn() error {
	return Wrap(errors.New("test"))
}

type testWrapObj struct{}

func (testWrapObj) Fn() error {
	return Wrap(errors.New("test"))
}

func (testWrapObj) FnCb() error {
	return func() error {
		return Wrap(errors.New("test"))
	}()
}

func (*testWrapObj) PointerFn() error {
	return Wrap(errors.New("test"))
}

func (*testWrapObj) PointerFnCb() error {
	return func() error {
		return Wrap(errors.New("test"))
	}()
}

func TestWrap(t *testing.T) {
	obj := testWrapObj{}
	tests := []struct {
		name string
		got  error
		want string
	}{
		{
			name: "pkg call",
			got:  testWrapVal,
			want: "caller: test",
		},
		{
			name: "pkg call cb",
			got:  testWrapCb,
			want: "caller: test",
		},
		{
			name: "call",
			got:  Wrap(errors.New("test")),
			want: "caller.TestWrap: test",
		},
		{
			name: "func",
			got:  testWrapFn(),
			want: "caller.testWrapFn: test",
		},
		{
			name: "object func",
			got:  obj.Fn(),
			want: "caller.testWrapObj.Fn: test",
		},
		{
			name: "object func",
			got:  obj.FnCb(),
			want: "caller.testWrapObj.FnCb: test",
		},
		{
			name: "object pointer func",
			got:  obj.PointerFn(),
			want: "caller.testWrapObj.PointerFn: test",
		},
		{
			name: "object pointer func cb",
			got:  obj.PointerFnCb(),
			want: "caller.testWrapObj.PointerFnCb: test",
		},
		{
			name: "call back call of object pointer func",
			got: func() error {
				return obj.PointerFn()
			}(),
			want: "caller.testWrapObj.PointerFn: test",
		},
		{
			name: "call back call of object pointer func cb",
			got: func() error {
				return obj.PointerFnCb()
			}(),
			want: "caller.testWrapObj.PointerFnCb: test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.Error() != tt.want {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}
