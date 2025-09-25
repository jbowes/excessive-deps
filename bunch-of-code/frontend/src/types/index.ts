export type User = {
  id: string
  email: string
  name: string
  createdAt: string
}

export type Todo = {
  id: string
  userId: string
  title: string
  completed: boolean
  createdAt: string
}
