import React from 'react'

//import UsersActions from '../actions/users'
import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'
import PasswordInput from '../../core/components/PasswordInput.jsx'

const UpdatePassword = () => {
    return (
        <div className="middle-box plain resource-owner-password">
            <Row>
                <Columns className="small-12">
                    <form className="form-common" action="." method="post">
                        <p>Update your password with the required fields below</p>
                        <PasswordInput placeholder="New password" name="password_update" />
                        <PasswordInput placeholder="Confirm password" name="password_confirmation" />
                        <button type="submit"
                            className="button expand"
                            onClick={(e) => {
                                if (e) e.preventDefault()
                            }}
                            disabled={false}>Update password</button>
                    </form>
                </Columns>
            </Row>
        </div>
    )
}

const ownerUpdate = () => {
    return (
        <div className="resource-owner-update">
            <Row>
                <Columns className="small-12">
                    <UpdatePassword />
                </Columns>
            </Row>
        </div>
    )
}

export default ownerUpdate
