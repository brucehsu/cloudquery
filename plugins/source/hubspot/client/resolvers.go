package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func ResolveObjectType(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	cl := meta.(*Client)
	return r.Set("object_type", cl.ObjectType)
}
