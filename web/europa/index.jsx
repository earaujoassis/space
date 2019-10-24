import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

import UserStore from './stores/users'

import Settings from './components/Settings.jsx'
import Applications from './components/Applications'
import Profile from './components/Profile.jsx'

const europa = (
    <Router>
        <Switch>
            <Settings>
                <Route path="/applications" exact component={Applications} />
                <Route path="/profile" exact component={Profile} />
            </Settings>
        </Switch>
    </Router>
)

UserStore.loadData()

ReactDOM.render(europa, document.getElementById('application-context'))
