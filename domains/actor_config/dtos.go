package actor_config

type AdminListDto struct{}

type AdminDetailDto struct{}

type BusinessDetailDto struct{}

type UserDetailDto struct{}

func (e *Entity) ToAdminList() *AdminListDto {
	return &AdminListDto{}
}

func (e *Entity) ToAdminDetail() *AdminDetailDto {
	return &AdminDetailDto{}
}

func (e *Entity) ToBusinessDetail() *BusinessDetailDto {
	return &BusinessDetailDto{}
}

func (e *Entity) ToUserDetail() *UserDetailDto {
	return &UserDetailDto{}
}
