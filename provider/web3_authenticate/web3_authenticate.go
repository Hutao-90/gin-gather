package web3_authenticate

import (
	messagePackages "debox/message"
	"debox/provider/jwt"
	"encoding/hex"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func Authenticate(walletAddress string, signature, message string) (string, error) {

	// 去掉 0x 前缀
	if len(signature) > 2 && signature[:2] == "0x" {
		signature = signature[2:]
	}

	// 将签名从十六进制字符串转换为字节数组
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return "", errors.New(messagePackages.SignatureError)
	}

	// 确保签名长度为 65 字节
	if len(sigBytes) != 65 {
		return "", errors.New(messagePackages.SignatureError)
	}

	// 生成消息的哈希值
	hash := crypto.Keccak256Hash([]byte(message)).Bytes()

	// 从签名中恢复公钥
	sigPublicKey, err := crypto.SigToPub(hash, sigBytes)
	if err != nil {
		return "", err
	}

	// 从公钥中恢复钱包地址
	recoveredAddr := crypto.PubkeyToAddress(*sigPublicKey).Hex()

	// 验证签名
	if recoveredAddr != walletAddress {
		return "", errors.New(recoveredAddr)
	}

	// 生成 JWT
	token, err := jwt.GenerateJWT(walletAddress, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
