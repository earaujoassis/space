import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, IndexRoute, Link, browserHistory } from 'react-router'

import UserStore from './stores/users'

import Settings from './components/Settings.jsx'

import Applications from './components/Applications.jsx'
import Profile from './components/Profile.jsx'

const europa = (
    <Router history={browserHistory}>
        <Route path="/" component={Settings}>
            <IndexRoute component={Applications} name="Applications" />
            <Route path="profile" component={Profile} name="Profile" />
        </Route>
    </Router>
)

UserStore.loadData()

ReactDOM.render(europa, document.getElementById('application-context'))
