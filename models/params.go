package models

type ByUserSearch struct {
	Sender string
	Recip  string
}

func (bus *ByUserSearch) OnlySender() bool {
	return bus.Sender != "" && bus.Recip == ""
}

func (bus *ByUserSearch) OnlyRecip() bool {
	return bus.Sender == "" && bus.Recip != ""
}

func (bus *ByUserSearch) Both() bool {
	return bus.Sender != "" && bus.Recip != ""
}
