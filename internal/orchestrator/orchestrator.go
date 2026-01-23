package orchestrator

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"you/internal/agents"
	"you/internal/models"
	"you/internal/state"

	"github.com/google/uuid"
)

// Orchestrator manages the workflow of AI agents
type Orchestrator struct {
	scr         *state.SCR
	projectPath string
}

// New creates a new Orchestrator instance
func New(projectPath string) (*Orchestrator, error) {
	// Create .you directory for state management
	youPath := filepath.Join(projectPath, ".you")
	scr, err := state.NewSCR(youPath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize SCR: %w", err)
	}

	return &Orchestrator{
		scr:         scr,
		projectPath: projectPath,
	}, nil
}

// SetupPresets generates the initial setup files
func (o *Orchestrator) SetupPresets() error {
	fmt.Println("🚀 Setting up You orchestrator presets...")

	// 1. Create USER_INPUT.md template
	if err := o.createUserInputTemplate(); err != nil {
		return fmt.Errorf("failed to create USER_INPUT.md: %w", err)
	}

	// 2. Create .opencode/agents directory and agent definitions
	if err := o.createAgentDefinitions(); err != nil {
		return fmt.Errorf("failed to create agent definitions: %w", err)
	}

	// 3. Create .opencode/opencode.json config
	if err := o.createOpenCodeConfig(); err != nil {
		return fmt.Errorf("failed to create OpenCode config: %w", err)
	}

	// 4. Create professional skills for agents
	if err := o.createProfessionalSkills(); err != nil {
		return fmt.Errorf("failed to create professional skills: %w", err)
	}

	fmt.Println("✅ Preset files generated successfully!")
	fmt.Println("\n📝 Next steps:")
	fmt.Println("1. Edit USER_INPUT.md with your project requirements")
	fmt.Println("2. Run: you.exe --orchestrate")
	fmt.Println("3. The orchestrator will coordinate agents to build your project")

	return nil
}

// createUserInputTemplate creates the USER_INPUT.md file
func (o *Orchestrator) createUserInputTemplate() error {
	userInputPath := filepath.Join(o.projectPath, "USER_INPUT.md")

	template := `# User Input - Project Requirements

## What do you want to build?
<!-- Describe your project idea in detail. Be specific about what you want to accomplish. -->

Example: "Build a simple task management web application where users can create, update, delete, and mark tasks as complete. The app should have user authentication and a clean, modern UI."

---

## Your Project Description

[Write your project requirements here]

---

## Additional Details (Optional)

### Target Users
<!-- Who will use this application? -->

### Key Features
<!-- List the most important features -->
- Feature 1
- Feature 2
- Feature 3

### Technical Preferences (if any)
<!-- Do you have any preferences for technology stack? -->
- Programming Language:
- Framework:
- Database:

### Constraints
<!-- Any limitations or requirements? -->
- Budget:
- Timeline:
- Performance requirements:

---

## Success Criteria
<!-- How will you know the project is successful? -->

---

## Notes
<!-- Any other information the agents should know -->
`

	if err := os.WriteFile(userInputPath, []byte(template), 0644); err != nil {
		return err
	}

	fmt.Printf("✓ Created USER_INPUT.md at %s\n", userInputPath)
	return nil
}

// createAgentDefinitions creates all agent markdown files in .opencode/agents
func (o *Orchestrator) createAgentDefinitions() error {
	agentsPath := filepath.Join(o.projectPath, ".opencode", "agents")
	if err := os.MkdirAll(agentsPath, 0755); err != nil {
		return err
	}

	templates := agents.GetAllAgentTemplates()
	for _, template := range templates {
		agentFile := filepath.Join(agentsPath, fmt.Sprintf("%s.md", template.Name))
		content := template.ToMarkdown()

		if err := os.WriteFile(agentFile, []byte(content), 0644); err != nil {
			return err
		}

		fmt.Printf("✓ Created agent: %s\n", template.Name)
	}

	return nil
}

// createOpenCodeConfig creates the opencode.json configuration file
func (o *Orchestrator) createOpenCodeConfig() error {
	configPath := filepath.Join(o.projectPath, ".opencode", "opencode.json")

	config := `{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "ceo": {
      "description": "Orchestrates the entire workflow, delegates to PM and reviews final output",
      "mode": "primary",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": false,
        "edit": false,
        "bash": false,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/ceo.md}"
    },
    "product-manager": {
      "description": "Defines requirements, creates PRDs, and validates acceptance criteria",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/product-manager.md}"
    },
    "product-designer": {
      "description": "Creates UI/UX designs, user flows, and design specifications",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.4,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/product-designer.md}"
    },
    "solution-architect": {
      "description": "Designs system architecture, tech stack, and data models",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/solution-architect.md}"
    },
    "lead-engineer": {
      "description": "Breaks architecture into tasks, reviews code, manages releases",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/lead-engineer.md}"
    },
    "software-engineer": {
      "description": "Implements features, writes tests, and submits code for review",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/software-engineer.md}"
    },
    "qa-engineer": {
      "description": "Performs automated testing, validates requirements, reports bugs",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.1,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/qa-engineer.md}"
    },
    "security-engineer": {
      "description": "Conducts security audits, identifies vulnerabilities, ensures secure coding practices",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.1,
      "tools": {
        "write": true,
        "edit": false,
        "bash": true,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "deny",
        "bash": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/security-engineer.md}"
    },
    "devops-sre": {
      "description": "Manages CI/CD pipelines, infrastructure, deployment, and observability",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "ask",
        "task": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/devops-sre.md}"
    },
    "technical-writer": {
      "description": "Creates documentation, API references, user guides, and changelogs",
      "mode": "subagent",
      "model": "github-copilot/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },
      "skill": true,
      "permission": {
        "edit": "allow",
        "webfetch": "allow",
        "skill": "allow"
      },
      "prompt": "{file:agents/technical-writer.md}"
    }
  }
}
`

	if err := os.WriteFile(configPath, []byte(config), 0644); err != nil {
		return err
	}

	fmt.Printf("✓ Created OpenCode config at %s\n", configPath)
	return nil
}

// Orchestrate reads USER_INPUT.md and starts the workflow
func (o *Orchestrator) Orchestrate() error {
	fmt.Println("🎭 Starting orchestration workflow...")

	// 1. Read USER_INPUT.md
	userInput, err := o.readUserInput()
	if err != nil {
		return fmt.Errorf("failed to read USER_INPUT.md: %w", err)
	}

	// 2. Create a Goal from user input
	goal := &models.Goal{
		Description:  userInput,
		Priority:     1,
		Stakeholders: []string{"user"},
		Scope:        "Full project implementation",
		Acceptance: []string{
			"All features implemented as specified",
			"All tests passing",
			"Documentation complete",
		},
	}

	if err := o.scr.SaveGoal(goal); err != nil {
		return fmt.Errorf("failed to save goal: %w", err)
	}

	fmt.Printf("✓ Created goal: %s\n", goal.ID)

	// 3. Initialize workflow state
	workflowState := &models.WorkflowState{
		GoalID:       goal.ID,
		CurrentPhase: "INCEPTION",
		Metadata: map[string]interface{}{
			"started_at": goal.CreatedAt,
		},
	}

	if err := o.scr.SaveWorkflowState(workflowState); err != nil {
		return fmt.Errorf("failed to save workflow state: %w", err)
	}

	// 4. Generate orchestration instructions
	if err := o.generateOrchestrationInstructions(goal); err != nil {
		return fmt.Errorf("failed to generate orchestration instructions: %w", err)
	}

	fmt.Println("\n✅ Orchestration setup complete!")
	fmt.Println("\n� Launching OpenCode with CEO agent...")

	// 5. Automatically launch OpenCode with CEO agent
	if err := o.launchOpenCodeWithCEO(); err != nil {
		fmt.Printf("\n⚠️  Could not automatically launch OpenCode: %v\n", err)
		fmt.Println("\n📋 Manual steps:")
		fmt.Println("1. Run: opencode")
		fmt.Println("2. Tell @ceo: 'Read USER_INPUT.md and orchestrate the team to build this project'")
		return nil
	}

	return nil
}

// launchOpenCodeWithCEO launches OpenCode server and orchestrates via HTTP API
func (o *Orchestrator) launchOpenCodeWithCEO() error {
	// Start OpenCode server in background
	serverCmd := exec.Command("opencode", "serve", "--port", "4096")
	serverCmd.Dir = o.projectPath

	// Don't connect stdin (headless server), but show server logs
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr

	fmt.Println("🔧 Starting OpenCode server...")
	if err := serverCmd.Start(); err != nil {
		return fmt.Errorf("failed to start opencode server: %w", err)
	}

	// Store server process for cleanup
	defer func() {
		if serverCmd.Process != nil {
			fmt.Println("\n🛑 Stopping OpenCode server...")
			serverCmd.Process.Kill()
		}
	}()

	// Wait for server to be ready with health check
	baseURL := "http://localhost:4096"
	fmt.Println("⏳ Waiting for OpenCode server to be ready...")
	time.Sleep(5 * time.Second) // Initial wait

	// Verify server is responding
	healthClient := &http.Client{Timeout: 5 * time.Second}
	for i := 0; i < 6; i++ { // Try for up to 30 more seconds (6 * 5 = 30)
		resp, err := healthClient.Get(baseURL + "/")
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				fmt.Println("✓ OpenCode server is ready!")
				break
			}
		}
		if i < 5 {
			time.Sleep(5 * time.Second)
		} else {
			return fmt.Errorf("opencode server did not become healthy after 35 seconds")
		}
	}

	// Orchestrate via HTTP API
	if err := o.orchestrateViaAPI(baseURL); err != nil {
		return err
	}

	return nil
}

// orchestrateViaAPI orchestrates the project using OpenCode HTTP API
func (o *Orchestrator) orchestrateViaAPI(baseURL string) error {
	// No timeout - let the orchestration run as long as needed
	client := &http.Client{}

	// 1. Create a new session
	fmt.Println("📝 Creating orchestration session...")
	sessionID, err := o.createSession(client, baseURL)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	fmt.Printf("✓ Session created: %s\n\n", sessionID)

	// 2. Start event stream in background to show real-time progress
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventsDone := make(chan error, 1)
	go func() {
		eventsDone <- o.streamEvents(ctx, baseURL)
	}()

	// Give event stream a moment to connect
	time.Sleep(500 * time.Millisecond)

	// 3. Send initial prompt to CEO agent in background (fire and forget)
	prompt := "Read USER_INPUT.md and orchestrate the team to build this project. Follow the workflow phases defined in ORCHESTRATION_GUIDE.md. Start by delegating to @product-manager to create a comprehensive PRD."

	fmt.Println("🎭 Sending initial message to session...")

	// Send message in goroutine - don't wait for response
	// The event stream will show all the activity
	go func() {
		// No timeout - OpenCode can take time to process the message queue
		sendClient := &http.Client{}
		if err := o.sendPromptAsync(sendClient, baseURL, sessionID, prompt); err != nil {
			fmt.Printf("\n⚠️  Warning: Failed to send message: %v\n", err)
			fmt.Println("   But event stream is still running - check for activity")
		}
	}()

	fmt.Println("✓ Message queued! CEO agent will receive it and start orchestrating.")
	fmt.Println()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📡 Streaming real-time events (press Ctrl+C to stop):")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	// 4. Wait for events or user interruption
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		fmt.Println("\n\n⚠️  Received interrupt signal...")
		cancel()
	case err := <-eventsDone:
		if err != nil && err != context.Canceled {
			return fmt.Errorf("event stream error: %w", err)
		}
	}

	fmt.Println("\n✅ Orchestration completed!")
	return nil
}

// createSession creates a new OpenCode session via HTTP API
func (o *Orchestrator) createSession(client *http.Client, baseURL string) (string, error) {
	reqBody := map[string]interface{}{
		"title": "You Orchestrator - Autonomous Build",
	}

	jsonBody, _ := json.Marshal(reqBody)
	resp, err := client.Post(baseURL+"/session", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var session struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return "", err
	}

	return session.ID, nil
}

// sendPromptAsync sends a message to the session
// OpenCode routes messages to primary agents automatically
func (o *Orchestrator) sendPromptAsync(client *http.Client, baseURL, sessionID, message string) error {
	// Generate IDs with proper prefixes (OpenCode requirement)
	// messageID must start with "msg" (like sessionID starts with "ses")
	messageID := "msg_" + uuid.New().String()
	partID := "prt_" + uuid.New().String()

	// OpenCode message format (from opencode-web reference implementation)
	reqBody := map[string]interface{}{
		"messageID":  messageID,
		"providerID": "github-copilot",
		"modelID":    "github-copilot/gpt-5-mini", // Fixed typo: was githhub
		"mode":       "build",
		"parts": []map[string]interface{}{
			{
				"id":        partID,
				"sessionID": sessionID,
				"messageID": messageID,
				"type":      "text",
				"text":      message,
			},
		},
	}

	jsonBody, _ := json.Marshal(reqBody)
	// Correct endpoint: /session/{id}/message NOT /prompt_async
	url := fmt.Sprintf("%s/session/%s/message", baseURL, sessionID)

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// OpenCode /message returns 200 with the message response, not 204
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// streamEvents connects to the SSE event stream and displays events in real-time
func (o *Orchestrator) streamEvents(ctx context.Context, baseURL string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/event", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	scanner := bufio.NewScanner(resp.Body)
	// Increase buffer size to handle large events (default 64KB is too small)
	const maxTokenSize = 10 * 1024 * 1024 // 10MB
	buffer := make([]byte, maxTokenSize)
	scanner.Buffer(buffer, maxTokenSize)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line := scanner.Text()

			// SSE format: "data: {json}"
			if strings.HasPrefix(line, "data: ") {
				eventData := strings.TrimPrefix(line, "data: ")

				// Pretty print event data
				var event map[string]interface{}
				if err := json.Unmarshal([]byte(eventData), &event); err == nil {
					o.displayEvent(event)
				} else {
					// Fallback: just print the raw data
					fmt.Println(eventData)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// displayEvent formats and displays an event from the SSE stream
func (o *Orchestrator) displayEvent(event map[string]interface{}) {
	eventType, _ := event["type"].(string)
	properties, _ := event["properties"].(map[string]interface{})

	switch eventType {
	case "server.connected":
		// Silent - just connected
		return

	case "message.updated":
		// Message metadata - show which agent is working
		if info, ok := properties["info"].(map[string]interface{}); ok {
			role, _ := info["role"].(string)
			agent, _ := info["agent"].(string)

			// Only show if we have meaningful info
			if agent != "" {
				emoji := "💬"
				if role == "assistant" {
					emoji = "🤖"
				}
				fmt.Printf("\n%s [%s] %s\n", emoji, role, agent)
			}
		}

	case "message.part.updated":
		// Message part updated (tool call, text, etc.)
		if part, ok := properties["part"].(map[string]interface{}); ok {
			partType, _ := part["type"].(string)

			switch partType {
			case "text":
				// Streaming text delta
				if delta, ok := properties["delta"].(string); ok && delta != "" {
					fmt.Print(delta)
				}

			case "tool":
				// Tool call
				tool, _ := part["tool"].(string)
				state, _ := part["state"].(map[string]interface{})
				status, _ := state["status"].(string)

				switch status {
				case "running":
					fmt.Printf("   🔧 Tool: %s (running...)\n", tool)
				case "completed":
					title, _ := state["title"].(string)
					if title != "" {
						fmt.Printf("   ✓ Tool %s: %s\n", tool, title)
					} else {
						fmt.Printf("   ✓ Tool %s completed\n", tool)
					}
				}

			case "agent":
				// Agent delegation
				agentName, _ := part["name"].(string)
				fmt.Printf("   👤 Delegating to: @%s\n", agentName)
			}
		}

	case "file.edited":
		// File was created or modified
		if file, ok := properties["file"].(string); ok {
			fmt.Printf("   📄 modified: %s\n", file)
		}

	case "session.status":
		// Session status change (idle, busy, retry)
		if status, ok := properties["type"].(string); ok {
			switch status {
			case "busy":
				// Silent - normal working state
			case "idle":
				fmt.Println("\n✓ Session completed!")
			case "retry":
				message, _ := properties["message"].(string)
				fmt.Printf("\n⚠️  Retry: %s\n", message)
			}
		}

	default:
		// For debugging: uncomment to see all events
		// jsonBytes, _ := json.MarshalIndent(event, "", "  ")
		// fmt.Printf("🔔 %s: %s\n", eventType, string(jsonBytes))
	}
}

// readUserInput reads and returns the content of USER_INPUT.md
func (o *Orchestrator) readUserInput() (string, error) {
	userInputPath := filepath.Join(o.projectPath, "USER_INPUT.md")

	data, err := os.ReadFile(userInputPath)
	if err != nil {
		return "", fmt.Errorf("could not read USER_INPUT.md: %w (did you run --presets first?)", err)
	}

	content := string(data)
	if len(content) == 0 {
		return "", fmt.Errorf("USER_INPUT.md is empty. Please fill it with your project requirements")
	}

	return content, nil
}

// generateOrchestrationInstructions creates a guide for the CEO agent
func (o *Orchestrator) generateOrchestrationInstructions(goal *models.Goal) error {
	instructionsPath := filepath.Join(o.projectPath, "ORCHESTRATION_GUIDE.md")

	instructions := fmt.Sprintf(`# Orchestration Guide for You

This guide explains how the You orchestrator coordinates AI agents to build your project.

## Goal ID: %s

## Workflow Phases

### 1. INCEPTION (CEO → Product Manager)
- **CEO Agent** reads USER_INPUT.md and understands the user's goal
- **CEO** delegates to **@product-manager** to create a PRD
- **Output**: Product Requirements Document (PRD)

### 2. REQUIREMENTS (Product Manager → Product Designer)
- **Product Manager** creates detailed PRD with user stories and acceptance criteria
- **Product Manager** delegates to **@product-designer** for UX work
- **Output**: PRD artifact saved to .you/artifacts/

### 3. DESIGN (Product Designer → Solution Architect)
- **Product Designer** creates user flows and design specifications
- **Product Designer** delegates to **@solution-architect** for technical design
- **Output**: DESIGN_DOC artifact

### 4. ARCHITECTURE (Solution Architect → Lead Engineer)
- **Solution Architect** designs system architecture, tech stack, data models
- **Solution Architect** delegates to **@lead-engineer** to break down into tasks
- **Output**: ARCH_DOC artifact

### 5. PLANNING (Lead Engineer → Software Engineers)
- **Lead Engineer** breaks architecture into concrete, actionable tasks
- **Lead Engineer** assigns tasks to **@software-engineer** agents
- **Output**: TASK_LIST artifact with 10-20 tasks

### 6. IMPLEMENTATION (Software Engineers)
- Multiple **Software Engineers** implement features in parallel
- Each SWE writes code, tests, and submits for review
- **Lead Engineer** reviews code quality
- **Output**: CODE artifacts for each task

### 7. SECURITY REVIEW (Security Engineer)
- **Security Engineer** audits code for vulnerabilities
- **Security Engineer** reports any security issues
- **Output**: Security audit report or BUG_REPORT artifacts

### 8. TESTING (QA Engineer)
- **QA Engineer** validates all PRD requirements
- **QA Engineer** runs automated tests and manual validation
- If bugs found, creates BUG_REPORT and sends back to SWE
- **Output**: TEST_REPORT artifact

### 9. DEPLOYMENT (DevOps/SRE)
- **DevOps** sets up CI/CD pipelines
- **DevOps** configures infrastructure and deployment
- **Output**: Deployment configurations, runbooks

### 10. DOCUMENTATION (Technical Writer)
- **Technical Writer** creates README, API docs, user guides
- **Technical Writer** updates CHANGELOG
- **Output**: Complete documentation

### 11. FINAL REVIEW (CEO)
- **CEO** reviews all artifacts
- **CEO** validates against original user goal
- **CEO** approves or requests changes
- **Output**: Final approval and project completion

## How to Use This

### Start the workflow in OpenCode:
`+"```"+`
opencode
`+"```"+`

### Invoke the CEO agent:
`+"```"+`
@ceo Read USER_INPUT.md and orchestrate the team to build this project. 
Follow the workflow phases to ensure all requirements are met.
`+"```"+`

The CEO will then automatically:
1. Invoke @product-manager
2. Monitor progress
3. Ensure proper handoffs between agents
4. Validate final output

## Monitoring Progress

All artifacts and state are tracked in the .you/ directory:
- **.you/artifacts/**: All deliverables (PRDs, code, docs, test reports)
- **.you/workflows/**: Goal and workflow state

**Note**: Task distribution and agent communication are managed internally by OpenCode.

## Agent Communication

Agents communicate using the **Task tool** in OpenCode:
- @ceo can invoke any agent
- @product-manager can invoke @product-designer
- @solution-architect can invoke @lead-engineer
- @lead-engineer can invoke @software-engineer and @qa-engineer
- And so on...

Each agent has specific permissions defined in .opencode/opencode.json

## Tips for Success

1. **Be Specific**: The more detail in USER_INPUT.md, the better the results
2. **Monitor Artifacts**: Check .you/artifacts/ to see what's been created
3. **Review Workflow**: Check .you/workflows/ for goal and session state
3. **Iterate**: Agents can refine their work based on feedback
4. **Trust the Process**: Let agents work through the phases systematically

## Troubleshooting

- **Agent stuck?**: Monitor the event stream for retry events and error messages
- **Wrong output?**: Provide feedback and ask agent to iterate
- **Missing dependencies?**: Ensure all agents are defined in .opencode/agents/

---

Generated for goal: %s
Created: %s
`, goal.ID, goal.ID, goal.CreatedAt.Format("2006-01-02 15:04:05"))

	if err := os.WriteFile(instructionsPath, []byte(instructions), 0644); err != nil {
		return err
	}

	fmt.Printf("✓ Created orchestration guide at %s\n", instructionsPath)
	return nil
}

// createProfessionalSkills creates production-ready skill definitions for professional software development workflows
// Skills follow industry standards and Roger Pressman's Software Engineering principles
func (o *Orchestrator) createProfessionalSkills() error {
	skillsPath := filepath.Join(o.projectPath, ".opencode", "skills")

	skills := map[string]string{
		"create-prd": `---
name: create-prd
description: Create comprehensive Product Requirement Documents following industry standards
license: MIT
compatibility: opencode
metadata:
  audience: product-manager
  workflow: requirements
---

## What I do

- Gather requirements from user input
- Define user stories with acceptance criteria
- Specify functional and non-functional requirements
- Create feature specifications with priority levels
- Document assumptions and constraints

## When to use me

Use this skill when you need to create a PRD from initial project requirements. 
Ask clarifying questions about:
- Target users and personas
- Success metrics (KPIs)
- Technical constraints
- Timeline expectations

## Output Format

Produce a structured PRD with these sections:
1. Executive Summary
2. User Stories & Personas
3. Feature Specifications
4. Non-Functional Requirements
5. Success Criteria
6. Dependencies & Risks
`,
		"code-review": `---
name: code-review
description: Perform systematic code reviews following engineering best practices
license: MIT
compatibility: opencode
metadata:
  audience: lead-engineer
  workflow: review
---

## What I do

- Review code for correctness, style, and maintainability
- Check for security vulnerabilities and performance issues
- Verify test coverage and edge case handling
- Ensure adherence to team coding standards
- Suggest improvements and optimizations

## When to use me

Use this skill before merging code or when quality review is needed.
Focus on:
- Logic errors and edge cases
- Code smells and anti-patterns
- Security best practices (OWASP Top 10)
- Performance bottlenecks
- Test adequacy

## Review Checklist

- [ ] Code follows project style guide
- [ ] Functions are single-responsibility
- [ ] Error handling is comprehensive
- [ ] Tests cover happy path and edge cases
- [ ] No hardcoded secrets or credentials
- [ ] Performance is acceptable for scale
- [ ] Documentation is clear and accurate
`,
		"security-audit": `---
name: security-audit
description: Conduct security assessments and vulnerability analysis
license: MIT
compatibility: opencode
metadata:
  audience: security-engineer
  workflow: security
---

## What I do

- Identify security vulnerabilities (OWASP Top 10)
- Review authentication and authorization logic
- Check for common attack vectors (SQL injection, XSS, CSRF)
- Verify secure dependency versions
- Assess data protection and privacy compliance

## When to use me

Use this skill for:
- Pre-deployment security review
- Dependency upgrade validation
- Authentication/authorization implementation
- API endpoint security verification

## Security Checklist

- [ ] Input validation on all user inputs
- [ ] Parameterized queries (no SQL injection)
- [ ] HTTPS enforced for all connections
- [ ] Secrets stored in environment variables
- [ ] CSRF protection enabled
- [ ] XSS protection (output encoding)
- [ ] Rate limiting on APIs
- [ ] Security headers configured
- [ ] Dependencies up-to-date (no CVEs)
- [ ] Principle of least privilege enforced
`,
		"api-design": `---
name: api-design
description: Design RESTful and GraphQL APIs following industry best practices
license: MIT
compatibility: opencode
metadata:
  audience: solution-architect
  workflow: architecture
---

## What I do

- Design consistent, intuitive API endpoints
- Define request/response schemas
- Specify authentication and authorization strategies
- Document error handling and status codes
- Plan versioning strategy

## When to use me

Use this skill when designing:
- New API endpoints
- API versioning strategy
- GraphQL schemas
- WebSocket/real-time APIs

## Design Principles

- **REST**: Use proper HTTP verbs (GET, POST, PUT, PATCH, DELETE)
- **Naming**: Use plural nouns for resources (/users, /tasks)
- **Status Codes**: 200 OK, 201 Created, 400 Bad Request, 401 Unauthorized, 404 Not Found, 500 Server Error
- **Pagination**: Include limit/offset or cursor-based pagination
- **Filtering**: Support query parameters for filtering/sorting
- **Versioning**: Use /v1/ prefix or Accept header versioning
- **Documentation**: OpenAPI/Swagger specification
`,
		"deployment-checklist": `---
name: deployment-checklist
description: Comprehensive pre-deployment validation and go-live procedures
license: MIT
compatibility: opencode
metadata:
  audience: devops-sre
  workflow: deployment
---

## What I do

- Validate all pre-deployment requirements
- Verify environment configurations
- Check monitoring and alerting setup
- Confirm rollback procedures
- Ensure documentation is complete

## When to use me

Use this skill before any production deployment or major release.

## Pre-Deployment Checklist

### Code & Testing
- [ ] All tests passing (unit, integration, e2e)
- [ ] Code review approved
- [ ] Security scan completed (no high/critical issues)
- [ ] Performance tested under expected load

### Infrastructure
- [ ] Environment variables configured
- [ ] Database migrations tested
- [ ] SSL certificates valid
- [ ] CDN/caching configured
- [ ] Backup strategy in place

### Monitoring
- [ ] Application logs configured
- [ ] Error tracking enabled (Sentry, etc.)
- [ ] Performance monitoring active (APM)
- [ ] Uptime monitoring configured
- [ ] Alerts configured for critical metrics

### Documentation
- [ ] Deployment runbook updated
- [ ] Rollback procedure documented
- [ ] API documentation current
- [ ] Changelog updated
- [ ] Stakeholders notified

### Rollback Plan
- [ ] Previous version tagged
- [ ] Rollback tested in staging
- [ ] Database rollback strategy defined
- [ ] Rollback triggers identified
`,
	}

	for skillName, content := range skills {
		skillDir := filepath.Join(skillsPath, skillName)
		if err := os.MkdirAll(skillDir, 0755); err != nil {
			return err
		}

		skillFile := filepath.Join(skillDir, "SKILL.md")
		if err := os.WriteFile(skillFile, []byte(content), 0644); err != nil {
			return err
		}

		fmt.Printf("✓ Created skill: %s\n", skillName)
	}

	return nil
}
