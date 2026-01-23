package orchestrator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"you/internal/agents"
	"you/internal/models"
	"you/internal/state"
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
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": false,
        "edit": false,
        "bash": false,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/ceo.md}"
    },
    "product-manager": {
      "description": "Defines requirements, creates PRDs, and validates acceptance criteria",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/product-manager.md}"
    },
    "product-designer": {
      "description": "Creates UI/UX designs, user flows, and design specifications",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.4,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/product-designer.md}"
    },
    "solution-architect": {
      "description": "Designs system architecture, tech stack, and data models",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/solution-architect.md}"
    },
    "lead-engineer": {
      "description": "Breaks architecture into tasks, reviews code, manages releases",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/lead-engineer.md}"
    },
    "software-engineer": {
      "description": "Implements features, writes tests, and submits code for review",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/software-engineer.md}"
    },
    "qa-engineer": {
      "description": "Performs automated testing, validates requirements, reports bugs",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.1,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "allow",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/qa-engineer.md}"
    },
    "security-engineer": {
      "description": "Conducts security audits, identifies vulnerabilities, ensures secure coding practices",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.1,
      "tools": {
        "write": true,
        "edit": false,
        "bash": true,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "deny",
        "bash": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/security-engineer.md}"
    },
    "devops-sre": {
      "description": "Manages CI/CD pipelines, infrastructure, deployment, and observability",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.2,
      "tools": {
        "write": true,
        "edit": true,
        "bash": true,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "bash": "ask",
        "task": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/devops-sre.md}"
    },
    "technical-writer": {
      "description": "Creates documentation, API references, user guides, and changelogs",
      "mode": "subagent",
      "model": "github-models/gpt-5-mini",
      "temperature": 0.3,
      "tools": {
        "write": true,
        "edit": true,
        "bash": false,
        "webfetch": true
      },\n        "skill": true,
      "permission": {
        "edit": "allow",
        "webfetch": "allow",\n        "skill": "allow"
      },
      "prompt": "{file:./.opencode/agents/technical-writer.md}"
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

// launchOpenCodeWithCEO launches OpenCode with the CEO agent and intelligent auto-response
func (o *Orchestrator) launchOpenCodeWithCEO() error {
	// Prepare the prompt for the CEO agent
	prompt := "Read USER_INPUT.md and orchestrate the team to build this project. " +
		"Follow the workflow phases defined in ORCHESTRATION_GUIDE.md. " +
		"Start by delegating to @product-manager to create a comprehensive PRD."

	// Create decision log file
	decisionLog := filepath.Join(o.projectPath, ".you", "decisions.log")
	logFile, err := os.OpenFile(decisionLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create decision log: %w", err)
	}
	defer logFile.Close()

	logDecision := func(question, response string) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		entry := fmt.Sprintf("[%s]\nQ: %s\nA: %s\n\n", timestamp, question, response)
		logFile.WriteString(entry)
		fmt.Printf("\n🤖 Auto-Response: %s\n", response)
	}

	// Use exec.Command to run OpenCode
	cmd := exec.Command("opencode", "run",
		"--agent", "ceo",
		"--prompt", prompt,
		"--dir", o.projectPath,
	)

	// Set working directory
	cmd.Dir = o.projectPath

	// Create pipes for stdin/stdout to intercept and auto-respond
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	// Start OpenCode process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start opencode: %w", err)
	}

	// Create channels for output handling
	done := make(chan error)
	needsResponse := make(chan string, 10)

	// Monitor stdout for questions and auto-respond
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		var recentOutput strings.Builder
		
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line) // Pass through to terminal
			
			// Accumulate recent output for context
			recentOutput.WriteString(line + "\n")
			if recentOutput.Len() > 2000 {
				// Keep only last ~2000 chars
				output := recentOutput.String()
				recentOutput.Reset()
				recentOutput.WriteString(output[len(output)-2000:])
			}
			
			// Detect questions or decision points
			if o.isQuestionOrDecisionPoint(line) {
				needsResponse <- recentOutput.String()
			}
		}
	}()

	// Monitor stderr (pass through)
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Fprintln(os.Stderr, scanner.Text())
		}
	}()

	// Handle auto-responses
	go func() {
		for context := range needsResponse {
			response := o.generateAutoResponse(context)
			logDecision(o.extractQuestion(context), response)
			
			// Send response to OpenCode
			io.WriteString(stdinPipe, response+"\n")
		}
	}()

	// Wait for completion
	go func() {
		done <- cmd.Wait()
	}()

	// Wait for process to complete
	err = <-done
	close(needsResponse)
	
	if err != nil {
		return fmt.Errorf("opencode process failed: %w", err)
	}

	fmt.Println("\n✅ OpenCode session completed!")
	fmt.Printf("📊 Decision log: %s\n", decisionLog)
	return nil
}

// isQuestionOrDecisionPoint detects if output contains a question or requires decision
func (o *Orchestrator) isQuestionOrDecisionPoint(line string) bool {
	line = strings.ToLower(line)
	
	// Pattern matching for questions and decision points
	patterns := []string{
		`\?$`,                           // Ends with question mark
		`which one`,                     // Choice questions
		`should i`,                      // Permission questions
		`do you want`,                   // Preference questions
		`would you like`,                // Polite questions
		`please (choose|select|decide)`, // Decision requests
		`confirm`,                       // Confirmation requests
		`clarify|clarification`,         // Clarification needs
		`(option|choice) [a-z]`,         // Multiple choice
	}
	
	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, line)
		if matched {
			return true
		}
	}
	
	return false
}

// extractQuestion extracts the actual question from context
func (o *Orchestrator) extractQuestion(context string) string {
	lines := strings.Split(context, "\n")
	
	// Find the last line with a question mark or decision keyword
	for i := len(lines) - 1; i >= 0 && i >= len(lines)-10; i-- {
		line := strings.TrimSpace(lines[i])
		if strings.Contains(line, "?") || 
		   strings.Contains(strings.ToLower(line), "which") ||
		   strings.Contains(strings.ToLower(line), "should") {
			return line
		}
	}
	
	// Return last non-empty line
	for i := len(lines) - 1; i >= 0; i-- {
		if line := strings.TrimSpace(lines[i]); line != "" {
			return line
		}
	}
	
	return "Decision point"
}

// generateAutoResponse creates intelligent CEO-level responses
func (o *Orchestrator) generateAutoResponse(context string) string {
	contextLower := strings.ToLower(context)
	
	// Priority/sequencing questions
	if strings.Contains(contextLower, "which one") && 
	   (strings.Contains(contextLower, "first") || strings.Contains(contextLower, "start")) {
		return "Start with the most foundational and critical component that other parts depend on. Follow the natural dependency order."
	}
	
	// Technical choice questions
	if strings.Contains(contextLower, "should i use") || strings.Contains(contextLower, "which technology") {
		return "Use the technology that best matches our requirements, has strong community support, and aligns with modern best practices. Research if needed using webfetch."
	}
	
	// Architecture decisions
	if strings.Contains(contextLower, "architecture") || strings.Contains(contextLower, "design pattern") {
		return "Follow SOLID principles, keep it simple and maintainable. Prefer proven patterns over experimental approaches."
	}
	
	// Testing questions
	if strings.Contains(contextLower, "test") && strings.Contains(contextLower, "should") {
		return "Yes, implement comprehensive tests. Prioritize: unit tests for business logic, integration tests for critical paths, and E2E for user workflows."
	}
	
	// Documentation questions
	if strings.Contains(contextLower, "document") && strings.Contains(contextLower, "should") {
		return "Yes, document all public APIs, architecture decisions, and setup instructions. Keep documentation close to the code."
	}
	
	// Security questions
	if strings.Contains(contextLower, "security") || strings.Contains(contextLower, "encrypt") {
		return "Always prioritize security. Use industry-standard practices, never roll your own crypto, and follow the principle of least privilege."
	}
	
	// Performance questions
	if strings.Contains(contextLower, "optim") || strings.Contains(contextLower, "performance") {
		return "Optimize for readability first, then measure before optimizing for performance. Focus on algorithmic improvements over micro-optimizations."
	}
	
	// Multiple options (A/B/C)
	if regexp.MustCompile(`option [a-c]|choice [a-c]`).MatchString(contextLower) {
		// Extract options and choose the most comprehensive one
		if strings.Contains(contextLower, "comprehensive") || strings.Contains(contextLower, "complete") {
			return "Choose the most comprehensive option that delivers maximum value."
		}
		return "Option A - proceed with the first viable approach and iterate if needed."
	}
	
	// Confirmation requests
	if strings.Contains(contextLower, "confirm") || strings.Contains(contextLower, "proceed") {
		return "Yes, proceed. Make progress and we can iterate based on results."
	}
	
	// Clarification requests
	if strings.Contains(contextLower, "clarify") || strings.Contains(contextLower, "unclear") {
		return "Make a reasonable assumption based on best practices and industry standards. Document your assumption and proceed."
	}
	
	// Error handling questions
	if strings.Contains(contextLower, "error") && strings.Contains(contextLower, "handle") {
		return "Implement graceful error handling with clear error messages. Log errors appropriately and fail fast for unrecoverable errors."
	}
	
	// Default intelligent response
	return "Use your professional judgment and industry best practices. Make the decision that delivers the most value while maintaining code quality and maintainability."
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
- **.you/tasks/**: Individual task definitions and status
- **.you/workflows/**: Goal and workflow state
- **.you/communications/**: Agent-to-agent message logs

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
3. **Iterate**: Agents can refine their work based on feedback
4. **Trust the Process**: Let agents work through the phases systematically

## Troubleshooting

- **Agent stuck?**: Check .you/workflows/ for current phase and blockers
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
