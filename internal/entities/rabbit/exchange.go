package rabbitEntities

type Exchange string

const (
	LeaCentralWashExchange Exchange = "lea_central_wash_service"
	SbpClientExchange      Exchange = "sbp_client_service"
	AdminsExchange         Exchange = "admins_exchange"
	WashBonusExchange      Exchange = "wash_bonus_service"
)
