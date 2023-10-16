package campaign

type CreateCampaignRequest struct {
	CampaignName string `json:"campaignName"`
	CampaignDesc string `json:"campaignDesc"`
}

// func (req *CreateCampaignRequest) validate() error {
// 	if utf8.RuneCountInString(req.CampaignName) == 0 {
// 		return errors.New("campaign is required field")
// 	}
// 	return nil
// }
