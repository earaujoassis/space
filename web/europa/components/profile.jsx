import React from 'react'

import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'
import { Entry } from '../../core/components/form.jsx'

import UserStore from '../stores/users'
import UsersActions from '../actions/users'

export default class Profile extends React.Component {
    constructor() {
        super()
        this.state = {loading: true}
        this._updateFromStore = this._updateFromStore.bind(this)
    }

    componentDidMount() {
        UserStore.addChangeListener(this._updateFromStore)
        UsersActions.fetchProfile()
    }

    componentWillUnmount() {
        UserStore.removeChangeListener(this._updateFromStore)
    }

    render() {
        if (this.state.loading) {
            return (
                <Row>
                    <Columns className="large-offset-1 large-10 end">
                        <p className="text-center">Loading...</p>
                    </Columns>
                </Row>
            )
        }
        return (
            <Row>
                <Columns className="large-offset-1 large-10 end">
                    <Row className="profile">
                        <Columns className="large-12">
                            <Entry field="Name" value={`${this.state.first_name} ${this.state.last_name}`} />
                            <Entry field="Username" value={this.state.username} />
                            <Entry field="Email" value={this.state.email} />
                            <Entry field="Timezone" value={this.state.timezone_identifier} />
                        </Columns>
                    </Row>
                </Columns>
            </Row>
        )
    }

    _updateFromStore() {
        if (UserStore.success()) {
            let state = Object.assign({}, UserStore.getState().payload || {}, {loading: false})
            this.setState(state)
        }
    }
}
