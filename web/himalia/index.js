import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'

import App from '@app'

import store from '@stores'
import '@core/styles/core.scss'
import './globals.css'

const app = (
  <Provider store={store}>
    <App />
  </Provider>
)

document.addEventListener('DOMContentLoaded', () => {
  ReactDOM.render(app, document.getElementById('application-context'))
})
