package bls

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var blsMgr = NewBlsManager()

func TestBlsMgr_GenerateKey(t *testing.T) {
	sk, pk := blsMgr.GenerateKey()
	assert.NotEmpty(t, sk, "gen key fail")
	assert.NotEmpty(t, pk, "gen key fail")
	pk1, err := sk.PubKey()
	assert.NoError(t, err, "expect no error")
	bpk := pk.Compress()
	bpk1 := pk1.Compress()
	fmt.Println(len(bpk))
	assert.EqualValues(t, bpk, bpk1, "public key not equal", bpk, bpk1)
	//t.Log(bpk)
	//t.Log(bpk1)

	sk2, _ := blsMgr.GenerateKey()
	bsk, bsk2 := sk.Compress(), sk2.Compress()
	assert.NotEqual(t, bsk, bsk2, "should not generate two same key", bsk, bsk2)
	//t.Log(bsk)
	//t.Log(bsk2)
}

func TestSingleSignAndVerify(t *testing.T) {
	sk, pk := blsMgr.GenerateKey()
	m1 := Message("message to be signed. 将要做签名的消息")
	//pair sign and verify
	sig1 := sk.Sign(m1)
	fmt.Println("len:", len(sig1.Compress().Bytes()))
	err := pk.Verify(m1, sig1)
	assert.NoError(t, err)

	//different message should have different signature
	m2 := Message("message to be signed. 将要做签名的消息.")
	sig2 := sk.Sign(m2)
	fmt.Println(len(sig2.Compress()))
	assert.NotEqual(t, sig1, sig2, "different message got the same signature", sig1, sig2)

	//different key should have different signature for a same message.
	sk2, _ := blsMgr.GenerateKey()
	sig12 := sk2.Sign(m1)
	err = pk.Verify(m1, sig12)
	assert.Error(t, err)
}

// contains test case of Compress and Decompress both for secret key and public key
func TestBlsManager_Decompress(t *testing.T) {
	sk, pk := blsMgr.GenerateKey()
	bsk, bpk := sk.Compress(), pk.Compress()
	dsk, err := blsMgr.DecSecretKey(bsk.Bytes())
	assert.NoError(t, err)
	dpk, err := blsMgr.DecPublicKey(bpk.Bytes())
	assert.NoError(t, err)

	//decompress string
	dskFromHex, err := blsMgr.DecSecretKeyHex(bsk.String())
	assert.NoError(t, err)
	assert.Equal(t, sk.Compress(), dskFromHex.Compress())

	dpkFromHex, err := blsMgr.DecPublicKeyHex(bpk.String())
	assert.NoError(t, err)
	assert.Equal(t, pk.Compress(), dpkFromHex.Compress())

	//cross sign and verify
	m1 := Message("message to be signed. 将要做签名的消息")
	sig1 := sk.Sign(m1)
	err = dpk.Verify(m1, sig1)
	assert.NoError(t, err)
	sig2 := dsk.Sign(m1)
	assert.EqualValues(t, sig1, sig2)

	dsig, err := blsMgr.DecSignature(sig1.Compress().Bytes())
	assert.NoError(t, err)
	err = dpk.Verify(m1, dsig)
	assert.NoError(t, err)
	dsigFromHex, err := blsMgr.DecSignatureHex(sig1.Compress().String())
	assert.NoError(t, err)
	err = dpk.Verify(m1, dsigFromHex)
	assert.NoError(t, err)

	//prefix string
	skstr := "0x" + bsk.String()
	pkstr := "0x" + bpk.String()
	sigstr := "0x" + sig1.Compress().String()
	dskFromHex, err = blsMgr.DecSecretKeyHex(skstr)
	assert.NoError(t, err)
	assert.Equal(t, sk.Compress(), dskFromHex.Compress())

	dpkFromHex, err = blsMgr.DecPublicKeyHex(pkstr)
	assert.NoError(t, err)
	assert.Equal(t, pk.Compress(), dpkFromHex.Compress())
	dsigFromHex, err = blsMgr.DecSignatureHex(sigstr)
	assert.NoError(t, err)
	err = dpk.Verify(m1, dsigFromHex)
	assert.NoError(t, err)

	//invalid hex
	skstr += "00"
	pkstr += pkstr[:len(pkstr)-1]
	//sigcp := sig1.Compress()
	//fmt.Println(sigcp.String())
	//sigcp[0] = sigcp[0] + 1
	//sigstr = sigcp.String()
	//fmt.Println(sigstr)
	_, err = blsMgr.DecSecretKeyHex(skstr)
	assert.Error(t, err)
	_, err = blsMgr.DecPublicKeyHex(pkstr)
	assert.Error(t, err)
	//dsigFromHex, err = blsMgr.DecSignatureHex(sigstr)
	//assert.Error(t, err)
}

func TestDebug(t *testing.T) {
	//	//valid sig: 82051c87397e54313c98ad614e3f2085e43cd2c8cb5262c8d6cc27c871b91efabbbfd13033938e8bb95fdb3da5973dfd
	//	//panic sig: 83051c87397e54313c98ad614e3f2085e43cd2c8cb5262c8d6cc27c871b91efabbbfd13033938e8bb95fdb3da5973dfd
	//	sigstr := "83051c87397e54313c98ad614e3f2085e43cd2c8cb5262c8d6cc27c871b91efabbbfd13033938e8bb95fdb3da5973dfd"
	//	_, err := blsMgr.DecSignatureHex(sigstr)
	//	assert.NoError(t, err)
}

func TestECDSA(t *testing.T) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Println(privKey.PublicKey.X.BitLen() / 8)

	plainText := []byte("12fsdfsdfsdfsdfsdfsasdasdasddfsdfsdfsdfsdfsdfdsfssdfds3")
	myHash := sha256.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(nil)
	_, _, err = ecdsa.Sign(rand.Reader, privKey, hashText)
	if err != nil {
		panic(err)
	}

	//rText,err :=r.MarshalText()
	//if err!=nil{
	//	panic(err)
	//}
	//sText,err :=s.MarshalText()
	//if err!=nil{
	//	panic(err)
	//}

	//for i:= 1 ; i <= 100 ; i++{
	//	len1,_,_ := GetCntBlockGroup(i)
	//	len2, _ := GetPreBlockGroup(i)
	//	fmt.Print(i, " ", len2, " ",len1)
	//	fmt.Println()
	//}

	//for i := 1; i <= 100; i++ {
	//	_, _, time1 := GetPreBlockGroup(10, i*100, 64)
	//	_, _, _, time2 := GetCntBlockGroup(10, i*100, 64)
	//	fmt.Print(i*100, " ", time1, " ", time2)
	//	fmt.Println()
	//}
	//_,_, time1 := GetPreBlockGroup(3, 1, 64)
	//fmt.Println(time1)
	//
	//_,_,_, time2 := GetCntBlockGroup(3, 1, 64)
	//fmt.Println(time2)

	for i := 0; i <= 1000; i += 2 {
		size1 := preNoTimeLine(7, 2500, 200, i*10)
		size4 := createTimeLine(2, 7, 2500, 200, i*10, 100, 100)
		fmt.Println(i*10, " ", size1/8, " ", size4/8, " ", float64(size1)/float64(size4))
	}

	//for i := 0; i <= 25; i++ {
	//	size1 := createTimeLine(2,7, 2500, 200, i * 10, 50, 50)
	//	size2 := createTimeLine(2,7, 2500, 200, i * 10, 50, 100)
	//	size3 := createTimeLine(2,7, 2500, 200, i * 10, 100, 50)
	//	size4 := createTimeLine(2,7, 2500, 200, i * 10, 100, 100)
	//	size5 := noTimeLine(7,2500, 200, i * 10)
	//	fmt.Println( i * 10, " ",size1/1024, " ", size2/1024, " ", size3/1024, " ", size4/1024, " ", size5/1024)
	//}
	//size1 := createTimeLine(1,4, 2500, 64, 100, 100, 100)
	//fmt.Println(size1)
	//
	//size2 := noTimeLine(4,2500,64,100)
	//fmt.Println( size2)
	//fmt.Println( 1.0* float64(size2)/ float64(size1))
}

func preNoTimeLine(n int, num int, length int, blockLen int) int {
	size2, _, _ := GetPreBlockGroup(n, num, length)
	return size2 * n * blockLen
}

func noTimeLine(n int, num int, length int, blockLen int) int {
	size2, _, _, _ := GetCntBlockGroup(n, num, length)
	return size2 * n * blockLen
}

func createTimeLine(f int, n int, txNum int, txLen int, blockLen int, hot int, warm int) int {
	size1, _, bg, _ := GetCntBlockGroup(n, txNum, txLen)
	if blockLen < hot {
		return size1 * blockLen * n
	}

	totalSize := 0
	totalSize += hot * size1 * n
	blockLen -= hot

	if blockLen < warm {
		var blTmp int = blockLen
		var cacheNum = 2*f + 1
		for blTmp >= 0 {
			for i := 0; i < warm/(2*f+1); i++ {
				totalSize += cacheNum * size1
				blTmp--
			}
			cacheNum--
		}
		totalSize += (size1 - SizeStruct(bg.Blocks)) * blockLen * n

		totalSize += SizeStruct(bg.Blocks[0].Timestamp) * blockLen * len(bg.Blocks) * n
		totalSize += SizeStruct(bg.Blocks[0].BlockNum) * blockLen * len(bg.Blocks) * n
		totalSize += SizeStruct(bg.Blocks[0].MerkleRoot) * blockLen * len(bg.Blocks) * n

		totalSize += SizeStruct((bg.Blocks[0].Transactions)) * len(bg.Blocks) * blockLen / (n - 2*f) * n

		return totalSize
	}

	totalSize += (2*f + 1) * (2*f + 2) / 2 * warm / (2*f + 1) * size1
	totalSize += (size1 - SizeStruct(bg.Blocks)) * warm * n
	totalSize += SizeStruct(bg.Blocks[0].Timestamp) * warm * len(bg.Blocks) * n
	totalSize += SizeStruct(bg.Blocks[0].BlockNum) * warm * len(bg.Blocks) * n
	totalSize += SizeStruct(bg.Blocks[0].MerkleRoot) * warm * len(bg.Blocks) * n
	totalSize += SizeStruct((bg.Blocks[0].Transactions)) * len(bg.Blocks) * warm / (n - 2*f) * n

	blockLen -= warm
	totalSize += (size1 - SizeStruct(bg.Blocks)) * blockLen * n
	totalSize += SizeStruct(bg.Blocks[0].Timestamp) * blockLen * len(bg.Blocks) * n
	totalSize += SizeStruct(bg.Blocks[0].BlockNum) * blockLen * len(bg.Blocks) * n
	totalSize += SizeStruct(bg.Blocks[0].MerkleRoot) * blockLen * len(bg.Blocks) * n
	totalSize += SizeStruct((bg.Blocks[0].Transactions)) * len(bg.Blocks) * blockLen / (n - 2*f) * n

	return totalSize
}

func TestBlsManager_Aggregate(t *testing.T) {
	m := Message("message to be signed. 将要做签名的消息")
	n := 8
	sks := make([]SecretKey, 0, n)
	pubs := make([]PublicKey, 0, n)
	sigs := make([]Signature, 0, n) //signatures for the same message
	msgs := make([]Message, 0, n)
	dsigs := make([]Signature, 0, n) //signatures for each (key,message) pair
	for i := 0; i < n; i++ {
		sk, pk := blsMgr.GenerateKey()
		sks = append(sks, sk)
		pubs = append(pubs, pk)
		sigs = append(sigs, sk.Sign(m))

		msgi := append(m, byte(i))
		msgs = append(msgs, msgi)
		dsigs = append(dsigs, sk.Sign(msgi))
	}
	fmt.Println(len(pubs[0].Compress().Bytes()))
	fmt.Println(len(sks[0].Compress()))

	asig, err := blsMgr.Aggregate(sigs)
	assert.NoError(t, err)
	// One
	err = blsMgr.VerifyAggregatedOne(pubs, m, asig)
	assert.NoError(t, err)

	apub, err := blsMgr.AggregatePublic(pubs)
	assert.NoError(t, err)

	err = apub.Verify(m, asig)
	assert.NoError(t, err)

	// N
	adsig, err := blsMgr.Aggregate(dsigs)
	assert.NoError(t, err)

	err = blsMgr.VerifyAggregatedN(pubs, msgs, adsig)
	assert.NoError(t, err)

	//lose some messages will cause an error
	err = blsMgr.VerifyAggregatedN(pubs, msgs[1:], adsig)
	assert.Error(t, err)

	//with out-of-order public keys, will has no effect on VerifyAggregatedOne, but DO effects VerifyAggregatedN
	pubs[0], pubs[1] = pubs[1], pubs[0]
	err = blsMgr.VerifyAggregatedOne(pubs, m, asig)
	assert.NoError(t, err)

	err = blsMgr.VerifyAggregatedN(pubs, msgs, adsig)
	assert.Error(t, err)

	//invalid length
	_, err = blsMgr.Aggregate(nil)
	assert.Error(t, err)
	_, err = blsMgr.AggregatePublic(make([]PublicKey, 0))
	assert.Error(t, err)
}

func TestRogueKey(t *testing.T) {
	rcpstr := "2be5782d6ecb09f0e4f6c995dc9ad470d552d650cafa960b9aca51005362d005e008aa6c7f348f3a6d238f275e9cfc1d2204c35c97b31c4d24a02d6c13bba5bd1249042ca7b83fd0be56a0001ab378afb03e90a9bbf6f08a63f74ffd267a0358"
	_, err := blsMgr.DecPublicKeyHex(rcpstr)
	assert.Error(t, err)
	t.Log(err)
}

func TestSignature_Compress(t *testing.T) {
	sk, _ := blsMgr.GenerateKey()
	m1 := Message("message to be signed. 将要做签名的消息")
	sig1 := sk.Sign(m1)

	cb := sig1.Compress()
	cb[0] = cb[0] + 0x01
	origin := sig1.Compress()
	assert.NotEqual(t, cb, origin)
}

//benchmark

func BenchmarkBLSAggregateSignature(b *testing.B) {
	msg := Message(">16 character identical message")
	n := 200
	sigs := make([]Signature, 0, n) //signatures for the same message
	for i := 0; i < n; i++ {
		sk, _ := blsMgr.GenerateKey()
		sigs = append(sigs, sk.Sign(msg))

	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blsMgr.Aggregate(sigs) //nolint:errcheck
	}
}

func BenchmarkBLSSign(b *testing.B) {
	sks := make([]SecretKey, b.N)
	msgs := make([]Message, 0, b.N)
	for i := range sks {
		sks[i], _ = blsMgr.GenerateKey()
		msgs = append(msgs, Message(fmt.Sprintf("Hello world! 16 characters %d", i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sks[i].Sign(msgs[i])
	}
}

func BenchmarkBLSVerify(b *testing.B) {
	sk, pk := blsMgr.GenerateKey()
	m := Message(">16 character identical message")
	sig := sk.Sign(m)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pk.Verify(m, sig) //nolint:errcheck
	}
}

func BenchmarkBlsManager_VerifyAggregatedOne(b *testing.B) {
	m := Message("message to be signed. 将要做签名的消息")
	n := 100
	//sks := make([]SecretKey, 0, n)
	pubs := make([]PublicKey, 0, n)
	sigs := make([]Signature, 0, n) //signatures for the same message
	for i := 0; i < n; i++ {
		sk, pk := blsMgr.GenerateKey()
		//sks = append(sks, sk)
		pubs = append(pubs, pk)
		sigs = append(sigs, sk.Sign(m))
	}
	asig, _ := blsMgr.Aggregate(sigs)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blsMgr.VerifyAggregatedOne(pubs, m, asig) //nolint:errcheck
	}
}

func BenchmarkBlsManager_VerifyAggregatedN(b *testing.B) {
	m := Message("message to be signed. 将要做签名的消息00000000000000000000000000000000000000000000000000000000000000000000000000000000")
	n := 100
	//sks := make([]SecretKey, 0, n)
	pubs := make([]PublicKey, 0, n)
	sigs := make([]Signature, 0, n)
	msgs := make([]Message, 0, n)
	for i := 0; i < n; i++ {
		mi := append(m, byte(i))
		sk, pk := blsMgr.GenerateKey()
		//sks = append(sks, sk)
		pubs = append(pubs, pk)
		sigs = append(sigs, sk.Sign(mi))
		msgs = append(msgs, mi)
	}
	asig, _ := blsMgr.Aggregate(sigs)
	fmt.Println(asig.Compress().Bytes())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blsMgr.VerifyAggregatedN(pubs, msgs, asig) //nolint:errcheck
	}
}

func BenchmarkBlsDecompressPublicKey(b *testing.B) {
	_, pk := blsMgr.GenerateKey()
	cpk := pk.Compress()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blsMgr.DecPublicKey(cpk.Bytes()) //nolint:errcheck
	}
}

func BenchmarkBlsManager_DecompressSignature(b *testing.B) {
	sk, _ := blsMgr.GenerateKey()
	m1 := Message("message to be signed. 将要做签名的消息")
	sig1 := sk.Sign(m1)
	cb := sig1.Compress()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blsMgr.DecSignature(cb.Bytes()) //nolint:errcheck
	}
}

func BenchmarkSignature_Compress(b *testing.B) {
	sk, _ := blsMgr.GenerateKey()
	m1 := Message("message to be signed. 将要做签名的消息")
	sig1 := sk.Sign(m1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sig1.Compress() //nolint:errcheck
	}
}

func BenchmarkSignature_serialize(b *testing.B) {
	sk, _ := blsMgr.GenerateKey()
	m1 := Message("message to be signed. 将要做签名的消息")
	sig1 := sk.Sign(m1)
	osig := sig1.(*signature)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		osig.serialize() //nolint:errcheck
	}
}
