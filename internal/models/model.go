package models

type Item struct {
	Id                        int    `json:"id" gorm:"primaryKey"`
	Uid                       string `json:"uid"`
	Domain                    string `json:"domain"`
	Cn                        string `json:"cn"`
	Department                string `json:"department"`
	Title                     string `json:"title"`
	Who                       string `json:"who"`
	LogonCount                int    `json:"logon_count"`
	NumLogons7                int    `json:"num_logons7"`
	NumShare7                 int    `json:"num_share7"`
	NumFile7                  int    `json:"num_file7"`
	NumAd7                    int    `json:"num_ad7"`
	NumN7                     int    `json:"num_n7"`
	NumLogons14               int    `json:"num_logons14"`
	NumShare14                int    `json:"num_share14"`
	NumFile14                 int    `json:"num_file14"`
	NumAd14                   int    `json:"num_ad14"`
	NumN14                    int    `json:"num_n14"`
	NumLogons30               int    `json:"num_logons30"`
	NumShare30                int    `json:"num_share30"`
	NumFile30                 int    `json:"num_file30"`
	NumAd30                   int    `json:"num_ad30"`
	NumN30                    int    `json:"num_n30"`
	NumLogons150              int    `json:"num_logons150"`
	NumShare150               int    `json:"num_share150"`
	NumFile150                int    `json:"num_file150"`
	NumAd150                  int    `json:"num_ad150"`
	NumN150                   int    `json:"num_n150"`
	NumLogons365              int    `json:"num_logons365"`
	NumShare365               int    `json:"num_share365"`
	NumFile365                int    `json:"num_file365"`
	NumAd365                  int    `json:"num_ad365"`
	NumN365                   int    `json:"num_n365"`
	HasUserPrincipalName      int    `json:"has_user_principal_name"`
	HasMail                   int    `json:"has_mail"`
	HasPhone                  int    `json:"has_phone"`
	FlagDisabled              int    `json:"flag_disabled"`
	FlagLockout               int    `json:"flag_lockout"`
	FlagPasswordNotRequired   int    `json:"flag_password_not_required"`
	FlagPasswordCantChange    int    `json:"flag_password_cant_change"`
	FlagDontExpirePassword    int    `json:"flag_dont_expire_password"`
	OwnedFiles                int    `json:"owned_files"`
	NumMailboxes              int    `json:"num_mailboxes"`
	NumMemberOfGroups         int    `json:"num_member_of_groups"`
	NumMemberOfIndirectGroups int    `json:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIds string `json:"member_of_indirect_groups_ids"`
	MemberOfGroupsIds         string `json:"member_of_groups_ids"`
	IsAdmin                   int    `json:"is_admin" gorm:"default:0;not null"`
	IsService                 int    `json:"is_service" gorm:"default:0;not null"`
}
