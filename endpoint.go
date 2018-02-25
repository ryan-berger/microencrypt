package microencrypt

import (
	"net/http"
	"io/ioutil"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
	"encoding/base64"
)

type AsymmetricEncrypt struct {
	Key string
}

func (asymmetric *AsymmetricEncrypt) ServeHTTP(r *http.Request, rw http.ResponseWriter) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"error": "Error reading body"}`))
		return
	}

	if len(body) == 0 {
		rw.WriteHeader(400)
		rw.Write([]byte(`{"error": "Body must not be empty"}`))
		return
	}

	key, err := nacl.Load(asymmetric.Key)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"error": "Internal server Error""}`))
		return
	}

	message := base64.StdEncoding.EncodeToString(secretbox.EasySeal(body, key))
	rw.WriteHeader(200)
	rw.Write([]byte(`{"message": "` + message + `" }`))
}

func NewAsymmetricEncrypt(key string) *AsymmetricEncrypt {
	return &AsymmetricEncrypt{Key: key}
}
