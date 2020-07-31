/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package invite

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

//  Open Url from given path
//
//  Url/
//  Input url (http://...., https://....., ftp://...., sftp://.....
//
//  The output are the call response  (to close), the input reader and the error, if any error occurs during the stream opening operation.
func OpenUrlStream(url string) (resp *http.Response, reader io.Reader, err error) {
	if url == "" {
		return nil, nil, errors.New(fmt.Sprint("Empty http/ftp url"))
	}
	if url[:7] != "http://" && url[:8] != "https://" && url[:6] != "ftp://" && url[:7] != "sftp://" {
		return nil, nil, errors.New(fmt.Sprintf("Invalid http/ftp url: %s", url))
	}
	resp, err = http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	reader = resp.Body
	return resp, reader, err
}

//  Open File from given file path
//
//  File/
//  File or Pipe absolute or relative path
//
//  The output are the the input reader and the error, if any error occurs during the stream opening operation.
func OpenFileStream(file string) (*os.File, error) {
	if file == "" {
		return nil, errors.New(fmt.Sprint("Empty file name"))
	}
	return os.Open(file)
}

//  Open Tcp Server at port from given path
//
//  TcpUrl/
//  Input url (tcp://host:port => tcp://192.168.0.1:19099 or tcp://my-dna-name:19099
//
//  The output are the call connection (to close), the input reader and the error, if any error occurs during the stream opening operation.
func OpenTcpStream(tcpUrl string) (conn net.Conn, reader io.Reader, err error) {
	if tcpUrl == "" {
		return nil, nil, errors.New(fmt.Sprint("Empty tcp url"))
	}
	if tcpUrl[:6] != "tcp://" {
		return nil, nil, errors.New(fmt.Sprintf("Invalid tcp protocol in url: %s", tcpUrl))
	}
	if len(tcpUrl) < 16 {
		return nil, nil, errors.New(fmt.Sprintf("Invalid tcp url: %s", tcpUrl))
	}
	var url = tcpUrl[6:]
	conn, err = net.Dial("tcp", url)
	if err != nil {
		return nil, nil, err
	}
	reader = conn
	return conn, reader, err
}

//  Open Udp Server at port from given path
//
//  UdpUrl/
//  Input url (udp://host:port => udp://192.168.0.1:19099 or udp://my-dna-name:19099
//
//  The output are the call connection (to close), the input reader and the error, if any error occurs during the stream opening operation.
func OpenUdpStream(udpUrl string) (conn net.Conn, reader io.Reader, err error) {
	if udpUrl == "" {
		return nil, nil, errors.New(fmt.Sprint("Empty udp url"))
	}
	if udpUrl[:6] != "udp://" {
		return nil, nil, errors.New(fmt.Sprintf("Invalid udp protocol in url: %s", udpUrl))
	}
	if len(udpUrl) < 16 {
		return nil, nil, errors.New(fmt.Sprintf("Invalid udp url: %s", udpUrl))
	}
	var url = udpUrl[6:]
	conn, err = net.Dial("udp", url)
	if err != nil {
		return nil, nil, err
	}
	reader = conn
	return conn, reader, err
}
