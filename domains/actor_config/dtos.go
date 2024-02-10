package actor_config

type AdminListDto struct{}

func (e *Entity) ToAdminList() *AdminListDto {
	return &AdminListDto{}
}
