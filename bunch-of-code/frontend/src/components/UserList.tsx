import React from 'react'
import { User } from '@/types'

type Props = { users: User[] }

export const UserList: React.FC<Props> = ({ users }) => {
  return (
    <section>
      <h2>Users</h2>
      <div className="list">
        {users.map(u => (
          <div className="card" key={u.id}>
            <div><strong>{u.name}</strong></div>
            <div className="small">{u.email}</div>
          </div>
        ))}
      </div>
    </section>
  )
}
