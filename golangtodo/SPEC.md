# Todo app

This contains the spec for the TODO app. Not really sure what should
be in a spec, but I'll try my best to make it look good?

## Initial spec
For a todo app that _i_ would like, it has to have the following
features:
- Have named lists.
- Add task.
- remove task.
- edit task.
- mark a task as complete.
- archive task.
- Save list.
- View lists.
- Search list for keywords.
- Add tags to entries.
- Show all entries with a given tag.

### My user stories
Stories that I would have when making this.

#### Shopping list
I am going to the store. I would like to make a list of items I need
to buy, and then mark them off in the list when I pick them up.

#### Longer tasks
At work, I have to complete some long task that has 3 steps. I would
like to make a separate list for these 3 tasks and then write the
tasks down as entries in the list.  When I complete the tasks I mark
them as done.

#### Save work
I completed a long list of tasks, but I would like to keep it around
so I have a record of what I did and I can refer back to it when I'm
curious.

#### Evaluate work
I have a lot of tasks to work on. I would like to have an overview of
what the state of my work is. So I would like to view all lists side
by side of work to do.

#### Edit planned work
A task that I planned to do no longer needs to be done. So I would
like to remove it from my list. I also Want to edit the title of a
task because of a typo. I would like to be able to edit each
individual task.

### Reqs after reading stories

- We need to have a collection of tasks be represented as a list of tasks.
- A task is a record with a field of text for the summary of the task along with
  task status.
- These lists of tasks must be able to be persisted, along with status.
- A task can be edited, deleted, or appended to a list.
- We must be able to make new lists.
- We must be able to delete lists.
- We must be able to rename lists.
- Must be able to get all tasks in a list by name.
- Each task can belong to ONLY ONE list.

Operations on list:
- Create empty list.
- Read list contents.
- Update list name.
- Delete list.
- Archive list.

Operations relating to tasks:
- Add task to a list.
- Edit task summary.
- Edit task status.
- Delete task.
- Read task.

### Interface

List methods:
- ListLists()
- CreateList(listName string)
- UpdateList(listName, newName )
- GetList(listName)
- DeleteList(listName)
- ArchiveList(listName)
- AddTaskToList(listName, taskSpec)

Task methods:
- GetTask(listName, taskId) -- Maybe just taskId?
- EditTask(listName, taskId, newSpecs)
- DeleteTask(listName, taskId)


### Models
ListModel
```
	listName: string
	listId: int
	tasks: repeated string
```

TaskModel
```
	taskLabel: string
	taskId: string
	taskStatus: STATUS
	(maybe) listId?
```

Status Enum
```
	ACTIVE
	DONE
```
