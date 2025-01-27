package notebooks

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func fetchNotebooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.NotebooksAPI.ListNotebooks(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
