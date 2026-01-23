package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"you/internal/models"

	"github.com/google/uuid"
)

// SCR (Shared Certified Repository) manages the orchestrator's state
type SCR struct {
	basePath string
}

// NewSCR creates a new Shared Certified Repository
func NewSCR(basePath string) (*SCR, error) {
	scr := &SCR{basePath: basePath}
	if err := scr.initialize(); err != nil {
		return nil, err
	}
	return scr, nil
}

// initialize creates the necessary directories for the SCR
func (s *SCR) initialize() error {
	directories := []string{
		filepath.Join(s.basePath, "artifacts"),
		filepath.Join(s.basePath, "tasks"),
		filepath.Join(s.basePath, "communications"),
		filepath.Join(s.basePath, "workflows"),
	}

	for _, dir := range directories {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

// SaveGoal saves a goal to the SCR
func (s *SCR) SaveGoal(goal *models.Goal) error {
	if goal.ID == "" {
		goal.ID = uuid.New().String()
	}
	if goal.CreatedAt.IsZero() {
		goal.CreatedAt = time.Now()
	}

	filePath := filepath.Join(s.basePath, "workflows", fmt.Sprintf("goal_%s.json", goal.ID))
	return s.saveJSON(filePath, goal)
}

// LoadGoal loads a goal from the SCR
func (s *SCR) LoadGoal(goalID string) (*models.Goal, error) {
	var goal models.Goal
	filePath := filepath.Join(s.basePath, "workflows", fmt.Sprintf("goal_%s.json", goalID))
	if err := s.loadJSON(filePath, &goal); err != nil {
		return nil, err
	}
	return &goal, nil
}

// SaveArtifact saves an artifact to the SCR
func (s *SCR) SaveArtifact(artifact *models.Artifact) error {
	if artifact.ID == "" {
		artifact.ID = uuid.New().String()
	}
	if artifact.CreatedAt.IsZero() {
		artifact.CreatedAt = time.Now()
	}
	artifact.UpdatedAt = time.Now()

	filePath := filepath.Join(s.basePath, "artifacts", fmt.Sprintf("%s_%s.json", artifact.Type, artifact.ID))
	return s.saveJSON(filePath, artifact)
}

// LoadArtifact loads an artifact from the SCR
func (s *SCR) LoadArtifact(artifactID string) (*models.Artifact, error) {
	// Search for the artifact file
	pattern := filepath.Join(s.basePath, "artifacts", fmt.Sprintf("*_%s.json", artifactID))
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("artifact %s not found", artifactID)
	}

	var artifact models.Artifact
	if err := s.loadJSON(matches[0], &artifact); err != nil {
		return nil, err
	}
	return &artifact, nil
}

// ListArtifactsByType lists all artifacts of a specific type
func (s *SCR) ListArtifactsByType(artifactType models.ArtifactType) ([]*models.Artifact, error) {
	pattern := filepath.Join(s.basePath, "artifacts", fmt.Sprintf("%s_*.json", artifactType))
	return s.listArtifacts(pattern)
}

// ListArtifactsByGoal lists all artifacts for a specific goal
func (s *SCR) ListArtifactsByGoal(goalID string) ([]*models.Artifact, error) {
	pattern := filepath.Join(s.basePath, "artifacts", "*.json")
	allArtifacts, err := s.listArtifacts(pattern)
	if err != nil {
		return nil, err
	}

	var goalArtifacts []*models.Artifact
	for _, artifact := range allArtifacts {
		if artifact.GoalID == goalID {
			goalArtifacts = append(goalArtifacts, artifact)
		}
	}
	return goalArtifacts, nil
}

// listArtifacts is a helper to load artifacts matching a pattern
func (s *SCR) listArtifacts(pattern string) ([]*models.Artifact, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	var artifacts []*models.Artifact
	for _, match := range matches {
		var artifact models.Artifact
		if err := s.loadJSON(match, &artifact); err != nil {
			continue // Skip invalid files
		}
		artifacts = append(artifacts, &artifact)
	}
	return artifacts, nil
}

// SaveTask saves a task to the SCR
func (s *SCR) SaveTask(task *models.Task) error {
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	if task.CreatedAt.IsZero() {
		task.CreatedAt = time.Now()
	}
	task.UpdatedAt = time.Now()

	filePath := filepath.Join(s.basePath, "tasks", fmt.Sprintf("task_%s.json", task.ID))
	return s.saveJSON(filePath, task)
}

// LoadTask loads a task from the SCR
func (s *SCR) LoadTask(taskID string) (*models.Task, error) {
	var task models.Task
	filePath := filepath.Join(s.basePath, "tasks", fmt.Sprintf("task_%s.json", taskID))
	if err := s.loadJSON(filePath, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

// ListTasksByGoal lists all tasks for a specific goal
func (s *SCR) ListTasksByGoal(goalID string) ([]*models.Task, error) {
	pattern := filepath.Join(s.basePath, "tasks", "task_*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, match := range matches {
		var task models.Task
		if err := s.loadJSON(match, &task); err != nil {
			continue
		}
		if task.GoalID == goalID {
			tasks = append(tasks, &task)
		}
	}
	return tasks, nil
}

// ListTasksByStatus lists all tasks with a specific status
func (s *SCR) ListTasksByStatus(status models.TaskStatus) ([]*models.Task, error) {
	pattern := filepath.Join(s.basePath, "tasks", "task_*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, match := range matches {
		var task models.Task
		if err := s.loadJSON(match, &task); err != nil {
			continue
		}
		if task.Status == status {
			tasks = append(tasks, &task)
		}
	}
	return tasks, nil
}

// SaveCommunication saves a communication log to the SCR
func (s *SCR) SaveCommunication(comm *models.Communication) error {
	if comm.ID == "" {
		comm.ID = uuid.New().String()
	}
	if comm.CreatedAt.IsZero() {
		comm.CreatedAt = time.Now()
	}

	filePath := filepath.Join(s.basePath, "communications", fmt.Sprintf("comm_%s.json", comm.ID))
	return s.saveJSON(filePath, comm)
}

// SaveWorkflowState saves the current workflow state
func (s *SCR) SaveWorkflowState(state *models.WorkflowState) error {
	state.UpdatedAt = time.Now()
	filePath := filepath.Join(s.basePath, "workflows", fmt.Sprintf("state_%s.json", state.GoalID))
	return s.saveJSON(filePath, state)
}

// LoadWorkflowState loads the workflow state for a goal
func (s *SCR) LoadWorkflowState(goalID string) (*models.WorkflowState, error) {
	var state models.WorkflowState
	filePath := filepath.Join(s.basePath, "workflows", fmt.Sprintf("state_%s.json", goalID))
	if err := s.loadJSON(filePath, &state); err != nil {
		return nil, err
	}
	return &state, nil
}

// saveJSON is a helper to save data as JSON
func (s *SCR) saveJSON(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// loadJSON is a helper to load data from JSON
func (s *SCR) loadJSON(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
	}

	return nil
}
