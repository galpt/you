# Auto-Response System

## Overview

The **You** orchestrator includes an intelligent auto-response system that makes it a truly autonomous software company. When OpenCode agents ask questions or need decisions, the system automatically provides CEO-level responses without requiring human intervention.

## How It Works

### 1. Output Monitoring

```
OpenCode Output → Scanner → Pattern Matcher → Decision Engine → Auto-Response → OpenCode Input
```

The system monitors all OpenCode output in real-time, scanning each line for:
- Question marks (`?`)
- Decision keywords (`which`, `should`, `confirm`, `clarify`)
- Multiple choice indicators (`option A`, `choice B`)

### 2. Pattern Detection

| Pattern | Matches |
|---------|---------|
| `\?$` | Lines ending with question mark |
| `which one` | Choice questions |
| `should i` | Permission/approval questions |
| `do you want` | Preference questions |
| `please (choose\|select\|decide)` | Explicit decision requests |
| `confirm` | Confirmation requests |
| `clarify\|clarification` | Clarification needs |

### 3. Intelligent Response Generation

The system provides context-aware responses based on software engineering best practices:

#### Priority & Sequencing
```
Q: "Which task should I do first?"
A: "Start with the most foundational and critical component that other parts depend on. Follow the natural dependency order."
```

#### Technical Choices
```
Q: "Should I use React or Vue?"
A: "Use the technology that best matches our requirements, has strong community support, and aligns with modern best practices. Research if needed using webfetch."
```

#### Architecture Decisions
```
Q: "How should I structure the database schema?"
A: "Follow SOLID principles, keep it simple and maintainable. Prefer proven patterns over experimental approaches."
```

#### Testing
```
Q: "Should I write unit tests for this?"
A: "Yes, implement comprehensive tests. Prioritize: unit tests for business logic, integration tests for critical paths, and E2E for user workflows."
```

#### Security
```
Q: "How should I handle API authentication?"
A: "Always prioritize security. Use industry-standard practices, never roll your own crypto, and follow the principle of least privilege."
```

#### Performance
```
Q: "Should I optimize this loop?"
A: "Optimize for readability first, then measure before optimizing for performance. Focus on algorithmic improvements over micro-optimizations."
```

#### Documentation
```
Q: "Do we need API documentation?"
A: "Yes, document all public APIs, architecture decisions, and setup instructions. Keep documentation close to the code."
```

#### Clarification
```
Q: "I need clarification on the authentication flow"
A: "Make a reasonable assumption based on best practices and industry standards. Document your assumption and proceed."
```

#### Confirmation
```
Q: "Should I proceed with deployment?"
A: "Yes, proceed. Make progress and we can iterate based on results."
```

### 4. Decision Logging

Every automated decision is logged to `.you/decisions.log`:

```
[2026-01-23 14:30:22]
Q: Which component should I implement first?
A: Start with the most foundational and critical component that other parts depend on. Follow the natural dependency order.

[2026-01-23 14:35:18]
Q: Should I implement authentication using JWT or sessions?
A: Use the technology that best matches our requirements, has strong community support, and aligns with modern best practices. Research if needed using webfetch.
```

## Benefits

### True Autonomy
No human intervention needed during execution. Describe your project once, then let the AI company build it.

### Transparency
Full audit trail of all decisions in `.you/decisions.log`

### Intelligent Defaults
Responses based on:
- Software engineering best practices
- SOLID principles
- Industry standards
- Security-first mindset
- Maintainability over cleverness

### Continuous Progress
Agents never get stuck waiting for human input. The workflow continues until completion.

## Implementation Details

### Code Location
`internal/orchestrator/orchestrator.go`

### Key Functions
- `launchOpenCodeWithCEO()` - Main orchestration with pipes
- `isQuestionOrDecisionPoint()` - Pattern matching
- `extractQuestion()` - Context extraction
- `generateAutoResponse()` - Response generation

### Architecture
```go
// Monitor stdout
go func() {
    scanner := bufio.NewScanner(stdoutPipe)
    for scanner.Scan() {
        line := scanner.Text()
        if isQuestionOrDecisionPoint(line) {
            needsResponse <- recentOutput.String()
        }
    }
}()

// Auto-respond
go func() {
    for context := range needsResponse {
        response := generateAutoResponse(context)
        logDecision(extractQuestion(context), response)
        io.WriteString(stdinPipe, response+"\n")
    }
}()
```

## Customization

To modify response patterns, edit the `generateAutoResponse()` function in `internal/orchestrator/orchestrator.go`.

Example adding a new pattern:

```go
// Custom deployment strategy
if strings.Contains(contextLower, "deploy") && strings.Contains(contextLower, "strategy") {
    return "Deploy to staging first, run smoke tests, then promote to production with blue-green deployment."
}
```

## Monitoring

During execution, auto-responses are displayed in the terminal with a 🤖 prefix:

```
[Agent output...]
🤖 Auto-Response: Start with the most foundational and critical component that other parts depend on. Follow the natural dependency order.
[Agent continues...]
```

## Best Practices

1. **Review `.you/decisions.log` after completion** to understand what decisions were made
2. **Use specific requirements in USER_INPUT.md** to guide better auto-responses
3. **Monitor terminal output** during execution to see real-time decisions
4. **Intervene manually if needed** by addressing agents directly in OpenCode

## Fallback

If OpenCode CLI is not available, the system gracefully falls back to manual instructions:

```
⚠️  Could not automatically launch OpenCode: exec: "opencode": executable file not found

📋 Manual steps:
1. Run: opencode
2. Tell @ceo: 'Read USER_INPUT.md and orchestrate the team to build this project'
```

---

**The result**: A truly autonomous AI software company that handles all decision-making automatically while providing full transparency and auditability.
