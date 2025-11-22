package repl

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/vintlang/vintlang/docs"
)

var (
	sidebarStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#aa6f5a")).
			Padding(1, 2)

	contentStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#aa6f5a")).
			Padding(1, 2)

	docsStatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#1a1a1a")).
			Padding(0, 1).
			MarginTop(1)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#aa6f5a")).
			Bold(true).
			Padding(0, 1)
)

// Key bindings for documentation navigation
type DocsKeyMap struct {
	Back       key.Binding
	ToggleHelp key.Binding
	Quit       key.Binding
	Up         key.Binding
	Down       key.Binding
}

func (k DocsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.ToggleHelp, k.Quit}
}

func (k DocsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Back, k.ToggleHelp},
		{k.Quit},
	}
}

var DocsKeys = DocsKeyMap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back to list"),
	),
	ToggleHelp: key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("f1", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "ctrl+q"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "navigate up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "navigate down"),
	),
}

type documentation struct {
	id           string
	sidebar      list.Model
	content      viewport.Model
	ready        bool
	filename     string
	fileContent  []byte
	docRenderer  *glamour.TermRenderer
	keys         DocsKeyMap
	showHelp     bool
	statusMsg    string
	statusTimer  *time.Timer
	windowWidth  int
	windowHeight int
	inDocView    bool
}

func newDocumentation() documentation {
	return documentation{
		keys:      DocsKeys,
		statusMsg: "Welcome to Vint Documentation!",
	}
}

func (d documentation) Init() tea.Cmd {
	return nil
}

func (d *documentation) setStatus(message string, duration time.Duration) {
	d.statusMsg = message
	if d.statusTimer != nil {
		d.statusTimer.Stop()
	}
	d.statusTimer = time.AfterFunc(duration, func() {
		d.statusMsg = "Ready"
	})
}

func (d documentation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		sidebarCmd tea.Cmd
		contentCmd tea.Cmd
	)

	d.sidebar, sidebarCmd = d.sidebar.Update(msg)
	d.content, contentCmd = d.content.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, d.keys.Back):
			if d.inDocView {
				d.inDocView = false
				d.setStatus("Back to documentation list", 2*time.Second)
			}
		case key.Matches(msg, d.keys.ToggleHelp):
			d.showHelp = !d.showHelp
			if d.showHelp {
				d.setStatus("Help displayed", 2*time.Second)
			} else {
				d.setStatus("Help hidden", 2*time.Second)
			}
		case key.Matches(msg, d.keys.Quit):
			return d, tea.Quit
		case msg.Type == tea.KeyEnter:
			if !d.inDocView {
				// Load selected documentation
				i, ok := d.sidebar.SelectedItem().(docs.Item)
				if ok {
					d.loadDocumentation(i.Filename())
					d.inDocView = true
					d.setStatus(fmt.Sprintf("Viewing: %s", i.Filename()), 2*time.Second)
				}
			}
		}

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			if zone.Get(d.id+"sidebar").InBounds(msg) && !d.inDocView {
				// Handle click in sidebar to select and open doc
				i, ok := d.sidebar.SelectedItem().(docs.Item)
				if ok {
					d.loadDocumentation(i.Filename())
					d.inDocView = true
					d.setStatus(fmt.Sprintf("Viewing: %s", i.Filename()), 2*time.Second)
				}
			}
		}

	case tea.WindowSizeMsg:
		return d.handleResize(msg)
	}

	return d, tea.Batch(sidebarCmd, contentCmd)
}

func (d *documentation) loadDocumentation(filename string) {
	d.filename = filename
	content, err := docs.Docs.ReadFile(d.filename)
	if err != nil {
		d.content.SetContent(fmt.Sprintf("Error loading documentation: %s\n\nFile: %s", err.Error(), d.filename))
		return
	}

	d.fileContent = content
	str, err := d.docRenderer.Render(string(d.fileContent))
	if err != nil {
		d.content.SetContent(fmt.Sprintf("Error rendering documentation: %s", err.Error()))
	} else {
		d.content.SetContent(str)
		d.content.GotoTop()
	}
}

func (d *documentation) handleResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	if !d.ready {
		d.initializeComponents(msg)
		d.ready = true
	} else {
		d.updateComponents(msg)
	}
	return d, nil
}

func (d *documentation) initializeComponents(msg tea.WindowSizeMsg) {
	// Initialize sidebar (documentation list)
	items := getDocItems()
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("#aa6f5a")).
		BorderForeground(lipgloss.Color("#aa6f5a"))
	
	d.sidebar = list.New(items, delegate, msg.Width/3, msg.Height-4)
	d.sidebar.Title = "Vint Documentation"
	d.sidebar.Styles.Title = titleStyle
	d.sidebar.SetShowStatusBar(false)
	d.sidebar.SetFilteringEnabled(true)

	// Initialize content viewport
	d.content = viewport.New(2*msg.Width/3-6, msg.Height-4)
	d.content.KeyMap = viewport.KeyMap{
		Up:            key.NewBinding(key.WithKeys("up", "k")),
		Down:          key.NewBinding(key.WithKeys("down", "j")),
		PageUp:        key.NewBinding(key.WithKeys("pgup")),
		PageDown:      key.NewBinding(key.WithKeys("pgdown")),
		HalfPageUp:    key.NewBinding(key.WithKeys("u")),
		HalfPageDown:  key.NewBinding(key.WithKeys("d")),
	}

	// Initialize markdown renderer
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(2*msg.Width/3-10),
	)
	if err == nil {
		d.docRenderer = renderer
	}

	d.windowWidth = msg.Width
	d.windowHeight = msg.Height

	// Set initial content
	d.content.SetContent("Select a documentation topic from the sidebar to get started.\n\nUse ↑/↓ to navigate, Enter to select, Esc to go back.")
}

func (d *documentation) updateComponents(msg tea.WindowSizeMsg) {
	d.sidebar.SetSize(msg.Width/3, msg.Height-4)
	d.content.Width = 2*msg.Width/3 - 6
	d.content.Height = msg.Height - 4

	// Update renderer with new width
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(2*msg.Width/3-10),
	)
	if err == nil {
		d.docRenderer = renderer
		// Re-render current content if we're in doc view
		if d.inDocView && len(d.fileContent) > 0 {
			str, err := d.docRenderer.Render(string(d.fileContent))
			if err == nil {
				d.content.SetContent(str)
			}
		}
	}

	d.windowWidth = msg.Width
	d.windowHeight = msg.Height
}

func (d documentation) View() string {
	if !d.ready {
		return "\n  Initializing Vint Documentation..."
	}

	// Sidebar panel
	sidebarPanel := zone.Mark(d.id+"sidebar", 
		sidebarStyle.
			Width(d.windowWidth/3-4).
			Height(d.windowHeight-6).
			Render(d.sidebar.View()),
	)

	// Content panel
	var contentPanel string
	if d.inDocView {
		contentPanel = zone.Mark(d.id+"content", 
			contentStyle.
				Width(2*d.windowWidth/3-8).
				Height(d.windowHeight-6).
				Render(d.content.View()),
		)
	} else {
		contentPanel = contentStyle.
			Width(2*d.windowWidth/3-8).
			Height(d.windowHeight-6).
			Render("Select a documentation topic from the sidebar to view its content.\n\nNavigation:\n• Use arrow keys or j/k to navigate the list\n• Press Enter or click to open a document\n• Press Esc to return to the list\n• Use F1 to toggle this help")
	}

	// Status bar
	statusBar := docsStatusStyle.Width(d.windowWidth).Render(d.statusMsg)

	// Help panel
	var helpPanel string
	if d.showHelp {
		helpPanel = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#aa6f5a")).
			Padding(1).
			Width(d.windowWidth).
			Render(
				lipgloss.JoinVertical(lipgloss.Left,
					titleStyle.Render("Vint Documentation - Help"),
					"",
					"Navigation:",
					"  ↑/k      - Navigate up",
					"  ↓/j      - Navigate down",
					"  Enter    - Open selected document",
					"  Esc      - Back to document list",
					"  F1       - Toggle this help",
					"  Ctrl+Q   - Quit",
					"",
					"Viewing:",
					"  PageUp/Down - Scroll page",
					"  u/d         - Half page scroll",
				),
			)
	}

	// Main layout - sidebar on left, content on right
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, sidebarPanel, contentPanel)
	
	if d.showHelp {
		return zone.Scan(lipgloss.JoinVertical(lipgloss.Left, mainContent, helpPanel, statusBar))
	}
	return zone.Scan(lipgloss.JoinVertical(lipgloss.Left, mainContent, statusBar))
}

// Helper function to get documentation items
func getDocItems() []list.Item {
	// This would typically come from your docs package
	// For now, using a placeholder - replace with actual doc items
	var items []list.Item
	docFiles := []string{
		"introduction.md",
		"getting-started.md", 
		"syntax.md",
		"types.md",
		"functions.md",
		"control-flow.md",
		"standard-library.md",
	}

	for _, file := range docFiles {
		// Convert filename to readable title
		title := strings.TrimSuffix(file, ".md")
		title = strings.ReplaceAll(title, "-", " ")
		title = strings.Title(title)
		
		items = append(items, docs.Item{
			Title: title,
			Desc:  fmt.Sprintf("Documentation for %s", title),
			File:  file,
		})
	}
	return items
}

// StartDocumentation starts the documentation viewer
func StartDocumentation() {
	zone.NewGlobal()
	
	d := newDocumentation()
	d.id = zone.NewPrefix()

	program := tea.NewProgram(d, 
		tea.WithMouseAllMotion(), 
		tea.WithAltScreen(),
	)
	
	if _, err := program.Run(); err != nil {
		log.Fatal("Error running documentation viewer:", err)
	}
}