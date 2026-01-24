package main

import (
	"fmt"
	"os"
	"path/filepath"
	"you/internal/orchestrator"
)

const version = "0.2.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	// Get current working directory as project path
	projectPath, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to get current directory: %v\n", err)
		os.Exit(1)
	}

	switch command {
	case "--presets":
		handlePresets(projectPath)
	case "--orchestrate":
		handleOrchestrate(projectPath)
	case "--version", "-v":
		fmt.Printf("You orchestrator v%s\n", version)
	case "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func handlePresets(projectPath string) {
	fmt.Println("╔═══════════════════════════════════════════════╗")
	fmt.Println("║   You - Agentic Orchestrator (Beta v" + version + ")   ║")
	fmt.Println("╚═══════════════════════════════════════════════╝")
	fmt.Println()

	orch, err := orchestrator.New(projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to initialize orchestrator: %v\n", err)
		os.Exit(1)
	}

	if err := orch.SetupPresets(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleOrchestrate(projectPath string) {
	fmt.Println("╔═══════════════════════════════════════════════╗")
	fmt.Println("║   You - Agentic Orchestrator (Beta v" + version + ")   ║")
	fmt.Println("╚═══════════════════════════════════════════════╝")
	fmt.Println()

	// Check if presets have been run
	userInputPath := filepath.Join(projectPath, "USER_INPUT.md")
	if _, err := os.Stat(userInputPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: USER_INPUT.md not found.\n")
		fmt.Fprintf(os.Stderr, "Please run 'you.exe --presets' first to generate required files.\n")
		os.Exit(1)
	}

	orch, err := orchestrator.New(projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to initialize orchestrator: %v\n", err)
		os.Exit(1)
	}

	if err := orch.Orchestrate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("You - Agentic Orchestrator")
	fmt.Println("A fully autonomous AI software company that builds projects from a single prompt")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  you.exe --presets       Generate preset files (USER_INPUT.md and agent configs)")
	fmt.Println("  you.exe --orchestrate   Start autonomous workflow (launches OpenCode automatically)")
	fmt.Println("  you.exe --version       Show version information")
	fmt.Println("  you.exe --help          Show this help message")
	fmt.Println()
	fmt.Println("Workflow:")
	fmt.Println("  1. Run 'you.exe --presets' to create initial setup files")
	fmt.Println("  2. Edit USER_INPUT.md with your project requirements")
	fmt.Println("  3. Run 'you.exe --orchestrate' - sit back and let AI agents build it!")
	fmt.Println("     → Auto-launches OpenCode with CEO agent")
	fmt.Println("     → Auto-responds to all agent questions")
	fmt.Println("     → Logs all decisions to .you/decisions.log")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  🤖 Fully autonomous - no human intervention needed")
	fmt.Println("  📊 Decision audit trail in .you/decisions.log")
	fmt.Println("  🌐 Web research via webfetch tool")
	fmt.Println("  ⚡ 10 specialized AI agents working together")
	fmt.Println()
	fmt.Println("Documentation:")
	fmt.Println("  README.md              - Full documentation")
	fmt.Println("  AUTO-RESPONSE.md       - How auto-response works")
	fmt.Println("  QUICKSTART.md          - 5-minute getting started guide")
	fmt.Println("  ARCHITECTURE.md        - System architecture")
	fmt.Println()
	fmt.Println("Learn more:")
	fmt.Println("  GitHub: https://github.com/galpt/you")
	fmt.Println("  OpenCode: https://opencode.ai/docs/agents/")
}
