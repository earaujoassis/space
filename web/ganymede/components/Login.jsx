import React from 'react'
import ReactDOM from 'react-dom'

import SessionStore from '../stores/sessions'
import SessionsActions from '../actions/sessions'
import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

const StepsOrder = ['access', 'secrecy', 'code']
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
        name: "passcode",
        type: "text",
        placeholder: "Code"

    }
}

export default class Login extends React.Component {
    constructor() {
        super()
        this.state = {
              currentStepIndex: 0,
              disableSubmit: false,
              form: {}
        }
        this._updateStep = this._updateStep.bind(this)
        this._updateStepValue = this._updateStepValue.bind(this)
    }

    render() {
        let step = StepsData[StepsOrder[this.state.currentStepIndex]]
        return (
            <div className="login-content">
                <Row>
                    <Columns className="large-12">
                        <div className={`user-avatar ${step.className}`}></div>
                        <form action="." method="post">
                            <input ref="input" type={step.type}
                                name={step.name}
                                placeholder={step.placeholder}
                                value={this.state.form[step.name]}
                                onChange={this._updateStepValue}
                                required={true}
                                disabled={this.state.disableSubmit} />
                            <button type="submit" className="button expand"
                                disabled={this.state.disableSubmit}
                                onClick={this._updateStep}>Continue</button>
                        </form>
                        <p className="upper-box">2<sub>min</sub> to attempt a login</p>
                    </Columns>
                </Row>
            </div>
        )
    }

    _updateStepValue(e) {
        let form = this.state.form
        form[e.target.name] = e.target.value
        this.setState({form: form})
    }

    _updateStep(e) {
        if (e) e.preventDefault()
        let state = {}
        let form = this.state.form
        let name = StepsData[StepsOrder[this.state.currentStepIndex]].name
        form[name] = ReactDOM.findDOMNode(this.refs.input).value
        if (this.state.currentStepIndex + 1 < 3) {
            state = Object.assign({}, this.state, {form: form, currentStepIndex: this.state.currentStepIndex + 1})
        } else {
            state = Object.assign({}, this.state, {form: form, disableSubmit: true})
            let formData = new FormData()
            Array.prototype.forEach.call(StepsOrder, (stepKey) => {
                let name = StepsData[stepKey].name
                formData.append(name, form[name])
            })
            SessionsActions.signIn(formData)
        }
        this.setState(state)
    }
}
