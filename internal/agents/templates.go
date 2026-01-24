package agents

import (
	"fmt"
	"you/internal/models"
)

// AgentTemplate represents the configuration for an OpenCode agent
type AgentTemplate struct {
	Name        string
	Role        models.AgentRole
	Description string
	Mode        string // "primary" or "subagent"
	Model       string
	Temperature float64
	Tools       map[string]bool
	Permissions map[string]string
	Prompt      string
}

// GetAllAgentTemplates returns all agent templates for the orchestrator
func GetAllAgentTemplates() []AgentTemplate {
	return []AgentTemplate{
		getCEOAgentTemplate(),
		getPMAgentTemplate(),
		getDesignerAgentTemplate(),
		getArchitectAgentTemplate(),
		getLeadEngineerAgentTemplate(),
		getSWEAgentTemplate(),
		getQAAgentTemplate(),
		getSecurityAgentTemplate(),
		getDevOpsAgentTemplate(),
		getTechnicalWriterAgentTemplate(),
	}
}

func getCEOAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "ceo",
		Role:        models.AgentRoleCEO,
		Description: "Orchestrates the entire workflow, delegates to PM and reviews final output",
		Mode:        "primary",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.2,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"task":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateCEOPrompt(),
	}
}

func getPMAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "product-manager",
		Role:        models.AgentRolePM,
		Description: "Defines requirements, creates PRDs, and validates acceptance criteria",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.3,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generatePMPrompt(),
	}
}

func getDesignerAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "product-designer",
		Role:        models.AgentRoleDesigner,
		Description: "Creates UI/UX designs, user flows, and design specifications",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.4,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateDesignerPrompt(),
	}
}

func getArchitectAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "solution-architect",
		Role:        models.AgentRoleArchitect,
		Description: "Designs system architecture, tech stack, and data models",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.2,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateArchitectPrompt(),
	}
}

func getLeadEngineerAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "lead-engineer",
		Role:        models.AgentRoleLeadEngineer,
		Description: "Breaks architecture into tasks, reviews code, manages releases",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.2,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"bash":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateLeadEngineerPrompt(),
	}
}

func getSWEAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "software-engineer",
		Role:        models.AgentRoleSWE,
		Description: "Implements features, writes tests, and submits code for review",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.3,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"bash":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateSWEPrompt(),
	}
}

func getQAAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "qa-engineer",
		Role:        models.AgentRoleQA,
		Description: "Performs automated testing, validates requirements, reports bugs",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.1,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"bash":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateQAPrompt(),
	}
}

func getSecurityAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "security-engineer",
		Role:        models.AgentRoleSecurity,
		Description: "Conducts security audits, identifies vulnerabilities, ensures secure coding practices",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.1,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"bash":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateSecurityPrompt(),
	}
}

func getDevOpsAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "devops-sre",
		Role:        models.AgentRoleDevOps,
		Description: "Manages CI/CD pipelines, infrastructure, deployment, and observability",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.2,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"bash":     "ask",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateDevOpsPrompt(),
	}
}

func getTechnicalWriterAgentTemplate() AgentTemplate {
	return AgentTemplate{
		Name:        "technical-writer",
		Role:        models.AgentRoleTechnicalWriter,
		Description: "Creates documentation, API references, user guides, and changelogs",
		Mode:        "subagent",
		Model:       "github-copilot/gpt-5-mini",
		Temperature: 0.3,
		Tools: map[string]bool{
			"write":    true,
			"edit":     true,
			"bash":     true,
			"webfetch": true,
			"skill":    true,
		},
		Permissions: map[string]string{
			"edit":     "allow",
			"webfetch": "allow",
			"skill":    "allow",
		},
		Prompt: generateTechnicalWriterPrompt(),
	}
}

// ToMarkdown converts an AgentTemplate to OpenCode markdown format
func (a *AgentTemplate) ToMarkdown() string {
	md := fmt.Sprintf(`---
description: "%s"
mode: %s
model: %s
temperature: %.1f
tools:
`, a.Description, a.Mode, a.Model, a.Temperature)

	for tool, enabled := range a.Tools {
		md += fmt.Sprintf("  %s: %t\n", tool, enabled)
	}

	if len(a.Permissions) > 0 {
		md += "permission:\n"
		for perm, value := range a.Permissions {
			md += fmt.Sprintf("  %s: %s\n", perm, value)
		}
	}

	md += "---\n\n"
	md += a.Prompt

	return md
}
