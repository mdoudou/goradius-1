package goradius

import (
	"log"
	"net"
)

func Server(conn *net.UDPConn) {

	defer conn.Close()
	buf := make([]byte, 4096)

	_, udpaddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	ipp := udpaddr.IP.String()
	secret := GetSecretAndDiff(ipp)
	if secret == nil {
		return
	}
	pakage, _ := Parse(buf, secret)
	userName := UserName_GetString(pakage)
	uerPassword := UserPassword_GetString(pakage)
	log.Println("username=", userName, "userpassword=", uerPassword, "nasIP=", ipp)

	if GetUserPasswd(uerPassword) == uerPassword {
		res := pakage.Response(CodeAccessAccept)
		var vl = []byte{'o', 'k'}
		ReplyMessage_Add(res, vl)
		udpp, _ := res.Encode()
		conn.WriteToUDP(udpp, udpaddr)
	} else {
		res := pakage.Response(CodeAccessReject)
		var vl = []byte{'n', 'o'}
		ReplyMessage_Add(res, vl)
		udpp, _ := res.Encode()
		conn.WriteToUDP(udpp, udpaddr)
	}

}
