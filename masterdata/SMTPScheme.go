package masterdata

import "time"

// Metadata used to mail the EDS's Report.
type SMTPScheme struct {
	// Email Trigger Frequency for the EDS report, cron job format            
	EmailTriggerFrequency                                          *string    `json:"EmailTriggerFrequency,omitempty"`
	// Unique name given to a SMTP scheme.                                    
	Name                                                           string     `json:"Name"`
	// End date till report is required.                                      
	ReportEndDate                                                  *time.Time `json:"ReportEndDate,omitempty"`
	// Start date from which report is required.                              
	ReportStartDate                                                *time.Time `json:"ReportStartDate,omitempty"`
	// Key for the SMTP host.                                                 
	SMTPHostKeyName                                                *string    `json:"SmtpHostKeyName,omitempty"`
	// Key for the SMTP password.                                             
	SMTPPasswordKeyName                                            *string    `json:"SmtpPasswordKeyName,omitempty"`
	// SMTP Port for the connection,its different for SSL/StartTLS.           
	SMTPPort                                                       *int64     `json:"SmtpPort,omitempty"`
	// List of receiver's mail addresses.                                     
	SMTPReceiverMail                                               []string   `json:"SmtpReceiverMail,omitempty"`
	// Limit to retry for the connection.                                     
	SMTPRetryLimit                                                 *int64     `json:"SmtpRetryLimit,omitempty"`
	// Sender's mail address.                                                 
	SMTPSenderMail                                                 *string    `json:"SmtpSenderMail,omitempty"`
	// Type of protocol for connection.                                       
	SMTPSSL                                                        *bool      `json:"SmtpSSL,omitempty"`
	// Type of protocol for connection.                                       
	SMTPStartTLS                                                   *bool      `json:"SmtpStartTLS,omitempty"`
	// Timeout time for the SMTP server.                                      
	SMTPTimeOut                                                    *int64     `json:"SmtpTimeOut,omitempty"`
	// Key for the SMTP user.                                                 
	SMTPUserKeyName                                                *string    `json:"SmtpUserKeyName,omitempty"`
}
