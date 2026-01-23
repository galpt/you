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
- [Professional Skills](#professional-skills)
- [Usage Examples](#usage-examples)
- [Design Notes](#design-notes)
- [Limitations & Roadmap](#limitations--roadmap)
- [Contributing](#contributing)
- [License](#license)

## Additional Documentation
- [TECHNICAL_DETAILS.md](TECHNICAL_DETAILS.md) - HTTP API integration and autonomous orchestration deep dive
- [IMPLEMENTATION.md](IMPLEMENTATION.md) - Complete implementation summary and architecture overview
- [QUICKSTART.md](QUICKSTART.md) - 5-minute getting started guide
- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines

---

## Status

**Current Version**: 0.1.0-beta

This is an early beta release. The orchestrator can:
- ✅ Generate OpenCode agent configurations
- ✅ Set up workflow templates and professional skills
- ✅ Track artifacts and state in a Shared Certified Repository (SCR)
- ✅ **Fully autonomous orchestration via HTTP API**
- ✅ **Real-time event streaming** via Server-Sent Events
- ✅ **Re-orchestration support** with session isolation
- 🚧 Multi-project concurrent orchestration
- 🚧 Web UI dashboard for monitoring

---

## Features

- **Fully Autonomous Operation**: Zero human intervention required - agents make intelligent decisions automatically via HTTP API
- **Real-time Event Streaming**: Watch all agent activity, file changes, and decisions in real-time via Server-Sent Events
- **Re-orchestration Support**: Edit requirements and re-run for fresh start with session isolation
- **Single-Prompt Orchestration**: Describe your project once; agents handle the rest
- **10 Specialized Agents**: CEO, Product Manager, Designer, Architect, Lead Engineer, SWE, QA, Security, DevOps, Technical Writer
- **Professional Skills**: Reusable workflows for PRD creation, code review, security audits, API design, deployment checklists
- **Web Research Capabilities**: All agents can browse the internet to verify documentation and latest syntax before writing code
- **Smart Agent System**: Agents research best practices and make autonomous decisions based on industry standards
- **OpenCode HTTP Integration**: Seamless integration with OpenCode's server API for programmatic control
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
cd you

# Install dependencies
go mod download

# Build the binary
# Windows:
go build -o you.exe .
# Linux/Mac:
go build -o you .

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
# Windows:
you.exe --presets
# Linux/Mac:
you --presets
```

This creates:
- `USER_INPUT.md` - Template for your project requirements
- `.opencode/agents/*.md` - 10 specialized agent definitions
- `.opencode/opencode.json` - OpenCode configuration
- `.opencode/skills/*/SKILL.md` - 5 professional skill definitions
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
# Windows:
you.exe --orchestrate
# Linux/Mac:
you --orchestrate
```

**What happens:**
1. Creates a Goal from your USER_INPUT.md
2. Initializes workflow state in `.you/`
3. Generates `ORCHESTRATION_GUIDE.md`
4. **Starts OpenCode server** (headless HTTP API)
5. **Automatically sends initial prompt to CEO agent**
6. **Streams all agent activity in real-time to your terminal**

**Fully Autonomous:**
- No human interaction required
- Agents orchestrate themselves via CEO delegation
- Real-time event streaming shows progress
- Files appear as they're created
- Press `Ctrl+C` anytime to stop

**Example output:**
```
🔧 Starting OpenCode server...
📝 Creating orchestration session...
✓ Session created: abc123

🎭 Sending initial prompt to CEO agent...
✓ CEO agent is now orchestrating the team!

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📡 Streaming real-time events (press Ctrl+C to stop):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💬 [assistant] ceo
   ⚙️  text...
Reading USER_INPUT.md... Delegating to @product-manager...
   📄 created: requirements/PRD.md
💬 [assistant] product-manager
   ⚙️  text...
Creating comprehensive Product Requirements Document...
   📄 modified: requirements/PRD.md
   🔧 Tool: delegate
...
```

### 4. Re-orchestration (if errors occur)

If something goes wrong or you want to refine the requirements:

```bash
# 1. Stop the current session (Ctrl+C)
# 2. Edit USER_INPUT.md with updated requirements
# 3. Re-run orchestration
you.exe --orchestrate
```

**What happens:**
- Creates a new goal with updated requirements
- Starts a fresh OpenCode session
- CEO agent receives the new prompt
- Team rebuilds with corrected specifications

**Note:** Each orchestration creates a new session. Previous sessions are preserved in `.you/goals/` for audit trail.

### 5. Agent Workflow (Automatic)

The CEO agent orchestrates the entire team **without human intervention**:

1. `@ceo` → `@product-manager` - Create PRD
2. `@product-manager` → `@product-designer` - UI/UX design
3. `@product-designer` → `@solution-architect` - System architecture
4. `@solution-architect` → `@lead-engineer` - Task breakdown
5. `@lead-engineer` → `@software-engineer` - Implementation
6. `@software-engineer` → `@qa-engineer` - Testing
7. `@qa-engineer` → `@security-engineer` - Security audit
8. `@security-engineer` → `@devops-sre` - Deployment setup
9. `@devops-sre` → `@technical-writer` - Documentation
10. `@technical-writer` → `@ceo` - Final approval

**All delegation happens automatically via the `delegate` tool.** You just watch the stream!
9. `@technical-writer` creates documentation
10. CEO approves final deliverable

All artifacts and state are tracked in the `.you/` directory.

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

### Autonomous Orchestration via HTTP API

**You** achieves full automation by:

1. **Starting OpenCode Server** (`opencode serve --port 4096`)
   - Headless HTTP server running in background
   - Exposes RESTful API for programmatic control
   - No TUI interaction required

2. **Creating a Session** (`POST /session`)
   - Each orchestration gets a new session ID
   - Sessions are isolated and independent
   - Multiple re-orchestrations = multiple sessions

3. **Sending Asynchronous Prompt** (`POST /session/:id/prompt_async`)
   - Sends initial message to CEO agent
   - Doesn't wait for response (truly async)
   - CEO agent immediately starts orchestrating

4. **Streaming Real-time Events** (`GET /event` - Server-Sent Events)
   - Subscribes to OpenCode event stream
   - Receives updates in real-time:
     - `message.created` - New agent message
     - `message.part.delta` - Streaming text chunks
     - `file.changed` - File created/modified/deleted
     - `tool.call` - Tool invocations (delegate, webfetch, etc.)
   - Formats and displays events in terminal

**Result**: Fully autonomous operation with complete visibility!

**Re-orchestration Flow**:
```
Error detected → User edits USER_INPUT.md → Ctrl+C to stop
  ↓
you.exe --orchestrate
  ↓
New session created → Fresh start → New goal ID
  ↓
Previous session preserved in .you/goals/{old-goal-id}/
```

### Agent Communication

Agents use OpenCode's **delegate tool** to pass work:

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
├── .github/
│   └── workflows/         # CI/CD automation
├── internal/
│   ├── agents/            # Agent templates and prompts
│   │   ├── templates.go   # Agent configurations
│   │   └── prompts.go     # Role-specific system prompts
│   ├── models/            # Data models (Goal, Artifact, Task, etc.)
│   ├── orchestrator/      # Core orchestration logic
│   └── state/             # SCR (Shared Certified Repository)
├── main.go                # CLI entry point
├── go.mod                 # Go module definition
└── README.md
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

**All agents can load professional skills** via the `skill` tool to access standardized workflows for their role-specific tasks (PRD creation, code reviews, security audits, etc.).

Each agent has a custom system prompt based on the fine-tuned prompts in [copilot-agent-modes/chatmode-v3.1/](https://github.com/galpt/copilot-agent-modes/tree/main/chatmode-v3.1).
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
you --presets
# Edit USER_INPUT.md with: "Build a blog platform with user authentication"
you --orchestrate
# OpenCode launches automatically with CEO agent orchestrating the build
```

### Example 2: CLI Tool

```bash
you --presets
# Edit USER_INPUT.md with: "Build a CLI tool for managing TODO lists"
you --orchestrate
# OpenCode launches automatically with CEO agent orchestrating the build
```

### Example 3: API Service

```bash
you --presets
# Edit USER_INPUT.md with: "Build a REST API for a book library"
you --orchestrate
# OpenCode launches automatically with CEO agent orchestrating the build
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
- One orchestration at a time (single OpenCode server instance)
- No automatic retry on OpenCode rate limits
- Terminal-only interface (no web UI)
- Fixed port 4096 for OpenCode server

### Roadmap (v0.2.0)
- [ ] Multi-project concurrent orchestration (multiple OpenCode servers on different ports)
- [ ] Smart retry logic for rate limits with exponential backoff
- [ ] Web UI dashboard for monitoring workflow progress
- [ ] Health check polling instead of fixed sleep on server startup
- [ ] Event filtering and customization (filter by agent, file type, etc.)
- [ ] Integration with GitHub/GitLab for issue tracking
- [ ] Cost tracking and budget limits
- [ ] Agent performance metrics and analytics

### Future Ideas
- [ ] Custom agent creation via CLI
- [ ] Human-in-the-loop approval system for dangerous operations
- [ ] Resume orchestration from checkpoints
- [ ] Agent skill learning and improvement
- [ ] VS Code extension for orchestrator control

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
