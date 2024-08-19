// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePacksResultListDataSourceEnvelope struct {
	Result *[]*CertificatePacksResultDataSourceModel `json:"result,computed"`
}

type CertificatePacksDataSourceModel struct {
	ZoneID   types.String                              `tfsdk:"zone_id" path:"zone_id"`
	Status   types.String                              `tfsdk:"status" query:"status"`
	MaxItems types.Int64                               `tfsdk:"max_items"`
	Result   *[]*CertificatePacksResultDataSourceModel `tfsdk:"result"`
}

type CertificatePacksResultDataSourceModel struct {
}
