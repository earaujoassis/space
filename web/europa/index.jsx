import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, IndexRoute, Link, browserHistory } from 'react-router'

import UserStore from './stores/users'

import Settings from './components/Settings.jsx'

import Applications from './components/Applications.jsx'
import Profile from './components/Profile.jsx'
import AccountLog from './components/AccountLog.jsx'

const europa = (
    <Router history={browserHistory}>
        <Route path="/" component={Settings}>
            <IndexRoute component={Applications} name="Applications" />
            <Route path="profile" component={Profile} name="Profile" />
            <Route path="log" component={AccountLog} name="Account log" />
        </Route>
    </Router>
)

UserStore.loadData()

ReactDOM.render(europa, document.getElementById('application-context'))
