package user

import (
	"github.com/samber/mo"
	"github.com/taniko/event-sourcing/internal/domain/model/user/vo"
)

// ChangeProfile プロフィールを変更
type ChangeProfile struct {
	Name mo.Option[vo.Name]
}
