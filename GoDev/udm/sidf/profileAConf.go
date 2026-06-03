package sidf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/udm/sidf/ecies"
	"os"
	"path/filepath"
)

var pkiMap = make(map[uint8]ecies.KeyPair)

// TS23.003
// Figure 2.2B-3:
// Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile A
// key pair
type HomeNetworkPKI struct {
	// Home network public key identifier (PKI)
	// Home network PKI value (1-254)
	HomeNetworkPKI uint8
	HomeNetworkKeyPair
}

// key pair
type HomeNetworkKeyPair struct {
	// The ECC ephemeral public key is formatted as 64 hexadecimal digits,
	// which allows to encode 256 bits.
	// Home network public key:Hexadecimal value
	HnPublicKey []byte
	//Home Network Private Key:Hexadecimal value
	HnPrivteKey []byte
}

type HomeNetworkPKIStr struct {
	// Home network public key identifier (PKI)
	// Home network PKI value (1-254)
	HomeNetworkPKI uint8 `json:"Home network public key identifier"`
	HomeNetworkKeyPairStr
}
type HomeNetworkKeyPairStr struct {
	// The ECC ephemeral public key is formatted as 64 hexadecimal digits,
	// which allows to encode 256 bits.
	// Home network public key:Hexadecimal value
	HnPublicKey string `json:"Home network public key"`
	//Home Network Private Key:Hexadecimal value
	HnPrivteKey string `json:"Home network private key"`
}

//var PKISet []HomeNetworkPKI

func Init() {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Load home network public private key pair")
	dir, _ := filepath.Abs(filepath.Dir(types.DefConfFileSidfKey))
	_, file := filepath.Split(types.DefConfFileSidfKey)
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Load directory", filepath.Join(dir, file))
	err := JsonFileRead(types.DefConfFileSidfKey)
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Failed to load home network public private key pair.", err)
	}
}

func JsonFileRead(filename string) error {
	//filename = "config/pki.key"
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	//rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil,  "%v", len(content))
	pkiSetStr := []HomeNetworkPKIStr{}
	err = json.Unmarshal(content, &pkiSetStr)
	if err != nil {
		return err
	}
	//fmt.Printf("%v\n\n", pkiSetStr)

	// HomeNetworkPKIStr 转存到 map[PKI]KeyPair
	//pkiMap := make(map[uint8]ecies.KeyPair)
	// 获取PKI，并验证公私钥匹配情况
	for _, kpi := range pkiSetStr {
		kPair := ecies.KeyPair{}
		ecies.HexToArray(kpi.HnPublicKey, &kPair.PublicKey)
		ecies.HexToArray(kpi.HnPrivteKey, &kPair.PrivateKey)

		ok := ecies.VerifyKeyPair(kPair)
		if ok != true {
			return fmt.Errorf("Public key does not match private key,"+
				"Home network public key identifier:%v", kpi.HomeNetworkPKI)
		}
		pkiMap[kpi.HomeNetworkPKI] = kPair
		rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "pki:%d,public key:%x", kpi.HomeNetworkPKI, kPair.PublicKey)
		//fmt.Printf("pki:%d，public key:%x\n", kpi.HomeNetworkPKI, kPair.PublicKey)
	}

	//for k, v := range pkiMap {
	//	fmt.Printf("pki：%d，private key:%x\n", k, v.PrivateKey)
	//}
	return nil
}

func GetHNwPrivateKey(id uint8) ([]byte, error) {
	//test set 0
	if id == 0 {
		return ecies.HnPrivteKey, nil
	}

	KeyPair, ok := pkiMap[id]
	if !ok {
		return nil, ErrInvalidPubKeyId
	}

	return KeyPair.PrivateKey[:], nil
}
