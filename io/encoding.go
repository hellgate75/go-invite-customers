/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package io

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/hellgate75/go-invite-customers/model"
	"gopkg.in/yaml.v2"
	"strings"
)

type Encoding string

const (
	JsonEncoding    Encoding = "json"
	YamlEncoding    Encoding = "yaml"
	XmlEncoding     Encoding = "xml"
	TextEncoding    Encoding = "text"
	UnknownEncoding Encoding = "unknown"
)

var InputEncoding = []string{"json", "yaml", "xml"}
var OutputEncoding = []string{"text", "json", "yaml", "xml"}

//  Convert text to Encoding or return an unknown Encoding error.
//
//  In/
//  input text to be converted to Encoding type enumeration
//
//  The output are the encoding element and the error, if the encoding text is not known.
func ToEncoding(in string) (enc Encoding, err error) {
	switch strings.ToLower(in) {
	case "json":
		enc = JsonEncoding
		break
	case "yaml":
		enc = YamlEncoding
		break
	case "xml":
		enc = XmlEncoding
		break
	case "text":
		enc = TextEncoding
		break
	default:
		enc = UnknownEncoding
		err = errors.New(fmt.Sprintf("Unknown encoding text: %s", in))
	}
	return enc, err
}

//  Read the input bytes and decode in the wanted format the wanted model.CustomerOffice input
//  data type, or report the arisen error.
//
//  Data/
//  bytes that defines the input data to be un-marshalled from the given encoding format
//
//  Enc/
//  Encoding format, accordingly to the type io.Encoding
//
//  The output are the decoded object and the error, if occurred during the decoding operations.
func ReadCustomerOffice(data []byte, enc Encoding) (customer model.CustomerOffice, err error) {
	switch enc {
	case JsonEncoding:
		err = json.Unmarshal(data, &customer)
	case YamlEncoding:
		err = yaml.Unmarshal(data, &customer)
	case XmlEncoding:
		err = xml.Unmarshal(data, &customer)
	default:
		err = errors.New(fmt.Sprintf("Unknown encoding format %v", enc))
	}
	return customer, err
}

//  Read the input bytes and decode in the wanted format the wanted model.CustomerOffice input
//  data type, or report the arisen error.
//
//  Data/
//  bytes that defines the input data to be un-marshalled from the given encoding format
//
//  Enc/
//  Encoding format, accordingly to the type io.Encoding
//
//  The output are the decoded object and the error, if occurred during the decoding operations.
func ReadCustomerOfficeList(data []byte, enc Encoding) (customer model.CustomerOfficeList, err error) {
	switch enc {
	case JsonEncoding:
		err = json.Unmarshal(data, &customer)
	case YamlEncoding:
		err = yaml.Unmarshal(data, &customer)
	case XmlEncoding:
		err = xml.Unmarshal(data, &customer)
	default:
		err = errors.New(fmt.Sprintf("Unknown encoding format %v", enc))
	}
	return customer, err
}

//  Encode the model.InviteList output data type, reporting any error arisen during the encoding
//
//  Invite/
//  The model.InviteList data type instance to be converted in the given encoding format
//
//  Enc/
//  Encoding format, accordingly to the type io.Encoding
//
//  The output are the byte array and the error, if occurred during the encoding operations.
func EncodeCustomerInvite(invite model.InviteList, enc Encoding) (data []byte, err error) {
	data = make([]byte, 0)
	switch enc {
	case JsonEncoding:
		data, err = json.Marshal(&invite)
	case YamlEncoding:
		data, err = yaml.Marshal(&invite)
	case XmlEncoding:
		data, err = xml.Marshal(&invite)
	case TextEncoding:
		data, err = textEncodeInviteList(invite)
	default:
		err = errors.New(fmt.Sprintf("Unknown encoding format %v", enc))
	}
	return data, err
}

//  Encode the model.CompleteInviteList output data type, reporting any error arisen during the encoding
//
//  Invite/
//  The model.CompleteInviteList data type instance to be converted in the given encoding format
//
//  Enc/
//  Encoding format, accordingly to the type io.Encoding
//
//  The output are the byte array and the error, if occurred during the encoding operations.
func EncodeCustomerDetailedInvite(invite model.CompleteInviteList, enc Encoding) (data []byte, err error) {
	data = make([]byte, 0)
	switch enc {
	case JsonEncoding:
		data, err = json.Marshal(&invite)
	case YamlEncoding:
		data, err = yaml.Marshal(&invite)
	case XmlEncoding:
		data, err = xml.Marshal(&invite)
	case TextEncoding:
		data, err = textEncodeCompleteInviteList(invite)
	default:
		err = errors.New(fmt.Sprintf("Unknown encoding format %v", enc))
	}
	return data, err
}

func textEncodeInviteList(list model.InviteList) (out []byte, err error) {
	out = make([]byte, 0)
	text := ""
	for _, c := range list.CustomerIds {
		text += fmt.Sprintf("[%v] %s\n", c.UserId, c.Name)
	}
	if len(text) == 0 {
		text = "No customer selected"
	}
	text = "Invite Summary:\n" + text
	out = append(out, []byte(text)...)
	return out, err
}

func textEncodeCompleteInviteList(list model.CompleteInviteList) (out []byte, err error) {
	out = make([]byte, 0)
	text1 := ""
	for _, c := range list.MatchingCustomerIds {
		text1 += fmt.Sprintf("[%v] %s\n", c.UserId, c.Name)
	}
	if len(text1) == 0 {
		text1 = "No customer selected"
	}
	text1 = "Invite Summary:\n" + text1
	out = append(out, []byte(text1)...)
	text2 := ""
	for _, c := range list.MatchingCustomerIds {
		text2 += fmt.Sprintf("[%v] %s\n", c.UserId, c.Name)
	}
	if len(text2) == 0 {
		text2 = "No customer excluded"
	}
	text2 = "Exclusion Summary:\n" + text2
	out = append(out, []byte(text2)...)
	return out, err
}
