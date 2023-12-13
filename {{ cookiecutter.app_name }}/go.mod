module {{ cookiecutter.app_path }}

go {{ cookiecutter.golang_version }}

require (
	github.com/charmbracelet/bubbles latest
	github.com/charmbracelet/bubbletea latest
	github.com/charmbracelet/lipgloss latest
  github.com/charmbracelet/log latest
)
