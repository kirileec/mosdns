package remove_ecs

import (
	"context"

	"github.com/IrineSistiana/mosdns/v5/pkg/query_context"
	"github.com/IrineSistiana/mosdns/v5/plugin/executable/sequence"
	"github.com/miekg/dns"
)

const PluginType = "remove_ecs"

func init() {
	sequence.MustRegExecQuickSetup(PluginType, QuickSetup)
}

type RemoveECS struct{}

func QuickSetup(_ sequence.BQ, _ string) (any, error) {
	return RemoveECS{}, nil
}

func (r RemoveECS) Exec(_ context.Context, qCtx *query_context.Context) error {
	opt := qCtx.QOpt()
	out := opt.Option[:0]
	for _, o := range opt.Option {
		if o.Option() != dns.EDNS0SUBNET {
			out = append(out, o)
		}
	}
	opt.Option = out
	return nil
}
