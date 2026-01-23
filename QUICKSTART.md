# Quick Start Guide

Get started with You orchestrator in 5 minutes.

## Prerequisites

1. **Install OpenCode**: https://opencode.ai/docs/
2. **Set up LLM Provider**: GitHub Copilot Pro (recommended) or any OpenCode provider
   ```bash
   # In OpenCode, run:
   /connect
   # For GitHub Copilot: Select 'github-models' (unlimited GPT-5 Mini with Copilot Pro)
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

## Step 4: Start Orchestration

```bash
# Windows:
you.exe --orchestrate
# Linux/Mac:
you --orchestrate
```

This will:
- Create a Goal from your USER_INPUT.md
- Initialize workflow state in `.you/workflows/`
- Generate `ORCHESTRATION_GUIDE.md`
- **Automatically launch OpenCode with the CEO agent**
- **Auto-respond to all agent questions** - no manual prompting needed!

The entire AI software company runs autonomously. When agents ask questions like "Which feature first?" or "Should I add tests?", the system provides intelligent CEO-level responses automatically. Check `.you/decisions.log` for the decision audit trail.

## Step 5: Run OpenCode

```bash
opencode
```

In OpenCode, invoke the CEO agent:

```
@ceo Read USER_INPUT.md and orchestrate the team to build this project.
```

## What Happens Next?

The CEO agent will:
1. Read your requirements
2. Delegate to `@product-manager` to create a PRD
3. PM delegates to `@product-designer` for UX
4. Designer delegates to `@solution-architect` for architecture
5. Architect delegates to `@lead-engineer` for task breakdown
6. Lead assigns tasks to `@software-engineer` agents
7. `@qa-engineer` validates the implementation
8. `@security-engineer` performs security audit
9. `@devops-sre` sets up deployment
10. `@technical-writer` creates documentation
11. CEO reviews and approves

## Monitoring Progress

All artifacts are saved in `.you/`:

```bash
# View created artifacts
ls -la .you/artifacts/

# Check workflow state
cat .you/workflows/state_<goal-id>.json

# View tasks
ls -la .you/tasks/
```

## Example Workflow

Here's a real example:

```bash
# 1. Setup
mkdir task-manager-app
cd task-manager-app
you --presets

# 2. Edit USER_INPUT.md
echo "Build a task manager web app with Next.js and Tailwind" > USER_INPUT.md

# 3. Orchestrate
you --orchestrate

# 4. Start OpenCode
opencode

# 5. In OpenCode chat:
@ceo Read USER_INPUT.md and build this project. 
Make sure to follow all workflow phases and create high-quality code.
```

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

### 4. Monitor the .you/ Directory
- Check artifacts as they're created
- Review task progress
- Verify workflow state
- **Check `.you/decisions.log`** to see all automated decisions

### 5. Iterate with Agents (If Needed)
The system is fully autonomous, but you can intervene if needed:
```
@lead-engineer The authentication approach in the architecture seems overcomplicated.
Can we simplify it to use JWT instead of sessions?
```

## Autonomous Decision-Making

**You** handles all agent questions automatically:

```
[Agent]: "Which component should I build first - auth or database layer?"
[Auto-Response]: "Start with the database layer as authentication depends on it."

[Agent]: "Should I add input validation?"  
[Auto-Response]: "Yes, implement comprehensive input validation for security."

[Agent]: "Which testing framework - testify or standard testing?"
[Auto-Response]: "Use technology with strong community support and best practices. Research using webfetch."
```

All decisions appear in your terminal with a 🤖 prefix and are logged to `.you/decisions.log`.

## Troubleshooting

### Agent doesn't respond
- Check OpenCode API key configuration
- Verify agent files in `.opencode/agents/` exist
- Check `.opencode/opencode.json` syntax

### Wrong tech stack chosen
- Be more specific in USER_INPUT.md
- Tell the `@solution-architect` directly:
  ```
  @solution-architect Please use Go with Gin framework, not Node.js
  ```

### Tasks not being completed
- Check `.you/tasks/` for task status
- Ask `@lead-engineer` for status:
  ```
  @lead-engineer What's the status of all tasks?
  ```

### Build fails
- Make sure Go 1.21+ is installed
- Run `go mod download` to fetch dependencies
- Check for syntax errors: `go vet ./...`

## Next Steps

- Read the full [README.md](README.md)
- Check [ORCHESTRATION_GUIDE.md](ORCHESTRATION_GUIDE.md) in your project
- Explore agent prompts in `.opencode/agents/`
- Join the community (Discord/GitHub discussions)

## Need Help?

- 📖 Documentation: See [README.md](README.md)
- 🐛 Issues: https://github.com/galpt/you/issues
- 💬 Discord: [Coming soon]
- 📧 Email: galpt@v.recipes

---

**Happy orchestrating! Let the AI agents do the heavy lifting.** 🚀
