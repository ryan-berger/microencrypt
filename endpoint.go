package microencrypt

import (
	"net/http"
	"io/ioutil"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
	"encoding/base64"
)

type MicroEncrypt struct {
	Key string
}

func (asymmetric *MicroEncrypt) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	rw.Header().Add("Content-Type", "application/json")

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

func NewMicroEncrypt(key string) *MicroEncrypt {
	return &MicroEncrypt{Key: key}
}
