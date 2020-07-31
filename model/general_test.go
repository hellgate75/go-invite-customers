/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package model

import (
	"reflect"
	"testing"
)

func TestCompleteInviteList_AddExcluded(t *testing.T) {
	type fields struct {
		MatchingCustomerIds   []CustomerDetails
		UnMatchingCustomerIds []CustomerDetails
	}
	type args struct {
		customerId *CustomerDetails
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test insert nil excluded customer element",
			fields: fields{
				make([]CustomerDetails, 0),
				make([]CustomerDetails, 0),
			},
			args: args{
				nil,
			},
			want: false,
		},
		{
			name: "Test insert sample excluded customer element",
			fields: fields{
				make([]CustomerDetails, 0),
				nil,
			},
			args: args{
				ToInviteData(&CustomerOffice{
					1,
					"James Barrett",
					"10.27654343",
					"-5.2653335",
				}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &CompleteInviteList{
				MatchingCustomerIds:   tt.fields.MatchingCustomerIds,
				UnMatchingCustomerIds: tt.fields.UnMatchingCustomerIds,
			}
			if got := il.AddExcluded(tt.args.customerId); got != tt.want {
				t.Errorf("AddExcluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompleteInviteList_AddInvited(t *testing.T) {
	type fields struct {
		MatchingCustomerIds   []CustomerDetails
		UnMatchingCustomerIds []CustomerDetails
	}
	type args struct {
		customerId *CustomerDetails
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test insert nil included customer element",
			fields: fields{
				make([]CustomerDetails, 0),
				make([]CustomerDetails, 0),
			},
			args: args{
				nil,
			},
			want: false,
		},
		{
			name: "Test insert sample included customer element",
			fields: fields{
				nil,
				make([]CustomerDetails, 0),
			},
			args: args{
				ToInviteData(&CustomerOffice{
					1,
					"James Barrett",
					"10.27654343",
					"-5.2653335",
				}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &CompleteInviteList{
				MatchingCustomerIds:   tt.fields.MatchingCustomerIds,
				UnMatchingCustomerIds: tt.fields.UnMatchingCustomerIds,
			}
			if got := il.AddInvited(tt.args.customerId); got != tt.want {
				t.Errorf("AddInvited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInviteList_Add(t *testing.T) {
	type fields struct {
		CustomerIds []CustomerDetails
	}
	type args struct {
		customerId *CustomerDetails
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test insert nil included customer element",
			fields: fields{
				make([]CustomerDetails, 0),
			},
			args: args{
				nil,
			},
			want: false,
		},
		{
			name: "Test insert sample included customer element",
			fields: fields{
				nil,
			},
			args: args{
				ToInviteData(&CustomerOffice{
					1,
					"James Barrett",
					"10.27654343",
					"-5.2653335",
				}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &InviteList{
				CustomerIds: tt.fields.CustomerIds,
			}
			if got := il.Add(tt.args.customerId); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCompleteInviteList(t *testing.T) {
	tests := []struct {
		name string
		want *CompleteInviteList
	}{
		{
			name: "Test Creation of a detailed output data",
			want: &CompleteInviteList{
				make([]CustomerDetails, 0),
				make([]CustomerDetails, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompleteInviteList(); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("NewCompleteInviteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInviteList(t *testing.T) {
	tests := []struct {
		name string
		want *InviteList
	}{
		{
			name: "Test Creation of a simple output data",
			want: &InviteList{
				make([]CustomerDetails, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInviteList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInviteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInviteData(t *testing.T) {
	type args struct {
		customerData *CustomerOffice
	}
	tests := []struct {
		name string
		args args
		want *CustomerDetails
	}{
		{
			name: "Test input to output data transformation, for given data",
			args: args{
				&CustomerOffice{
					UserId:    1,
					Name:      "Thomas Barrett",
					Latitude:  "10.58889684",
					Longitude: "2.345355",
				},
			},
			want: &CustomerDetails{
				UserId: 1,
				Name:   "Thomas Barrett",
			},
		},
		{
			name: "Test input to output data transformation, for nil input data",
			args: args{
				nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInviteData(tt.args.customerData); got != nil && tt.want != nil && !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("ToInviteData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerOffice_GetLatitude(t *testing.T) {
	type fields struct {
		UserId    int64
		Name      string
		Latitude  string
		Longitude string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "Test Correct Latitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.2345534535",
				Latitude:  "10.2345534535",
			},
			wantErr: false,
			want:    10.2345534535,
		},
		{
			name: "Test Wrong Latitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.2345534535",
				Latitude:  "10.22N",
			},
			wantErr: true,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerOffice{
				UserId:    tt.fields.UserId,
				Name:      tt.fields.Name,
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			got, err := c.GetLatitude()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatitude() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLatitude() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerOffice_GetLongitude(t *testing.T) {
	type fields struct {
		UserId    int64
		Name      string
		Latitude  string
		Longitude string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "Test Correct Latitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.2345534535",
				Latitude:  "10.2345534535",
			},
			wantErr: false,
			want:    10.2345534535,
		},
		{
			name: "Test Wrong Longitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.22W",
				Latitude:  "10.2345534535",
			},
			wantErr: true,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerOffice{
				UserId:    tt.fields.UserId,
				Name:      tt.fields.Name,
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			got, err := c.GetLongitude()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLongitude() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLongitude() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerOffice_IsValid(t *testing.T) {
	type fields struct {
		UserId    int64
		Name      string
		Latitude  string
		Longitude string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test Correct Latitude and Longitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.2345534535",
				Latitude:  "10.2345534535",
			},
			want: true,
		},
		{
			name: "Test Wrong Longitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.22W",
				Latitude:  "10.2345534535",
			},
			want: false,
		},
		{
			name: "Test Wrong Latitude",
			fields: fields{
				UserId:    1,
				Name:      "Thomas Barrett",
				Longitude: "10.2345534535",
				Latitude:  "10.22W",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerOffice{
				UserId:    tt.fields.UserId,
				Name:      tt.fields.Name,
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
