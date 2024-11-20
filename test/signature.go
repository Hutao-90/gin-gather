package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// 生成私钥
func generatePrivateKey() (*ecdsa.PrivateKey, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return privateKey, fmt.Sprintf("%x", privateKeyBytes)
}

func main() {
	// 假设有一个私钥
	//privateKeyHex := "your_private_key_hex" // 这里替换为你的私钥（32字节的hex格式）
	privateKey, privateKeyHex := generatePrivateKey()
	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// 要签署的消息
	message := "This is a test message"

	// 生成消息哈希
	messageHash := crypto.Keccak256Hash([]byte(message))

	// 使用私钥对消息哈希签名
	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 打印签名
	fmt.Printf("Signature: 0x%x\n", signature)
}
