package main
import ("encoding/base32"
        "fmt")

func main() {
    //var enc = base32.NewEncoding(base32.encodeHex)
    s:=base32.HexEncoding.EncodeToString([]byte("violetsss~~"))
    fmt.Printf("E: %s",s)
    r := base32.HexEncoding.DecodeString(s)
    if (e!=nil) {
        fmt.Print("error")
    }
    fmt.Printf("D: %s", r)
}
