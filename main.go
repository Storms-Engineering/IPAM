package main

// An example demonstrating an application with multiple views.
//
// Note that this example was produced before the Bubbles progress component
// was available (github.com/charmbracelet/bubbles/progress) and thus, we're
// implementing a progress bar from scratch here.

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

const (
	ip_addr = iota
	ip_desc
	ip_type
)

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// General stuff for styling the view
var (
	keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#005ce6"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(green).Align(lipgloss.Left)
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle()
)

const (
	green    = lipgloss.Color("#009900")
	darkGray = lipgloss.Color("#767676")
)

// Global Variables
var (
	inputStyle    = lipgloss.NewStyle().Foreground(green).Background(lipgloss.Color("#a09595ff")).Align(lipgloss.Left)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type (
	errMsg error
)

// List types
type item string
type itemDelegate struct{}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Border(lipgloss.ASCIIBorder(), true).Background(lipgloss.Color(green))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

func (i item) FilterValue() string { return "" }

const listHeight = 14

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func main() {

	//Ip type list
	items := []list.Item{
		item("Normal"),
		item("Reserved"),
		item("DHCP Range"),
		item("Gateway"),
	}

	const defaultWidth = 100

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Type:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	//Table Init
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "São Paulo", "Brazil", "22,429,800"},
		{"6", "Mexico City", "Mexico", "22,085,140"},
		{"7", "Cairo", "Egypt", "21,750,020"},
		{"8", "Beijing", "China", "21,333,332"},
		{"9", "Mumbai", "India", "20,961,472"},
		{"10", "Osaka", "Japan", "19,059,856"},
		{"11", "Chongqing", "China", "16,874,740"},
		{"12", "Karachi", "Pakistan", "16,839,950"},
		{"13", "Istanbul", "Turkey", "15,636,243"},
		{"14", "Kinshasa", "DR Congo", "15,628,085"},
		{"15", "Lagos", "Nigeria", "15,387,639"},
		{"16", "Buenos Aires", "Argentina", "15,369,919"},
		{"17", "Kolkata", "India", "15,133,888"},
		{"18", "Manila", "Philippines", "14,406,059"},
		{"19", "Tianjin", "China", "14,011,828"},
		{"20", "Guangzhou", "China", "13,964,637"},
		{"21", "Rio De Janeiro", "Brazil", "13,634,274"},
		{"22", "Lahore", "Pakistan", "13,541,764"},
		{"23", "Bangalore", "India", "13,193,035"},
		{"24", "Shenzhen", "China", "12,831,330"},
		{"25", "Moscow", "Russia", "12,640,818"},
		{"26", "Chennai", "India", "11,503,293"},
		{"27", "Bogota", "Colombia", "11,344,312"},
		{"28", "Paris", "France", "11,142,303"},
		{"29", "Jakarta", "Indonesia", "11,074,811"},
		{"30", "Lima", "Peru", "11,044,607"},
		{"31", "Bangkok", "Thailand", "10,899,698"},
		{"32", "Hyderabad", "India", "10,534,418"},
		{"33", "Seoul", "South Korea", "9,975,709"},
		{"34", "Nagoya", "Japan", "9,571,596"},
		{"35", "London", "United Kingdom", "9,540,576"},
		{"36", "Chengdu", "China", "9,478,521"},
		{"37", "Nanjing", "China", "9,429,381"},
		{"38", "Tehran", "Iran", "9,381,546"},
		{"39", "Ho Chi Minh City", "Vietnam", "9,077,158"},
		{"40", "Luanda", "Angola", "8,952,496"},
		{"41", "Wuhan", "China", "8,591,611"},
		{"42", "Xi An Shaanxi", "China", "8,537,646"},
		{"43", "Ahmedabad", "India", "8,450,228"},
		{"44", "Kuala Lumpur", "Malaysia", "8,419,566"},
		{"45", "New York City", "United States", "8,177,020"},
		{"46", "Hangzhou", "China", "8,044,878"},
		{"47", "Surat", "India", "7,784,276"},
		{"48", "Suzhou", "China", "7,764,499"},
		{"49", "Hong Kong", "Hong Kong", "7,643,256"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"51", "Shenyang", "China", "7,527,975"},
		{"52", "Baghdad", "Iraq", "7,511,920"},
		{"53", "Dongguan", "China", "7,511,851"},
		{"54", "Foshan", "China", "7,497,263"},
		{"55", "Dar Es Salaam", "Tanzania", "7,404,689"},
		{"56", "Pune", "India", "6,987,077"},
		{"57", "Santiago", "Chile", "6,856,939"},
		{"58", "Madrid", "Spain", "6,713,557"},
		{"59", "Haerbin", "China", "6,665,951"},
		{"60", "Toronto", "Canada", "6,312,974"},
		{"61", "Belo Horizonte", "Brazil", "6,194,292"},
		{"62", "Khartoum", "Sudan", "6,160,327"},
		{"63", "Johannesburg", "South Africa", "6,065,354"},
		{"64", "Singapore", "Singapore", "6,039,577"},
		{"65", "Dalian", "China", "5,930,140"},
		{"66", "Qingdao", "China", "5,865,232"},
		{"67", "Zhengzhou", "China", "5,690,312"},
		{"68", "Ji Nan Shandong", "China", "5,663,015"},
		{"69", "Barcelona", "Spain", "5,658,472"},
		{"70", "Saint Petersburg", "Russia", "5,535,556"},
		{"71", "Abidjan", "Ivory Coast", "5,515,790"},
		{"72", "Yangon", "Myanmar", "5,514,454"},
		{"73", "Fukuoka", "Japan", "5,502,591"},
		{"74", "Alexandria", "Egypt", "5,483,605"},
		{"75", "Guadalajara", "Mexico", "5,339,583"},
		{"76", "Ankara", "Turkey", "5,309,690"},
		{"77", "Chittagong", "Bangladesh", "5,252,842"},
		{"78", "Addis Ababa", "Ethiopia", "5,227,794"},
		{"79", "Melbourne", "Australia", "5,150,766"},
		{"80", "Nairobi", "Kenya", "5,118,844"},
		{"81", "Hanoi", "Vietnam", "5,067,352"},
		{"82", "Sydney", "Australia", "5,056,571"},
		{"83", "Monterrey", "Mexico", "5,036,535"},
		{"84", "Changsha", "China", "4,809,887"},
		{"85", "Brasilia", "Brazil", "4,803,877"},
		{"86", "Cape Town", "South Africa", "4,800,954"},
		{"87", "Jiddah", "Saudi Arabia", "4,780,740"},
		{"88", "Urumqi", "China", "4,710,203"},
		{"89", "Kunming", "China", "4,657,381"},
		{"90", "Changchun", "China", "4,616,002"},
		{"91", "Hefei", "China", "4,496,456"},
		{"92", "Shantou", "China", "4,490,411"},
		{"93", "Xinbei", "Taiwan", "4,470,672"},
		{"94", "Kabul", "Afghanistan", "4,457,882"},
		{"95", "Ningbo", "China", "4,405,292"},
		{"96", "Tel Aviv", "Israel", "4,343,584"},
		{"97", "Yaounde", "Cameroon", "4,336,670"},
		{"98", "Rome", "Italy", "4,297,877"},
		{"99", "Shijiazhuang", "China", "4,285,135"},
		{"100", "Montreal", "Canada", "4,276,526"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	//Input init
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[ip_addr] = textinput.New()
	inputs[ip_addr].Placeholder = "IP"
	inputs[ip_addr].Focus()
	inputs[ip_addr].CharLimit = 20
	inputs[ip_addr].Width = 30
	inputs[ip_addr].Prompt = "10.25.66." //Make this dynamic for which subnet we are in
	//TODO add these later inputs[ccn].Validate = ccnValidator

	inputs[ip_desc] = textinput.New()
	inputs[ip_desc].Placeholder = "FS Switch "
	inputs[ip_desc].CharLimit = 50
	inputs[ip_desc].Width = 50
	inputs[ip_desc].Prompt = ""
	//inputs[exp].Validate = expValidator

	//Possibly do this with suggestions instead of a list picker?
	inputs[ip_type] = textinput.New()
	inputs[ip_type].CharLimit = 10
	inputs[ip_type].Width = 10
	inputs[ip_type].Prompt = ""

	//inputs[cvv].Validate = cvvValidator
	//Init Model 		Choice,MainMenu, State, Loaded, Quitting, table,Selected,inputs,focused, list, ChoiceSubnetModel, err,
	initialModel := model{0, true, MainMenu, false, false, t, false, inputs, 0, l, "Normal", nil}
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type ContextState int

// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on.
// Source: https://gobyexample.com/enums
const (
	MainMenu ContextState = iota
	Search
	SubnetTable
	EditValue
	PickSubnetType
)

type model struct {
	Choice           int
	MainMenu         bool
	State            ContextState
	Loaded           bool
	Quitting         bool
	table            table.Model
	Selected         bool
	inputs           []textinput.Model
	focused          int
	list             list.Model
	ChoiceSubnetType string
	err              error
}

func (m model) Init() tea.Cmd {
	//We can only do this here which is annoying :(
	m.inputs[ip_type].SetValue(m.ChoiceSubnetType)
	return nil
}

// Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
		if k == "esc" {
			//Go back one Menu -- Theoretically
			m.State--
			if m.State <= 0 {
				m.State = MainMenu
				m.MainMenu = true
			}
		}
	}
	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.MainMenu {
		return updateMainMenu(msg, m)
	}
	return updateTasks(msg, m)
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.Quitting {
		return mainStyle.Render("\n  Follow the white rabbit\n\n")
	}
	if m.MainMenu {
		s = mainMenuView(m)
	} else {
		s = taskView(m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

// Sub-update functions

// Update loop for the main menu
func updateMainMenu(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.MainMenu = false
			m.State = ContextState(m.Choice + 1)
			return m, nil
		}

	}

	return m, nil
}

// Update loop for the second view after a choice has been made
func updateTasks(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//SubnetTable View
		if m.Choice == 1 && m.State == SubnetTable {
			m.table, cmd = m.table.Update(msg)
			switch msg.String() {
			case "p":
				if m.table.Focused() {
					m.table.Blur()
				} else {
					m.table.Focus()
				}
			case "enter":
				m.State = EditValue
				m.Selected = true
				return m, nil
			}
		}
		//If we were in SubnetTable View and now editing a value
		if m.Choice == 1 && m.State == EditValue {
			return updateInputs(msg, m)
		}
		//Subnet Type List View
		if m.Choice == 1 && m.State == PickSubnetType {
			m.list, cmd = m.list.Update(msg)
			if msg.String() == "enter" {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.ChoiceSubnetType = string(i)
					m.State = EditValue
					return m, nil
				} else {
					m.err = errors.New("Selected list failed fantastically!")
					return m, tea.Quit
				}
			}

		}
	}
	//Catch all return
	return m, cmd
}

func updateInputs(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				if m.State == EditValue {
					m.State = PickSubnetType
				} else {
					m.State = EditValue
				}
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
	//TODO add a special case before we loop around once we reach the end to change the style of the continue "button"
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

// Sub-views

func mainMenuView(m model) string {
	c := m.Choice

	tpl := titleStyle.Render("What to do today?") + "\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q: quit") + dotStyle +
		subtleStyle.Render("esc: go back")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Search", c == 0),
		checkbox("View Table (Subnets future)", c == 1),
		checkbox("Read something", c == 2),
		checkbox("See friends", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}

func taskView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("Carrot planting?\n\nCool, we'll need %s and %s...", keywordStyle.Render("libgarden"), keywordStyle.Render("vegeutils"))
	//SubnetTable View
	case 1:
		if m.State == SubnetTable {
			msg += tableStyle.Render(m.table.View())
			if m.Selected == true {
				msg += fmt.Sprintf("\n\nLets go to %s", m.table.SelectedRow()[1])
			} else {
				msg += "Not Selected"
			}
			return msg
		}
		if m.State == EditValue {
			ipText := "Ip Address:"
			descText := "Description:"
			typeText := "Type:"
			m.inputs[ip_type].SetValue(m.ChoiceSubnetType)
			//Input fields
			msg += fmt.Sprintf(`Edit IP Address
%s%s
%s%s
%s%s
%s`,
				inputStyle.Width(len(ipText)).Render(ipText),
				m.inputs[ip_addr].View(),
				inputStyle.Width(len(descText)).Render(descText),
				m.inputs[ip_desc].View(),
				inputStyle.Width(len(typeText)).Render(typeText),
				m.inputs[ip_type].View(),
				continueStyle.Render("Continue ->"),
			) + "\n"
		}
		if m.State == PickSubnetType {
			msg += "\n" + m.list.View()
		}

	case 2:
		msg = fmt.Sprintf("Reading time?\n\nOkay, cool, then we’ll need a library. Yes, an %s.", keywordStyle.Render("actual library"))
	default:
		msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", keywordStyle.Render("social-skills"), keywordStyle.Render("conversationutils"))
	}
	return msg
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
