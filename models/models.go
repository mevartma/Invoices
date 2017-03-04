package models

type Invoice struct {
	Id             int64      `json:"id"`
	Amount         int64      `json:"amount"`
	CompanyName    string     `json:"company_name"`
	ClientName     string     `json:"client_name"`
	InvoicePath    string     `json:"invoice_path"`
	InvoiceDate    string     `json:"invoice_date"`
	InvoiceNumber  string     `json:"invoice_number"`
	Issuer         string     `json:"issuer"`
	DepartmentName Department `json:"department_name"`
	Approved       bool       `json:"approved"`
	ApprovalsID    []int64    `json:"approvals_id"`
	TicketID       string     `json:"ticket_id"`
	Created_at     string     `json:"created_at"`
	Updated_at     string     `json:"updated_at"`
}

type Department struct {
	Id              int64      `json:"id"`
	Office          string     `json:"office"`
	DeparmentName   string     `json:"deparment_name"`
	ApprovalDetails []Approval `json:"approval_details"`
	BrandName       Brand      `json:"brand_name"`
}

type Brand struct {
	Id        int64  `json:"id"`
	BrandName string `json:"brand_name"`
}

/*
 Approval Levels
 0 - first approval (yahav)
 1 - second approval (avi fatal, asaf dagan)
 2 - third approval (assaf lantiziono/)
*/

type Approval struct {
	Id            int64  `json:"id"`
	Role          string `json:"role"`
	Name          string `json:"name"`
	EamilAddress  string `json:"eamil_address"`
	ApprovalLevel int64  `json:"approval_level"`
}
