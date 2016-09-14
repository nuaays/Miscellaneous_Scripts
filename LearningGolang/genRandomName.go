package main

import (
    "fmt"
    "io"
    "crypto/rand"
    "encoding/hex"
    "strings"
    "github.com/docker/libnetwork/types"
    "github.com/vishvananda/netlink"
)



func GenerateRandomName(prefix string, size int) (string, error) {
        id := make([]byte, 32)
        if _, err := io.ReadFull(rand.Reader, id); err != nil {
                return "", err
        }
        return prefix + hex.EncodeToString(id)[:size], nil
}


func GenerateIfaceName(prefix string, len int) (string, error) {
	for i := 0; i < 3; i++ {
		name, err := GenerateRandomName(prefix, len)
		if err != nil {
			continue
		}
		if res, err := netlink.LinkByName(name); err != nil {
                        fmt.Println("!!!", res)
			if strings.Contains(err.Error(), "not found") {
				return name, nil
			}
			return "", err
		}
	}
	return "", types.InternalErrorf("could not generate interface name")
}






func main(){

        res, _ := GenerateRandomName("Prefix_", 10)
        fmt.Println(res)
        ret, err := GenerateIfaceName("Prefix_", 10)
        fmt.Println(ret, err)
 
}
