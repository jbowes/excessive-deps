import React from 'react'
import { Todo, User } from '@/types'
import { TodoItem } from './TodoItem'

type Props = { todos: Todo[], users: User[] }

export const TodoList: React.FC<Props> = ({ todos, users }) => {
  const userById = new Map(users.map(u => [u.id, u]))
  return (
    <section>
      <h2>Todos</h2>
      <div className="list">
        {todos.map(t => (
          <TodoItem key={t.id} todo={t} user={userById.get(t.userId)} />
        ))}
      </div>
    </section>
  )
}
