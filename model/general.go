/*
 * Copyright (c) 2020. This application code is under GNU Lesser General Public License, available here:
 * https://www.gnu.org/licenses/lgpl-3.0-standalone.html
 *
 * Any change or alterations are forbidden under the name of the author without any prior authorization, any abuse will be persecuted accordingly to the International Copyright Laws.
 * You can contact the author Fabrizio Torelli via email: hellgate75@gmail.com or using LinkedIn profile: https://www.linkedin.com/in/fabriziotorelli
 */

package model

import (
	"strconv"
	"sync"
)

// Describe input customer office information
type CustomerOffice struct {
	UserId    int64  `json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user-id,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Latitude  string `json:"latitude,omitempty" yaml:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty" xml:"longitude,omitempty"`
}

// Describe input customer office information
type CustomerOfficeList struct {
	List []CustomerOffice `json:"customers,omitempty" yaml:"customers,omitempty" xml:"customers,omitempty"`
}

func (c *CustomerOffice) GetLatitude() (float64, error) {
	return strconv.ParseFloat(c.Latitude, 64)
}

func (c *CustomerOffice) GetLongitude() (float64, error) {
	return strconv.ParseFloat(c.Longitude, 64)
}

func (c *CustomerOffice) IsValid() bool {
	_, err1 := c.GetLongitude()
	_, err2 := c.GetLatitude()
	return err1 == nil && err2 == nil
}

// Describe Output Customer details unit
type CustomerDetails struct {
	UserId int64  `json:"user_id" yaml:"user_id" xml:"user-id"`
	Name   string `json:"name" yaml:"name" xml:"name"`
}

// Describe standard output list
type InviteList struct {
	m           sync.Mutex
	CustomerIds []CustomerDetails `json:"customers_list" yaml:"customers_list" xml:"customers-list"`
}

// Add a new customer id to the invited customers list
func (il *InviteList) Add(customerId *CustomerDetails) bool {
	defer func() {
		_ = recover()
		il.m.Unlock()
	}()
	il.m.Lock()
	if il.CustomerIds == nil {
		il.CustomerIds = make([]CustomerDetails, 0)
	}
	if customerId == nil {
		return false
	}
	il.CustomerIds = append(il.CustomerIds, *customerId)
	return true
}

// Describe detailed output list
type CompleteInviteList struct {
	m1                    sync.Mutex
	m2                    sync.Mutex
	MatchingCustomerIds   []CustomerDetails `json:"customers_list" yaml:"customers_list" xml:"customers-list"`
	UnMatchingCustomerIds []CustomerDetails `json:"exclusions_list" yaml:"exclusions_list" xml:"exclusions-list"`
}

// Add a new customer id to the invited customers list
func (il *CompleteInviteList) AddInvited(customerId *CustomerDetails) bool {
	defer func() {
		_ = recover()
		il.m1.Unlock()
	}()
	il.m1.Lock()
	if il.MatchingCustomerIds == nil {
		il.MatchingCustomerIds = make([]CustomerDetails, 0)
	}
	if customerId == nil {
		return false
	}
	il.MatchingCustomerIds = append(il.MatchingCustomerIds, *customerId)
	return true
}

// Add a new customer id to the excluded customers list
func (il *CompleteInviteList) AddExcluded(customerId *CustomerDetails) bool {
	defer func() {
		_ = recover()
		il.m2.Unlock()
	}()
	il.m2.Lock()
	if il.UnMatchingCustomerIds == nil {
		il.UnMatchingCustomerIds = make([]CustomerDetails, 0)
	}
	if customerId == nil {
		return false
	}
	il.UnMatchingCustomerIds = append(il.UnMatchingCustomerIds, *customerId)
	return true
}

// Transform data from input to output data type
func ToInviteData(customerData *CustomerOffice) *CustomerDetails {
	if customerData == nil {
		return nil
	}
	return &CustomerDetails{
		UserId: customerData.UserId,
		Name:   customerData.Name,
	}
}

// Creates a simple output invitation list bucket pointer
func NewInviteList() *InviteList {
	return &InviteList{
		m:           sync.Mutex{},
		CustomerIds: make([]CustomerDetails, 0),
	}
}

// Creates a detailed output invitation list bucket pointer
func NewCompleteInviteList() *CompleteInviteList {
	return &CompleteInviteList{
		sync.Mutex{},
		sync.Mutex{},
		make([]CustomerDetails, 0),
		make([]CustomerDetails, 0),
	}
}
