module {{ cookiecutter.app_path }}

go {{ cookiecutter.golang_version }}

require (
	github.com/charmbracelet/bubbles v0.16.1
	github.com/charmbracelet/bubbletea v0.24.2
	github.com/charmbracelet/lipgloss v0.8.0
)
