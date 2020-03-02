package protosmart

import (
	"reflect"

	"google.golang.org/grpc/encoding"
)

// This codec overrides the proto codec and if you are passing []byte it assumes it's already been encoded and
// will not attempt to re-encode the data

type codec struct {
	name           string
	overridenCodec encoding.Codec
}

var protoSmartInstalled bool

func OverrideCodec(name string) {
	// This will panic if the codec doesn't exist, what you gonna do?!
	overriddenCodec := encoding.GetCodec(name)

	encoding.RegisterCodec(&codec{
		overridenCodec: overriddenCodec,
		name:           name,
	})

	// Global just to let other libraries know that depend on this that it's registered
	protoSmartInstalled = true
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	// If it's already bytes, don't try to remarshall it
	if b, ok := v.([]byte); ok {
		return b, nil
	}
	return c.overridenCodec.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	// If you provided a *[]byte, don't unmarshal it, just copy it
	if _, ok := v.(*[]byte); ok {
		reflect.ValueOf(v).Elem().SetBytes(data)
		return nil
	}
	return c.overridenCodec.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return c.name
}
