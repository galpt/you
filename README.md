# You - Agentic Orchestrator

An autonomous AI agent orchestrator that coordinates an entire software company within [OpenCode](https://opencode.ai/).

This project aims to supercharge AI agents without burdening ***you*** as the user to prompt after each turn. Instead of manually managing multiple agents, **You** orchestrates a complete software company where AI agents act as Product Managers, Architects, Engineers, QA, DevOps, and more—all working together from a single user input.

---

## Table of Contents
- [Status](#status)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [How It Works](#how-it-works)
- [Architecture](#architecture)
- [Agent Roles](#agent-roles)
- [Usage Examples](#usage-examples)
- [Design Notes](#design-notes)
- [Limitations & Roadmap](#limitations--roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Status

**Current Version**: 0.1.0-beta

This is an early beta release. The orchestrator can:
- ✅ Generate OpenCode agent configurations
- ✅ Set up workflow templates
- ✅ Track artifacts and state in a Shared Certified Repository (SCR)
- ✅ Provide orchestration guidance for AI agents
- 🚧 Automated agent-to-agent handoffs (manual via OpenCode Task tool)
- 🚧 Automatic error recovery and retry logic

---

## Features

- **Single-Prompt Orchestration**: Describe your project once; agents handle the rest
- **10 Specialized Agents**: CEO, Product Manager, Designer, Architect, Lead Engineer, SWE, QA, Security, DevOps, Technical Writer
- **Reusable Skills**: Agents can load standardized workflows via OpenCode skills (PRD creation, code review, security audits, etc.)
- **Web Research Capabilities**: All agents can browse the internet to verify documentation and latest syntax before writing code
- **Obsolescence Protocol**: Agents automatically research external libraries/APIs to prevent outdated implementations
- **OpenCode Integration**: Seamless integration with OpenCode's agent system
- **State Management**: Tracks goals, tasks, artifacts, and communications in a centralized SCR
- **Production-Ready Code**: Clean architecture following SOLID principles
- **Workflow Automation**: Agents follow a structured workflow from requirements to deployment

---

## Requirements

- **Go**: 1.21 or higher
- **OpenCode**: Latest version ([install here](https://opencode.ai/docs/))
- **GitHub Copilot Pro**: Recommended for unlimited GPT-5 Mini access
- **Alternative**: Any OpenCode-supported LLM provider (Anthropic, OpenAI, etc.)

---

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/galpt/you.git
cd you/push-to-github

# Install dependencies
go mod download

# Build the binary
go build -o you.exe .

# (Optional) Add to PATH for global access
# Windows: Copy you.exe to a directory in your PATH
# Linux/Mac: sudo mv you /usr/local/bin/
```

### From Release

Download the latest binary from the [Releases](https://github.com/galpt/you/releases) page.

---

## Quick Start

### 1. Generate Preset Files

```bash
cd your-new-project
you.exe --presets
```

This creates:
- `USER_INPUT.md` - Template for your project requirements
- `.opencode/agents/*.md` - 10 specialized agent definitions
- `.opencode/opencode.json` - OpenCode configuration
- `.you/` - State management directory

### 2. Define Your Project

Edit `USER_INPUT.md` with your project idea:

```markdown
# User Input - Project Requirements

## What do you want to build?

Build a simple task management web application where users can create, 
update, delete, and mark tasks as complete. The app should have user 
authentication and a clean, modern UI using Next.js and Tailwind CSS.

### Key Features
- User registration and login
- Create/edit/delete tasks
- Mark tasks as complete
- Filter tasks (all, active, completed)
- Responsive design
```

### 3. Start Orchestration

```bash
you.exe --orchestrate
```

This:
- Creates a Goal from your user input
- Initializes workflow state
- Generates `ORCHESTRATION_GUIDE.md`

### 4. Run OpenCode

```bash
opencode
```

In OpenCode, invoke the CEO agent:

```
@ceo Read USER_INPUT.md and orchestrate the team to build this project.
```

The CEO agent will then:
1. Delegate to `@product-manager` to create a PRD
2. Monitor as PM delegates to `@product-designer`
3. Designer delegates to `@solution-architect`
4. Architect delegates to `@lead-engineer`
5. Lead Engineer assigns tasks to `@software-engineer` agents
6. `@qa-engineer` validates the implementation
7. `@security-engineer` performs security audit
8. `@devops-sre` sets up deployment
9. `@technical-writer` creates documentation
10. CEO approves final deliverable

---

## How It Works

### Workflow Phases

```
User Goal → CEO Agent
            ↓
         Product Manager (PRD)
            ↓
         Product Designer (UI/UX)
            ↓
         Solution Architect (Architecture)
            ↓
         Lead Engineer (Task Breakdown)
            ↓
         Software Engineers (Implementation)
            ↓
         QA Engineer (Testing) ←→ Bug Reports → Back to SWE
            ↓
         Security Engineer (Audit)
            ↓
         DevOps/SRE (Deployment)
            ↓
         Technical Writer (Documentation)
            ↓
         CEO (Final Review & Approval)
```

### State Management (SCR)

All artifacts are tracked in `.you/`:

```
.you/
├── artifacts/       # PRDs, code, test reports, docs
├── tasks/           # Individual task definitions
├── workflows/       # Goals and workflow state
└── communications/  # Agent-to-agent messages
```

### Agent Communication

Agents use OpenCode's **Task tool** to delegate work:

```
@ceo → invokes → @product-manager
@product-manager → invokes → @product-designer
@solution-architect → invokes → @lead-engineer
@lead-engineer → invokes → @software-engineer
```

Each agent has specific permissions and tools defined in `.opencode/opencode.json`.

---

## Architecture

```
you/
├── cmd/                    # CLI entry points (future expansion)
├── internal/
│   ├── agents/            # Agent templates and prompts
│   │   ├── templates.go   # Agent configurations
│   │   └── prompts.go     # Role-specific system prompts
│   ├── models/            # Data models (Goal, Artifact, Task, etc.)
│   ├── orchestrator/      # Core orchestration logic
│   └── state/             # SCR (Shared Certified Repository)
├── templates/             # File templates
├── main.go                # CLI entry point
└── go.mod
```

**Design Principles**:
- **Separation of Concerns**: Each package has a single responsibility
- **Dependency Injection**: Orchestrator depends on interfaces, not concrete types
- **Clean Architecture**: Business logic is independent of frameworks
- **SOLID Principles**: Single responsibility, open/closed, Liskov substitution, etc.

---

## Agent Roles

| Agent | Role | Responsibilities |
|-------|------|------------------|
| **CEO** | Orchestrator | High-level decision making, delegates to PM, final approval |
| **Product Manager** | Requirements | Creates PRDs, user stories, acceptance criteria |
| **Product Designer** | UI/UX | User flows, design systems, component specifications |
| **Solution Architect** | Architecture | Tech stack, data models, API design, system architecture |
| **Lead Engineer** | Task Management | Breaks architecture into tasks, code review, releases |
| **Software Engineer** | Implementation | Writes code, unit tests, implements features |
| **QA Engineer** | Quality Assurance | Automated testing, validation, bug reporting |
| **Security Engineer** | Security | Security audits, vulnerability scanning, compliance |
| **DevOps/SRE** | Infrastructure | CI/CD, deployment, monitoring, observability |
| **Technical Writer** | Documentation | READMEs, API docs, user guides, changelogs |

**All agents have web browsing capabilities** to research documentation, verify latest library syntax, and prevent outdated implementations. Before writing code for external dependencies, agents automatically use `webfetch` to validate current best practices.


---

## Professional Skills

The orchestrator includes 5 production-ready skills following industry standards and Software Engineering best practices (Roger Pressman):

| Skill | Description | Primary Users | Standards Applied |
|-------|-------------|---------------|-------------------|
| **create-prd** | Comprehensive PRD creation with user stories, acceptance criteria, and requirements | Product Manager | Requirements Engineering, Stakeholder Analysis |
| **code-review** | Systematic code review checklist covering correctness, security, and maintainability | Lead Engineer, SWE | Quality Assurance, Code Inspection, Best Practices |
| **security-audit** | Security assessment following OWASP Top 10, dependency scanning, and compliance checks | Security Engineer | Security Engineering, Risk Management |
| **api-design** | RESTful and GraphQL API design with versioning, documentation, and best practices | Solution Architect | Interface Design, Software Architecture |
| **deployment-checklist** | Pre-deployment validation covering testing, infrastructure, monitoring, and rollback | DevOps/SRE | Process Models, Configuration Management, Quality Control |

These skills mirror real software company workflows and are production-ready—not examples. Each follows established software engineering principles including:
- **Process Discipline**: Structured workflows with clear phases
- **Quality Assurance**: Built-in validation and verification steps
- **Risk Management**: Explicit consideration of dependencies and constraints
- **Documentation Standards**: Consistent format and completeness requirements

Skills are stored in `.opencode/skills/<skill-name>/SKILL.md` and can be extended with your organization's specific methodologies.

**How Agents Use Skills:**
```
@lead-engineer Review the user authentication code using the code-review skill
```

OpenCode presents available skills to agents, who load them on-demand for structured, professional guidance.

---

## Usage Examples

### Example 1: Web Application

```bash
you.exe --presets
# Edit USER_INPUT.md with: "Build a blog platform with user authentication"
you.exe --orchestrate
opencode
# In OpenCode: @ceo Read USER_INPUT.md and build this project
```

### Example 2: CLI Tool

```bash
you.exe --presets
# Edit USER_INPUT.md with: "Build a CLI tool for managing TODO lists"
you.exe --orchestrate
opencode
# In OpenCode: @ceo Read USER_INPUT.md and build this project
```

### Example 3: API Service

```bash
you.exe --presets
# Edit USER_INPUT.md with: "Build a REST API for a book library"
you.exe --orchestrate
opencode
# In OpenCode: @ceo Read USER_INPUT.md and build this project
```

---

## Design Notes

### Why OpenCode?

OpenCode provides:
- Native agent support with custom prompts
- Task delegation between agents
- File operations and terminal access
- LLM provider flexibility

### Shared Certified Repository (SCR)

Inspired by Kubernetes' etcd, the SCR is:
- **Immutable**: All changes create new versions
- **Auditable**: Every decision is logged
- **State-driven**: Agents react to state changes

### Persistence Protocol

Agents follow a "never stop until complete" protocol:
1. Agent receives task
2. Agent works autonomously
3. Agent only stops when acceptance criteria met
4. If blocked, agent researches or escalates

---

## Limitations & Roadmap

### Current Limitations
- Agents communicate via manual Task tool invocation (not fully automated)
- No automatic retry on OpenCode rate limits
- No persistent chat history between orchestration runs
- Limited error recovery

### Roadmap (v0.2.0)
- [ ] Automatic agent-to-agent handoffs
- [ ] Smart retry logic for rate limits
- [ ] Web UI for monitoring workflow progress
- [ ] Support for multiple concurrent goals
- [ ] Integration with GitHub/GitLab for issue tracking
- [ ] Cost tracking and budget limits
- [ ] Agent performance metrics

### Future Ideas
- [ ] Custom agent creation via CLI
- [ ] Multi-project orchestration
- [ ] Human-in-the-loop decision points
- [ ] Agent skill learning and improvement

---

## Contributing

Contributions are welcome! This is a beta project looking for feedback.

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** (follow Go best practices)
4. **Run tests**: `go test ./...`
5. **Commit**: `git commit -m 'Add amazing feature'`
6. **Push**: `git push origin feature/amazing-feature`
7. **Open a Pull Request**

### Code Standards

- Follow Go idioms and conventions
- Use `gofmt` and `golint`
- Add tests for new features
- Update documentation

### Reporting Issues

Found a bug? Have a suggestion? [Open an issue](https://github.com/galpt/you/issues)!

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- Inspired by [MetaGPT](https://github.com/geekan/MetaGPT) and hierarchical agent systems
- Built for [OpenCode](https://opencode.ai/)
- Follows software engineering principles from Roger Pressman's *Software Engineering* textbook
- Agent personality fine-tuning inspired by Anthropic's prompt engineering guide

---

**Built with ❤️ for developers who want AI agents to do the heavy lifting.**
