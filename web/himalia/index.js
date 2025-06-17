import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { configureStore } from '@reduxjs/toolkit'
import thunk from 'redux-thunk'

import App from '@app'
import reducers from '@stores/reducers'

import '@core/styles/core.scss'
import './globals.css'

const store = configureStore({
    reducer: {
        root: reducers,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(thunk),
    devTools: process.env.NODE_ENV !== 'production'
})

const app = (
    <Provider store={store}>
        <App />
    </Provider>
)

document.addEventListener('DOMContentLoaded', () => {
    ReactDOM.render(app, document.getElementById('application-context'))
})
