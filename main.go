package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

func authHash() string {
    curr_time := time.Now().Unix()
    num := make([]byte, 8)
    num = binary.LittleEndian.AppendUint64(num, uint64(curr_time))
    sha := sha256.New()
    sha.Write(num)
    hashed := sha.Sum(nil)
    return hex.EncodeToString(hashed)
}

func main() {
    fmt.Println(authHash())
}
