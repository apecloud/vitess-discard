package global

const ApeCloud = "apecloud"

// Keyspace
const (
	DefaultKeyspace = "_vt"
	DefaultShard    = "0"

	//// VtDbPrefix + keyspace is the default name for databases.
	VtDbPrefix = "vt_"
	//// EmptyDbPrefix is the default database prefix for apecloud databases.
	EmptyDbPrefix = ""
)

var ApeCloudFeaturesEnable = true

// Planner
var (
	ApeCloudDbDDLPlugin = ApeCloudFeaturesEnable
	UnshardEnabled      = false
	DbPrefix            = func() string {
		if ApeCloudDbDDLPlugin {
			return EmptyDbPrefix
		}
		return VtDbPrefix
	}()
)