package cdklabscdkconstructconnectdatalake


// Amazon Connect analytics dataset types.
//
// These enum values represent the available dataset types that can be associated with Amazon Connect instances
// for data lake integration. The dataset types are updated periodically as new analytics capabilities are added
// to Amazon Connect. String values can be used for dataset types not included in this enum
//
// Regional availability varies by dataset type.
// Experimental.
type DataType string

const (
	// Experimental.
	DataType_CONTACT_RECORD DataType = "CONTACT_RECORD"
	// Experimental.
	DataType_AGENT_QUEUE_STATISTIC_RECORD DataType = "AGENT_QUEUE_STATISTIC_RECORD"
	// Experimental.
	DataType_AGENT_STATISTIC_RECORD DataType = "AGENT_STATISTIC_RECORD"
	// Experimental.
	DataType_CONTACT_STATISTIC_RECORD DataType = "CONTACT_STATISTIC_RECORD"
	// Experimental.
	DataType_CONTACT_EVALUATION_RECORD DataType = "CONTACT_EVALUATION_RECORD"
	// Experimental.
	DataType_CONTACT_LENS_CONVERSATIONAL_ANALYTICS DataType = "CONTACT_LENS_CONVERSATIONAL_ANALYTICS"
	// Experimental.
	DataType_CONTACT_FLOW_EVENTS DataType = "CONTACT_FLOW_EVENTS"
	// Experimental.
	DataType_OUTBOUND_CAMPAIGN_EVENTS DataType = "OUTBOUND_CAMPAIGN_EVENTS"
	// Experimental.
	DataType_FORECAST_GROUPS DataType = "FORECAST_GROUPS"
	// Experimental.
	DataType_LONG_TERM_FORECASTS DataType = "LONG_TERM_FORECASTS"
	// Experimental.
	DataType_SHORT_TERM_FORECASTS DataType = "SHORT_TERM_FORECASTS"
	// Experimental.
	DataType_STAFF_SHIFTS DataType = "STAFF_SHIFTS"
	// Experimental.
	DataType_STAFF_SHIFT_ACTIVITIES DataType = "STAFF_SHIFT_ACTIVITIES"
	// Experimental.
	DataType_STAFF_TIMEOFFS DataType = "STAFF_TIMEOFFS"
	// Experimental.
	DataType_STAFF_TIMEOFF_INTERVALS DataType = "STAFF_TIMEOFF_INTERVALS"
	// Experimental.
	DataType_STAFF_TIMEOFF_BALANCE_CHANGES DataType = "STAFF_TIMEOFF_BALANCE_CHANGES"
	// Experimental.
	DataType_SHIFT_ACTIVITIES DataType = "SHIFT_ACTIVITIES"
	// Experimental.
	DataType_SHIFT_PROFILES DataType = "SHIFT_PROFILES"
	// Experimental.
	DataType_STAFF_SCHEDULING_PROFILE DataType = "STAFF_SCHEDULING_PROFILE"
	// Experimental.
	DataType_STAFFING_GROUP_FORECAST_GROUPS DataType = "STAFFING_GROUP_FORECAST_GROUPS"
	// Experimental.
	DataType_STAFFING_GROUP_SUPERVISORS DataType = "STAFFING_GROUP_SUPERVISORS"
	// Experimental.
	DataType_STAFFING_GROUPS DataType = "STAFFING_GROUPS"
	// Experimental.
	DataType_BOT_CONVERSATIONS DataType = "BOT_CONVERSATIONS"
	// Experimental.
	DataType_BOT_INTENTS DataType = "BOT_INTENTS"
	// Experimental.
	DataType_BOT_SLOTS DataType = "BOT_SLOTS"
	// Experimental.
	DataType_INTRADAY_FORECASTS DataType = "INTRADAY_FORECASTS"
	// Experimental.
	DataType_USERS DataType = "USERS"
	// Experimental.
	DataType_ROUTING_PROFILES DataType = "ROUTING_PROFILES"
	// Experimental.
	DataType_AGENT_HIERARCHY_GROUPS DataType = "AGENT_HIERARCHY_GROUPS"
	// Experimental.
	DataType_AI_AGENT DataType = "AI_AGENT"
	// Experimental.
	DataType_AI_TOOL DataType = "AI_TOOL"
	// Experimental.
	DataType_AI_PROMPT DataType = "AI_PROMPT"
	// Experimental.
	DataType_AI_SESSION DataType = "AI_SESSION"
	// Experimental.
	DataType_AI_AGENT_KNOWLEDGE_BASE DataType = "AI_AGENT_KNOWLEDGE_BASE"
	// Experimental.
	DataType_STAFFING_GROUP_DEMAND_GROUP DataType = "STAFFING_GROUP_DEMAND_GROUP"
	// Experimental.
	DataType_DEMAND_GROUP DataType = "DEMAND_GROUP"
	// Experimental.
	DataType_DEMAND_GROUP_DEFINITIONS DataType = "DEMAND_GROUP_DEFINITIONS"
	// Experimental.
	DataType_STAFF_DEMAND_GROUP DataType = "STAFF_DEMAND_GROUP"
	// Experimental.
	DataType_STAFF_SHIFT_ACTIVITY_ALLOCATIONS DataType = "STAFF_SHIFT_ACTIVITY_ALLOCATIONS"
	// Experimental.
	DataType_CONNECT_TEST_CASE_RESULTS DataType = "CONNECT_TEST_CASE_RESULTS"
)

