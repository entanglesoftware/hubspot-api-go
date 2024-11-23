package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

const MaxAllowedTimestamp = 300000

type Options struct {
	Signature        string
	ClientSecret     string
	RequestBody      string
	SignatureVersion string
	URL              string
	Method           string
	Timestamp        int64
}

type Signature struct{}

// IsValid checks the validity of a signature
func (s Signature) IsValid(opts Options) (bool, error) {
	hash, err := s.GetSignature(opts.Method, opts.SignatureVersion, opts)
	if err != nil {
		return false, err
	}

	if opts.SignatureVersion == "v3" {
		currentTime := time.Now().UnixNano() / int64(time.Millisecond)
		if opts.Timestamp == 0 || currentTime-opts.Timestamp > MaxAllowedTimestamp {
			return false, errors.New("timestamp is invalid, reject request")
		}
	}

	return opts.Signature == hash, nil
}

// GetSignature generates a signature based on method, version, and options
func (s Signature) GetSignature(method string, signatureVersion string, opts Options) (string, error) {
	var sourceString string

	switch signatureVersion {
	case "v1":
		sourceString = opts.ClientSecret + opts.RequestBody
		hash := sha256.New()
		hash.Write([]byte(sourceString))
		return hex.EncodeToString(hash.Sum(nil)), nil
	case "v2":
		sourceString = opts.ClientSecret + method + opts.URL + opts.RequestBody
		hash := sha256.New()
		hash.Write([]byte(sourceString))
		return hex.EncodeToString(hash.Sum(nil)), nil
	case "v3":
		sourceString = method + opts.URL + opts.RequestBody + strconv.FormatInt(opts.Timestamp, 10)
		h := hmac.New(sha256.New, []byte(opts.ClientSecret))
		h.Write([]byte(sourceString))
		return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
	default:
		return "", errors.New("not supported signature version: " + signatureVersion)
	}
}
