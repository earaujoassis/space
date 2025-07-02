import React from 'react'
import { Link } from 'react-router-dom'

const submenu = ({ activeAction, editingClient, editingScopes }) => {
  var clientEditingLink

  if (editingClient) {
    clientEditingLink = (
      <li
        className={activeAction === 'edit-client' ? 'submenu__list-active' : ''}
      >
        <span>Edit client</span>
      </li>
    )
  } else {
    clientEditingLink = null
  }

  if (editingScopes) {
    clientEditingLink = (
      <li
        className={activeAction === 'edit-scopes' ? 'submenu__list-active' : ''}
      >
        <span>Select client scopes</span>
      </li>
    )
  } else {
    clientEditingLink = null
  }

  return (
    <>
      <ul className="submenu__list">
        <li
          className={
            activeAction === 'all-clients' ? 'submenu__list-active' : ''
          }
        >
          <Link to="/clients">All clients</Link>
        </li>
        <li
          className={
            activeAction === 'new-client' ? 'submenu__list-active' : ''
          }
        >
          <Link to="/clients/new">New client</Link>
        </li>
        {clientEditingLink}
      </ul>
    </>
  )
}

export default submenu
