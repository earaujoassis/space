import React from 'react';

import Row from '../../core/components/row.jsx';
import Columns from '../../core/components/columns.jsx';

const StepsOrder = ['access', 'secrecy', 'code'];
const StepsData = {
    access: {
        className: "unknown",
        name: "holder",
        type: "text",
        placeholder: "Access holder"
    },
    secrecy: {
        className: "secrecy",
        name: "password",
        type: "password",
        placeholder: "Passphrase"
    },
    code: {
        className: "code",
        name: "code",
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
                        <form action="." method="post">
                            <input type={step.type} name={step.name} placeholder={step.placeholder} />
                            <button type="submit" className="button expand" onClick={this._updateStep}>Continue</button>
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
