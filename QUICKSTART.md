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
cd you/push-to-github

# Build the binary
go build -o you.exe .

# Test installation
./you.exe --version
```

## Step 2: Create a New Project

```bash
# Create and navigate to your project directory
mkdir my-awesome-project
cd my-awesome-project

# Generate preset files
/path/to/you.exe --presets
```

This creates:
- `USER_INPUT.md` - Your project requirements
- `.opencode/agents/*.md` - 10 AI agent definitions
- `.opencode/opencode.json` - OpenCode configuration
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
/path/to/you.exe --orchestrate
```

This will:
- Create a Goal from your USER_INPUT.md
- Initialize workflow state in `.you/workflows/`
- Generate `ORCHESTRATION_GUIDE.md`

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
you.exe --presets

# 2. Edit USER_INPUT.md
echo "Build a task manager web app with Next.js and Tailwind" > USER_INPUT.md

# 3. Orchestrate
you.exe --orchestrate

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

### 5. Iterate with Agents
If something isn't right, just tell the agent:
```
@lead-engineer The authentication approach in the architecture seems overcomplicated.
Can we simplify it to use JWT instead of sessions?
```

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
