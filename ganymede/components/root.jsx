import React from 'react';

import Login from './Login.jsx';

export default class Root extends React.Component {
    constructor() {
        super();
        this.state = { user: null };
        this._afterLogin = this._afterLogin.bind(this);
    }

    render() {
        let component;
        if (!this.state.user) {
            component = <Login afterLogin={this._afterLogin} />;
        } else {
            component = null;
        }
        return component;
    }

    _afterLogin(user) {
        this.setState({user: user});
    }
};
