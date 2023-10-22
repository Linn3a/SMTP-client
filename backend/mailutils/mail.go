package mailutils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"mail/model"
	"net"
)

var (
	mailServer  = "smtp.163.com"
	serverPort  = 25
	UserName    = "cos"
	User        = "co_sin_o@163.com"
	Password    = "****"
	fromAddress = "co_sin_o@163.com"
)

func SendMail(
	to []*model.User,
	subject string, content string) {
	// Connect to the SMTP server
	client, err := net.Dial("tcp", fmt.Sprintf("%s:%d", mailServer, serverPort))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(client net.Conn) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)

	reader := bufio.NewReader(client)
	receiveResponse(reader)

	// Send HELO command
	sendCommand(client, "HELO localhost")

	// Authenticate
	sendCommand(client, "AUTH LOGIN")
	sendCommand(client, base64.StdEncoding.EncodeToString([]byte(User)))
	sendCommand(client, base64.StdEncoding.EncodeToString([]byte(Password)))

	// Send MAIL FROM command
	sendCommand(client, fmt.Sprintf("MAIL FROM: <%s>", fromAddress))

	// Send RCPT TO command
	for _, User := range to {
		fmt.Printf("User:%v\n", User)
		sendCommand(client, fmt.Sprintf("RCPT TO: <%s>", User.Address))
		// Send DATA command
		sendCommand(client, "DATA")

		// Send email message
		endmsg := "\r\n.\r\n"

		msg := fmt.Sprintf("From: %s\r\n", UserName)
		msg += fmt.Sprintf("To: %s\r\n", User.Username)
		msg += fmt.Sprintf("Subject: %s\r\n", subject)
		msg += "Content-Type: text/plain\t\n"
		msg += "MIME-Version: 1.0\r\n"
		msg += fmt.Sprintf("\r\n%s\r\n", content)
		msg += "."

		sendCommand(client, msg)

		// End the email message
		sendCommand(client, endmsg)
	}

	// Quit the session
	sendCommand(client, "QUIT")
}

func receiveResponse(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(line)
		if line[3] == ' ' {
			break
		}
	}
}

func sendCommand(client net.Conn, command string) {
	_, err := client.Write([]byte(command + "\r\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	receiveResponse(bufio.NewReader(client))
}
