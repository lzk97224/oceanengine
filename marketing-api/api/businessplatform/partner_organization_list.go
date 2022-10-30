package businessplatform

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/businessplatform"
)

// PartnerOrganizationList 查询合作组织
func PartnerOrganizationList(clt *core.SDKClient, accessToken string, req *businessplatform.PartnerOrganizationListRequest) (*businessplatform.PartnerOrganizationListData, error) {
	var resp businessplatform.PartnerOrganizationListResponse
	if err := clt.Get("v3.0/business_platform/partner_organization/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
