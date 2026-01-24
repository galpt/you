# Quick Start Guide

Get started with You orchestrator in 5 minutes.

## Prerequisites

1. **Install OpenCode**: https://opencode.ai/docs/
2. **Set up LLM Provider**: GitHub Copilot Pro (recommended) or any OpenCode provider
   ```bash
   # In OpenCode, run:
   /connect
   # For GitHub Copilot: Select 'github-copilot' (unlimited GPT-5 Mini with Copilot Pro)
   # Or select 'anthropic', 'openai', etc.
   ```

## Step 1: Install You

```bash
# Clone the repository
git clone https://github.com/galpt/you.git
cd you

# Build the binary
# Windows:
go build -o you.exe .
# Linux/Mac:
go build -o you .

# Test installation
# Windows:
.\you.exe --version
# Linux/Mac:
./you --version
```

## Step 2: Create a New Project

```bash
# Create and navigate to your project directory
mkdir my-awesome-project
cd my-awesome-project

# Generate preset files
# Windows:
you.exe --presets
# Linux/Mac:
you --presets
```

This creates:
- `USER_INPUT.md` - Your project requirements
- `.opencode/agents/*.md` - 10 AI agent definitions
- `.opencode/opencode.json` - OpenCode configuration
- `.opencode/skills/*/SKILL.md` - 5 professional skill definitions
- `.you/` - State management directory

## Step 3: Define Your Project

Edit `USER_INPUT.md`:

```markdown
# User Input - Project Requirements

## What do you want to build?

Build a REST API for a blog platform with the following features:
- User authentication (JWT)
- CRUD operations for blog posts
- Comments system
- Tag-based filtering
- PostgreSQL database

### Technical Preferences
- Programming Language: Go
- Framework: Gin
- Database: PostgreSQL
- Deployment: Docker

### Success Criteria
- All endpoints documented with Swagger
- 80%+ test coverage
- README with setup instructions
```

## Step 4: Start Autonomous Orchestration

```bash
# Windows:
you.exe --orchestrate
# Linux/Mac:
you --orchestrate
```

**What happens (fully automated):**
1. Creates a Goal from your USER_INPUT.md
2. Initializes workflow state in `.you/workflows/`
3. Generates `ORCHESTRATION_GUIDE.md`
4. **Starts OpenCode HTTP server in background** (`opencode serve --port 4096`)
5. **Creates HTTP session** for this orchestration
6. **Sends initial prompt to CEO agent** (async, no waiting!)
7. **Streams all agent activity in real-time** to your terminal

**Real-time output example:**
```
🔧 Starting OpenCode server...
📝 Creating orchestration session...
✓ Session created: ses_414190ba65ffeRPG3382HKeVEZH

🎭 Sending initial message to session...
✓ Message queued! CEO agent will receive it and start orchestrating.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📡 Streaming real-time events (press Ctrl+C to stop):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💬 [assistant] ceo
Reading USER_INPUT.md... Analyzing requirements...
   👤 Delegating to: @product-manager
   🔧 Tool: task (running...)
   ✓ Tool task: Delegated to product-manager

💬 [assistant] product-manager
Creating comprehensive PRD...
   🔧 Tool: webfetch (running...)
   ✓ Tool webfetch: Researched best practices
   📄 modified: requirements/PRD.md
   ✓ Tool write completed
   👤 Delegating to: @product-designer

✓ Session completed!
```

**Fully autonomous** - no human interaction needed! Press `Ctrl+C` anytime to stop.

## Step 5: Re-orchestration (if needed)

If something goes wrong or requirements change:

```bash
# 1. Stop current session (Ctrl+C)
# 2. Edit USER_INPUT.md with updated requirements
# 3. Re-run orchestration
you.exe --orchestrate
```

**What happens:**
- Creates new goal with fresh session ID
- Starts new OpenCode server
- CEO agent receives updated requirements
- Builds from scratch with corrected specs
- Previous attempt preserved in `.you/goals/{old-id}/`

## What Happens Next?

The orchestration runs **completely autonomously**:

1. CEO agent reads requirements and creates project plan
2. Delegates to `@product-manager` to create comprehensive PRD
3. PM delegates to `@product-designer` for UX design
4. Designer delegates to `@solution-architect` for system architecture
5. Architect delegates to `@lead-engineer` for task breakdown
6. Lead Engineer delegates to `@software-engineer` for implementation
7. Software Engineer uses `webfetch` to research best practices
8. `@qa-engineer` validates implementation with comprehensive tests
9. `@security-engineer` performs security audit
10. `@devops-sre` sets up deployment pipeline
11. `@technical-writer` creates documentation
12. CEO reviews and approves final deliverable

**All delegation happens automatically** - you just watch the stream!

**Resilience Features (v0.1.11+):**
- Automatic retry with exponential backoff on rate limits (429 errors)
- Up to 5 retry attempts with backoff from 2s to 2 minutes
- Activity monitoring warns if stuck for 10+ minutes
- CEO as primary agent, all others as subagent for correct routing
- Detailed error logging for diagnostics

## Monitoring Progress

All artifacts are saved in `.you/`:

```bash
# View created artifacts
ls -la .you/artifacts/

# Check workflow state
cat .you/workflows/state_<goal-id>.json

# View goals and their sessions
ls -la .you/goals/
```

**Real-time monitoring:** Just watch the terminal! All file operations, agent conversations, and tool calls stream live.

## Example Workflow

Here's a complete example:

```bash
# 1. Setup
mkdir task-manager-app
cd task-manager-app
you --presets

# 2. Edit USER_INPUT.md with requirements
# (Open in your editor and write detailed project specs)

# 3. Start autonomous orchestration
you --orchestrate

# That's it! Watch the magic happen in real-time:
# - CEO agent reads requirements
# - PM creates PRD
# - Designer creates UI mockups
# - Architect designs system
# - Engineers implement features
# - QA tests everything
# - DevOps sets up deployment
# - Tech Writer documents it all
```

**No manual OpenCode interaction needed!** The entire workflow is automated via HTTP API.

## Tips for Success

### 1. Be Specific in USER_INPUT.md
❌ Bad: "Build a website"
✅ Good: "Build a blog platform with user auth, CRUD posts, comments, using Next.js"

### 2. Provide Technical Preferences
Specify your preferred:
- Programming languages
- Frameworks
- Database
- Deployment platform

### 3. Define Success Criteria
- "All tests passing"
- "80%+ code coverage"
- "API documentation complete"
- "Deployed to staging"

### 4. Monitor the Real-time Stream
Watch the terminal output to see:
- Agent conversations and decisions
- Files being created/modified
- Tool invocations (webfetch, delegate, etc.)
- Progress through workflow phases

### 5. Intervene if Needed (Optional)
The system is fully autonomous, but you can stop and restart:
```bash
# Press Ctrl+C to stop
# Edit USER_INPUT.md with changes
you --orchestrate  # Starts fresh with new requirements
```

## How Agents Make Decisions Autonomously

Agents are configured to be **self-sufficient** through:

1. **Smart System Prompts**: Instructed to make reasonable decisions based on best practices
2. **webfetch Tool**: Can research current best practices and documentation
3. **delegate Tool**: Can ask other specialist agents for expertise
4. **Best Practice Guidelines**: Configured with SOLID principles, security-first mindset

**Example autonomous decisions:**
- "Should I use React or Vue?" → Agent uses webfetch to research, picks based on community support
- "Which database schema?" → Follows normalization principles and SOLID design
- "Add tests?" → Yes, implements comprehensive test coverage (best practice)
- "Authentication approach?" → JWT with industry-standard security practices

**No human intervention required!**

## Troubleshooting

### OpenCode server fails to start
- Verify OpenCode is installed: `opencode --version`
- Check port 4096 is available
- Try manually: `opencode serve --port 4096`

### No events streaming / Stuck after "Streaming real-time events"
**Most likely cause:** Using an older version with incorrect agent modes (versions before 0.1.11).

**Solution:** 
- Re-run `you --presets` to regenerate agent configs (CEO must be primary in v0.1.11+)
- Or manually edit `.opencode/opencode.json` and `.opencode/agents/ceo.md`:
  ```yaml
  mode: primary  # CEO must be primary
  ```

**Why this matters:** OpenCode requires at least one primary agent to start the workflow. CEO as primary agent ensures proper orchestration routing. All other agents (PM, Designer, Architect, etc.) remain as subagents.

### No agent responses (messages sent but no replies)
- Verify GitHub Copilot is connected: Run `opencode` TUI and check `/connect` status
- Check if the model provider has API limits or authentication issues
- Try using a different provider in `.opencode/opencode.json`

### Wrong tech stack chosen
- Be more specific in USER_INPUT.md about technical preferences
- Include preferred languages, frameworks, and tools
- Specify architectural patterns if needed

### Want to restart orchestration
```bash
Ctrl+C  # Stop current session
# Edit USER_INPUT.md if needed
you --orchestrate  # Fresh start with new session
```

## Next Steps

- Read the full [README.md](README.md) for comprehensive documentation
- Check [TECHNICAL_DETAILS.md](TECHNICAL_DETAILS.md) for HTTP API architecture deep dive
- Review [IMPLEMENTATION.md](IMPLEMENTATION.md) for codebase overview
- Explore agent prompts in `.opencode/agents/` to understand autonomous behavior and anti-scope-creep protocol
- Join the community (Discord/GitHub discussions)

---

**Happy orchestrating! Watch AI agents build your software autonomously.** 🚀
