package mnger

import (
	"xserver/model"
	"xserver/service"
)

var organizaMapper = make(map[string]*model.OprOrganization)

func Company(guid string) *model.OprOrganization {
	v, ok := organizaMapper[guid]
	if ok {
		return v
	}
	data, err := service.OprPrimaryOrganization(guid)
	if err != nil {
		return nil
	}
	organizaMapper[guid] = &data
	return &data
}
