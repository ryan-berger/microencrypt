package microencrypt

import (
	"testing"
	"net/http"
	"io/ioutil"
	"bytes"
)

type HttpWriteTester struct {
	test          *testing.T
	onWriteHeader func(int)
}

func (writeTest *HttpWriteTester) Header() http.Header {
	return http.Header{}
}

func (writeTest *HttpWriteTester) Write([]byte) (int, error) {
	return 0, nil
}

func (writeTest *HttpWriteTester) WriteHeader(header int) {
	writeTest.onWriteHeader(header)
}

func TestAsymmetricEncrypt_ServeHTTPBadKey(t *testing.T) {
	asymmetricEncrypt := NewMicroEncrypt("")
	tester := &HttpWriteTester{
		test: t,
		onWriteHeader: func(header int) {
			if header != 500 {
				t.Errorf("Key is not valid, throw 500")
			}
		},
	}
	asymmetricEncrypt.ServeHTTP(&http.Request{Body: ioutil.NopCloser(bytes.NewReader([]byte("HI")))}, tester)
}

func TestAsymmetricEncrypt_ServeHTTPEmpty(t *testing.T) {
	asymmetricEncrypt := NewMicroEncrypt("")
	tester := &HttpWriteTester{
		test: t,
		onWriteHeader: func(header int) {
			if header != 400 {
				t.Errorf("Body is empty, throw 400")
			}
		},
	}
	asymmetricEncrypt.ServeHTTP(&http.Request{Body: ioutil.NopCloser(bytes.NewReader([]byte("")))}, tester)
}

func TestAsymmetricEncrypt_ServeHTTPSuccess(t *testing.T) {
	asymmetricEncrypt := NewMicroEncrypt("c5781adee495dfa51cfb8e2d357a0e90ba7be0a6f55fe557b89800ae7240df3b")

	tester := &HttpWriteTester {
		test: t,
		onWriteHeader: func(header int) {
			if header != 200 {
				t.Errorf("Key and body are valid, it should pass")
			}
		},
	}

	asymmetricEncrypt.ServeHTTP(&http.Request{Body: ioutil.NopCloser(bytes.NewReader([]byte("HI")))}, tester)
}
