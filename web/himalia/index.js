import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { configureStore } from '@reduxjs/toolkit'
import { applyMiddleware, compose, combineReducers } from 'redux'
import thunk from 'redux-thunk'

import App from '@app'
import reducers from '@stores/reducers'

import 'public-css/core.css'
import './globals.css'

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

const rootReducer = combineReducers({
    root: reducers
})

const store = configureStore({
    reducer: {
        rootReducer,
        composeEnhancers: composeEnhancers(
            applyMiddleware(thunk)
        )
    }
})

const app = (
    <Provider store={store}>
        <App />
    </Provider>
)

document.addEventListener('DOMContentLoaded', () => {
    ReactDOM.render(app, document.getElementById('application-context'))
})
