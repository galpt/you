# Architecture Diagram

## System Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                         USER                                     │
│                           ↓                                      │
│                    USER_INPUT.md                                 │
│                  (Project Requirements)                          │
└─────────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────────┐
│                    You Orchestrator (CLI)                        │
│                                                                  │
│  Commands:                                                       │
│  • you.exe --presets     → Generate agent configs                │
│  • you.exe --orchestrate → Initialize workflow                   │
└─────────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────────┐
│                       OpenCode                                   │
│                                                                  │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │              AI Agents (Subagents)                      │   │
│  │                                                         │   │
│  │  ┌──────────┐         ┌──────────────┐                │   │
│  │  │   CEO    │────────→│ Product Mgr  │                │   │
│  │  └──────────┘         └──────────────┘                │   │
│  │       ↓                      ↓                         │   │
│  │  ┌──────────────┐     ┌──────────────┐                │   │
│  │  │ Final Review │     │   Designer   │                │   │
│  │  └──────────────┘     └──────────────┘                │   │
│  │                              ↓                         │   │
│  │                       ┌──────────────┐                │   │
│  │                       │  Architect   │                │   │
│  │                       └──────────────┘                │   │
│  │                              ↓                         │   │
│  │                       ┌──────────────┐                │   │
│  │                       │Lead Engineer │                │   │
│  │                       └──────────────┘                │   │
│  │                              ↓                         │   │
│  │     ┌────────────────────────┴────────────┐           │   │
│  │     ↓                    ↓                ↓           │   │
│  │  ┌─────┐            ┌─────┐          ┌─────┐         │   │
│  │  │ SWE │            │ SWE │          │ SWE │         │   │
│  │  └─────┘            └─────┘          └─────┘         │   │
│  │     ↓                    ↓                ↓           │   │
│  │     └────────────────────┬────────────────┘           │   │
│  │                          ↓                            │   │
│  │                   ┌──────────────┐                    │   │
│  │                   │  QA Engineer │                    │   │
│  │                   └──────────────┘                    │   │
│  │                          ↓                            │   │
│  │            ┌─────────────┴─────────────┐              │   │
│  │            ↓                           ↓              │   │
│  │     ┌────────────┐              ┌────────────┐       │   │
│  │     │  Security  │              │  DevOps    │       │   │
│  │     └────────────┘              └────────────┘       │   │
│  │            ↓                           ↓              │   │
│  │            └────────────┬──────────────┘              │   │
│  │                         ↓                             │   │
│  │                  ┌──────────────┐                     │   │
│  │                  │ Tech Writer  │                     │   │
│  │                  └──────────────┘                     │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────────┐
│            Shared Certified Repository (SCR)                     │
│                      (.you/ directory)                           │
│                                                                  │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────┐   │
│  │ Goals    │  │Artifacts │  │  Tasks   │  │Communications│   │
│  └──────────┘  └──────────┘  └──────────┘  └──────────────┘   │
│                                                                  │
│  Artifact Types:                                                 │
│  • PRD (Product Requirements)                                    │
│  • DESIGN_DOC (UI/UX Specifications)                            │
│  • ARCH_DOC (Architecture)                                       │
│  • TASK_LIST (Engineering Tasks)                                │
│  • CODE (Implementation)                                         │
│  • TEST_REPORT (QA Results)                                     │
│  • BUG_REPORT (Issues)                                          │
└─────────────────────────────────────────────────────────────────┘
```

## Data Flow

```
1. USER writes requirements
   └─→ USER_INPUT.md

2. You CLI generates setup
   ├─→ .opencode/agents/*.md (Agent definitions)
   ├─→ .opencode/opencode.json (Configuration)
   ├─→ USER_INPUT.md (Template)
   └─→ .you/ (State directory)

3. User runs OpenCode
   └─→ @ceo agent starts workflow

4. CEO delegates to PM
   └─→ PM creates PRD artifact

5. PM delegates to Designer
   └─→ Designer creates DESIGN_DOC artifact

6. Designer delegates to Architect
   └─→ Architect creates ARCH_DOC artifact

7. Architect delegates to Lead Engineer
   └─→ Lead creates TASK_LIST artifact

8. Lead assigns tasks to SWEs
   └─→ Each SWE creates CODE artifacts

9. SWEs complete → QA validates
   └─→ QA creates TEST_REPORT or BUG_REPORT artifacts

10. QA passes → Security audits
    └─→ Security creates audit report

11. Security passes → DevOps deploys
    └─→ DevOps creates deployment configs

12. DevOps completes → Tech Writer documents
    └─→ Tech Writer creates README, API docs

13. All complete → CEO reviews
    └─→ CEO approves → Project complete ✅
```

## Module Dependencies

```
main.go
  │
  ├─→ internal/orchestrator
  │     │
  │     ├─→ internal/agents
  │     │     │
  │     │     └─→ internal/models
  │     │
  │     ├─→ internal/models
  │     │
  │     └─→ internal/state
  │           │
  │           └─→ internal/models
  │
  └─→ github.com/google/uuid
```

## File System Structure

```
project-root/
│
├── USER_INPUT.md              # User's project requirements
├── ORCHESTRATION_GUIDE.md     # Generated workflow guide
│
├── .opencode/
│   ├── opencode.json         # OpenCode configuration
│   └── agents/               # Agent definitions (10 files)
│       ├── ceo.md
│       ├── product-manager.md
│       ├── product-designer.md
│       ├── solution-architect.md
│       ├── lead-engineer.md
│       ├── software-engineer.md
│       ├── qa-engineer.md
│       ├── security-engineer.md
│       ├── devops-sre.md
│       └── technical-writer.md
│
└── .you/                     # State management (SCR)
    ├── artifacts/            # Deliverables (JSON files)
    │   ├── PRD_*.json
    │   ├── DESIGN_DOC_*.json
    │   ├── ARCH_DOC_*.json
    │   ├── TASK_LIST_*.json
    │   ├── CODE_*.json
    │   ├── TEST_REPORT_*.json
    │   └── BUG_REPORT_*.json
    │
    ├── tasks/                # Task tracking (JSON files)
    │   └── task_*.json
    │
    ├── workflows/            # Goal and state tracking
    │   ├── goal_*.json
    │   └── state_*.json
    │
    ├── communications/       # Agent message logs
    │   └── comm_*.json
    │
    └── decisions.log         # Audit trail of automated decisions
```

## Auto-Response System Architecture

```
OpenCode Process
     ↓
  stdout pipe
     ↓
  ┌────────────────────────────┐
  │  Output Monitor (goroutine)│
  │  • Scans each line          │
  │  • Pattern matching         │
  │  • Question detection       │
  └────────────────────────────┘
     ↓
  Question Detected?
     ↓
  ┌────────────────────────────┐
  │  Decision Engine            │
  │  • Extract context          │
  │  • Generate CEO response    │
  │  • Log decision             │
  └────────────────────────────┘
     ↓
  stdin pipe
     ↓
  OpenCode receives response
     ↓
  Agent continues working
```

### Decision Pattern Matching

The system detects these question patterns:

| Pattern | Example | Auto-Response Strategy |
|---------|---------|------------------------|
| `?$` | "Should I add tests?" | Context-aware yes/no with reasoning |
| `which one` | "Which component first?" | Dependency-based prioritization |
| `should i` | "Should I use React?" | Best practices + research suggestion |
| `do you want` | "Do you want me to optimize?" | Strategic decision based on phase |
| `please (choose\|select)` | "Please choose A or B" | Value-based selection |
| `confirm` | "Confirm deployment?" | Phase-appropriate approval |
| `clarify` | "Need clarification" | Assumption + documentation |

### Response Generation Logic

```go
func generateAutoResponse(context string) string {
    // 1. Priority/Sequencing
    if containsPattern(context, "which.*first") {
        return "Start with foundational components"
    }
    
    // 2. Technical Choices
    if containsPattern(context, "should i use|which technology") {
        return "Use proven tech with strong community + research"
    }
    
    // 3. Architecture Decisions
    if containsPattern(context, "architecture|design pattern") {
        return "Follow SOLID, keep it maintainable"
    }
    
    // 4. Testing
    if containsPattern(context, "test.*should") {
        return "Yes, comprehensive testing required"
    }
    
    // ... 10+ more decision patterns
    
    // Default
    return "Use professional judgment and best practices"
}
```

### Decision Audit Trail

Every automated decision is logged to `.you/decisions.log`:

```
[2026-01-23 14:30:22]
Q: Which component should I implement first - authentication or database models?
A: Start with the most foundational and critical component that other parts depend on. Follow the natural dependency order.

[2026-01-23 14:35:18]
Q: Should I add input validation to the API endpoints?
A: Yes, implement comprehensive input validation for security and data integrity.

[2026-01-23 14:42:05]
Q: Which testing framework should I use - testify or standard testing?
A: Use the technology that best matches our requirements, has strong community support, and aligns with modern best practices. Research if needed using webfetch.
```

This provides:
- **Transparency**: Full visibility into automated decisions
- **Auditability**: Track reasoning for post-mortem analysis
- **Learning**: Understand decision patterns over time

## Agent Communication Protocol

```
Agent A wants to delegate to Agent B:

1. Agent A uses OpenCode Task tool
   ┌─────────────────────────────────────────┐
   │ @agent-b Please [task description]       │
   │                                          │
   │ Context: [Artifact IDs, requirements]    │
   │ Expected output: [Artifact type]         │
   └─────────────────────────────────────────┘

2. Agent B receives task and processes

3. Agent B saves artifact to SCR
   .you/artifacts/[TYPE]_[UUID].json

4. Agent B updates task status
   .you/tasks/task_[UUID].json

5. Agent B reports completion or delegates further
```

## Artifact Lifecycle

```
┌─────────┐
│ DRAFT   │ ← Created by agent
└────┬────┘
     ↓
┌──────────────┐
│PENDING_REVIEW│ ← Submitted for review
└──────┬───────┘
       ↓
  ┌────┴────┐
  ↓         ↓
┌──────┐  ┌──────────┐
│APPROVED│  │REJECTED│
└────┬───┘  └────┬────┘
     ↓           ↓
  [Next Phase]  [Revise]
```

## Technology Stack

```
┌─────────────────────────────────────────┐
│          Application Layer               │
│                                          │
│  • Go 1.21+ (Main language)             │
│  • Standard library (io, os, json)      │
│  • github.com/google/uuid               │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│         Integration Layer                │
│                                          │
│  • OpenCode (Agent runtime)             │
│  • GitHub Models GPT-5 Mini (default)   │
│  • File system (State persistence)      │
└─────────────────────────────────────────┘
```

---

## Key Design Patterns

1. **Command Pattern**: CLI commands encapsulate operations
2. **Repository Pattern**: SCR provides data access abstraction
3. **Factory Pattern**: Agent templates create agent instances
4. **State Pattern**: WorkflowState tracks orchestration phases
5. **Strategy Pattern**: Different agents = different strategies

## Concurrency Model

```
User Process (you.exe)
  │
  ├─ Generate files (sequential)
  └─ Exit

OpenCode Process
  │
  ├─ Agent 1 (sequential within agent)
  │   ├─ Think
  │   ├─ Act (save artifact)
  │   └─ Delegate (invoke next agent)
  │
  ├─ Agent 2 (started by Agent 1)
  │   └─ ...
  │
  └─ Multiple SWE agents (could run in parallel if OpenCode supports)
```

Currently single-threaded within You orchestrator.
Parallelism happens at OpenCode agent level (future enhancement).

---

*This diagram represents v0.1.0-beta architecture*
