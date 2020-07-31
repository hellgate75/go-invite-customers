/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package io

import (
	"github.com/hellgate75/go-invite-customers/model"
	"reflect"
	"testing"
)

func TestEncodeCustomerDetailedInvite(t *testing.T) {
	type args struct {
		invite model.CompleteInviteList
		enc    Encoding
	}
	inviteList := *model.NewCompleteInviteList()
	inviteList.MatchingCustomerIds = []model.CustomerDetails{{1, "Thomas Barret"}}
	inviteList.UnMatchingCustomerIds = []model.CustomerDetails{{2, "Michael Barret"}}

	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name: "Encode a valid model.CompleteInviteList to JSON format",
			args: args{
				invite: inviteList,
				enc:    JsonEncoding,
			},
			wantErr:  false,
			wantData: []byte("{\"customers_list\":[{\"user_id\":1,\"name\":\"Thomas Barret\"}],\"exclusions_list\":[{\"user_id\":2,\"name\":\"Michael Barret\"}]}"),
		},
		{
			name: "Encode a valid model.CompleteInviteList to YAML format",
			args: args{
				invite: inviteList,
				enc:    YamlEncoding,
			},
			wantErr: false,
			wantData: []byte(`customers_list:
- user_id: 1
  name: Thomas Barret
exclusions_list:
- user_id: 2
  name: Michael Barret
`),
		},
		{
			name: "Encode a valid model.CompleteInviteList to XML format",
			args: args{
				invite: inviteList,
				enc:    XmlEncoding,
			},
			wantErr:  false,
			wantData: []byte("<CompleteInviteList><customers-list><user-id>1</user-id><name>Thomas Barret</name></customers-list><exclusions-list><user-id>2</user-id><name>Michael Barret</name></exclusions-list></CompleteInviteList>"),
		},
		{
			name: "Encode a valid model.CompleteInviteList to Text format",
			args: args{
				invite: inviteList,
				enc:    TextEncoding,
			},
			wantErr: false,
			wantData: []byte(`Invite Summary:
[1] Thomas Barret
Exclusion Summary:
[1] Thomas Barret
`),
		},
		{
			name: "Not Encode a valid model.CompleteInviteList to Unknown format",
			args: args{
				invite: inviteList,
				enc:    UnknownEncoding,
			},
			wantErr:  true,
			wantData: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := EncodeCustomerDetailedInvite(tt.args.invite, tt.args.enc)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeCustomerDetailedInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("EncodeCustomerDetailedInvite() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestEncodeCustomerInvite(t *testing.T) {
	type args struct {
		invite model.InviteList
		enc    Encoding
	}
	inviteList := *model.NewInviteList()
	inviteList.CustomerIds = []model.CustomerDetails{{1, "Thomas Barret"}}

	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			name: "Encode a valid model.CompleteInviteList to JSON format",
			args: args{
				invite: inviteList,
				enc:    JsonEncoding,
			},
			wantErr:  false,
			wantData: []byte("{\"customers_list\":[{\"user_id\":1,\"name\":\"Thomas Barret\"}]}"),
		},
		{
			name: "Encode a valid model.CompleteInviteList to YAML format",
			args: args{
				invite: inviteList,
				enc:    YamlEncoding,
			},
			wantErr: false,
			wantData: []byte(`customers_list:
- user_id: 1
  name: Thomas Barret
`),
		},
		{
			name: "Encode a valid model.CompleteInviteList to XML format",
			args: args{
				invite: inviteList,
				enc:    XmlEncoding,
			},
			wantErr:  false,
			wantData: []byte("<InviteList><customers-list><user-id>1</user-id><name>Thomas Barret</name></customers-list></InviteList>"),
		},
		{
			name: "Encode a valid model.CompleteInviteList to Text format",
			args: args{
				invite: inviteList,
				enc:    TextEncoding,
			},
			wantErr: false,
			wantData: []byte(`Invite Summary:
[1] Thomas Barret
`),
		},
		{
			name: "Not Encode a valid model.CompleteInviteList to Unknown format",
			args: args{
				invite: inviteList,
				enc:    UnknownEncoding,
			},
			wantErr:  true,
			wantData: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := EncodeCustomerInvite(tt.args.invite, tt.args.enc)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeCustomerInvite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("EncodeCustomerInvite() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestReadCustomerOffice(t *testing.T) {
	type args struct {
		data []byte
		enc  Encoding
	}
	customer := model.CustomerOffice{
		UserId:    1,
		Name:      "Thomas Barrett",
		Latitude:  "10.123456",
		Longitude: "-5.98765",
	}
	tests := []struct {
		name         string
		args         args
		wantCustomer model.CustomerOffice
		wantErr      bool
	}{
		{
			name: "Test Json Import single model.CustomerOffice data",
			args: args{
				data: []byte("{\"user_id\":1,\"name\":\"Thomas Barrett\",\"latitude\":\"10.123456\",\"longitude\":\"-5.98765\"}"),
				enc:  JsonEncoding,
			},
			wantCustomer: customer,
			wantErr:      false,
		},
		{
			name: "Test Yaml Import single model.CustomerOffice data",
			args: args{
				data: []byte(`
user_id: 1
name: "Thomas Barrett"
latitude: "10.123456"
longitude: "-5.98765"
`),
				enc: YamlEncoding,
			},
			wantCustomer: customer,
			wantErr:      false,
		},
		{
			name: "Test Xml Import single model.CustomerOffice data",
			args: args{
				data: []byte("<CustomerOffice><user-id>1</user-id><name>Thomas Barrett</name><latitude>10.123456</latitude><longitude>-5.98765</longitude></CustomerOffice>"),
				enc:  XmlEncoding,
			},
			wantCustomer: customer,
			wantErr:      false,
		},
		{
			name: "Test Unknown Import single model.CustomerOffice data",
			args: args{
				data: []byte{},
				enc:  UnknownEncoding,
			},
			wantCustomer: model.CustomerOffice{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCustomer, err := ReadCustomerOffice(tt.args.data, tt.args.enc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCustomerOffice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("ReadCustomerOffice() gotCustomer = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func TestToEncoding(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantEnc Encoding
		wantErr bool
	}{
		{
			name: "Transform correct encoding format text",
			args: args{
				in: "json",
			},
			wantEnc: JsonEncoding,
			wantErr: false,
		},
		{
			name: "Transform correct case sensitive encoding format text",
			args: args{
				in: "YaMl",
			},
			wantEnc: YamlEncoding,
			wantErr: false,
		},
		{
			name: "Transform correct xml encoding format text",
			args: args{
				in: "xml",
			},
			wantEnc: XmlEncoding,
			wantErr: false,
		},
		{
			name: "Transform correct text encoding format text",
			args: args{
				in: "text",
			},
			wantEnc: TextEncoding,
			wantErr: false,
		},
		{
			name: "Transform incorrect encoding format text",
			args: args{
				in: "bin",
			},
			wantEnc: UnknownEncoding,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEnc, err := ToEncoding(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEnc != tt.wantEnc {
				t.Errorf("ToEncoding() gotEnc = %v, want %v", gotEnc, tt.wantEnc)
			}
		})
	}
}

func Test_textEncodeCompleteInviteList(t *testing.T) {
	inviteList := model.CompleteInviteList{
		MatchingCustomerIds:   []model.CustomerDetails{{1, "Thomas Barret"}},
		UnMatchingCustomerIds: []model.CustomerDetails{{2, "Michael Barret"}},
	}
	type args struct {
		list model.CompleteInviteList
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{
			name: "Test Text Encode model.CompleteInviteList in Json format",
			args: args{
				list: inviteList,
			},
			wantErr: false,
			wantOut: []byte(`Invite Summary:
[1] Thomas Barret
Exclusion Summary:
[1] Thomas Barret
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := textEncodeCompleteInviteList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("textEncodeCompleteInviteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("textEncodeCompleteInviteList() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_textEncodeInviteList(t *testing.T) {
	inviteList := model.InviteList{
		CustomerIds: []model.CustomerDetails{{1, "Thomas Barret"}},
	}
	type args struct {
		list model.InviteList
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{
			name: "Test Text Encode model.CompleteInviteList in Json format",
			args: args{
				list: inviteList,
			},
			wantErr: false,
			wantOut: []byte(`Invite Summary:
[1] Thomas Barret
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := textEncodeInviteList(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("textEncodeInviteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("textEncodeInviteList() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestReadCustomerOfficeList(t *testing.T) {
	type args struct {
		data []byte
		enc  Encoding
	}
	customer := model.CustomerOffice{
		UserId:    1,
		Name:      "Thomas Barrett",
		Latitude:  "10.123456",
		Longitude: "-5.98765",
	}
	customerList := model.CustomerOfficeList{
		List: []model.CustomerOffice{customer},
	}
	tests := []struct {
		name         string
		args         args
		wantCustomer model.CustomerOfficeList
		wantErr      bool
	}{
		{
			name: "Test Json Import single model.CustomerOfficeList data",
			args: args{
				data: []byte("{\"customers\":[{\"user_id\":1,\"name\":\"Thomas Barrett\",\"latitude\":\"10.123456\",\"longitude\":\"-5.98765\"}]}"),
				enc:  JsonEncoding,
			},
			wantCustomer: customerList,
			wantErr:      false,
		},
		{
			name: "Test Yaml Import single model.CustomerOfficeList data",
			args: args{
				data: []byte(`
customers:
- user_id: 1
  name: Thomas Barrett
  latitude: "10.123456"
  longitude: "-5.98765"
`),
				enc: YamlEncoding,
			},
			wantCustomer: customerList,
			wantErr:      false,
		},
		{
			name: "Test Xml Import single model.CustomerOfficeList data",
			args: args{
				data: []byte("<CustomerOfficeList><customers><user-id>1</user-id><name>Thomas Barrett</name><latitude>10.123456</latitude><longitude>-5.98765</longitude></customers></CustomerOfficeList>"),
				enc:  XmlEncoding,
			},
			wantCustomer: customerList,
			wantErr:      false,
		},
		{
			name: "Test Unknown Import single model.CustomerOfficeList data",
			args: args{
				data: []byte{},
				enc:  UnknownEncoding,
			},
			wantCustomer: model.CustomerOfficeList{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCustomer, err := ReadCustomerOfficeList(tt.args.data, tt.args.enc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCustomerOfficeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("ReadCustomerOfficeList() gotCustomer = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}
