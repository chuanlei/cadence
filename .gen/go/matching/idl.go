// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package matching

import (
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "matching",
	Package:  "github.com/uber/cadence/.gen/go/matching",
	FilePath: "matching.thrift",
	SHA1:     "fd8da2f11240e27375e4fc9d55252b17d3bde9c6",
	Includes: []*thriftreflect.ThriftModule{
		shared.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\ninclude \"shared.thrift\"\n\nnamespace java com.uber.cadence.matching\n\nstruct PollForDecisionTaskRequest {\n  10: optional string domainUUID\n  15: optional string pollerID\n  20: optional shared.PollForDecisionTaskRequest pollRequest\n}\n\nstruct PollForDecisionTaskResponse {\n  10: optional binary taskToken\n  20: optional shared.WorkflowExecution workflowExecution\n  30: optional shared.WorkflowType workflowType\n  40: optional i64 (js.type = \"Long\") previousStartedEventId\n  50: optional i64 (js.type = \"Long\") startedEventId\n  51: optional i64 (js.type = \"Long\") attempt\n  60: optional i64 (js.type = \"Long\") nextEventId\n  65: optional i64 (js.type = \"Long\") backlogCountHint\n  70: optional bool stickyExecutionEnabled\n  80: optional shared.WorkflowQuery query\n  90: optional shared.TransientDecisionInfo decisionInfo\n}\n\nstruct PollForActivityTaskRequest {\n  10: optional string domainUUID\n  15: optional string pollerID\n  20: optional shared.PollForActivityTaskRequest pollRequest\n}\n\nstruct AddDecisionTaskRequest {\n  10: optional string domainUUID\n  20: optional shared.WorkflowExecution execution\n  30: optional shared.TaskList taskList\n  40: optional i64 (js.type = \"Long\") scheduleId\n  50: optional i32 scheduleToStartTimeoutSeconds\n}\n\nstruct AddActivityTaskRequest {\n  10: optional string domainUUID\n  20: optional shared.WorkflowExecution execution\n  30: optional string sourceDomainUUID\n  40: optional shared.TaskList taskList\n  50: optional i64 (js.type = \"Long\") scheduleId\n  60: optional i32 scheduleToStartTimeoutSeconds\n}\n\nstruct QueryWorkflowRequest {\n  10: optional string domainUUID\n  20: optional shared.TaskList taskList\n  30: optional shared.QueryWorkflowRequest queryRequest\n}\n\nstruct RespondQueryTaskCompletedRequest {\n  10: optional string domainUUID\n  20: optional shared.TaskList taskList\n  30: optional string taskID\n  40: optional shared.RespondQueryTaskCompletedRequest completedRequest\n}\n\nstruct CancelOutstandingPollRequest {\n  10: optional string domainUUID\n  20: optional i32 taskListType\n  30: optional shared.TaskList taskList\n  40: optional string pollerID\n}\n\nstruct DescribeTaskListRequest {\n  10: optional string domainUUID\n  20: optional shared.DescribeTaskListRequest descRequest\n}\n\n/**\n* MatchingService API is exposed to provide support for polling from long running applications.\n* Such applications are expected to have a worker which regularly polls for DecisionTask and ActivityTask.  For each\n* DecisionTask, application is expected to process the history of events for that session and respond back with next\n* decisions.  For each ActivityTask, application is expected to execute the actual logic for that task and respond back\n* with completion or failure.\n**/\nservice MatchingService {\n  /**\n  * PollForDecisionTask is called by frontend to process DecisionTask from a specific taskList.  A\n  * DecisionTask is dispatched to callers for active workflow executions, with pending decisions.\n  **/\n  PollForDecisionTaskResponse PollForDecisionTask(1: PollForDecisionTaskRequest pollRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n    )\n\n  /**\n  * PollForActivityTask is called by frontend to process ActivityTask from a specific taskList.  ActivityTask\n  * is dispatched to callers whenever a ScheduleTask decision is made for a workflow execution.\n  **/\n  shared.PollForActivityTaskResponse PollForActivityTask(1: PollForActivityTaskRequest pollRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n    )\n\n  /**\n  * AddDecisionTask is called by the history service when a decision task is scheduled, so that it can be dispatched\n  * by the MatchingEngine.\n  **/\n  void AddDecisionTask(1: AddDecisionTaskRequest addRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * AddActivityTask is called by the history service when a decision task is scheduled, so that it can be dispatched\n  * by the MatchingEngine.\n  **/\n  void AddActivityTask(1: AddActivityTaskRequest addRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.ServiceBusyError serviceBusyError,\n    )\n\n  /**\n  * QueryWorkflow is called by frontend to query a workflow.\n  **/\n  shared.QueryWorkflowResponse QueryWorkflow(1: QueryWorkflowRequest queryRequest)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n      4: shared.QueryFailedError queryFailedError,\n    )\n\n  /**\n  * RespondQueryTaskCompleted is called by frontend to respond query completed.\n  **/\n  void RespondQueryTaskCompleted(1: RespondQueryTaskCompletedRequest request)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n      3: shared.EntityNotExistsError entityNotExistError,\n    )\n\n  /**\n    * CancelOutstandingPoll is called by frontend to unblock long polls on matching for zombie pollers.\n    * Our rpc stack does not support context propagation, so when a client connection goes away frontend sees\n    * cancellation of context for that handler, but any corresponding calls (long-poll) to matching service does not\n    * see the cancellation propagated so it can unblock corresponding long-polls on its end.  This results is tasks\n    * being dispatched to zombie pollers in this situation.  This API is added so everytime frontend makes a long-poll\n    * api call to matching it passes in a pollerID and then calls this API when it detects client connection is closed\n    * to unblock long polls for this poller and prevent tasks being sent to these zombie pollers.\n    **/\n  void CancelOutstandingPoll(1: CancelOutstandingPollRequest request)\n    throws (\n      1: shared.BadRequestError badRequestError,\n      2: shared.InternalServiceError internalServiceError,\n    )\n\n  /**\n  * DescribeTaskList returns information about the target tasklist, right now this API returns the\n  * pollers which polled this tasklist in last few minutes.\n  **/\n  shared.DescribeTaskListResponse DescribeTaskList(1: DescribeTaskListRequest request)\n    throws (\n        1: shared.BadRequestError badRequestError,\n        2: shared.InternalServiceError internalServiceError,\n        3: shared.EntityNotExistsError entityNotExistError,\n      )\n}\n"
