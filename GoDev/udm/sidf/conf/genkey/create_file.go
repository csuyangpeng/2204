package genkey

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lite5gc/udm/sidf/ecies"
	"os"
	"path/filepath"
)

// Public-private key pair of profile A using the curve25519 algorithm
func CreateJsonFile(sum uint, filename string) error {
	// Home network PKI value (1-254)
	if sum < 1 || sum > 254 {
		return fmt.Errorf("sum (1-254)")
	}
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}
	//keypair, _ := ecies.GenerateKeyPair()
	PKISet = PKISet[:0]
	for i := 0; i < int(sum); i++ {
		keypair, _ := ecies.GenerateKeyPair()
		PKISet = append(PKISet,
			HomeNetworkPKI{HomeNetworkPKI: uint8(i + 1),
				HomeNetworkKeyPair: HomeNetworkKeyPair{HnPrivteKey: keypair.PrivateKey[:],
					HnPublicKey: keypair.PublicKey[:]},
			},
		)
	}

	// []byte 转换为hex编码字符，在文本中可见
	var PkiSetStr []HomeNetworkPKIStr
	for _, v := range PKISet {
		//fmt.Printf("HnPublicKey:\n %s\n", hex.EncodeToString(v.HnPublicKey))
		//vv, _ := hex.DecodeString(hex.EncodeToString(v.HnPublicKey))
		//fmt.Printf("HnPublicKey:\n %#x\n", vv)
		Pki := HomeNetworkPKIStr{}
		Pki.HomeNetworkPKI = v.HomeNetworkPKI
		//Pki. := HomeNetworkKeyPairStr{}
		Pki.HnPublicKey = hex.EncodeToString(v.HnPublicKey)
		Pki.HnPrivteKey = hex.EncodeToString(v.HnPrivteKey)

		PkiSetStr = append(PkiSetStr, Pki)
	}
	//data, err := json.Marshal(PKISet)
	data, err := json.MarshalIndent(PkiSetStr, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("data:\n %s\n", data)
	//if filename == "" {
	//	filename = "pki.key"
	//}
	//fmt.Println("Remove", os.Remove(filename))
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		return err
	}
	//fmt.Println("dir", dir, err)

	exist, err := PathExists(filename)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("file already exist")
	}
	//fmt.Println("PathExists", exist, err)
	f, err := os.Create(filename)
	if err != nil {
		//fmt.Println("Create", err)
		return err
	}
	//fmt.Println("Truncate", f.Truncate(0))
	//fmt.Println("f.Sync", f.Sync())
	_, err = f.Write(data) //f.WriteString(data)
	if err != nil {
		f.Close()
		return err
	}
	f.Sync()
	f.Close()
	//fmt.Println("f.Sync", f.Sync())
	//fmt.Println("f.Close", f.Close())
	//defer os.Remove("testjsonWithArray.json")

	// 修改文件属性
	info, err := os.Stat(filename) //Stat获取文件属性
	if err != nil {
		//fmt.Println("os.Stat err =", err)
		return err
	}
	//fmt.Println("获取文件属性")
	//fmt.Println("name =", info.Name())
	//fmt.Println("size =", info.Size())
	//fmt.Println("mode =", info.Mode())
	//fmt.Println("modtime =", info.ModTime())
	//fmt.Println("isDir =", info.IsDir())
	//fmt.Println("sys =", info.Sys())

	os.Chmod(filename, 0444)
	info, err = os.Stat(filename) //Stat获取文件属性
	if err != nil {
		//fmt.Println("os.Stat err =", err)
		return err
	}
	fmt.Println("File generated successfully")
	fmt.Println("name      =", info.Name())
	fmt.Println("sum       =", sum)
	fmt.Println("directory =", dir)

	fmt.Println("size      =", info.Size())
	fmt.Println("mode      =", info.Mode())
	fmt.Println("modtime   =", info.ModTime())
	//fmt.Println("isDir =", info.IsDir())
	//fmt.Println("sys =", info.Sys())
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
