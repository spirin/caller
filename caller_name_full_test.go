package caller

import "testing"

var testNameFullVal = NameFull()

var testNameFullCb = func() string {
	return NameFull()
}()

func testNameFullFn() string {
	return NameFull()
}

type testNameFullObj struct{}

func (testNameFullObj) Fn() string {
	return NameFull()
}

func (testNameFullObj) FnCb() string {
	return func() string {
		return NameFull()
	}()
}

func (*testNameFullObj) PointerFn() string {
	return NameFull()
}

func (*testNameFullObj) PointerFnCb() string {
	return func() string {
		return NameFull()
	}()
}

func TestNameFull(t *testing.T) {
	obj := testNameFullObj{}
	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "pkg call",
			got:  testNameFullVal,
			want: "github.com/spirin/caller",
		},
		{
			name: "pkg call cb",
			got:  testNameFullCb,
			want: "github.com/spirin/caller",
		},
		{
			name: "call",
			got:  NameFull(),
			want: "TestNameFull @ github.com/spirin/caller",
		},
		{
			name: "func",
			got:  testNameFullFn(),
			want: "testNameFullFn @ github.com/spirin/caller",
		},
		{
			name: "object func",
			got:  obj.Fn(),
			want: "testNameFullObj.Fn @ github.com/spirin/caller",
		},
		{
			name: "object func",
			got:  obj.FnCb(),
			want: "testNameFullObj.FnCb @ github.com/spirin/caller",
		},
		{
			name: "object pointer func",
			got:  obj.PointerFn(),
			want: "testNameFullObj.PointerFn @ github.com/spirin/caller",
		},
		{
			name: "object pointer func cb",
			got:  obj.PointerFnCb(),
			want: "testNameFullObj.PointerFnCb @ github.com/spirin/caller",
		},
		{
			name: "call back call of object pointer func",
			got: func() string {
				return obj.PointerFn()
			}(),
			want: "testNameFullObj.PointerFn @ github.com/spirin/caller",
		},
		{
			name: "call back call of object pointer func cb",
			got: func() string {
				return obj.PointerFnCb()
			}(),
			want: "testNameFullObj.PointerFnCb @ github.com/spirin/caller",
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
