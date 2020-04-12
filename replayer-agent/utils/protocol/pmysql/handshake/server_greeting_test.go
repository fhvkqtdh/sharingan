package handshake

import (
	"bytes"
	"testing"

	"github.com/modern-go/parse"
	"github.com/stretchr/testify/require"
)

func TestDecodeServerGreeting(t *testing.T) {
	var testCase = []struct {
		rawBytes []byte
		expect   *ServerGreeting
		err      error
	}{
		{
			rawBytes: []byte{0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x37, 0x2e, 0x32, 0x30, 0x00,
				0x13, 0x00, 0x00, 0x00, 0x36, 0x17, 0x03, 0x14, 0x0b, 0x11, 0x03, 0x59,
				0x00, 0xff, 0xf7, 0x08, 0x02, 0x00, 0xff, 0x81, 0x15, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x35, 0x36, 0x43, 0x46,
				0x05, 0x19, 0x3c, 0x05, 0x13, 0x27, 0x78, 0x00, 0x6d, 0x79, 0x73, 0x71,
				0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73,
				0x73, 0x77, 0x6f, 0x72, 0x64, 0x00},
			expect: &ServerGreeting{
				ProtocolVersion:            10,
				MySQLVersion:               "5.7.20",
				ThreadID:                   19,
				Salt_1:                     []byte{0x36, 0x17, 0x03, 0x14, 0x0b, 0x11, 0x03, 0x59},
				Capabilities:               0xf7ff,
				ServerLanguage:             0x08,
				StatusFlag:                 0x02,
				ExtendedServerCapabilities: 0x81ff,
				AuthPluginLen:              21,
				Salt_2: []byte{0x31, 0x35, 0x36, 0x43, 0x46, 0x05, 0x19, 0x3c, 0x05, 0x13, 0x27, 0x78,
					0x00},
				AuthPlugin: "mysql_native_password",
			},
			err: nil,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		src, err := parse.NewSource(bytes.NewBuffer(tc.rawBytes), 20)
		should.NoError(err)
		actual, err := DecodeServerGreeting(src)
		should.Equal(tc.err, err, "case #%d fail", idx)
		should.Equal(tc.expect.String(), actual.String(), "case #%d fail", idx)
	}
}
