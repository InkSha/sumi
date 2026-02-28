import sys
from pathlib import Path

root = str(Path(__file__).parent.parent.parent)

sys.path.insert(0, root)

from lib.todos import TodoItem, TodoItemStatus, print_todos

# metadata start
Name        = 'Demo Plan'
Author      = 'Me'
Description = 'This is a demo plan'
CreateAt    = '2026-02-27 09:00:00'
ModityAt    = '2026-02-27 09:00:00'
# metadata end

# data start
Todos = [
    TodoItem('Todo 1', 'This is a todo item', TodoItemStatus.Draft),
    TodoItem('Todo 2', 'This is another todo item', TodoItemStatus.Todo),
]
# data end

if __name__ == '__main__':
    print_todos(Todos)
