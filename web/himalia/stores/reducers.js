import { combineReducers } from '@reduxjs/toolkit'

import internal from './resources/internal'
import workspace from './resources/workspace'
import user from './resources/user'
import requests from './resources/requests'
import sessions from './resources/sessions'
import emails from './resources/emails'
import settings from './resources/settings'
import clients from './resources/clients'
import services from './resources/services'

const reducer = combineReducers({
  internal,
  workspace,
  user,
  requests,
  sessions,
  emails,
  settings,
  clients,
  services,
})

export default reducer
