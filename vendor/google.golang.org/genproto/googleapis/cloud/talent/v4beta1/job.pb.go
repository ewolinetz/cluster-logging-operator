// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/job.proto

package talent

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/type/postaladdress"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A Job resource represents a job posting (also referred to as a "job listing"
// or "job requisition"). A job belongs to a
// [Company][google.cloud.talent.v4beta1.Company], which is the hiring entity
// responsible for the job.
type Job struct {
	// Required during job update.
	//
	// The resource name for the job. This is generated by the service when a
	// job is created.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}", for
	// example, "projects/api-test-project/tenants/foo/jobs/1234".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project/jobs/1234".
	//
	// Use of this field in job queries and API calls is preferred over the use of
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] since this
	// value is unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The resource name of the company listing the job.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project/companies/bar".
	Company string `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	// Required. The requisition ID, also referred to as the posting ID, is
	// assigned by the client to identify a job. This field is intended to be used
	// by clients for client identification and tracking of postings. A job isn't
	// allowed to be created if there is another job with the same
	// [company][google.cloud.talent.v4beta1.Job.name],
	// [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id].
	//
	// The maximum number of allowed characters is 255.
	RequisitionId string `protobuf:"bytes,3,opt,name=requisition_id,json=requisitionId,proto3" json:"requisition_id,omitempty"`
	// Required. The title of the job, such as "Software Engineer"
	//
	// The maximum number of allowed characters is 500.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Required. The description of the job, which typically includes a
	// multi-paragraph description of the company and related information.
	// Separate fields are provided on the job object for
	// [responsibilities][google.cloud.talent.v4beta1.Job.responsibilities],
	// [qualifications][google.cloud.talent.v4beta1.Job.qualifications], and other
	// job characteristics. Use of these separate job fields is recommended.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 100,000.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Optional but strongly recommended for the best service experience.
	//
	// Location(s) where the employer is looking to hire for this job posting.
	//
	// Specifying the full street address(es) of the hiring location enables
	// better API results, especially job searches by commute time.
	//
	// At most 50 locations are allowed for best search performance. If a job has
	// more locations, it is suggested to split it into multiple jobs with unique
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id]s (e.g.
	// 'ReqA' becomes 'ReqA-1', 'ReqA-2', and so on.) as multiple jobs with the
	// same [company][google.cloud.talent.v4beta1.Job.company],
	// [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] are not
	// allowed. If the original
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id] must be
	// preserved, a custom field should be used for storage. It is also suggested
	// to group the locations that close to each other in the same job for better
	// search experience.
	//
	// The maximum number of allowed characters is 500.
	Addresses []string `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// Optional. Job application information.
	ApplicationInfo *Job_ApplicationInfo `protobuf:"bytes,7,opt,name=application_info,json=applicationInfo,proto3" json:"application_info,omitempty"`
	// Optional. The benefits included with the job.
	JobBenefits []JobBenefit `protobuf:"varint,8,rep,packed,name=job_benefits,json=jobBenefits,proto3,enum=google.cloud.talent.v4beta1.JobBenefit" json:"job_benefits,omitempty"`
	// Optional. Job compensation information (a.k.a. "pay rate") i.e., the
	// compensation that will paid to the employee.
	CompensationInfo *CompensationInfo `protobuf:"bytes,9,opt,name=compensation_info,json=compensationInfo,proto3" json:"compensation_info,omitempty"`
	// Optional. A map of fields to hold both filterable and non-filterable custom
	// job attributes that are not covered by the provided structured fields.
	//
	// The keys of the map are strings up to 64 bytes and must match the
	// pattern: [a-zA-Z][a-zA-Z0-9_]*. For example, key0LikeThis or
	// KEY_1_LIKE_THIS.
	//
	// At most 100 filterable and at most 100 unfilterable keys are supported.
	// For filterable `string_values`, across all keys at most 200 values are
	// allowed, with each string no more than 255 characters. For unfilterable
	// `string_values`, the maximum total size of `string_values` across all keys
	// is 50KB.
	CustomAttributes map[string]*CustomAttribute `protobuf:"bytes,10,rep,name=custom_attributes,json=customAttributes,proto3" json:"custom_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. The desired education degrees for the job, such as Bachelors,
	// Masters.
	DegreeTypes []DegreeType `protobuf:"varint,11,rep,packed,name=degree_types,json=degreeTypes,proto3,enum=google.cloud.talent.v4beta1.DegreeType" json:"degree_types,omitempty"`
	// Optional. The department or functional area within the company with the
	// open position.
	//
	// The maximum number of allowed characters is 255.
	Department string `protobuf:"bytes,12,opt,name=department,proto3" json:"department,omitempty"`
	// Optional. The employment type(s) of a job, for example,
	// [full time][google.cloud.talent.v4beta1.EmploymentType.FULL_TIME] or
	// [part time][google.cloud.talent.v4beta1.EmploymentType.PART_TIME].
	EmploymentTypes []EmploymentType `protobuf:"varint,13,rep,packed,name=employment_types,json=employmentTypes,proto3,enum=google.cloud.talent.v4beta1.EmploymentType" json:"employment_types,omitempty"`
	// Optional. A description of bonus, commission, and other compensation
	// incentives associated with the job not including salary or pay.
	//
	// The maximum number of allowed characters is 10,000.
	Incentives string `protobuf:"bytes,14,opt,name=incentives,proto3" json:"incentives,omitempty"`
	// Optional. The language of the posting. This field is distinct from
	// any requirements for fluency that are associated with the job.
	//
	// Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){:
	// class="external" target="_blank" }.
	//
	// If this field is unspecified and
	// [Job.description][google.cloud.talent.v4beta1.Job.description] is present,
	// detected language code based on
	// [Job.description][google.cloud.talent.v4beta1.Job.description] is assigned,
	// otherwise defaults to 'en_US'.
	LanguageCode string `protobuf:"bytes,15,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. The experience level associated with the job, such as "Entry
	// Level".
	JobLevel JobLevel `protobuf:"varint,16,opt,name=job_level,json=jobLevel,proto3,enum=google.cloud.talent.v4beta1.JobLevel" json:"job_level,omitempty"`
	// Optional. A promotion value of the job, as determined by the client.
	// The value determines the sort order of the jobs returned when searching for
	// jobs using the featured jobs search call, with higher promotional values
	// being returned first and ties being resolved by relevance sort. Only the
	// jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH.
	//
	// Default value is 0, and negative values are treated as 0.
	PromotionValue int32 `protobuf:"varint,17,opt,name=promotion_value,json=promotionValue,proto3" json:"promotion_value,omitempty"`
	// Optional. A description of the qualifications required to perform the
	// job. The use of this field is recommended
	// as an alternative to using the more general
	// [description][google.cloud.talent.v4beta1.Job.description] field.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 10,000.
	Qualifications string `protobuf:"bytes,18,opt,name=qualifications,proto3" json:"qualifications,omitempty"`
	// Optional. A description of job responsibilities. The use of this field is
	// recommended as an alternative to using the more general
	// [description][google.cloud.talent.v4beta1.Job.description] field.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 10,000.
	Responsibilities string `protobuf:"bytes,19,opt,name=responsibilities,proto3" json:"responsibilities,omitempty"`
	// Optional. The job
	// [PostingRegion][google.cloud.talent.v4beta1.PostingRegion] (for example,
	// state, country) throughout which the job is available. If this field is
	// set, a [LocationFilter][google.cloud.talent.v4beta1.LocationFilter] in a
	// search query within the job region finds this job posting if an exact
	// location match isn't specified. If this field is set to
	// [PostingRegion.NATION][google.cloud.talent.v4beta1.PostingRegion.NATION] or
	// [PostingRegion.ADMINISTRATIVE_AREA][google.cloud.talent.v4beta1.PostingRegion.ADMINISTRATIVE_AREA],
	// setting job [Job.addresses][google.cloud.talent.v4beta1.Job.addresses] to
	// the same location level as this field is strongly recommended.
	PostingRegion PostingRegion `protobuf:"varint,20,opt,name=posting_region,json=postingRegion,proto3,enum=google.cloud.talent.v4beta1.PostingRegion" json:"posting_region,omitempty"`
	// Optional. The visibility of the job.
	//
	// Defaults to
	// [Visibility.ACCOUNT_ONLY][google.cloud.talent.v4beta1.Visibility.ACCOUNT_ONLY]
	// if not specified.
	Visibility Visibility `protobuf:"varint,21,opt,name=visibility,proto3,enum=google.cloud.talent.v4beta1.Visibility" json:"visibility,omitempty"`
	// Optional. The start timestamp of the job in UTC time zone. Typically this
	// field is used for contracting engagements. Invalid timestamps are ignored.
	JobStartTime *timestamp.Timestamp `protobuf:"bytes,22,opt,name=job_start_time,json=jobStartTime,proto3" json:"job_start_time,omitempty"`
	// Optional. The end timestamp of the job. Typically this field is used for
	// contracting engagements. Invalid timestamps are ignored.
	JobEndTime *timestamp.Timestamp `protobuf:"bytes,23,opt,name=job_end_time,json=jobEndTime,proto3" json:"job_end_time,omitempty"`
	// Optional. The timestamp this job posting was most recently published. The
	// default value is the time the request arrives at the server. Invalid
	// timestamps are ignored.
	PostingPublishTime *timestamp.Timestamp `protobuf:"bytes,24,opt,name=posting_publish_time,json=postingPublishTime,proto3" json:"posting_publish_time,omitempty"`
	// Optional but strongly recommended for the best service
	// experience.
	//
	// The expiration timestamp of the job. After this timestamp, the
	// job is marked as expired, and it no longer appears in search results. The
	// expired job can't be listed by the
	// [ListJobs][google.cloud.talent.v4beta1.JobService.ListJobs] API, but it can
	// be retrieved with the
	// [GetJob][google.cloud.talent.v4beta1.JobService.GetJob] API or updated with
	// the [UpdateJob][google.cloud.talent.v4beta1.JobService.UpdateJob] API or
	// deleted with the
	// [DeleteJob][google.cloud.talent.v4beta1.JobService.DeleteJob] API. An
	// expired job can be updated and opened again by using a future expiration
	// timestamp. Updating an expired job fails if there is another existing open
	// job with same [company][google.cloud.talent.v4beta1.Job.company],
	// [language_code][google.cloud.talent.v4beta1.Job.language_code] and
	// [requisition_id][google.cloud.talent.v4beta1.Job.requisition_id].
	//
	// The expired jobs are retained in our system for 90 days. However, the
	// overall expired job count cannot exceed 3 times the maximum number of
	// open jobs over previous 7 days. If this threshold is exceeded,
	// expired jobs are cleaned out in order of earliest expire time.
	// Expired jobs are no longer accessible after they are cleaned
	// out.
	//
	// Invalid timestamps are ignored, and treated as expire time not provided.
	//
	// If the timestamp is before the instant request is made, the job
	// is treated as expired immediately on creation. This kind of job can
	// not be updated. And when creating a job with past timestamp, the
	// [posting_publish_time][google.cloud.talent.v4beta1.Job.posting_publish_time]
	// must be set before
	// [posting_expire_time][google.cloud.talent.v4beta1.Job.posting_expire_time].
	// The purpose of this feature is to allow other objects, such as
	// [Application][google.cloud.talent.v4beta1.Application], to refer a job that
	// didn't exist in the system prior to becoming expired. If you want to modify
	// a job that was expired on creation, delete it and create a new one.
	//
	// If this value isn't provided at the time of job creation or is invalid,
	// the job posting expires after 30 days from the job's creation time. For
	// example, if the job was created on 2017/01/01 13:00AM UTC with an
	// unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC.
	//
	// If this value isn't provided on job update, it depends on the field masks
	// set by
	// [UpdateJobRequest.update_mask][google.cloud.talent.v4beta1.UpdateJobRequest.update_mask].
	// If the field masks include
	// [job_end_time][google.cloud.talent.v4beta1.Job.job_end_time], or the masks
	// are empty meaning that every field is updated, the job posting expires
	// after 30 days from the job's last update time. Otherwise the expiration
	// date isn't updated.
	PostingExpireTime *timestamp.Timestamp `protobuf:"bytes,25,opt,name=posting_expire_time,json=postingExpireTime,proto3" json:"posting_expire_time,omitempty"`
	// Output only. The timestamp when this job posting was created.
	PostingCreateTime *timestamp.Timestamp `protobuf:"bytes,26,opt,name=posting_create_time,json=postingCreateTime,proto3" json:"posting_create_time,omitempty"`
	// Output only. The timestamp when this job posting was last updated.
	PostingUpdateTime *timestamp.Timestamp `protobuf:"bytes,27,opt,name=posting_update_time,json=postingUpdateTime,proto3" json:"posting_update_time,omitempty"`
	// Output only. Display name of the company listing the job.
	CompanyDisplayName string `protobuf:"bytes,28,opt,name=company_display_name,json=companyDisplayName,proto3" json:"company_display_name,omitempty"`
	// Output only. Derived details about the job posting.
	DerivedInfo *Job_DerivedInfo `protobuf:"bytes,29,opt,name=derived_info,json=derivedInfo,proto3" json:"derived_info,omitempty"`
	// Optional. Options for job processing.
	ProcessingOptions    *Job_ProcessingOptions `protobuf:"bytes,30,opt,name=processing_options,json=processingOptions,proto3" json:"processing_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_c717d27462774c27, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Job) GetRequisitionId() string {
	if m != nil {
		return m.RequisitionId
	}
	return ""
}

func (m *Job) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Job) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Job) GetApplicationInfo() *Job_ApplicationInfo {
	if m != nil {
		return m.ApplicationInfo
	}
	return nil
}

func (m *Job) GetJobBenefits() []JobBenefit {
	if m != nil {
		return m.JobBenefits
	}
	return nil
}

func (m *Job) GetCompensationInfo() *CompensationInfo {
	if m != nil {
		return m.CompensationInfo
	}
	return nil
}

func (m *Job) GetCustomAttributes() map[string]*CustomAttribute {
	if m != nil {
		return m.CustomAttributes
	}
	return nil
}

func (m *Job) GetDegreeTypes() []DegreeType {
	if m != nil {
		return m.DegreeTypes
	}
	return nil
}

func (m *Job) GetDepartment() string {
	if m != nil {
		return m.Department
	}
	return ""
}

func (m *Job) GetEmploymentTypes() []EmploymentType {
	if m != nil {
		return m.EmploymentTypes
	}
	return nil
}

func (m *Job) GetIncentives() string {
	if m != nil {
		return m.Incentives
	}
	return ""
}

func (m *Job) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *Job) GetJobLevel() JobLevel {
	if m != nil {
		return m.JobLevel
	}
	return JobLevel_JOB_LEVEL_UNSPECIFIED
}

func (m *Job) GetPromotionValue() int32 {
	if m != nil {
		return m.PromotionValue
	}
	return 0
}

func (m *Job) GetQualifications() string {
	if m != nil {
		return m.Qualifications
	}
	return ""
}

func (m *Job) GetResponsibilities() string {
	if m != nil {
		return m.Responsibilities
	}
	return ""
}

func (m *Job) GetPostingRegion() PostingRegion {
	if m != nil {
		return m.PostingRegion
	}
	return PostingRegion_POSTING_REGION_UNSPECIFIED
}

func (m *Job) GetVisibility() Visibility {
	if m != nil {
		return m.Visibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

func (m *Job) GetJobStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.JobStartTime
	}
	return nil
}

func (m *Job) GetJobEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.JobEndTime
	}
	return nil
}

func (m *Job) GetPostingPublishTime() *timestamp.Timestamp {
	if m != nil {
		return m.PostingPublishTime
	}
	return nil
}

func (m *Job) GetPostingExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.PostingExpireTime
	}
	return nil
}

func (m *Job) GetPostingCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.PostingCreateTime
	}
	return nil
}

func (m *Job) GetPostingUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.PostingUpdateTime
	}
	return nil
}

func (m *Job) GetCompanyDisplayName() string {
	if m != nil {
		return m.CompanyDisplayName
	}
	return ""
}

func (m *Job) GetDerivedInfo() *Job_DerivedInfo {
	if m != nil {
		return m.DerivedInfo
	}
	return nil
}

func (m *Job) GetProcessingOptions() *Job_ProcessingOptions {
	if m != nil {
		return m.ProcessingOptions
	}
	return nil
}

// Application related details of a job posting.
type Job_ApplicationInfo struct {
	// Optional. Use this field to specify email address(es) to which resumes or
	// applications can be sent.
	//
	// The maximum number of allowed characters for each entry is 255.
	Emails []string `protobuf:"bytes,1,rep,name=emails,proto3" json:"emails,omitempty"`
	// Optional. Use this field to provide instructions, such as "Mail your
	// application to ...", that a candidate can follow to apply for the job.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 3,000.
	Instruction string `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// Optional. Use this URI field to direct an applicant to a website, for
	// example to link to an online application form.
	//
	// The maximum number of allowed characters for each entry is 2,000.
	Uris                 []string `protobuf:"bytes,3,rep,name=uris,proto3" json:"uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job_ApplicationInfo) Reset()         { *m = Job_ApplicationInfo{} }
func (m *Job_ApplicationInfo) String() string { return proto.CompactTextString(m) }
func (*Job_ApplicationInfo) ProtoMessage()    {}
func (*Job_ApplicationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c717d27462774c27, []int{0, 0}
}

func (m *Job_ApplicationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job_ApplicationInfo.Unmarshal(m, b)
}
func (m *Job_ApplicationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job_ApplicationInfo.Marshal(b, m, deterministic)
}
func (m *Job_ApplicationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job_ApplicationInfo.Merge(m, src)
}
func (m *Job_ApplicationInfo) XXX_Size() int {
	return xxx_messageInfo_Job_ApplicationInfo.Size(m)
}
func (m *Job_ApplicationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Job_ApplicationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Job_ApplicationInfo proto.InternalMessageInfo

func (m *Job_ApplicationInfo) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *Job_ApplicationInfo) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *Job_ApplicationInfo) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

// Output only.
//
// Derived details about the job posting.
type Job_DerivedInfo struct {
	// Structured locations of the job, resolved from
	// [Job.addresses][google.cloud.talent.v4beta1.Job.addresses].
	//
	// [locations][google.cloud.talent.v4beta1.Job.DerivedInfo.locations] are
	// exactly matched to
	// [Job.addresses][google.cloud.talent.v4beta1.Job.addresses] in the same
	// order.
	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	// Job categories derived from
	// [Job.title][google.cloud.talent.v4beta1.Job.title] and
	// [Job.description][google.cloud.talent.v4beta1.Job.description].
	JobCategories        []JobCategory `protobuf:"varint,3,rep,packed,name=job_categories,json=jobCategories,proto3,enum=google.cloud.talent.v4beta1.JobCategory" json:"job_categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Job_DerivedInfo) Reset()         { *m = Job_DerivedInfo{} }
func (m *Job_DerivedInfo) String() string { return proto.CompactTextString(m) }
func (*Job_DerivedInfo) ProtoMessage()    {}
func (*Job_DerivedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c717d27462774c27, []int{0, 1}
}

func (m *Job_DerivedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job_DerivedInfo.Unmarshal(m, b)
}
func (m *Job_DerivedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job_DerivedInfo.Marshal(b, m, deterministic)
}
func (m *Job_DerivedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job_DerivedInfo.Merge(m, src)
}
func (m *Job_DerivedInfo) XXX_Size() int {
	return xxx_messageInfo_Job_DerivedInfo.Size(m)
}
func (m *Job_DerivedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Job_DerivedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Job_DerivedInfo proto.InternalMessageInfo

func (m *Job_DerivedInfo) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Job_DerivedInfo) GetJobCategories() []JobCategory {
	if m != nil {
		return m.JobCategories
	}
	return nil
}

// Input only.
//
// Options for job processing.
type Job_ProcessingOptions struct {
	// Optional. If set to `true`, the service does not attempt to resolve a
	// more precise address for the job.
	DisableStreetAddressResolution bool `protobuf:"varint,1,opt,name=disable_street_address_resolution,json=disableStreetAddressResolution,proto3" json:"disable_street_address_resolution,omitempty"`
	// Optional. Option for job HTML content sanitization. Applied fields are:
	//
	// * description
	// * applicationInfo.instruction
	// * incentives
	// * qualifications
	// * responsibilities
	//
	// HTML tags in these fields may be stripped if sanitiazation isn't
	// disabled.
	//
	// Defaults to
	// [HtmlSanitization.SIMPLE_FORMATTING_ONLY][google.cloud.talent.v4beta1.HtmlSanitization.SIMPLE_FORMATTING_ONLY].
	HtmlSanitization     HtmlSanitization `protobuf:"varint,2,opt,name=html_sanitization,json=htmlSanitization,proto3,enum=google.cloud.talent.v4beta1.HtmlSanitization" json:"html_sanitization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Job_ProcessingOptions) Reset()         { *m = Job_ProcessingOptions{} }
func (m *Job_ProcessingOptions) String() string { return proto.CompactTextString(m) }
func (*Job_ProcessingOptions) ProtoMessage()    {}
func (*Job_ProcessingOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c717d27462774c27, []int{0, 2}
}

func (m *Job_ProcessingOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job_ProcessingOptions.Unmarshal(m, b)
}
func (m *Job_ProcessingOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job_ProcessingOptions.Marshal(b, m, deterministic)
}
func (m *Job_ProcessingOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job_ProcessingOptions.Merge(m, src)
}
func (m *Job_ProcessingOptions) XXX_Size() int {
	return xxx_messageInfo_Job_ProcessingOptions.Size(m)
}
func (m *Job_ProcessingOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Job_ProcessingOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Job_ProcessingOptions proto.InternalMessageInfo

func (m *Job_ProcessingOptions) GetDisableStreetAddressResolution() bool {
	if m != nil {
		return m.DisableStreetAddressResolution
	}
	return false
}

func (m *Job_ProcessingOptions) GetHtmlSanitization() HtmlSanitization {
	if m != nil {
		return m.HtmlSanitization
	}
	return HtmlSanitization_HTML_SANITIZATION_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Job)(nil), "google.cloud.talent.v4beta1.Job")
	proto.RegisterMapType((map[string]*CustomAttribute)(nil), "google.cloud.talent.v4beta1.Job.CustomAttributesEntry")
	proto.RegisterType((*Job_ApplicationInfo)(nil), "google.cloud.talent.v4beta1.Job.ApplicationInfo")
	proto.RegisterType((*Job_DerivedInfo)(nil), "google.cloud.talent.v4beta1.Job.DerivedInfo")
	proto.RegisterType((*Job_ProcessingOptions)(nil), "google.cloud.talent.v4beta1.Job.ProcessingOptions")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/job.proto", fileDescriptor_c717d27462774c27)
}

var fileDescriptor_c717d27462774c27 = []byte{
	// 1102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xd1, 0x72, 0x1b, 0x35,
	0x17, 0x1e, 0xc7, 0x4d, 0x1b, 0xcb, 0x89, 0x63, 0xab, 0x69, 0x7f, 0xfd, 0x6e, 0x48, 0x0d, 0x4c,
	0xa8, 0xa7, 0x80, 0x5d, 0x02, 0xc3, 0x30, 0xc0, 0x05, 0x89, 0x93, 0x81, 0x64, 0x32, 0x24, 0x6c,
	0x42, 0x2e, 0xca, 0xc5, 0x8e, 0x76, 0xf7, 0xc4, 0x91, 0xd9, 0x95, 0x14, 0x49, 0x6b, 0x30, 0x17,
	0x3c, 0x0c, 0xbc, 0x05, 0x2f, 0xc1, 0x2b, 0x31, 0x92, 0x76, 0x1d, 0xd7, 0xed, 0xd8, 0xb9, 0xd3,
	0x39, 0xfa, 0xbe, 0xef, 0x48, 0x67, 0x75, 0xce, 0x59, 0xb4, 0x3b, 0x14, 0x62, 0x98, 0x42, 0x3f,
	0x4e, 0x45, 0x9e, 0xf4, 0x0d, 0x4d, 0x81, 0x9b, 0xfe, 0xf8, 0x8b, 0x08, 0x0c, 0xfd, 0xac, 0x3f,
	0x12, 0x51, 0x4f, 0x2a, 0x61, 0x04, 0x7e, 0xe6, 0x61, 0x3d, 0x07, 0xeb, 0x79, 0x58, 0xaf, 0x80,
	0xb5, 0xb7, 0x0b, 0x0d, 0x2a, 0x59, 0x9f, 0x72, 0x2e, 0x0c, 0x35, 0x4c, 0x70, 0xed, 0xa9, 0xed,
	0xee, 0xa2, 0x08, 0xb1, 0xc8, 0x32, 0xc1, 0x0b, 0xe4, 0xf3, 0x02, 0xe9, 0xac, 0x28, 0xbf, 0xee,
	0x1b, 0x96, 0x81, 0x36, 0x34, 0x93, 0x05, 0x60, 0x67, 0x1e, 0xf0, 0x9b, 0xa2, 0x52, 0x82, 0x2a,
	0x43, 0x75, 0x8a, 0x7d, 0x33, 0x91, 0xd0, 0x97, 0x42, 0x1b, 0x9a, 0x86, 0x34, 0x49, 0x14, 0xe8,
	0x02, 0xf1, 0xc1, 0xbf, 0x18, 0x55, 0x4f, 0x44, 0x84, 0x31, 0x7a, 0xc0, 0x69, 0x06, 0xa4, 0xd2,
	0xa9, 0x74, 0x6b, 0x81, 0x5b, 0x63, 0x82, 0x1e, 0xc5, 0x22, 0x93, 0x94, 0x4f, 0xc8, 0x8a, 0x73,
	0x97, 0x26, 0xde, 0x45, 0x0d, 0x05, 0xb7, 0x39, 0xd3, 0xcc, 0x5e, 0x2c, 0x64, 0x09, 0xa9, 0x3a,
	0xc0, 0xc6, 0x8c, 0xf7, 0x38, 0xc1, 0x5b, 0x68, 0xd5, 0x30, 0x93, 0x02, 0x79, 0xe0, 0x76, 0xbd,
	0x81, 0x3b, 0xa8, 0x9e, 0x80, 0x8e, 0x15, 0x93, 0x16, 0x46, 0x56, 0xdd, 0xde, 0xac, 0x0b, 0x6f,
	0xa3, 0x5a, 0x71, 0x4a, 0xd0, 0xe4, 0x61, 0xa7, 0xda, 0xad, 0x05, 0x77, 0x0e, 0xfc, 0x0b, 0x6a,
	0x52, 0x29, 0x53, 0x16, 0x53, 0x1f, 0x9c, 0x5f, 0x0b, 0xf2, 0xa8, 0x53, 0xe9, 0xd6, 0xf7, 0x5e,
	0xf5, 0x16, 0x7c, 0x95, 0xde, 0x89, 0x88, 0x7a, 0xfb, 0x77, 0xc4, 0x63, 0x7e, 0x2d, 0x82, 0x4d,
	0xfa, 0xa6, 0x03, 0x9f, 0xa0, 0xf5, 0x91, 0x88, 0xc2, 0x08, 0x38, 0x5c, 0x33, 0xa3, 0xc9, 0x5a,
	0xa7, 0xda, 0x6d, 0xec, 0xbd, 0x58, 0x26, 0x7c, 0xe0, 0xf1, 0x41, 0x7d, 0x34, 0x5d, 0x6b, 0xfc,
	0x1a, 0xb5, 0x6c, 0xc2, 0x80, 0xeb, 0x99, 0x93, 0xd6, 0xdc, 0x49, 0x3f, 0x5d, 0x28, 0x38, 0x98,
	0x61, 0xb9, 0x63, 0x36, 0xe3, 0x39, 0x0f, 0x8e, 0x51, 0x2b, 0xce, 0xb5, 0x11, 0x59, 0x48, 0x8d,
	0x51, 0x2c, 0xca, 0x0d, 0x68, 0x82, 0x3a, 0xd5, 0x6e, 0x7d, 0xef, 0xcb, 0xa5, 0x59, 0x18, 0x38,
	0xe6, 0xfe, 0x94, 0x78, 0xc4, 0x8d, 0x9a, 0x04, 0xcd, 0x78, 0xce, 0x6d, 0x93, 0x91, 0xc0, 0x50,
	0x01, 0x84, 0xf6, 0x01, 0x69, 0x52, 0xbf, 0x47, 0x32, 0x0e, 0x1d, 0xe1, 0x72, 0x22, 0xc1, 0x7e,
	0xd3, 0x72, 0xad, 0xf1, 0x0e, 0x42, 0x09, 0x48, 0xaa, 0x4c, 0x06, 0xdc, 0x90, 0x75, 0xf7, 0xd1,
	0x67, 0x3c, 0xf8, 0x0a, 0x35, 0x21, 0x93, 0xa9, 0x98, 0x58, 0xab, 0x88, 0xb7, 0xe1, 0xe2, 0x7d,
	0xbc, 0x30, 0xde, 0xd1, 0x94, 0xe4, 0x62, 0x6e, 0xc2, 0x1b, 0xb6, 0x8b, 0xcb, 0x78, 0x0c, 0xdc,
	0xb0, 0x31, 0x68, 0xd2, 0xf0, 0x71, 0xef, 0x3c, 0xf8, 0x43, 0xb4, 0x91, 0x52, 0x3e, 0xcc, 0xe9,
	0x10, 0xc2, 0x58, 0x24, 0x40, 0x36, 0x1d, 0x64, 0xbd, 0x74, 0x0e, 0x44, 0x02, 0xf8, 0x00, 0xd5,
	0xec, 0xab, 0x48, 0x61, 0x0c, 0x29, 0x69, 0x76, 0x2a, 0xdd, 0xc6, 0xde, 0xee, 0xb2, 0x2c, 0x9f,
	0x5a, 0x70, 0xb0, 0x36, 0x2a, 0x56, 0xf8, 0x05, 0xda, 0x94, 0x4a, 0x64, 0xc2, 0x3d, 0x85, 0x31,
	0x4d, 0x73, 0x20, 0xad, 0x4e, 0xa5, 0xbb, 0x1a, 0x34, 0xa6, 0xee, 0x2b, 0xeb, 0xc5, 0x1f, 0xa1,
	0xc6, 0x6d, 0x4e, 0x53, 0x76, 0x5d, 0xbc, 0x4b, 0x4d, 0xb0, 0x3b, 0xd2, 0x9c, 0x17, 0xbf, 0x44,
	0x4d, 0x05, 0x5a, 0x0a, 0xae, 0x59, 0xc4, 0x52, 0x66, 0x18, 0x68, 0xf2, 0xd8, 0x21, 0xdf, 0xf2,
	0xe3, 0x9f, 0x50, 0xc3, 0x96, 0x3f, 0xe3, 0xc3, 0x50, 0xc1, 0xd0, 0x96, 0xdd, 0x96, 0xbb, 0xc5,
	0xcb, 0x85, 0xb7, 0x38, 0xf7, 0x94, 0xc0, 0x31, 0x82, 0x0d, 0x39, 0x6b, 0xe2, 0xef, 0x11, 0x1a,
	0xb3, 0x22, 0xc4, 0x84, 0x3c, 0x71, 0x72, 0x8b, 0x9f, 0xc6, 0xd5, 0x14, 0x1e, 0xcc, 0x50, 0xf1,
	0x77, 0xa8, 0x61, 0x93, 0xab, 0x0d, 0x55, 0x26, 0xb4, 0x1d, 0x8e, 0x3c, 0x75, 0x35, 0xd2, 0x2e,
	0xc5, 0xca, 0xee, 0xd6, 0xbb, 0x2c, 0xdb, 0x5f, 0x60, 0x8b, 0xf4, 0xc2, 0x12, 0xac, 0x0b, 0x7f,
	0xeb, 0x8b, 0x16, 0x78, 0xe2, 0xf9, 0xff, 0x5b, 0xca, 0x47, 0x23, 0x11, 0x1d, 0xf1, 0xc4, 0xb1,
	0x4f, 0xd1, 0x56, 0x99, 0x1b, 0x99, 0x47, 0x29, 0xd3, 0x37, 0x5e, 0x85, 0x2c, 0x55, 0xc1, 0x05,
	0xef, 0xdc, 0xd3, 0x9c, 0xda, 0x09, 0x7a, 0x5c, 0xaa, 0xc1, 0xef, 0x92, 0x29, 0xf0, 0x62, 0xff,
	0x5f, 0x2a, 0xd6, 0x2a, 0x68, 0x47, 0x8e, 0x35, 0xaf, 0x15, 0x2b, 0xa0, 0xa6, 0xd0, 0x6a, 0xdf,
	0x5b, 0x6b, 0xe0, 0x58, 0xf3, 0x5a, 0xb9, 0x4c, 0xa6, 0x5a, 0xcf, 0xee, 0xad, 0xf5, 0xb3, 0x63,
	0x39, 0xad, 0x57, 0x68, 0xab, 0x98, 0x04, 0x61, 0xc2, 0xb4, 0x4c, 0xe9, 0x24, 0x74, 0xc3, 0x63,
	0xdb, 0xbd, 0x3e, 0x5c, 0xec, 0x1d, 0xfa, 0xad, 0x1f, 0xed, 0x28, 0x39, 0xb3, 0x9d, 0x44, 0xb1,
	0x31, 0x24, 0xbe, 0x0b, 0xbe, 0xe7, 0xc2, 0x7e, 0xb2, 0xb4, 0x53, 0x1d, 0x7a, 0x92, 0x6b, 0x82,
	0xf5, 0xe4, 0xce, 0xc0, 0x14, 0x61, 0xa9, 0x44, 0x0c, 0x5a, 0xdb, 0x1b, 0x09, 0xe9, 0x0b, 0x65,
	0xc7, 0xc9, 0xee, 0x2d, 0x95, 0x3d, 0x9f, 0x52, 0xcf, 0x3c, 0x33, 0x68, 0xc9, 0x79, 0x57, 0x3b,
	0x44, 0x9b, 0x73, 0xe3, 0x02, 0x3f, 0x45, 0x0f, 0x21, 0xa3, 0x2c, 0xd5, 0xa4, 0xe2, 0xa6, 0x52,
	0x61, 0xd9, 0x91, 0xc6, 0xb8, 0x36, 0x2a, 0x8f, 0xdd, 0x48, 0xf3, 0xd3, 0x72, 0xd6, 0x65, 0xe7,
	0x6b, 0xae, 0x98, 0x26, 0x55, 0xc7, 0x73, 0xeb, 0xf6, 0xdf, 0x15, 0x54, 0x9f, 0xb9, 0x20, 0x1e,
	0xa0, 0x5a, 0x2a, 0xca, 0x9a, 0xaf, 0xb8, 0x5e, 0xbe, 0xb8, 0xcb, 0x9c, 0x16, 0xe8, 0xe0, 0x8e,
	0x87, 0xcf, 0x7c, 0x35, 0xc5, 0xd4, 0xc0, 0x50, 0x28, 0xdb, 0x13, 0xaa, 0xae, 0x8b, 0x76, 0x97,
	0x25, 0x65, 0xe0, 0x19, 0x93, 0x60, 0x63, 0x34, 0x35, 0x18, 0xe8, 0xf6, 0x3f, 0x15, 0xd4, 0x7a,
	0x2b, 0x5f, 0xf8, 0x18, 0xbd, 0x9f, 0x30, 0x4d, 0xa3, 0x14, 0x42, 0x6d, 0x14, 0x80, 0x29, 0xff,
	0x2b, 0x42, 0x05, 0x5a, 0xa4, 0xb9, 0xcb, 0x83, 0xfd, 0x99, 0x58, 0x0b, 0x76, 0x0a, 0xe0, 0x85,
	0xc3, 0xed, 0x7b, 0x58, 0x30, 0x45, 0xd9, 0x31, 0x79, 0x63, 0xb2, 0x34, 0xd4, 0x94, 0x33, 0xc3,
	0xfe, 0xa0, 0xd3, 0x14, 0x36, 0x96, 0x8c, 0xc9, 0x1f, 0x4c, 0x96, 0x5e, 0xcc, 0x90, 0x82, 0xe6,
	0xcd, 0x9c, 0xa7, 0x7d, 0x8b, 0x9e, 0xbc, 0x73, 0xd8, 0xe1, 0x26, 0xaa, 0xfe, 0x0a, 0x93, 0xe2,
	0x77, 0xc7, 0x2e, 0xf1, 0x01, 0x5a, 0xf5, 0x5d, 0x79, 0xe5, 0x1e, 0x6f, 0x73, 0x4e, 0x34, 0xf0,
	0xd4, 0xaf, 0x57, 0xbe, 0xaa, 0x1c, 0xfc, 0x89, 0x9e, 0xc7, 0x22, 0x5b, 0xc4, 0x3e, 0x68, 0x9e,
	0x88, 0xc8, 0x26, 0x20, 0x57, 0x31, 0x9c, 0xdb, 0x8a, 0x3b, 0xaf, 0xbc, 0xde, 0x2f, 0x08, 0x43,
	0x61, 0x27, 0x4f, 0x4f, 0xa8, 0x61, 0x7f, 0x08, 0xdc, 0xd5, 0x63, 0xdf, 0x6f, 0x51, 0xc9, 0xf4,
	0x3b, 0x7f, 0x1a, 0xbf, 0xf1, 0xe6, 0x5f, 0x2b, 0xd5, 0xc1, 0xe5, 0x45, 0xf4, 0xd0, 0x71, 0x3e,
	0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x69, 0x27, 0x27, 0xc9, 0x0a, 0x00, 0x00,
}