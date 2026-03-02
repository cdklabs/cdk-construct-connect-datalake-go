// Construct library for Amazon Connect Data Lake
package cdklabscdkconstructconnectdatalake

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-construct-connect-datalake.DataLakeAccess",
		reflect.TypeOf((*DataLakeAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_DataLakeAccess{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-construct-connect-datalake.DataLakeAccessProps",
		reflect.TypeOf((*DataLakeAccessProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/cdk-construct-connect-datalake.DataType",
		reflect.TypeOf((*DataType)(nil)).Elem(),
		map[string]interface{}{
			"CONTACT_RECORD": DataType_CONTACT_RECORD,
			"AGENT_QUEUE_STATISTIC_RECORD": DataType_AGENT_QUEUE_STATISTIC_RECORD,
			"AGENT_STATISTIC_RECORD": DataType_AGENT_STATISTIC_RECORD,
			"CONTACT_STATISTIC_RECORD": DataType_CONTACT_STATISTIC_RECORD,
			"CONTACT_EVALUATION_RECORD": DataType_CONTACT_EVALUATION_RECORD,
			"CONTACT_LENS_CONVERSATIONAL_ANALYTICS": DataType_CONTACT_LENS_CONVERSATIONAL_ANALYTICS,
			"CONTACT_FLOW_EVENTS": DataType_CONTACT_FLOW_EVENTS,
			"OUTBOUND_CAMPAIGN_EVENTS": DataType_OUTBOUND_CAMPAIGN_EVENTS,
			"FORECAST_GROUPS": DataType_FORECAST_GROUPS,
			"LONG_TERM_FORECASTS": DataType_LONG_TERM_FORECASTS,
			"SHORT_TERM_FORECASTS": DataType_SHORT_TERM_FORECASTS,
			"STAFF_SHIFTS": DataType_STAFF_SHIFTS,
			"STAFF_SHIFT_ACTIVITIES": DataType_STAFF_SHIFT_ACTIVITIES,
			"STAFF_TIMEOFFS": DataType_STAFF_TIMEOFFS,
			"STAFF_TIMEOFF_INTERVALS": DataType_STAFF_TIMEOFF_INTERVALS,
			"STAFF_TIMEOFF_BALANCE_CHANGES": DataType_STAFF_TIMEOFF_BALANCE_CHANGES,
			"SHIFT_ACTIVITIES": DataType_SHIFT_ACTIVITIES,
			"SHIFT_PROFILES": DataType_SHIFT_PROFILES,
			"STAFF_SCHEDULING_PROFILE": DataType_STAFF_SCHEDULING_PROFILE,
			"STAFFING_GROUP_FORECAST_GROUPS": DataType_STAFFING_GROUP_FORECAST_GROUPS,
			"STAFFING_GROUP_SUPERVISORS": DataType_STAFFING_GROUP_SUPERVISORS,
			"STAFFING_GROUPS": DataType_STAFFING_GROUPS,
			"BOT_CONVERSATIONS": DataType_BOT_CONVERSATIONS,
			"BOT_INTENTS": DataType_BOT_INTENTS,
			"BOT_SLOTS": DataType_BOT_SLOTS,
			"INTRADAY_FORECASTS": DataType_INTRADAY_FORECASTS,
			"USERS": DataType_USERS,
			"ROUTING_PROFILES": DataType_ROUTING_PROFILES,
			"AGENT_HIERARCHY_GROUPS": DataType_AGENT_HIERARCHY_GROUPS,
			"AI_AGENT": DataType_AI_AGENT,
			"AI_TOOL": DataType_AI_TOOL,
			"AI_PROMPT": DataType_AI_PROMPT,
			"AI_SESSION": DataType_AI_SESSION,
			"AI_AGENT_KNOWLEDGE_BASE": DataType_AI_AGENT_KNOWLEDGE_BASE,
			"STAFFING_GROUP_DEMAND_GROUP": DataType_STAFFING_GROUP_DEMAND_GROUP,
			"DEMAND_GROUP": DataType_DEMAND_GROUP,
			"DEMAND_GROUP_DEFINITIONS": DataType_DEMAND_GROUP_DEFINITIONS,
			"STAFF_DEMAND_GROUP": DataType_STAFF_DEMAND_GROUP,
			"STAFF_SHIFT_ACTIVITY_ALLOCATIONS": DataType_STAFF_SHIFT_ACTIVITY_ALLOCATIONS,
			"CONNECT_TEST_CASE_RESULTS": DataType_CONNECT_TEST_CASE_RESULTS,
		},
	)
}
