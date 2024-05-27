package caller

import "testing"

var testNameVal = Name()

var testNameCb = func() string {
	return Name()
}()

func testNameFn() string {
	return Name()
}

type testNameObj struct{}

func (testNameObj) Fn() string {
	return Name()
}

func (testNameObj) FnCb() string {
	return func() string {
		return Name()
	}()
}

func (*testNameObj) PointerFn() string {
	return Name()
}

func (*testNameObj) PointerFnCb() string {
	return func() string {
		return Name()
	}()
}

func TestName(t *testing.T) {
	obj := testNameObj{}
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "pkg call",
			got:  testNameVal,
			want: "caller",
		},
		{
			name: "pkg call cb",
			got:  testNameCb,
			want: "caller",
		},
		{
			name: "call",
			got:  Name(),
			want: "caller.TestName",
		},
		{
			name: "func",
			got:  testNameFn(),
			want: "caller.testNameFn",
		},
		{
			name: "object func",
			got:  obj.Fn(),
			want: "caller.testNameObj.Fn",
		},
		{
			name: "object func",
			got:  obj.FnCb(),
			want: "caller.testNameObj.FnCb",
		},
		{
			name: "object pointer func",
			got:  obj.PointerFn(),
			want: "caller.testNameObj.PointerFn",
		},
		{
			name: "object pointer func cb",
			got:  obj.PointerFnCb(),
			want: "caller.testNameObj.PointerFnCb",
		},
		{
			name: "call back call of object pointer func",
			got: func() string {
				return obj.PointerFn()
			}(),
			want: "caller.testNameObj.PointerFn",
		},
		{
			name: "call back call of object pointer func cb",
			got: func() string {
				return obj.PointerFnCb()
			}(),
			want: "caller.testNameObj.PointerFnCb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}
