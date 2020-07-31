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
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"testing"
)

func CreateTestFile() (*os.File, error) {
	var data = make([]byte, 0)
	data = append(data, []byte("{\"latitude\": \"53.339111\", \"user_id\": 12, \"name\": \"Thomas Barret\", \"longitude\": \"-6.257611\"}\n")...)
	data = append(data, []byte("{\"latitude\": \"50.339428\", \"user_id\": 1, \"name\": \"Michael Barret\", \"longitude\": \"-3.257664\"}\n")...)
	file, err := ioutil.TempFile("", uuid.New().String())
	if err != nil {
		return nil, err
	}
	file.Write(data)
	err = file.Sync()
	if err != nil {
		return nil, err
	}
	return file, err
}

func CreateTestListFile() (*os.File, error) {
	var data = []byte("{\"customers\":[{\"user_id\":1,\"name\":\"Thomas Barrett\",\"latitude\":\"10.123456\",\"longitude\":\"-5.98765\"}]}\n\n")
	file, err := ioutil.TempFile("", uuid.New().String())
	if err != nil {
		return nil, err
	}
	file.Write(data)
	err = file.Sync()
	if err != nil {
		return nil, err
	}
	return file, err
}

func DeleteTestFile(name string) error {
	if name != "" {
		return os.Remove(name)
	}
	return errors.New(fmt.Sprintf("Invalid file path: %s", name))
}

func TestOpenTcpStream(t *testing.T) {
	type args struct {
		tcpUrl string
	}
	tests := []struct {
		name       string
		args       args
		wantConn   bool
		wantReader bool
		wantErr    bool
	}{
		{
			name: "Test valid tcp url",
			args: args{
				tcpUrl: "tcp://8.8.8.8:53",
			},
			wantErr:    false,
			wantConn:   true,
			wantReader: true,
		},
		{
			name: "Test empty tcp url",
			args: args{
				tcpUrl: "",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
		{
			name: "Test Invalid tcp url",
			args: args{
				tcpUrl: "tcp://8.8.8.8:100012",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
		{
			name: "Test Invalid tcp url format",
			args: args{
				tcpUrl: "udp://8.8.8.8:53",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConn, gotReader, err := OpenTcpStream(tt.args.tcpUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenUdpStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && gotConn != nil {
				defer func() {
					_ = gotConn.Close()
				}()
			}
			if (gotConn != nil) != tt.wantConn {
				t.Errorf("OpenUdpStream() gotConn = %v, want %v", gotConn, tt.wantConn)
			}
			if (gotReader != nil) != tt.wantReader {
				t.Errorf("OpenUdpStream() gotReader = %v, want %v", gotReader, tt.wantReader)
			}
		})
	}
}

func TestOpenUdpStream(t *testing.T) {
	type args struct {
		udpUrl string
	}
	tests := []struct {
		name       string
		args       args
		wantConn   bool
		wantReader bool
		wantErr    bool
	}{
		{
			name: "Test Google dns url",
			args: args{
				udpUrl: "udp://8.8.8.8:53",
			},
			wantErr:    false,
			wantConn:   true,
			wantReader: true,
		},
		{
			name: "Test empty udp url",
			args: args{
				udpUrl: "",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
		{
			name: "Test Invalid udp url",
			args: args{
				udpUrl: "udp://8.8.8.8:100012",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
		{
			name: "Test Invalid udp url format",
			args: args{
				udpUrl: "tcp://8.8.8.8:53",
			},
			wantErr:    true,
			wantConn:   false,
			wantReader: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConn, gotReader, err := OpenUdpStream(tt.args.udpUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenUdpStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && gotConn != nil {
				defer func() {
					_ = gotConn.Close()
				}()
			}
			if (gotConn != nil) != tt.wantConn {
				t.Errorf("OpenUdpStream() gotConn = %v, want %v", gotConn, tt.wantConn)
			}
			if (gotReader != nil) != tt.wantReader {
				t.Errorf("OpenUdpStream() gotReader = %v, want %v", gotReader, tt.wantReader)
			}
		})
	}
}

func TestOpenUrlStream(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name       string
		args       args
		wantResp   bool
		wantReader bool
		wantErr    bool
	}{
		{
			name: "Test open valid url data stream",
			args: args{
				url: "https://s3.amazonaws.com/intercom-take-home-test/customers.txt",
			},
			wantErr:    false,
			wantReader: true,
			wantResp:   true,
		},
		{
			name: "Test open empty url data stream",
			args: args{
				url: "",
			},
			wantErr:    true,
			wantReader: false,
			wantResp:   false,
		},
		{
			name: "Test open invalid url data stream",
			args: args{
				url: "https://myserver1234/stream",
			},
			wantErr:    true,
			wantReader: false,
			wantResp:   false,
		},
		{
			name: "Test open invalid url data stream format",
			args: args{
				url: "htp://www.google.com",
			},
			wantErr:    true,
			wantReader: false,
			wantResp:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotReader, err := OpenUrlStream(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFileStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && gotResp != nil {
				defer func() {
					_ = gotResp.Body.Close()
				}()
			}
			if (gotResp != nil) != tt.wantResp {
				t.Errorf("OpenUrlStream() gotResp = %v, wantResp %v", gotResp, tt.wantResp)
			}
			if (gotReader != nil) != tt.wantReader {
				t.Errorf("OpenUrlStream() gotReader = %v, wantReader %v", gotReader, tt.wantReader)
			}
		})
	}
}

func TestOpenFileStream(t *testing.T) {
	type args struct {
		file string
	}
	file, err := CreateTestFile()
	if err != nil {
		t.Errorf("OpenFileStream() error = %v, opening main file stream", err)
		return
	}
	name := file.Name()
	defer func() {
		_ = DeleteTestFile(name)
	}()
	_ = file.Close()
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test valid file path stream opening",
			args: args{
				file: name,
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "Test empty file path stream opening",
			args: args{
				file: "",
			},
			wantErr: true,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenFileStream(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFileStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != nil {
				defer func() {
					_ = got.Close()
				}()
			}
			if (got != nil) != tt.want {
				t.Errorf("OpenFileStream() got = %v, want %v", got, tt.want)
			}
		})
	}
}
