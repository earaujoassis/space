import React from 'react';

import SessionStore from '../stores/sessions';
import SessionsActions from '../actions/sessions';
import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';

import { getParameterByName } from '../../core/utils/url';

const StepsOrder = ['access', 'secrecy', 'code'];
const StepsData = {
  access: {
    className: 'unknown',
    name: 'holder',
    type: 'text',
    placeholder: 'Access holder'
  },
  secrecy: {
    className: 'secrecy',
    name: 'password',
    type: 'password',
    placeholder: 'Passphrase'
  },
  code: {
    className: 'code',
    name: 'passcode',
    type: 'text',
    placeholder: 'Code'

  }
};

export default class SignIn extends React.Component {
  constructor () {
    super();
    this.state = {
      currentStepIndex: 0,
      disableSubmit: false,
      form: this._initialForm()
    };
    this._updateStep = this._updateStep.bind(this);
    this._updateStepValue = this._updateStepValue.bind(this);
    this._updateFromStore = this._updateFromStore.bind(this);
    this._setFormSubmitTimeout = this._setFormSubmitTimeout.bind(this);
    this._triggerFormLock = this._triggerFormLock.bind(this);
    this._triggerFormUnlock = this._triggerFormUnlock.bind(this);
  }

  componentDidMount () {
    SessionStore.addChangeListener(this._updateFromStore);
    this._triggerFormLock();
  }

  componentWillUnmount () {
    SessionStore.removeChangeListener(this._updateFromStore);
    clearTimeout(this.securityTimeoutID);
  }

  render () {
    const step = StepsData[StepsOrder[this.state.currentStepIndex]];
    return (
            <div className="middle-box signin-content">
                <Row>
                    <Columns className="small-12">
                        <div className={`user-avatar ${step.className}`}></div>
                        {
                            this.state.failed
                              ? (
                                <p className="error-message">Authentication failed</p>
                                )
                              : null
                        }
                        {
                            this.state.blocked
                              ? (
                                <p className="error-message">Authentication blocked for user</p>
                                )
                              : null
                        }
                        <form action="." method="post">
                            <input ref={(r) => { this.inputRef = r; }} type={step.type}
                                name={step.name}
                                placeholder={step.placeholder}
                                value={this.state.form[step.name]}
                                onChange={this._updateStepValue}
                                required
                                disabled={this.state.disableSubmit} />
                            <button type="submit" className="button expand"
                                disabled={this.state.disableSubmit}
                                onClick={this._updateStep}>Continue</button>
                        </form>
                        <p className="upper-box">2<sub>min</sub> to attempt a sign-in</p>
                    </Columns>
                </Row>
            </div>
    );
  }

  _setFormSubmitTimeout (bool, delay) {
    const value = bool;
    this.securityTimeoutID = setTimeout(() => {
      const state = { disableSubmit: value };
      if (!value) {
        state.currentStepIndex = 0;
        state.failed = undefined;
      }
      this.setState(state);
      this.securityTimeoutID = null;
    }, delay);
  }

  _triggerFormLock () {
    this._setFormSubmitTimeout(true, 2 * 60000);
  }

  _triggerFormUnlock () {
    this._setFormSubmitTimeout(false, 5 * 1000);
  }

  _updateStepValue (e) {
    const form = this.state.form;
    form[e.target.name] = e.target.value;
    this.setState({ form: form });
  }

  _initialForm () {
    return { holder: '', password: '', passcode: '' };
  }

  _updateStep (e) {
    if (e) e.preventDefault();
    let state = {};
    const form = this.state.form;
    const name = StepsData[StepsOrder[this.state.currentStepIndex]].name;
    form[name] = this.inputRef.value;
    if (this.state.currentStepIndex + 1 < 3) {
      state = Object.assign({}, this.state, { form: form, currentStepIndex: this.state.currentStepIndex + 1 });
    } else {
      state = Object.assign({}, this.state, { form: form, disableSubmit: true });
      const formData = new FormData();
      Array.prototype.forEach.call(StepsOrder, (stepKey) => {
        const name = StepsData[stepKey].name;
        formData.append(name, form[name]);
      });
      SessionsActions.signIn(formData);
    }
    this.setState(state);
  }

  _updateFromStore () {
    if (SessionStore.success()) {
      const r = SessionStore.getState().payload || {};
      const next = getParameterByName('_');
      let location = `${r.redirect_uri}?client_id=${r.client_id}&code=${r.code}&grant_type=${r.grant_type}&scope=${r.scope}&state=${r.state}`;
      if (next && next) {
        location += `&_=${encodeURI(next)}`;
      }
      window.location.href = location;
    } else {
      const error = SessionStore.getState().payload;
      if (error.attempts === 'blocked') {
        this.setState({ blocked: true, form: this._initialForm(), disableSubmit: true, currentStepIndex: 0 });
        return;
      }
      this.setState({ failed: true, form: this._initialForm() });
      this._triggerFormUnlock();
    }
  }
}
