import React from 'react';

import Row from './reusable/row.jsx';
import Columns from './reusable/columns.jsx';

const StepsOrder = ['access', 'secrecy', 'code'];
const StepsData = {
    access: {
        className: "unknown",
        type: "text",
        placeholder: "Access holder"
    },
    secrecy: {
        className: "secrecy",
        type: "password",
        placeholder: "Passphrase"
    },
    code: {
        className: "code",
        type: "text",
        placeholder: "Code"

    }
};

export default class Login extends React.Component {
    constructor() {
        super();
        this.state = {
              currentStepIndex: 0
        }
        this._updateStep = this._updateStep.bind(this);
    }

    static get propTypes() {
        return {
            afterLogin: React.PropTypes.func
        }
    }

    render() {
        let step = StepsData[StepsOrder[this.state.currentStepIndex]];
        return (
            <div className="login-content">
                <Row>
                    <Columns className="large-12">
                        <div className={`user-avatar ${step.className}`}></div>
                        <form action="" method="POST">
                            <input type={step.type} placeholder={step.placeholder} />
                            <button name="" className="expand" onClick={this._updateStep}>Continue</button>
                        </form>
                        <p className="upper-box">2<sub>min</sub> to attempt a login</p>
                    </Columns>
                </Row>
            </div>
        );
    }

    _updateStep(e) {
        if (e) e.preventDefault();
        if (this.state.currentStepIndex + 1 < 3) {
            this.setState({currentStepIndex: this.state.currentStepIndex + 1});
        } else {
            this.props.afterLogin(1); // TODO implement user/login callback
        }
    }
};
