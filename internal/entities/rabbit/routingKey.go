package rabbitEntities

type RoutingKey string

const (
	SbpClientRoutingKey RoutingKey = "sbp_client"
	WashBonusRoutingKey RoutingKey = "wash_bonus"
)
