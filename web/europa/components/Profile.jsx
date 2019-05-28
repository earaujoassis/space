import React from 'react'

import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'
import { Entry } from '../../core/components/Form.jsx'

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
                    <Columns className="small-offset-1 small-10 end">
                        <p className="text-center">Loading...</p>
                    </Columns>
                </Row>
            )
        }

        const { user } = this.state

        return (
            <Row>
                <Columns className="small-offset-1 small-10 end">
                    <Row className="profile">
                        <Columns className="small-12">
                            <Entry field="Name" value={`${user.first_name} ${user.last_name}`} />
                            <Entry field="Username" value={user.username} />
                            <Entry field="Email" value={user.email} />
                            <Entry field="Timezone" value={user.timezone_identifier} />
                            {!UserStore.isCurrentUserAdmin() && UserStore.isFeatureGateActive('user.adminify') && (
                                <Row className="profile-entry">
                                    <Columns className="columns small-11 small-offset-1 field">
                                        <input className="thin-input"
                                            ref={(r) => this.inputKey = r}
                                            type="text"
                                            name="application_key"
                                            placeholder="Application Key"
                                            required />
                                        <button className="button-anchor"
                                            onClick={() => UsersActions.adminify(this.inputKey.value) }>
                                            Make me an admin
                                        </button>
                                    </Columns>
                                </Row>
                            )}
                            {UserStore.isCurrentUserAdmin() && (
                                <Entry field="Role" value="Administrator" />
                            )}
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
