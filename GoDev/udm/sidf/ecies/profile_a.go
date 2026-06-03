package ecies

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/curve25519"
	"hash"
	"io"
)

// 3GPP TS 33.501 V15.5.0 (2019-06)
/*The ME and SIDF shall implement this profile. The ECIES parameters for this profile shall be the following:
ME和SIDF应实现此配置文件。 此配置文件的ECIES参数应如下：
-	EC domain parameters							: Curve25519 [46]
-	EC Diffie-Hellman primitive					    : X25519 [46]
-	point compression								: N/A
-	KDF												: ANSI-X9.63-KDF [29]
-	Hash												: SHA-256
-	SharedInfo1										:  (the ephemeral public key octet string – see [29] section 5.1.3)
-	MAC												: HMAC–SHA-256
-	mackeylen										: 32 octets (256 bits)
-	maclen											: 8 octets (64 bits)
-	SharedInfo2										: the empty string
-	ENC												: AES–128 in CTR mode
-	enckeylen											: 16 octets (128 bits)
-	icblen												: 16 octets (128 bits)
-	backwards compatibility mode					: false*/

const (
	hashlen   = 32 //byte
	mackeylen = 32 //byte
	maclen    = 8  //byte
	enckeylen = 16 //byte
	icblen    = 16 //byte
)

var ueSchemeOutput UeSchemeOutputA

// example:
//Home Network Private Key:
//'c53c22208b61860b06c62e5406a7b330c2b577aa5558981510d128247d38bd1d'
var HnPrivteKey, _ = hex.DecodeString("c53c22208b61860b06c62e5406a7b330c2b577aa5558981510d128247d38bd1d")

//Home Network Public Key:
//'5a8d38864820197c3394b92613b20b91633cbd897119273bf8e4a6f4eec0a650'
var HnPublicKey, _ = hex.DecodeString("5a8d38864820197c3394b92613b20b91633cbd897119273bf8e4a6f4eec0a650")

// The Figure C.3.3-1 illustrates the home network's steps.
// input:
// public key of UE
// private key of HN
// Ciphertext
// MAC tag
type ProfileAInput struct {
	UePubKey   [32]byte
	HnPriKey   [32]byte
	Ciphertext []byte
	MacTag     []byte
}

// output:
// plaintext
type ProfileAOutput struct {
	plaintext []byte
}

// Value in process
type InternalKey struct {
	EphSharedKey [32]byte
	EphEncKey    []byte
	ICB          []byte
	EphMacKey    []byte
}

// Decryption based on ECIES at home network
type HNEciesA struct {
	ProfileAInput
	ProfileAOutput
	InternalKey
}

func NewHnEciesA() *HNEciesA {
	return new(HNEciesA)
}

// kdf result
type KeyMaterial struct {
	CipherKey []byte
	MacKey    []byte
	IV        []byte
}

// ECIES Profile A ,Curve25519
// 生成共享秘钥
func (d *HNEciesA) KeyAgreement() error {
	////B:Home Network交互Their shared secret KEY
	dst := &d.EphSharedKey
	in := &d.HnPriKey
	base := &d.UePubKey
	if (dst != nil) && (in != nil) && (base != nil) {
		curve25519.ScalarMult(dst, in, base)
		return nil
	}

	return ErrInvalidParams
}

// ANSI-X9.63-KDF
func (d *HNEciesA) KeyDerivation() error {
	info := d.UePubKey[:]
	sk := d.EphSharedKey[:]
	//enckeylen + icblen + mackeylen
	//    -	enckeylen			: 16 octets (128 bits)
	//    -	icblen				: 16 octets (128 bits)
	//    -	mackeylen		    : 32 octets (256 bits)
	outputLen := 16 + 16 + 32
	dKey, err := DeriveSecretsV2(sk, info, outputLen)
	if err != nil {
		return ErrDerivationFailed
	}

	var kdfKey KeyMaterial
	kdfKey.CipherKey = dKey[:16]
	kdfKey.IV = dKey[16:32]
	kdfKey.MacKey = dKey[32:]

	d.EphEncKey = kdfKey.CipherKey
	d.ICB = kdfKey.IV
	d.EphMacKey = kdfKey.MacKey

	return nil
}

// return plaintext
func (d *HNEciesA) SymDecrytion() ([]byte, error) {
	key := d.EphEncKey
	iv := d.ICB
	plaintext := make([]byte, len(d.Ciphertext))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrDecryptionBlock
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, d.Ciphertext)
	d.plaintext = plaintext

	return plaintext, nil
}

func (d *HNEciesA) KeyMacVerify() bool {
	var ECIES_AES128_SHA256 = ECIESParams{
		Hash:      sha256.New,
		Cipher:    aes.NewCipher,
		BlockSize: aes.BlockSize,
		KeyLen:    16,
	}
	MACTagTmp := messageTag(ECIES_AES128_SHA256.Hash, d.EphMacKey, d.Ciphertext, nil)
	if bytes.Equal(d.MacTag, MACTagTmp[:8]) {
		return true
	}
	return false
}

type ECIESParams struct {
	Hash      func() hash.Hash // hash function
	hashAlgo  crypto.Hash
	Cipher    func([]byte) (cipher.Block, error) // symmetric cipher
	BlockSize int                                // block size of symmetric cipher
	KeyLen    int                                // length of symmetric key
}

type KeyPair struct {
	PrivateKey [32]byte
	PublicKey  [32]byte
}

// MAC Tag
func messageTag(hash func() hash.Hash, km, msg, shared []byte) []byte {
	mac := hmac.New(hash, km)
	mac.Write(msg)
	mac.Write(shared)
	tag := mac.Sum(nil)
	return tag
}

// 生成公私秘钥对
func GenerateKeyPair() (*KeyPair, error) {
	var privateKey, publicKey [32]byte
	//产生随机私钥
	if _, err := io.ReadFull(rand.Reader, privateKey[:]); err != nil {
		return nil, ErrGeneratePrivatelicKeyFailed
	}
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	//fmt.Println("A私钥", base64.StdEncoding.EncodeToString(privateKey[:]))

	return &KeyPair{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

// 秘钥对有效性验证
func VerifyKeyPair(keyPair KeyPair) bool {
	// private key test // 私钥验证公钥
	var publicTmp [32]byte
	//curve25519.ScalarBaseMult(&publicTmp, (*[32]byte)(unsafe.Pointer(&keyPair.PrivateKey[0])))
	curve25519.ScalarBaseMult(&publicTmp, &keyPair.PrivateKey)

	return bytes.Equal(keyPair.PublicKey[:], publicTmp[:])
}

func Keccak256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
		//fmt.Println(hex.EncodeToString(b))
	}
	return d.Sum(nil)
}

//输入：密钥派生功能的输入是：
//1.一个八位位组字符串Z，它是共享的机密值。(zj: shared key)
//2.整数keydatalen，它是要生成的密钥数据的八位字节长度。
//3.（可选）八位字节字符串SharedInfo，该字符串由一些实体共享的一些数据组成，这些数据旨在共享Z(shared key)。
//    -	SharedInfo1	:  (the ephemeral public key octet string – see [29] section 5.1.3)(zj:作为G点的public key，UePubKey)
//
//输出：键数据K，它是长度为keydatalen八位字节的八位字节串，或“无效”。
// inputKeyMaterial :EphSharedKey,info :UePubKey
func DeriveSecretsV2(inputKeyMaterial, info []byte, outputLength int) ([]byte, error) {

	counter := "00000001"
	ct, _ := hex.DecodeString(counter)
	//fmt.Println(ct)
	//fmt.Println("len:", len(inputKeyMaterial)+len(info)+len(ct)) //68

	var dgstlen int = 32
	//ct := []byte{0,0,0,1}
	out := make([]byte, outputLength)
	rlen := outputLength
	var length int = 0
	for rlen > 0 {
		d := Keccak256(inputKeyMaterial, ct, info)
		if rlen >= dgstlen {
			copy(out[length:], d[:dgstlen])
			length += dgstlen
		} else {
			copy(out[length:], d[:rlen])
			length += rlen
		}
		rlen -= dgstlen
		ct[3] += 1
	}

	return out, nil
}

func HexToArray(hexStr string, array *[32]byte) error {
	hexSlice, err := hex.DecodeString(hexStr)
	if err != nil {
		err = errors.New("invalid hex string")
		return err
	}
	if len(hexSlice) != 32 {
		err = errors.New("invalid hex length")
		return err
	}
	for i, _ := range array {
		array[i] = hexSlice[i]
	}
	return nil
}
