import React from 'react'

import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

export default class Success extends React.Component {
    render() {
        return (
            <div className="centered-message">
                <Row>
                    <Columns className="medium-8 medium-offset-2 end columns">
                        <h2>Congrats!</h2>
                        <p>You've just created a new account! If our sign up system is open, you will receive an email message in order
                        to activate your newly created account. Otherwise, we will contact you when your account is ready to be activated.</p>
                        <p>Thank you for your interest in joining us!</p>
                        <p>Following is the secret code generator for the sign-in process. Please add it to one-time password providers,
                        like Google Authenticator, OTP Auth etc. Two-factor authentication is a mandatory step in order to sign-in.</p>
                        <p className="attention-points">
                            <img src={`data:image/png;base64,${this.props.codeSecretImage}`} />
                        </p>
                        <p>Also, if somehow you loose access to your one-time password provider, you must use the following recovery code:</p>
                        <p className="attention-points">
                            <span className="recovery-code">
                            {
                                Array.prototype.map.call(this.props.recoverSecret.split(/\s*-\s*/), (piece) => {
                                    return (<span key={piece} className="piece">{piece}</span>)
                                })
                            }
                            </span>
                        </p>
                        <p>We hope to make your account as secure as possible using these settings.</p>
                        <p>Kindly,</p>
                        <p>The QuatroLabs Team.</p>
                    </Columns>
                </Row>
            </div>
        );
    }
}
