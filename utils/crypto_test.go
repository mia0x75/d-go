package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/mia0x75/dashboard-go/utils"
	"golang.org/x/crypto/scrypt"
)

var key = []byte("cryto_key_123456")

func TestSHA256(t *testing.T) {
	h := sha256.New()
	password := [...]string{
		"123456789",
		"12345678",
		"11111111",
		"00000000",
		"123123123",
		"1234567890",
		"88888888",
		"111111111",
		"147258369",
		"987654321",
		"aaaaaaaa",
		"1111111111",
		"66666666",
		"a123456789",
		`a123456789 a123456789 a123456789 ~!@#$%^&*()_+`,
	}
	for index, elem := range password {
		h.Write([]byte(elem))
		bs := h.Sum(nil)
		fmt.Printf("%d - %d - %v\n", index, len(bs), bs)
	}
}

func TestSHA384(t *testing.T) {
	h := sha512.New384()
	password := [...]string{
		"123456kuJrPEZ2",
		"123456789",
		"12345678",
		"11111111",
		"00000000",
		"123123123",
		"1234567890",
		"88888888",
		"111111111",
		"147258369",
		"987654321",
		"aaaaaaaa",
		"1111111111",
		"66666666",
		"a123456789",
		`a123456789 a123456789 a123456789 ~!@#$%^&*()_+`,
	}
	for index, elem := range password {
		h.Write([]byte(elem))
		bs := h.Sum(nil)
		fmt.Printf(hex.EncodeToString(bs))
		fmt.Printf("%d - %d - %v\n", index, len(bs), bs)
	}
}

func TestSHA512(t *testing.T) {
	h := sha512.New()
	password := [...]string{
		"123456789",
		"12345678",
		"11111111",
		"00000000",
		"123123123",
		"1234567890",
		"88888888",
		"111111111",
		"147258369",
		"987654321",
		"aaaaaaaa",
		"1111111111",
		"66666666",
		"a123456789",
		`a123456789 a123456789 a123456789 ~!@#$%^&*()_+`,
	}
	for index, elem := range password {
		h.Write([]byte(elem))
		bs := h.Sum(nil)
		fmt.Printf("%d - %d - %v\n", index, len(bs), bs)
	}
}

func TestScrypt(t *testing.T) {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	password := [...]string{
		"123456789",
		"12345678",
		"11111111",
		"00000000",
		"123123123",
		"1234567890",
		"88888888",
		"111111111",
		"147258369",
		"987654321",
		"aaaaaaaa",
		"1111111111",
		"66666666",
		"a123456789",
		`a123456789 a123456789 a123456789 ~!@#$%^&*()_+`,
	}
	for index, elem := range password {
		dk, _ := scrypt.Key([]byte(elem), salt, 1<<15, 8, 1, 32)
		fmt.Printf("%d - %d - %v\n", index, len(dk), dk)
	}

}

func TestEncryptCBC(t *testing.T) {
	plaintext := [...]string{
		``, // Empty String
		`1`,
		`12`,
		`123`,
		`1234`,
		`12345`,
		`123456`,
		`1234567`,
		`12345678`,
		`123456789`,
		`1234567890`,
		`12345678901`,
		`123456789012`,
		`1234567890123`,
		`12345678901234`,
		`123456789012345`,
		`1234567890123456`,
		`12345678901234567`,
		`123456789012345678`,
		`1234567890123456789`,
		`~!@#$%^&*()_`,
		`abcdefghijklmnopqrstuvwxyz`,
		`acb 123 $%^`,
	}

	for index, elem := range plaintext {
		if cipher, err := utils.EncryptCBC([]byte(elem), key); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%d - %v - %s\n", index, cipher, hex.EncodeToString(cipher))
		}
	}
}

func TestDecryptCBC(t *testing.T) {
	cipher := [...]string{
		"7f77229c1b60ee1c729bfefaa7dd0c1a28efbd64fe9333ce247f2537b1cac348",
		"9263f6ac5565deddf7c72d2b87253950f5ce7df0a63e1a01275d4edbee0f9bcf",
		"08d47ca6ff5a3bd1afdabd742de532d1b1609154dcdc607aa8db2966ccf0eca2",
		"b854c3b0a78d7d196b324c45596b841b6224c4d7df4e01fba7b1a992df82eefb",
		"5b992e8d8cf1fc1bd38e4549432aeef320d13548ecabd1bb7b02382deecd87aa",
		"6e3a2e1c9d0aeba6d9631dadedcda010cc1ceb614ccddfa263cf461b481d1cce",
		"17f101a4872be53a804dc509588cb25118c3c7deb7dd44556953ff1e26753ac9",
		"0b36d578918075359cfc935ad542f08c376ce0fb09af3056843695a75f774c94",
		"48c7017943129194404ead22195a1bc10949fdb63dafe5fc8f1522a9c64a0a56",
		"aec9cdbcffadb4b803c71592fc7b2e4596a766aab08c81f6e86eb57eba45233f",
		"74308107ed86e3458867763dceab2dc4ef651d479b99fae6f8e8db79f99a63aa",
		"4cd9d2b70b315499433b324f6091995672f77d7d3bf209e6c1cc46f0c8838d1b",
		"af50ace0ab25441fff6c7c63fc545dbc1dc154c55d9f7abd3784ec3b487d9022",
		"c9f65eaf134905160ca90917979f080487632e150f3b1b54ec22cb87f2cea505",
		"c57058d7790f983d6d8da0990db0c28f73fef7f264f151dd7e872692b982f0cb",
		"0d655038080d45a20cfe42c785ae0e20cc9cc946555b0c225814eb6a9a31817d",
		"45c490ed4805da509372ec89863fc9fd22294e3ea5ddd63c0499aedb67768a5170014fefaf4702d6dd5837f8dba975ff",
		"e36c1be2d5b59520ca50196989df573490adaf08f86af9d0561d780f6e19c647db6e1f4ef7577a1887421fd46d7389f7",
		"b162b99a437a3a88fb49b75949251174ab34a4cab197438a7b1d6fb0986bfdd60ce100d8007d5c0e43ea6515aee41511",
		"49322b78c06867e19f5724049c70ff65f7b0ec27d2d3ad6b8fbd470c0482958483ecdfc4888fd54fda64e4b84b3efa5f",
		"a9ca81f7011ba6832fa207bd5b84e4561c1748d530ed03a188edde349110a5c5",
		"b6c8b9aaa49d6cc1a46a6a902a75784e5c6e7788b8d74b70029a2b680ff51f13cb10e3bf5aa3f31c8dc4b1c72e377902",
		"8bf75ef6ea0374fe66d4151277714717d6e2d00c936af9d4a506a8745ceeeae7",
	}
	for index, elem := range cipher {
		data, _ := hex.DecodeString(elem)
		if plaintext, err := utils.DecryptCBC(data, key); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%d - %v - %s\n", index, plaintext, string(plaintext))
		}
	}
}
