package actor_config

type AdminListDto struct {
	UUID      string `json:"uuid"`
	Actor     Actor  `json:"actor"`
	Telegram  int    `json:"telegram"`
	Mail      int    `json:"mail"`
	SMS       int    `json:"sms"`
	UpdatedAt string `json:"updatedAt"`
}

type AdminDetailDto struct {
	*Entity
}

type BusinessDetailDto struct {
	*Entity
}

type UserDetailDto struct {
	*Entity
}

func (e *Entity) ToAdminList() *AdminListDto {
	return &AdminListDto{
		UUID:      e.UUID,
		Actor:     e.Actor,
		Telegram:  len(e.Telegram),
		Mail:      len(e.Mail),
		SMS:       len(e.SMS),
		UpdatedAt: e.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (e *Entity) ToAdminDetail() *AdminDetailDto {
	return &AdminDetailDto{
		Entity: e,
	}
}

func (e *Entity) ToBusinessDetail() *BusinessDetailDto {
	return &BusinessDetailDto{
		Entity: e,
	}
}

func (e *Entity) ToUserDetail() *UserDetailDto {
	return &UserDetailDto{
		Entity: e,
	}
}
