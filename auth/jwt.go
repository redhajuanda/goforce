package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
)

type JWTHeader struct {
	Algorithm string `json:"alg"`
}

type JWTClaims struct {
	Issuer         string `json:"iss"`
	Subject        string `json:"sub"`
	Audience       string `json:"aud"`
	ExpirationTime string `json:"exp"`
}

type JWT struct {
	JWTHeader
	JWTClaims
}

func (jwt *JWT) Encode(header JWTHeader, claims JWTClaims) string {
	headerEncoded := EncodeBase64(header)
	claimEncoded := EncodeBase64(claims)
	return headerEncoded + "." + claimEncoded
}

func EncodeBase64(claims interface{}) string {
	bytes, _ := json.Marshal(claims)
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

func LoadKey(keyLocation string) (*rsa.PrivateKey, error) {
	pk, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		return nil, err
	}
	pkp, _ := pem.Decode(pk)
	pks, err := x509.ParsePKCS1PrivateKey(pkp.Bytes)
	if err != nil {
		return nil, err
	}
	return pks, nil
}

func SignWithRSA(jwtEncoded []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	h := sha256.New()
	h.Write(jwtEncoded)
	a := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, a)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// func test() {
// 	// jwtClaims := JWTClaims{
// 	// 	Issuer:         "3MVG99OxTyEMCQ3gNp2PjkqeZKxnmAiG1xV4oHh9AKL_rSK.BoSVPGZHQukXnVjzRgSuQqGn75NL7yfkQcyy7",
// 	// 	Subject:        "my@email.com",
// 	// 	Audience:       "https://login.salesforce.com",
// 	// 	ExpirationTime: "1333685628",
// 	// }
// 	// byte, _ := json.Marshal(jwtClaims)

// 	// fmt.Println(string(byte))

// 	jwtHeader := JWTHeader{
// 		Algorithm: "RS256",
// 	}
// 	bytea, _ := json.Marshal(jwtHeader)

// 	stringss := base64.StdEncoding.EncodeToString(bytea)
// 	fmt.Println(stringss)
// }
