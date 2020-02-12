package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestClaim(t *testing.T) {
	jwtClaims := JWTClaims{
		Issuer:         "3MVG99OxTyEMCQ3gNp2PjkqeZKxnmAiG1xV4oHh9AKL_rSK.BoSVPGZHQukXnVjzRgSuQqGn75NL7yfkQcyy7",
		Subject:        "my@email.com",
		Audience:       "https://login.salesforce.com",
		ExpirationTime: "1333685628",
	}
	bytes, err := json.Marshal(jwtClaims)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	fmt.Println(string(bytes))
	stringss := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(stringss)
}
func TestHeader(t *testing.T) {
	jwtHeader := JWTHeader{
		Algorithm: "RS256",
	}
	bytea, _ := json.Marshal(jwtHeader)

	stringss := base64.StdEncoding.EncodeToString(bytea)
	fmt.Println(stringss)
}

func TestEncode(t *testing.T) {
	jwt := JWT{
		JWTHeader: JWTHeader{
			Algorithm: "RS256",
		},
		JWTClaims: JWTClaims{
			Issuer:         "3MVG9rnryk9FxFMX.Ol9n.__jZ6WA7LYsp6ioXCY0HRm5mnlenXdif3yylGEGCFvLIc.rZzh6OrfAfBEn7gU2",
			Subject:        "redha@amalan.org.developer",
			Audience:       "https://test.salesforce.com",
			ExpirationTime: strconv.FormatInt(time.Now().Add(time.Minute*3).Unix(), 10),
		},
	}

	// 	Create a string for the encoded JWT Header and the encoded JWT Claims Set in this format.
	// encoded_JWT_Header + "." + encoded_JWT_Claims_Set
	jwtEncoded := jwt.Encode(jwt.JWTHeader, jwt.JWTClaims)
	fmt.Println(jwtEncoded)
	// load private key
	privateKey, err := LoadKey("../certificates/server.key")
	if err != nil {
		t.Errorf("Err %v", err)
	}

	signature, err := SignWithRSA([]byte(jwtEncoded), privateKey)

	if err != nil {
		t.Errorf("Err %v", err)
	}

	encodedSig := base64.StdEncoding.EncodeToString(signature)

	result := jwtEncoded + "." + encodedSig
	fmt.Println(result)
}
