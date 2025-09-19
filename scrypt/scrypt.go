package scrypt

import (
	"crypto/rand"
	"encoding/base64"
	"reflect"

	"go.k6.io/k6/js/modules"
	"golang.org/x/crypto/scrypt"
)

type ScryptResult struct {
	Hash  string
	Salt  string
	Input string
}

type ScryptModule struct{}
type Exporter struct{}

type ScryptOptions struct {
	N       int
	r       int
	p       int
	keyLen  int
	saltLen int
}

func (m *ScryptModule) GenerateHash(input string, opts map[string]interface{}) (ScryptResult, error) {
	// Set defaults
	options := ScryptOptions{
		N:       16384,
		r:       8,
		p:       1,
		keyLen:  64,
		saltLen: 16,
	}
	// Override with provided opts
	v := reflect.ValueOf(&options).Elem()
	for k, val := range opts {
		field := v.FieldByName(k)
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.Int:
				switch v := val.(type) {
				case int:
					field.SetInt(int64(v))
				case float64:
					field.SetInt(int64(v)) // JS numbers are float64
				}
			}
		}
	}
	salt := make([]byte, options.saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return ScryptResult{}, err
	}
	dk, err := scrypt.Key([]byte(input), salt, options.N, options.r, options.p, options.keyLen)
	if err != nil {
		return ScryptResult{}, err
	}
	return ScryptResult{
		Hash:  base64.StdEncoding.EncodeToString(dk),
		Salt:  base64.StdEncoding.EncodeToString(salt),
		Input: input,
	}, nil
}

func (e *Exporter) Exports() modules.Exports {
	m := &ScryptModule{}
	named := map[string]any{
		"generateHash": m.GenerateHash,
	}
	return modules.Exports{Named: named}
}

// NewModuleInstance makes Exporter implement modules.Module so k6 will
// call Exports() via the Module/Instance contract.
func (e *Exporter) NewModuleInstance(vu modules.VU) modules.Instance {
	return e
}

func init() {
	// Register the module so k6 can import it as k6/x/scrypt
	modules.Register("k6/x/scrypt", new(Exporter))
}



