package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Core UI Styles
	TitleStyle   = lipgloss.NewStyle().Margin(1, 0).Foreground(lipgloss.Color("#aa6f5a")).Bold(true)
	VersionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9671")).Italic(true)
	AuthorStyle  = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#ff9671"))
	HelpStyle    = lipgloss.NewStyle().Italic(true).Faint(true).Foreground(lipgloss.Color("#ffe6d6"))
	PromptStyle  = ""

	// Error System Styles
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	WarningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Italic(true)
	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Faint(true)
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Bold(true)

	// Error Context Styles
	ErrorHeaderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true).Underline(true)
	ErrorMessageStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	ErrorLocationStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)
	ErrorPointerStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	ErrorCodeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Faint(true)
	SuggestionStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Italic(true)

	// Code Display Styles
	CodeStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Background(lipgloss.Color("235"))
	CodeLineNumberStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Faint(true)
	HighlightStyle      = lipgloss.NewStyle().Background(lipgloss.Color("52")).Foreground(lipgloss.Color("255"))

	// REPL Styles
	ReplStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Italic(true)
	ReplPromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	ReplOutputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	ReplErrorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Italic(true)

	// Debug and Development Styles
	DebugStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("213")).Faint(true)
	VerboseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Faint(true)
	TraceStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)

	// File and Path Styles
	FilenameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Underline(true)
	PathStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	// Syntax Highlighting Styles
	KeywordStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).Bold(true)
	IdentifierStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	StringStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("76"))
	NumberStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	CommentStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
	OperatorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("213"))

	// Semantic Styles
	TypeStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	FunctionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true)
	VariableStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	ConstantStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("213")).Bold(true)

	// Border and Container Styles
	BorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("39"))
	PanelStyle  = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("244")).Padding(1)
	BoxStyle    = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("39")).Padding(1, 2)

	// Status and Progress Styles
	LoadingStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Faint(true)
	CompletedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Bold(true)
	ProgressStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))

	// Documentation Styles
	DocTitleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true).Underline(true)
	DocSectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true)
	DocTextStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	DocExampleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Background(lipgloss.Color("235")).Padding(0, 1)
)
