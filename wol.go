package main

import (
    "encoding/hex"
    "net"
    "strings"
)

func sendWOL(macAddr string) error {
    macAddr = strings.Replace(macAddr, ":", "", -1)
    macAddr = strings.Replace(macAddr, "-", "", -1)

    macBytes, err := hex.DecodeString(macAddr)
    if err != nil {
        return err
    }

    var packet []byte
    // Create magic packet: 6 x 0xFF followed by MAC 16 times
    for i := 0; i < 6; i++ {
        packet = append(packet, 0xFF)
    }
    for i := 0; i < 16; i++ {
        packet = append(packet, macBytes...)
    }

    conn, err := net.Dial("udp", "255.255.255.255:9")
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Write(packet)
    return err
}
