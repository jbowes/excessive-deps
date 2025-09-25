import React, { useEffect, useState } from 'react'
import { Header } from '@/components/Header'
import { UserList } from '@/components/UserList'
import { TodoList } from '@/components/TodoList'
import { createTodo, createUser, getHealth, listTodos, listUsers } from '@/services/api'
import { Todo, User } from '@/types'

export default function App() {
  const [health, setHealth] = useState<string>('unknown')
  const [users, setUsers] = useState<User[]>([])
  const [todos, setTodos] = useState<Todo[]>([])
  const [newUser, setNewUser] = useState({ name: '', email: '' })
  const [newTodo, setNewTodo] = useState({ title: '', userId: '' })

  useEffect(() => {
    getHealth().then((h) => setHealth(h.status)).catch(() => setHealth('down'))
    listUsers().then(setUsers)
    listTodos().then(setTodos)
  }, [])

  const addUser = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newUser.name || !newUser.email) return
    const created = await createUser({ name: newUser.name, email: newUser.email })
    setUsers((u) => [...u, created])
    setNewUser({ name: '', email: '' })
  }

  const addTodo = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newTodo.title || !newTodo.userId) return
    const created = await createTodo({ title: newTodo.title, userId: newTodo.userId })
    setTodos((t) => [...t, created])
    setNewTodo({ title: '', userId: '' })
  }

  return (
    <div className="container">
      <Header status={health} />

      <section>
        <h2>Create User</h2>
        <form onSubmit={addUser} className="row">
          <input
            placeholder="Name"
            value={newUser.name}
            onChange={(e) => setNewUser({ ...newUser, name: e.target.value })}
          />
          <input
            placeholder="Email"
            value={newUser.email}
            onChange={(e) => setNewUser({ ...newUser, email: e.target.value })}
          />
          <button type="submit">Add</button>
        </form>
      </section>

      <UserList users={users} />

      <section>
        <h2>Create Todo</h2>
        <form onSubmit={addTodo} className="row">
          <input
            placeholder="Title"
            value={newTodo.title}
            onChange={(e) => setNewTodo({ ...newTodo, title: e.target.value })}
          />
          <select
            value={newTodo.userId}
            onChange={(e) => setNewTodo({ ...newTodo, userId: e.target.value })}
          >
            <option value="">Select user</option>
            {users.map((u) => (
              <option key={u.id} value={u.id}>{u.name}</option>
            ))}
          </select>
          <button type="submit">Add</button>
        </form>
      </section>

      <TodoList todos={todos} users={users} />
    </div>
  )
}
