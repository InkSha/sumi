#! /usr/bin/env python3
# ==============================================================================
#
# Demo Plan
# This is a demo plan
#
# Author is Me
#
# CreateAt 2026-02-27 09:00:00
# ModifyAt 2026-02-27 09:00:00
#
# ==============================================================================
#

import os
from datetime import datetime
from enum import Enum
from typing import List

class TextAssets:
    TodoItemStatus = (
      'Draft',
      'Todo',
      'Doing',
      'DONE',
      'Pause'
    )

class TodoItemStatus(Enum):
    Draft   = 0
    Todo    = 1
    Doing   = 2
    DONE    = 3
    Pause   = 4

    def getSymbol(self):
        match self.value:
            # case TodoItemStatus.Draft.value:
            #     return '#'
            case TodoItemStatus.Todo.value:
                return '='
            case TodoItemStatus.Doing.value:
                return '+'
            case TodoItemStatus.DONE.value:
                return '-'
            case TodoItemStatus.Pause.value:
                return '|'

        return '#'

class TodoItem:
    title     = ''
    content   = ''
    createAt  = datetime.now()
    status    = TodoItemStatus.Draft

    def __init__(self, title: str, content: str, status: TodoItemStatus = TodoItemStatus.Draft):
        self.__change__(title,content,status)

    def __change__(self,title: str, content: str, status: TodoItemStatus = TodoItemStatus.Draft):
      self.title = title
      self.content = content
      self.status = status
      self.modityAt = datetime.now()

    def get_print_content(self, width: int) -> str:
        symbol = self.status.getSymbol()

        contents = []

        contents.extend(self.__format_content__(width, self.title, f'{symbol} '))
        contents.append('\n')
        contents.extend(self.__format_content__(width, self.content, ' ' * (len(symbol) + 1)))

        return '\n'.join(contents)

    def __format_content__(self, width: int, content: str, prefix: str = '') -> List[str]:
        result: List[str] = []
        currentPtr = 0
        currents = content.split('\n')
        current = currents[currentPtr]
        csl = len(currents)

        while currentPtr < csl:
            pl = len(prefix)

            if len(current) <= 0:
                if (currentPtr + 1) < csl:
                    currentPtr += 1
                    current = currents[currentPtr]
                else:
                    break

            if len(result) == 0:
                line = prefix + current
            else:
                line = (' ' * pl) + current

            line = line[:width]

            current = current[width:]

            result.append(line)

        return result

def print_todos(todos: List[TodoItem]):
    width = int(os.get_terminal_size().columns / 2)

    maxLen = 0
    contents: List[str] = []

    for todo in todos:
        content = todo.get_print_content(width)
        maxLen = max(maxLen, max([len(line) for line in content.split('\n')]))
        contents.append(content)

    line = '-' * maxLen

    for content in contents:
      print(line)
      print(content)
    print(line)
