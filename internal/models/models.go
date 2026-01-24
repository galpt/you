package models

import "time"

// ArtifactType represents the type of artifact
type ArtifactType string

const (
	ArtifactTypePRD        ArtifactType = "PRD"
	ArtifactTypeArchDoc    ArtifactType = "ARCH_DOC"
	ArtifactTypeCode       ArtifactType = "CODE"
	ArtifactTypeBugReport  ArtifactType = "BUG_REPORT"
	ArtifactTypeTestReport ArtifactType = "TEST_REPORT"
	ArtifactTypeDesignDoc  ArtifactType = "DESIGN_DOC"
	ArtifactTypeTaskList   ArtifactType = "TASK_LIST"
)

// ArtifactStatus represents the lifecycle state of an artifact
type ArtifactStatus string

const (
	ArtifactStatusDraft         ArtifactStatus = "DRAFT"
	ArtifactStatusPendingReview ArtifactStatus = "PENDING_REVIEW"
	ArtifactStatusApproved      ArtifactStatus = "APPROVED"
	ArtifactStatusRejected      ArtifactStatus = "REJECTED"
)

// AgentRole represents the role of an agent in the software company
type AgentRole string

const (
	AgentRoleCEO               AgentRole = "CEO"
	AgentRolePM                AgentRole = "PRODUCT_MANAGER"
	AgentRoleDesigner          AgentRole = "PRODUCT_DESIGNER"
	AgentRoleArchitect         AgentRole = "SOLUTION_ARCHITECT"
	AgentRoleLeadEngineer      AgentRole = "LEAD_ENGINEER"
	AgentRoleSWE               AgentRole = "SOFTWARE_ENGINEER"
	AgentRoleQA                AgentRole = "QA_ENGINEER"
	AgentRoleSecurity          AgentRole = "SECURITY_ENGINEER"
	AgentRoleDevOps            AgentRole = "DEVOPS_SRE"
	AgentRoleTechnicalWriter   AgentRole = "TECHNICAL_WRITER"
	AgentRoleGuardrail         AgentRole = "GUARDRAIL"
)

// TaskStatus represents the state of a task
type TaskStatus string

const (
	TaskStatusNotStarted TaskStatus = "NOT_STARTED"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
	TaskStatusBlocked    TaskStatus = "BLOCKED"
)

// Goal represents a user's high-level objective
type Goal struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Stakeholders []string `json:"stakeholders"`
	Scope       string    `json:"scope"`
	Priority    int       `json:"priority"`
	Acceptance  []string  `json:"acceptance_criteria"`
	CreatedAt   time.Time `json:"created_at"`
}

// Artifact represents a deliverable produced by agents
type Artifact struct {
	ID        string         `json:"id"`
	Type      ArtifactType   `json:"type"`
	Version   string         `json:"version"`
	Content   interface{}    `json:"content"`
	Status    ArtifactStatus `json:"status"`
	GoalID    string         `json:"goal_id"`
	CreatedBy AgentRole      `json:"created_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	FilePath  string         `json:"file_path,omitempty"`
}

// Task represents a unit of work assigned to an agent
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AssignedTo  AgentRole  `json:"assigned_to"`
	ArtifactID  string     `json:"artifact_id"`
	GoalID      string     `json:"goal_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Agent represents an AI agent with specific capabilities
type Agent struct {
	Role         AgentRole `json:"role"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Capabilities []string  `json:"capabilities"`
	PromptFile   string    `json:"prompt_file"`
}

// Communication represents a message between agents
type Communication struct {
	ID             string    `json:"id"`
	FromAgent      AgentRole `json:"from_agent"`
	ToAgent        AgentRole `json:"to_agent"`
	Payload        string    `json:"payload"`
	ConversationID string    `json:"conversation_id"`
	GoalID         string    `json:"goal_id"`
	CreatedAt      time.Time `json:"created_at"`
}

// WorkflowState represents the current orchestration state
type WorkflowState struct {
	GoalID           string                 `json:"goal_id"`
	CurrentPhase     string                 `json:"current_phase"`
	ActiveTasks      []string               `json:"active_tasks"`
	CompletedTasks   []string               `json:"completed_tasks"`
	PendingArtifacts []string               `json:"pending_artifacts"`
	Metadata         map[string]interface{} `json:"metadata"`
	UpdatedAt        time.Time              `json:"updated_at"`
}
