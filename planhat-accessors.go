// Copyright 2021 The go-planhat AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package planhat

import (
	"time"
)

// GetCompanyID returns the CompanyID field if it's non-nil, zero value otherwise.
func (a *Asset) GetCompanyID() string {
	if a == nil || a.CompanyID == nil {
		return ""
	}
	return *a.CompanyID
}

// GetExternalID returns the ExternalID field if it's non-nil, zero value otherwise.
func (a *Asset) GetExternalID() string {
	if a == nil || a.ExternalID == nil {
		return ""
	}
	return *a.ExternalID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (a *Asset) GetID() string {
	if a == nil || a.ID == nil {
		return ""
	}
	return *a.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (a *Asset) GetName() string {
	if a == nil || a.Name == nil {
		return ""
	}
	return *a.Name
}

// GetSourceID returns the SourceID field if it's non-nil, zero value otherwise.
func (a *Asset) GetSourceID() string {
	if a == nil || a.SourceID == nil {
		return ""
	}
	return *a.SourceID
}

// GetCompanyID returns the CompanyID field if it's non-nil, zero value otherwise.
func (a *AssetListOptions) GetCompanyID() string {
	if a == nil || a.CompanyID == nil {
		return ""
	}
	return *a.CompanyID
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (a *AssetListOptions) GetLimit() int {
	if a == nil || a.Limit == nil {
		return 0
	}
	return *a.Limit
}

// GetOffset returns the Offset field if it's non-nil, zero value otherwise.
func (a *AssetListOptions) GetOffset() int {
	if a == nil || a.Offset == nil {
		return 0
	}
	return *a.Offset
}

// GetSelect returns the Select field if it's non-nil, zero value otherwise.
func (a *AssetListOptions) GetSelect() string {
	if a == nil || a.Select == nil {
		return ""
	}
	return *a.Select
}

// GetSort returns the Sort field if it's non-nil, zero value otherwise.
func (a *AssetListOptions) GetSort() string {
	if a == nil || a.Sort == nil {
		return ""
	}
	return *a.Sort
}

// GetCSMScore returns the CSMScore field if it's non-nil, zero value otherwise.
func (c *Company) GetCSMScore() int {
	if c == nil || c.CSMScore == nil {
		return 0
	}
	return *c.CSMScore
}

// GetCustomerFrom returns the CustomerFrom field if it's non-nil, zero value otherwise.
func (c *Company) GetCustomerFrom() time.Time {
	if c == nil || c.CustomerFrom == nil {
		return time.Time{}
	}
	return *c.CustomerFrom
}

// GetCustomerTo returns the CustomerTo field if it's non-nil, zero value otherwise.
func (c *Company) GetCustomerTo() time.Time {
	if c == nil || c.CustomerTo == nil {
		return time.Time{}
	}
	return *c.CustomerTo
}

// GetExternalID returns the ExternalID field if it's non-nil, zero value otherwise.
func (c *Company) GetExternalID() string {
	if c == nil || c.ExternalID == nil {
		return ""
	}
	return *c.ExternalID
}

// GetH returns the H field if it's non-nil, zero value otherwise.
func (c *Company) GetH() int {
	if c == nil || c.H == nil {
		return 0
	}
	return *c.H
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (c *Company) GetID() string {
	if c == nil || c.ID == nil {
		return ""
	}
	return *c.ID
}

// GetLastRenewal returns the LastRenewal field if it's non-nil, zero value otherwise.
func (c *Company) GetLastRenewal() time.Time {
	if c == nil || c.LastRenewal == nil {
		return time.Time{}
	}
	return *c.LastRenewal
}

// GetLastTouch returns the LastTouch field if it's non-nil, zero value otherwise.
func (c *Company) GetLastTouch() string {
	if c == nil || c.LastTouch == nil {
		return ""
	}
	return *c.LastTouch
}

// GetLastTouchType returns the LastTouchType field if it's non-nil, zero value otherwise.
func (c *Company) GetLastTouchType() string {
	if c == nil || c.LastTouchType == nil {
		return ""
	}
	return *c.LastTouchType
}

// GetLicenses returns the Licenses field if it's non-nil, zero value otherwise.
func (c *Company) GetLicenses() []License {
	if c == nil || c.Licenses == nil {
		return nil
	}
	return *c.Licenses
}

// GetMR returns the MR field.
func (c *Company) GetMR() *float64 {
	if c == nil {
		return nil
	}
	return c.MR
}

// GetMRR returns the MRR field.
func (c *Company) GetMRR() *float64 {
	if c == nil {
		return nil
	}
	return c.MRR
}

// GetMRRTotal returns the MRRTotal field.
func (c *Company) GetMRRTotal() *float64 {
	if c == nil {
		return nil
	}
	return c.MRRTotal
}

// GetMRTotal returns the MRTotal field.
func (c *Company) GetMRTotal() *float64 {
	if c == nil {
		return nil
	}
	return c.MRTotal
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (c *Company) GetName() string {
	if c == nil || c.Name == nil {
		return ""
	}
	return *c.Name
}

// GetNRR30 returns the NRR30 field.
func (c *Company) GetNRR30() *float64 {
	if c == nil {
		return nil
	}
	return c.NRR30
}

// GetNRRTotal returns the NRRTotal field.
func (c *Company) GetNRRTotal() *float64 {
	if c == nil {
		return nil
	}
	return c.NRRTotal
}

// GetPhase returns the Phase field if it's non-nil, zero value otherwise.
func (c *Company) GetPhase() string {
	if c == nil || c.Phase == nil {
		return ""
	}
	return *c.Phase
}

// GetPhaseSince returns the PhaseSince field if it's non-nil, zero value otherwise.
func (c *Company) GetPhaseSince() time.Time {
	if c == nil || c.PhaseSince == nil {
		return time.Time{}
	}
	return *c.PhaseSince
}

// GetProducts returns the Products field if it's non-nil, zero value otherwise.
func (c *Company) GetProducts() []string {
	if c == nil || c.Products == nil {
		return nil
	}
	return *c.Products
}

// GetRenewalDate returns the RenewalDate field if it's non-nil, zero value otherwise.
func (c *Company) GetRenewalDate() time.Time {
	if c == nil || c.RenewalDate == nil {
		return time.Time{}
	}
	return *c.RenewalDate
}

// GetRenewalDaysFromNow returns the RenewalDaysFromNow field if it's non-nil, zero value otherwise.
func (c *Company) GetRenewalDaysFromNow() int {
	if c == nil || c.RenewalDaysFromNow == nil {
		return 0
	}
	return *c.RenewalDaysFromNow
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (c *Company) GetStatus() string {
	if c == nil || c.Status == nil {
		return ""
	}
	return *c.Status
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (c *CompanyListOptions) GetLimit() int {
	if c == nil || c.Limit == nil {
		return 0
	}
	return *c.Limit
}

// GetOffset returns the Offset field if it's non-nil, zero value otherwise.
func (c *CompanyListOptions) GetOffset() int {
	if c == nil || c.Offset == nil {
		return 0
	}
	return *c.Offset
}

// GetSort returns the Sort field if it's non-nil, zero value otherwise.
func (c *CompanyListOptions) GetSort() string {
	if c == nil || c.Sort == nil {
		return ""
	}
	return *c.Sort
}

// GetExternalID returns the ExternalID field if it's non-nil, zero value otherwise.
func (l *LeanCompanyListOptions) GetExternalID() string {
	if l == nil || l.ExternalID == nil {
		return ""
	}
	return *l.ExternalID
}

// GetSourceID returns the SourceID field if it's non-nil, zero value otherwise.
func (l *LeanCompanyListOptions) GetSourceID() string {
	if l == nil || l.SourceID == nil {
		return ""
	}
	return *l.SourceID
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (l *LeanCompanyListOptions) GetStatus() string {
	if l == nil || l.Status == nil {
		return ""
	}
	return *l.Status
}

// GetDate returns the Date field if it's non-nil, zero value otherwise.
func (m *Metric) GetDate() string {
	if m == nil || m.Date == nil {
		return ""
	}
	return *m.Date
}

// GetDimensionID returns the DimensionID field if it's non-nil, zero value otherwise.
func (m *Metric) GetDimensionID() string {
	if m == nil || m.DimensionID == nil {
		return ""
	}
	return *m.DimensionID
}

// GetExternalID returns the ExternalID field if it's non-nil, zero value otherwise.
func (m *Metric) GetExternalID() string {
	if m == nil || m.ExternalID == nil {
		return ""
	}
	return *m.ExternalID
}

// GetModel returns the Model field if it's non-nil, zero value otherwise.
func (m *Metric) GetModel() string {
	if m == nil || m.Model == nil {
		return ""
	}
	return *m.Model
}

// GetValue returns the Value field.
func (m *Metric) GetValue() *float64 {
	if m == nil {
		return nil
	}
	return m.Value
}

// GetCID returns the CID field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetCID() string {
	if m == nil || m.CID == nil {
		return ""
	}
	return *m.CID
}

// GetDimID returns the DimID field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetDimID() string {
	if m == nil || m.DimID == nil {
		return ""
	}
	return *m.DimID
}

// GetFrom returns the From field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetFrom() int {
	if m == nil || m.From == nil {
		return 0
	}
	return *m.From
}

// GetLimit returns the Limit field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetLimit() int {
	if m == nil || m.Limit == nil {
		return 0
	}
	return *m.Limit
}

// GetOffset returns the Offset field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetOffset() int {
	if m == nil || m.Offset == nil {
		return 0
	}
	return *m.Offset
}

// GetTo returns the To field if it's non-nil, zero value otherwise.
func (m *MetricsListOptions) GetTo() int {
	if m == nil || m.To == nil {
		return 0
	}
	return *m.To
}

// GetCompanyFilter returns the CompanyFilter field if it's non-nil, zero value otherwise.
func (u *User) GetCompanyFilter() string {
	if u == nil || u.CompanyFilter == nil {
		return ""
	}
	return *u.CompanyFilter
}

// GetCompressedView returns the CompressedView field if it's non-nil, zero value otherwise.
func (u *User) GetCompressedView() bool {
	if u == nil || u.CompressedView == nil {
		return false
	}
	return *u.CompressedView
}

// GetCreateDate returns the CreateDate field if it's non-nil, zero value otherwise.
func (u *User) GetCreateDate() time.Time {
	if u == nil || u.CreateDate == nil {
		return time.Time{}
	}
	return *u.CreateDate
}

// GetDailyDigest returns the DailyDigest field if it's non-nil, zero value otherwise.
func (u *User) GetDailyDigest() bool {
	if u == nil || u.DailyDigest == nil {
		return false
	}
	return *u.DailyDigest
}

// GetDefaultMeetingLength returns the DefaultMeetingLength field if it's non-nil, zero value otherwise.
func (u *User) GetDefaultMeetingLength() int {
	if u == nil || u.DefaultMeetingLength == nil {
		return 0
	}
	return *u.DefaultMeetingLength
}

// GetEmail returns the Email field if it's non-nil, zero value otherwise.
func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}
	return *u.Email
}

// GetFirstName returns the FirstName field if it's non-nil, zero value otherwise.
func (u *User) GetFirstName() string {
	if u == nil || u.FirstName == nil {
		return ""
	}
	return *u.FirstName
}

// GetFollowerUpdate returns the FollowerUpdate field if it's non-nil, zero value otherwise.
func (u *User) GetFollowerUpdate() bool {
	if u == nil || u.FollowerUpdate == nil {
		return false
	}
	return *u.FollowerUpdate
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (u *User) GetID() string {
	if u == nil || u.ID == nil {
		return ""
	}
	return *u.ID
}

// GetInactive returns the Inactive field if it's non-nil, zero value otherwise.
func (u *User) GetInactive() bool {
	if u == nil || u.Inactive == nil {
		return false
	}
	return *u.Inactive
}

// GetInAppNotifications returns the InAppNotifications field if it's non-nil, zero value otherwise.
func (u *User) GetInAppNotifications() bool {
	if u == nil || u.InAppNotifications == nil {
		return false
	}
	return *u.InAppNotifications
}

// GetIsExposedAsSenderOption returns the IsExposedAsSenderOption field if it's non-nil, zero value otherwise.
func (u *User) GetIsExposedAsSenderOption() bool {
	if u == nil || u.IsExposedAsSenderOption == nil {
		return false
	}
	return *u.IsExposedAsSenderOption
}

// GetIsHidden returns the IsHidden field if it's non-nil, zero value otherwise.
func (u *User) GetIsHidden() bool {
	if u == nil || u.IsHidden == nil {
		return false
	}
	return *u.IsHidden
}

// GetLastName returns the LastName field if it's non-nil, zero value otherwise.
func (u *User) GetLastName() string {
	if u == nil || u.LastName == nil {
		return ""
	}
	return *u.LastName
}

// GetNickName returns the NickName field if it's non-nil, zero value otherwise.
func (u *User) GetNickName() string {
	if u == nil || u.NickName == nil {
		return ""
	}
	return *u.NickName
}

// GetPlayLogDisabled returns the PlayLogDisabled field if it's non-nil, zero value otherwise.
func (u *User) GetPlayLogDisabled() bool {
	if u == nil || u.PlayLogDisabled == nil {
		return false
	}
	return *u.PlayLogDisabled
}

// GetRadarOneLine returns the RadarOneLine field if it's non-nil, zero value otherwise.
func (u *User) GetRadarOneLine() bool {
	if u == nil || u.RadarOneLine == nil {
		return false
	}
	return *u.RadarOneLine
}

// GetRecentOpenPage returns the RecentOpenPage field if it's non-nil, zero value otherwise.
func (u *User) GetRecentOpenPage() string {
	if u == nil || u.RecentOpenPage == nil {
		return ""
	}
	return *u.RecentOpenPage
}

// GetRemoved returns the Removed field if it's non-nil, zero value otherwise.
func (u *User) GetRemoved() bool {
	if u == nil || u.Removed == nil {
		return false
	}
	return *u.Removed
}

// GetRevReportPeriodType returns the RevReportPeriodType field if it's non-nil, zero value otherwise.
func (u *User) GetRevReportPeriodType() string {
	if u == nil || u.RevReportPeriodType == nil {
		return ""
	}
	return *u.RevReportPeriodType
}

// GetSegment returns the Segment field if it's non-nil, zero value otherwise.
func (u *User) GetSegment() string {
	if u == nil || u.Segment == nil {
		return ""
	}
	return *u.Segment
}

// GetSplitLayoutDisabled returns the SplitLayoutDisabled field if it's non-nil, zero value otherwise.
func (u *User) GetSplitLayoutDisabled() bool {
	if u == nil || u.SplitLayoutDisabled == nil {
		return false
	}
	return *u.SplitLayoutDisabled
}

// GetTaskFilter returns the TaskFilter field if it's non-nil, zero value otherwise.
func (u *User) GetTaskFilter() string {
	if u == nil || u.TaskFilter == nil {
		return ""
	}
	return *u.TaskFilter
}

// GetV returns the V field if it's non-nil, zero value otherwise.
func (u *User) GetV() int {
	if u == nil || u.V == nil {
		return 0
	}
	return *u.V
}

// GetWorkflowFilter returns the WorkflowFilter field if it's non-nil, zero value otherwise.
func (u *User) GetWorkflowFilter() string {
	if u == nil || u.WorkflowFilter == nil {
		return ""
	}
	return *u.WorkflowFilter
}
