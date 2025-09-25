import React from 'react'

type Props = { status: string }

export const Header: React.FC<Props> = ({ status }) => {
  return (
    <header className="card">
      <h1>Sample App</h1>
      <div className="small">Backend status: {status}</div>
    </header>
  )
}
