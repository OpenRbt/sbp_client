package rabbitEntities

type Queue string

const (
	SBPStartupQueue   Queue = "sbp_startup"
	SBPAdminDataQueue Queue = "sbp_admin_data"
)
