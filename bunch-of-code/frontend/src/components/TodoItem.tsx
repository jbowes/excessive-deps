import React from 'react'
import { Todo, User } from '@/types'

type Props = { todo: Todo, user?: User }

export const TodoItem: React.FC<Props> = ({ todo, user }) => {
  return (
    <div className="card">
      <div><strong>{todo.title}</strong></div>
      <div className="small">Owner: {user ? user.name : 'unknown'}</div>
    </div>
  )
}
