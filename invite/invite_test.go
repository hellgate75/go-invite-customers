/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package invite

import (
	io2 "github.com/hellgate75/go-invite-customers/io"
	"github.com/hellgate75/go-invite-customers/model"
	"io"
	"testing"
)

func TestExecuteInviteScan(t *testing.T) {
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
	type args struct {
		input InputData
	}
	tests := []struct {
		name     string
		args     args
		wantOut  OutputData
		wantErrs []error
	}{
		{
			name: "Test url case",
			args: args{
				input: InputData{
					UseDetailedOutput: false,
					Distance:          100,
					MeasureUnit:       "K",
					HomeLongitude:     -6.257664,
					HomeLatitude:      53.339428,
					InputEncoding:     io2.JsonEncoding,
					OutputEncoding:    io2.JsonEncoding,
					SilentOutput:      true,
					FileOrStream:      "https://s3.amazonaws.com/intercom-take-home-test/customers.txt",
					UsePerLineInput:   true,
				},
			},
			wantOut: OutputData{
				IsDone:     true,
				IsComplete: false,
				Complete:   model.NewCompleteInviteList(),
				Simple: &model.InviteList{
					CustomerIds: []model.CustomerDetails{
						{UserId: 12, Name: "Christina McArdle"},
						{UserId: 8, Name: "Eoin Ahearn"},
						{UserId: 26, Name: "Stephen McArdle"},
						{UserId: 6, Name: "Theresa Enright"},
						{UserId: 4, Name: "Ian Kehoe"},
						{UserId: 5, Name: "Nora Dempsey"},
						{UserId: 11, Name: "Richard Finnegan"},
						{UserId: 31, Name: "Alan Behan"},
						{UserId: 13, Name: "Olive Ahearn"},
						{UserId: 15, Name: "Michael Ahearn"},
						{UserId: 17, Name: "Patricia Cahill"},
						{UserId: 39, Name: "Lisa Ahearn"},
						{UserId: 24, Name: "Rose Enright"},
						{UserId: 29, Name: "Oliver Ahearn"},
						{UserId: 30, Name: "Nick Enright"},
						{UserId: 23, Name: "Eoin Gallagher"},
					},
				},
			},
			wantErrs: make([]error, 0),
		},
		{
			name: "Test file case",
			args: args{
				input: InputData{
					UseDetailedOutput: false,
					Distance:          100,
					MeasureUnit:       "K",
					HomeLongitude:     -6.257664,
					HomeLatitude:      53.339428,
					InputEncoding:     io2.JsonEncoding,
					OutputEncoding:    io2.JsonEncoding,
					SilentOutput:      true,
					FileOrStream:      name,
					UsePerLineInput:   true,
				},
			},
			wantOut: OutputData{
				IsDone:     true,
				IsComplete: false,
				Complete:   model.NewCompleteInviteList(),
				Simple: &model.InviteList{
					CustomerIds: []model.CustomerDetails{
						{UserId: 1, Name: "Thomas Barret"},
					},
				},
			},
			wantErrs: make([]error, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, gotErrs := ExecuteInviteScan(tt.args.input)
			if len(gotErrs) != len(tt.wantErrs) {
				t.Errorf("ExecuteInviteScan() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
			if len(gotOut.Simple.CustomerIds) != len(tt.wantOut.Simple.CustomerIds) {
				t.Errorf("ExecuteInviteScan() gotOut Simple.CustomerIds = %+v, want %+v", gotOut.Simple.CustomerIds, tt.wantOut.Simple.CustomerIds)
			}
			if len(gotOut.Complete.MatchingCustomerIds) != len(tt.wantOut.Complete.MatchingCustomerIds) {
				t.Errorf("ExecuteInviteScan() gotOut Complete.MatchingCustomerIds = %+v, want %+v", gotOut.Complete.MatchingCustomerIds, tt.wantOut.Complete.MatchingCustomerIds)
			}
			if len(gotOut.Complete.UnMatchingCustomerIds) != len(tt.wantOut.Complete.UnMatchingCustomerIds) {
				t.Errorf("ExecuteInviteScan() gotOut Complete.UnMatchingCustomerIds = %+v, want %+v", gotOut.Complete.UnMatchingCustomerIds, tt.wantOut.Complete.UnMatchingCustomerIds)
			}
		})
	}
}

func Test_createChannelWriterFunc(t *testing.T) {
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
	type args struct {
		url string
	}
	tests := []struct {
		name         string
		args         args
		wantFunction bool
		wantErr      bool
	}{
		{
			name: "Test UDP Connect function",
			args: args{
				url: "udp://8.8.8.8:53",
			},
			wantErr:      false,
			wantFunction: true,
		},
		{
			name: "Test TCP Connect function",
			args: args{
				url: "tcp://8.8.8.8:53",
			},
			wantErr:      false,
			wantFunction: true,
		},
		{
			name: "Test http Connect function",
			args: args{
				url: "https://s3.amazonaws.com/intercom-take-home-test/customers.txt",
			},
			wantErr:      false,
			wantFunction: true,
		},
		{
			name: "Test File Connect function",
			args: args{
				url: name,
			},
			wantErr:      false,
			wantFunction: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunction, err := createChannelWriterFunc(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("createChannelWriterFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotFunction != nil) != tt.wantFunction {
				t.Errorf("createChannelWriterFunc() gotFunction = %v, want %v", gotFunction != nil, tt.wantFunction)
			}
		})
	}
}

func Test_parseAndServerList(t *testing.T) {
	file, err := CreateTestListFile()
	if err != nil {
		t.Errorf("OpenFileStream() error = %v, opening main file stream", err)
		return
	}
	name := file.Name()
	defer func() {
		_ = file.Close()
		_ = DeleteTestFile(name)
	}()
	ch := make(chan model.CustomerOffice, 1000)
	errCh := make(chan error, 1000)
	defer func() {
		close(ch)
		close(errCh)
	}()
	type args struct {
		r         io.Reader
		inputData InputData
		ch        chan model.CustomerOffice
		errCh     chan error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test reading from file stream",
			args: args{
				inputData: InputData{
					UseDetailedOutput: false,
					Distance:          100,
					MeasureUnit:       "K",
					HomeLongitude:     -6.257664,
					HomeLatitude:      53.339428,
					InputEncoding:     io2.JsonEncoding,
					OutputEncoding:    io2.JsonEncoding,
					SilentOutput:      true,
					FileOrStream:      name,
					UsePerLineInput:   false,
				},
				r:     file,
				ch:    ch,
				errCh: errCh,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseAndServerList(tt.args.r, tt.args.inputData, tt.args.ch, tt.args.errCh)
		})
	}
}

func Test_readLineByLine(t *testing.T) {
	file, err := CreateTestFile()
	if err != nil {
		t.Errorf("OpenFileStream() error = %v, opening main file stream", err)
		return
	}
	name := file.Name()
	defer func() {
		_ = file.Close()
		_ = DeleteTestFile(name)
	}()
	ch := make(chan model.CustomerOffice, 1000)
	errCh := make(chan error, 1000)
	defer func() {
		close(ch)
		close(errCh)
	}()
	type args struct {
		r         io.Reader
		inputData InputData
		ch        chan model.CustomerOffice
		errCh     chan error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test reading from file stream",
			args: args{
				inputData: InputData{
					UseDetailedOutput: false,
					Distance:          100,
					MeasureUnit:       "K",
					HomeLongitude:     -6.257664,
					HomeLatitude:      53.339428,
					InputEncoding:     io2.JsonEncoding,
					OutputEncoding:    io2.JsonEncoding,
					SilentOutput:      true,
					FileOrStream:      name,
					UsePerLineInput:   true,
				},
				r:     file,
				ch:    ch,
				errCh: errCh,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readLineByLine(tt.args.r, tt.args.inputData, tt.args.ch, tt.args.errCh)
		})
	}
}
