/*
 * Copyright (c) 2002-2018 "Neo4j,"
 * Neo4j Sweden AB [http://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package neo4j

import (
    "time"
)

// StatementType defines the type of the statement
type StatementType int

const (
    // StatementTypeUnknown identifies an unknown statement type
    StatementTypeUnknown StatementType = 0
    // StatementTypeReadOnly identifies a read-only statement
    StatementTypeReadOnly = 1
    // StatementTypeReadWrite identifies a read-write statement
    StatementTypeReadWrite = 2
    // StatementTypeWriteOnly identifies a write-only statement
    StatementTypeWriteOnly = 3
    // StatementTypeSchemaWrite identifies a schema-write statement
    StatementTypeSchemaWrite = 4
)

// ServerInfo contains basic information of the server
type ServerInfo struct {
    address string
    version string
}

func (server *ServerInfo) Address() string {
    return server.address
}

func (server *ServerInfo) Version() string {
    return server.version
}

// InputPosition contains information about a specific position in a statement
type InputPosition struct {
    offset int
    line   int
    column int
}

func (pos *InputPosition) Offset() int {
    return pos.offset
}

func (pos *InputPosition) Line() int {
    return pos.line
}

func (pos *InputPosition) Column() int {
    return pos.column
}

// Notification contains information about notifications generated by the server
type Notification struct {
    code        string
    title       string
    description string
    position    InputPosition
    severity    string
}

func (notification *Notification) Code() string {
    return notification.code
}

func (notification *Notification) Title() string {
    return notification.title
}

func (notification *Notification) Description() string {
    return notification.description
}

func (notification *Notification) Position() *InputPosition {
    return &notification.position
}

func (notification *Notification) Severity() string {
    return notification.severity
}

// Counters contains statistics about the changes made to the database made as part
// of the statement execution.
type Counters struct {
    nodesCreated         int
    nodesDeleted         int
    relationshipsCreated int
    relationshipsDeleted int
    propertiesSet        int
    labelsAdded          int
    labelsRemoved        int
    indexesAdded         int
    indexesRemoved       int
    constraintsAdded     int
    constraintsRemoved   int
}

func (counters *Counters) NodesCreated() int {
    return counters.nodesCreated
}

func (counters *Counters) NodesDeleted() int {
    return counters.nodesDeleted
}

func (counters *Counters) RelationshipsCreated() int {
    return counters.relationshipsCreated
}

func (counters *Counters) RelationshipsDeleted() int {
    return counters.relationshipsDeleted
}

func (counters *Counters) PropertiesSet() int {
    return counters.propertiesSet
}

func (counters *Counters) LabelsAdded() int {
    return counters.labelsAdded
}

func (counters *Counters) LabelsRemoved() int {
    return counters.labelsRemoved
}

func (counters *Counters) IndexesAdded() int {
    return counters.indexesAdded
}

func (counters *Counters) IndexesRemoved() int {
    return counters.indexesRemoved
}

func (counters *Counters) ConstraintsAdded() int {
    return counters.constraintsAdded
}

func (counters *Counters) ConstraintsRemoved() int {
    return counters.constraintsRemoved
}

// Plan describes the plan that the database planner produced
type Plan struct {
    operator    string
    arguments   map[string]interface{}
    identifiers []string
    children    []Plan
}

func (plan *Plan) Operator() string {
    return plan.operator
}

func (plan *Plan) Arguments() map[string]interface{} {
    return plan.arguments
}

func (plan *Plan) Identifiers() []string {
    return plan.identifiers
}

func (plan *Plan) Children() []Plan {
    return plan.children
}

// Profile describes the plan that the database planner produced and executed
type ProfiledPlan struct {
    operator    string
    arguments   map[string]interface{}
    identifiers []string
    dbHits      int64
    records     int64
    children    []ProfiledPlan
}

func (plan *ProfiledPlan) Operator() string {
    return plan.operator
}

func (plan *ProfiledPlan) Arguments() map[string]interface{} {
    return plan.arguments
}

func (plan *ProfiledPlan) Identifiers() []string {
    return plan.identifiers
}

func (plan *ProfiledPlan) DbHits() int64 {
    return plan.dbHits
}

func (plan *ProfiledPlan) Records() int64 {
    return plan.records
}

func (plan *ProfiledPlan) Children() []ProfiledPlan {
    return plan.children
}

// ResultSummary contains information about statement execution.
type ResultSummary struct {
    server               ServerInfo
    statement            Statement
    statementType        StatementType
    counters             Counters
    plan                 *Plan
    profile              *ProfiledPlan
    notifications        []Notification
    resultAvailableAfter time.Duration
    resultConsumedAfter  time.Duration
}

func (summary *ResultSummary) Server() *ServerInfo {
    return &summary.server
}

func (summary *ResultSummary) Statement() *Statement {
    return &summary.statement
}

func (summary *ResultSummary) StatementType() StatementType {
    return summary.statementType
}

func (summary *ResultSummary) Counters() *Counters {
    return &summary.counters
}

func (summary *ResultSummary) Plan() *Plan {
    return summary.plan
}

func (summary *ResultSummary) Profile() *ProfiledPlan {
    return summary.profile
}

func (summary *ResultSummary) Notifications() []Notification {
    return summary.notifications
}

func (summary *ResultSummary) ResultAvailableAfter() time.Duration {
    return summary.resultAvailableAfter
}

func (summary *ResultSummary) ResultConsumedAfter() time.Duration {
    return summary.resultConsumedAfter
}
