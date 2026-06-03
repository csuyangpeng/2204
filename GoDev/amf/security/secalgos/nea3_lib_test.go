package secalgos

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"testing"
)

// Equal returns true if the slices have equal lengths and
// all elements are numerically identical.
func Equal(s1, s2 []uint32) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, val := range s1 {
		if s2[i] != val {
			return false
		}
	}
	return true
}

func TestNea3Set1(t *testing.T) {
	Key, _ := hex.DecodeString("173d14ba5003731d7a60049470f00a29")
	Count := uint32(0x66035492)
	Bearer := uint32(0xf)
	Direction := uint32(0x0)
	Length := uint32(193) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("6cf65340735552ab0c9752fa6f9025fe0bd675d9005875b200000000")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext  len ", len(Ciphertext)*32)
	Eea3(Key, Count, Bearer, Direction, Length, text32, Ciphertext)

	testC, _ := hex.DecodeString("a6c85fc66afb8533aafc2518dfe784940ee1e4b030238cc800000000")
	test32, _ := BytesToUint32ArrayV1(testC)
	fmt.Printf("-----test32 %08x\n", test32)

	fmt.Printf("-----Ciphertext %08x\n", Ciphertext) // uint32 使用byte读有小端问题

	//// 解密
	//receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	////ctxt := BytesToUint32Array(Ciphertext)
	//Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	//fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	//fmt.Printf("-----receiveTxt  %08x\n", receiveTxt)
	//if !Equal(receiveTxt, text32) {
	//	panic("解密错误!")
	//}

}

func TestNea3Set1V1(t *testing.T) {
	Key, _ := hex.DecodeString("173d14ba5003731d7a60049470f00a29")
	Count := uint32(0x66035492)
	Bearer := uint32(0xf)
	Direction := uint32(0x0)
	Length := uint32(193) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("6cf65340735552ab0c9752fa6f9025fe0bd675d9005875b200000000")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext  len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	testC, _ := hex.DecodeString("a6c85fc66afb8533aafc2518dfe784940ee1e4b030238cc800000000")
	test32, _ := BytesToUint32ArrayV1(testC)
	// []uint32 转换为大端码流
	ct, _ := Uint32ArrayToBytesV1(Ciphertext)
	fmt.Printf("-----message     %08x\n", ct)

	if !Equal(Ciphertext, test32) {
		panic("加密错误!")
	}

	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt  %08x\n", receiveTxt)
	if !Equal(receiveTxt, text32) {
		panic("解密错误!")
	}

}

func TestNea3Set2V1(t *testing.T) {
	Key, _ := hex.DecodeString("e5bd3ea0eb55ade866c6ac58bd54302a")
	Count := uint32(0x56823)
	Bearer := uint32(0x18)
	Direction := uint32(0x1)
	Length := uint32(800) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("14a8ef693d678507bbe7270a7f67ff5006c3525b9807e467c4e56000ba338f5d" +
		"429559036751822246c80d3b38f07f4be2d8ff5805f5132229bde93bbbdcaf38" +
		"2bf1ee972fbf9977bada8945847a2a6c9ad34a667554e04d1f7fa2c33241bd8f" +
		"01ba220d")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext    len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	testC, _ := hex.DecodeString("131d43e0dea1be5c5a1bfd971d852cbf712d7b4f57961fea3208afa8bca433f4" +
		"56ad09c7417e58bc69cf8866d1353f74865e80781d202dfb3ecff7fcbc3b190f" +
		"e82a204ed0e350fc0f6f2613b2f2bca6df5a473a57a4a00d985ebad880d6f238" +
		"64a07b01")
	test32, _ := BytesToUint32ArrayV1(testC)

	if !Equal(Ciphertext, test32) {
		panic("加密错误!")
	}

	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt %08x\n", receiveTxt)
	if !Equal(receiveTxt, text32) {
		panic("解密错误!")
	}

}

func TestNea3Set3V1(t *testing.T) {
	Key, _ := hex.DecodeString("d4552a8fd6e61cc81a2009141a29c10b")
	Count := uint32(0x76452ec1)
	Bearer := uint32(0x2)
	Direction := uint32(0x1)
	Length := uint32(1570) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("38f07f4be2d8ff5805f5132229bde93bbbdcaf382bf1ee972fbf9977bada8945" +
		"847a2a6c9ad34a667554e04d1f7fa2c33241bd8f01ba220d3ca4ec41e074595f" +
		"54ae2b454fd971432043601965cca85c2417ed6cbec3bada84fc8a579aea7837" +
		"b0271177242a64dc0a9de71a8edee86ca3d47d033d6bf539804eca86c584a905" +
		"2de46ad3fced65543bd90207372b27afb79234f5ff43ea870820e2c2b78a8aae" +
		"61cce52a0515e348d196664a3456b182a07c406e4a20791271cfeda165d535ec" +
		"5ea2d4df40000000")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext    len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	testC, _ := hex.DecodeString("8383b0229fcc0b9d2295ec41c977e9c2bb72e220378141f9c8318f3a270dfbcd" +
		"ee6411c2b3044f176dc6e00f8960f97afacd131ad6a3b49b16b7babcf2a509eb" +
		"b16a75dcab14ff275dbeeea1a2b155f9d52c26452d0187c310a4ee55beaa78ab" +
		"4024615ba9f5d5adc7728f73560671f013e5e550085d3291df7d5fecedded559" +
		"641b6c2f585233bc71e9602bd2305855bbd25ffa7f17ecbc042daae38c1f57ad" +
		"8e8ebd37346f71befdbb7432e0e0bb2cfc09bcd96570cb0c0c39df5e29294e82" +
		"703a637f80000000")
	test32, _ := BytesToUint32ArrayV1(testC)

	if !Equal(Ciphertext, test32) {
		panic("加密错误!")
	}

	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt %08x\n", receiveTxt)
	if !Equal(receiveTxt, text32) {
		panic("解密错误!")
	}

}

func TestNea3Set4V1(t *testing.T) {
	Key, _ := hex.DecodeString("db84b4fbccda563b66227bfe456f0f77")
	Count := uint32(0xe4850fe1)
	Bearer := uint32(0x10)
	Direction := uint32(0x1)
	Length := uint32(2798) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("e539f3b8973240da03f2b8aa05ee0a00dbafc0e182055dfe3d7383d92cef" +
		"40e92928605d52d05f4f9018a1f189ae3997ce19155fb1221db8bb0951a853ad" +
		"852ce16cff07382c93a157de00ddb125c7539fd85045e4ee07e0c43f9e9d6f41" +
		"4fc4d1c62917813f74c00fc83f3e2ed7c45ba5835264b43e0b20afda6b3053bf" +
		"b6423b7fce25479ff5f139dd9b5b995558e2a56be18dd581cd017c735e6f0d0d" +
		"97c4ddc1d1da70c6db4a12cc92778e2fbbd6f3ba52af91c9c6b64e8da4f7a2c2" +
		"66d02d001753df08960393c5d56888bf49eb5c16d9a80427a416bcb597df5bfe" +
		"6f13890a07ee1340e6476b0d9aa8f822ab0fd1ab0d204f40b7ce6f2e136eb674" +
		"85e507804d504588ad37ffd816568b2dc40311dfb654cdead47e2385c3436203" +
		"dd836f9c64d97462ad5dfa63b5cfe08acb9532866f5ca787566fca93e6b1693e" +
		"e15cf6f7a2d689d9741798dc1c238e1be650733b18fb34ff880e16bbd21b47ac" +
		"0000")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext    len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	testC, _ := hex.DecodeString("4bbfa91ba25d47db9a9f190d962a19ab323926b351fbd39e351e05da8b8925e3" +
		"0b1cce0d1221101095815cc7cb6319509ec0d67940491987e13f0affac332aa6" +
		"aa64626d3e9a1917519e0b97b655c6a165e44ca9feac0790d2a321ad3d86b79c" +
		"5138739fa38d887ec7def449ce8abdd3e7f8dc4ca9e7b73314ad310f9025e619" +
		"46b3a56dc649ec0da0d63943dff592cf962a7efb2c8524e35a2a6e7879d62604" +
		"ef268695fa4003027e22e6083077522064bd4a5b906b5f531274f235ed506cff" +
		"0154c754928a0ce5476f2cb1020a1222d32c1455ecaef1e368fb344d1735bfbe" +
		"deb71d0a33a2a54b1da5a294e679144ddf11eb1a3de8cf0cc061917974f35c1d" +
		"9ca0ac81807f8fcce6199a6c7712da865021b04ce0439516f1a526ccda9fd9ab" +
		"bd53c3a684f9ae1e7ee6b11da138ea826c5516b5aadf1abbe36fa7fff92e3a11" +
		"76064e8d95f2e4882b5500b93228b2194a475c1a27f63f9ffd264989a1bc0000")
	test32, _ := BytesToUint32ArrayV1(testC)

	if !Equal(Ciphertext, test32) {
		panic("加密错误!")
	}

	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt %08x\n", receiveTxt)
	if !Equal(receiveTxt, text32) {
		panic("解密错误!")
	}

}

func TestNea3Set5V1(t *testing.T) {
	Key, _ := hex.DecodeString("e13fed21b46e4e7ec31253b2bb17b3e0")
	Count := uint32(0x2738cdaa)
	Bearer := uint32(0x1a)
	Direction := uint32(0x0)
	Length := uint32(4019) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("8d74e20d54894e06d3cb13cb3933065e8674be62adb1c72b3a646965ab63" +
		"cb7b7854dfdc27e84929f49c64b872a490b13f957b64827e71f41fbd4269a42c" +
		"97f824537027f86e9f4ad82d1df451690fdd98b6d03f3a0ebe3a312d6b840ba5" +
		"a1820b2a2c9709c090d245ed267cf845ae41fa975d3333ac3009fd40eba9eb5b" +
		"885714b768b697138baf21380eca49f644d48689e4215760b906739f0d2b3f09" +
		"1133ca15d981cbe401baf72d05ace05cccb2d297f4ef6a5f58d91246cfa77215" +
		"b892ab441d5278452795ccb7f5d79057a1c4f77f80d46db2033cb79bedf8e605" +
		"51ce10c667f62a97abafabbcd6772018df96a282ea737ce2cb331211f60d5354" +
		"ce78f9918d9c206ca042c9b62387dd709604a50af16d8d35a8906be484cf2e74" +
		"a9289940364353249b27b4c9ae29eddfc7da6418791a4e7baa0660fa64511f2d" +
		"685cc3a5ff70e0d2b74292e3b8a0cd6b04b1c790b8ead2703708540dea2fc09c" +
		"3da770f65449e84d817a4f551055e19ab85018a0028b71a144d96791e9a35779" +
		"33504eee0060340c69d274e1bf9d805dcbcc1a6faa976800b6ff2b671dc46365" +
		"2fa8a33ee50974c1c21be01eabb2167430269d72ee511c9dde30797c9a25d86c" +
		"e74f5b961be5fdfb6807814039e7137636bd1d7fa9e09efd2007505906a5ac45" +
		"dfdeed7757bbee745749c29633350bee0ea6f409df4580160000")
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext    len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	testC, _ := hex.DecodeString("94eaa4aa30a57137ddf09b97b25618a20a13e2f10fa5bf8161a879cc2ae7" +
		"97a6b4cf2d9df31debb9905ccfec97de605d21c61ab8531b7f3c9da5f03931f8" +
		"a0642de48211f5f52ffea10f392a047669985da454a28f080961a6c2b62daa17" +
		"f33cd60a4971f48d2d909394a55f48117ace43d708e6b77d3dc46d8bc017d4d1" +
		"abb77b7428c042b06f2f99d8d07c9879d99600127a31985f1099bbd7d6c1519e" +
		"de8f5eeb4a610b349ac01ea2350691756bd105c974a53eddb35d1d4100b012e5" +
		"22ab41f4c5f2fde76b59cb8b96d885cfe4080d1328a0d636cc0edc05800b76ac" +
		"ca8fef672084d1f52a8bbd8e0993320992c7ffbae17c408441e0ee883fc8a8b0" +
		"5e22f5ff7f8d1b48c74c468c467a028f09fd7ce91109a570a2d5c4d5f4fa18c5" +
		"dd3e4562afe24ef771901f59af645898acef088abae07e92d52eb2de55045bb1" +
		"b7c4164ef2d7a6cac15eeb926d7ea2f08b66e1f759f3aee44614725aa3c7482b" +
		"30844c143ff85b53f1e583c501257dddd096b81268daa303f17234c2333541f0" +
		"bb8e190648c5807c866d7193228609adb948686f7de294a802cc38f7fe5208f5" +
		"ea3196d0167b9bdd02f0d2a5221ca508f893af5c4b4bb9f4f520fd84289b3dbe" +
		"7e61497a7e2a584037ea637b6981127174af57b471df4b2768fd79c1540fb3ed" +
		"f2ea22cb69bec0cf8d933d9c6fdd645e850591cca3d62c0cc000")
	test32, _ := BytesToUint32ArrayV1(testC)

	if !Equal(Ciphertext, test32) {
		panic("加密错误!")
	}

	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt %08x\n", receiveTxt)
	if !Equal(receiveTxt, text32) {
		panic("解密错误!")
	}

}

// example
func TestNea3MsgV1(t *testing.T) {
	Key, _ := hex.DecodeString("173d14ba5003731d7a60049470f00a29")
	Count := uint32(0x66035492)
	Bearer := uint32(0xf)
	Direction := uint32(0x0)
	Length := uint32(0) //bits,Plaintext length

	Plaintext := []byte("hello!")
	Length = uint32(len(Plaintext) * 8)
	//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	textPadding := ZeroPadding(Plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐
	fmt.Println("Plaintext     len ", len(text32)*32)
	fmt.Printf("-----Plaintext  %08x\n", text32)
	//txtT := []uint32{0xc8a9595e}

	// 加密
	Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext  len ", len(Ciphertext)*32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)

	//testC, _ := hex.DecodeString("a6c85fc66afb8533aafc2518dfe784940ee1e4b030238cc800000000")
	//test32, _ := BytesToUint32ArrayV1(testC)
	// []uint32 转换为大端码流
	// 解密
	receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
	//ctxt := BytesToUint32Array(Ciphertext)
	Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)
	fmt.Printf("-----Ciphertext %08x\n", Ciphertext)
	fmt.Printf("-----receiveTxt  %08x\n", receiveTxt)
	ct, _ := Uint32ArrayToBytesV1(receiveTxt)
	fmt.Printf("-----message     %08x\n", ct)
	fmt.Printf("-----message     %s\n", ct)

	fmt.Printf("-----enc message len %d\n", len(ct))
	fmt.Printf("-----txt message len %d\n", len(Plaintext))
	ct = ct[:len(Plaintext)]
	fmt.Printf("-----message     %08x\n", ct)
	fmt.Printf("-----message     %s\n", ct)
	if !bytes.Equal(ct, Plaintext) {
		panic("解密错误!")
	}

}

func BenchmarkNEA3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Key, _ := hex.DecodeString("173d14ba5003731d7a60049470f00a29")
		Count := uint32(0x66035492)
		Bearer := uint32(0xf)
		Direction := uint32(0x0)
		Length := uint32(0) //bits,Plaintext length

		Plaintext := []byte("hello!")
		Length = uint32(len(Plaintext) * 8)
		//text32 := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
		textPadding := ZeroPadding(Plaintext, 4)
		text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐

		// 加密
		Ciphertext := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28

		Eea3V1(Key, Count, Bearer, Direction, Length, text32, Ciphertext)

		//testC, _ := hex.DecodeString("a6c85fc66afb8533aafc2518dfe784940ee1e4b030238cc800000000")
		//test32, _ := BytesToUint32ArrayV1(testC)
		// []uint32 转换为大端码流
		// 解密
		receiveTxt := make([]uint32, uint(math.Ceil(float64(Length)/32))) //28
		//ctxt := BytesToUint32Array(Ciphertext)
		Eea3V1(Key, Count, Bearer, Direction, Length, Ciphertext, receiveTxt)

		ct, _ := Uint32ArrayToBytesV1(receiveTxt)

		ct = ct[:len(Plaintext)]

		if !bytes.Equal(ct, Plaintext) {
			panic("解密错误!")
		}
		//	BenchmarkNEA3-4   	 1000000	      1969 ns/op
	}
}
