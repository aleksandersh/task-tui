package tasks

import (
	"context"
	"strings"
	"time"

	"github.com/aleksandersh/task-tui/internal/app/ui"
	"github.com/aleksandersh/task-tui/internal/domain"
	"github.com/aleksandersh/task-tui/internal/task"
)

type controllerState struct {
	filterActive bool
	filter       string
	currentList  *tasksViewList
}

type controller struct {
	ctx      context.Context
	task     *task.Task
	ui       ui.Controller
	view     *view
	fullList *tasksViewList
	state    controllerState
	filter   chan string
}

func newController(ctx context.Context, task *task.Task, ui ui.Controller, view *view, taskfile *domain.Taskfile) *controller {
	fullList := createTasksViewList(taskfile)
	state := controllerState{filterActive: false, filter: "", currentList: fullList}
	filter := make(chan string, 20)
	return &controller{ctx: ctx, task: task, ui: ui, view: view, fullList: fullList, state: state, filter: filter}
}

func (c *controller) start() {
	c.populateTasksView(c.fullList)
	go c.startDebouncedFilteringJob()
}

func (c *controller) onResetFilter() {
	if c.state.filterActive {
		return
	}
	if len(c.state.filter) == 0 {
		if c.view.tasks.GetCurrentItem() != 0 {
			c.view.tasks.SetCurrentItem(0)
		}
		return
	}

	c.resetFilter()
}

func (c *controller) onActivateFilter() {
	if c.state.filterActive {
		return
	}
	c.focusFilter()
}

func (c *controller) onCancelFilter() {
	if !c.state.filterActive {
		return
	}

	c.resetFilter()
	c.focusTasks()
}

func (c *controller) onFinishFilter() {
	if !c.state.filterActive {
		return
	}
	c.focusTasks()
}

func (c *controller) onFilterChanged(filter string) {
	if c.state.filter == filter {
		return
	}
	c.state.filter = filter
	c.filter <- filter
}

func (c *controller) onClickSummary() {
	if item := c.getCurrentTaskViewItem(); item != nil {
		c.ui.ShowTaskSummary(&item.Task)
	}
}

func (c *controller) onClickHelp() {
	c.ui.ShowHelp()
}

type filterResult struct {
	filter   string
	duration time.Duration
}

func (c *controller) startDebouncedFilteringJob() {
	ticker := make(chan bool, 20)
	onFinish := make(chan filterResult, 20)
	filter := c.state.filter
	appliedFilter := filter
	applying := false
	nextTime := time.Now()
	for {
		select {
		case <-c.ctx.Done():
			return
		case result := <-onFinish:
			nextTime = time.Now().Add(result.duration)
			applying = false
			appliedFilter = result.filter
			if filter != appliedFilter {
				delayFilter(result.duration, ticker)
			}
		case <-ticker:
			if filter == appliedFilter {
				return
			}
			if !applying && !nextTime.After(time.Now()) {
				applying = true
				c.applyFilter(filter, onFinish)
			}
		case f := <-c.filter:
			if filter == f {
				continue
			}
			filter = f
			if applying {
				continue
			}
			interval := nextTime.Sub(time.Now())
			if interval <= 100 {
				applying = true
				c.applyFilter(filter, onFinish)
			} else {
				delayFilter(interval, ticker)
			}
		}
	}
}

func delayFilter(interval time.Duration, ticker chan bool) {
	go func() {
		time.Sleep(interval)
		ticker <- true
	}()
}

func (c *controller) applyFilter(filter string, onFinish chan filterResult) {
	start := time.Now()
	tasks := filtered(c.fullList, filter)
	c.ui.PostDraw(func() {
		c.populateTasksView(tasks)
		go func() {
			onFinish <- filterResult{filter: filter, duration: time.Since(start)}
		}()
	})
}

func filtered(tasks *tasksViewList, text string) *tasksViewList {
	if len(text) == 0 {
		return tasks
	}
	filteredItems := make([]tasksViewItem, 0, len(tasks.items))
	for _, item := range tasks.items {
		if strings.Contains(item.Text, text) {
			filteredItems = append(filteredItems, item)
		}
	}
	return &tasksViewList{items: filteredItems}
}

func (c *controller) populateTasksView(tasks *tasksViewList) {
	focusedAbsoluteIndex := c.getAbsoluteTaskViewIndex()
	focusedIndex := 0
	c.state.currentList = tasks
	c.view.tasks.Clear()
	for index, item := range tasks.items {
		c.addTaskView(item)
		if item.Index == focusedAbsoluteIndex {
			focusedIndex = index
		}
	}
	c.view.tasks.SetCurrentItem(focusedIndex)
}

func (c *controller) addTaskView(item tasksViewItem) {
	c.view.tasks.AddItem(item.Task.Name, item.Task.Description, 0, func() {
		c.ui.Close()
		if len(item.Task.Aliases) > 0 {
			c.task.ExecuteTask(c.ctx, item.Task.Aliases[0])
		} else {
			c.task.ExecuteTask(c.ctx, item.Task.Name)
		}
	})
}

func (c *controller) focusFilter() {
	c.state.filterActive = true
	c.view.filter.SetDisabled(false)
	c.ui.Focus(c.view.filter)
}

func (c *controller) focusTasks() {
	c.state.filterActive = false
	c.view.filter.SetDisabled(true)
	c.ui.Focus(c.view.tasks)
}

func (c *controller) resetFilter() {
	c.state.filter = ""
	c.view.filter.SetText("", true)
	c.filter <- ""
}

func (c *controller) getAbsoluteTaskViewIndex() int {
	if item := c.getCurrentTaskViewItem(); item != nil {
		return item.Index
	}
	return -1
}

func (c *controller) getCurrentTaskViewItem() *tasksViewItem {
	currentItem := c.view.tasks.GetCurrentItem()
	if currentItem >= 0 && len(c.state.currentList.items) > 0 {
		return &c.state.currentList.items[currentItem]
	}
	return nil
}
