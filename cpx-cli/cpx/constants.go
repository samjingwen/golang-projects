package cpx

var BASE_URL = "http://localhost:8888"

var SERVICES = map[string]struct{}{
	"PermissionsService": {},
	"AuthService":        {},
	"MLService":          {},
	"StorageService":     {},
	"TimeService":        {},
	"GeoService":         {},
	"TicketService":      {},
	"RoleService":        {},
	"IdService":          {},
	"UserService":        {},
}

const COUNT_THRESHOLD = 2
const CPU_THRESHOLD = 80
const MEM_THRESHOLD = 80
const MAX_CONCURRENT_REQUESTS = 5
