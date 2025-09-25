import { Todo, User } from '@/types'

async function json<T>(res: Response): Promise<T> {
  if (!res.ok) throw new Error(`HTTP ${res.status}`)
  return res.json() as Promise<T>
}

const headers = { 'Content-Type': 'application/json', 'X-API-Key': 'dev' }

export function getHealth(): Promise<{ status: string }> {
  return fetch('/health').then(json)
}

export function listUsers(): Promise<User[]> {
  return fetch('/api/users', { headers }).then(json)
}

export function createUser(u: Pick<User, 'name' | 'email'>): Promise<User> {
  return fetch('/api/users', { method: 'POST', headers, body: JSON.stringify(u) }).then(json)
}

export function listTodos(): Promise<Todo[]> {
  return fetch('/api/todos', { headers }).then(json)
}

export function createTodo(t: Pick<Todo, 'title' | 'userId'>): Promise<Todo> {
  return fetch('/api/todos', { method: 'POST', headers, body: JSON.stringify(t) }).then(json)
}
