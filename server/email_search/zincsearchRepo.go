package email_search

import "context"

type ZincsearchRepo struct{}

func NewZincsearchRepository() Repo {
	return &ZincsearchRepo{}
}

// GetEmails implements Repo.
func (z *ZincsearchRepo) GetEmails(ctx context.Context, filter string) {
	panic("unimplemented")
}
