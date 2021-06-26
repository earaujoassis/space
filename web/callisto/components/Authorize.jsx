import React from 'react';

import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';

export default class Authorize extends React.Component {
  constructor () {
    super();
    this.state = this._loadData();
  }

  render () {
    let keyNumber = 0;
    return (
            <div className="authorize">
                <Row className="separator">
                    <Columns className="small-offset-2 small-4 user">
                        <h2>{`${this.state.first_name} ${this.state.last_name}`}</h2>
                    </Columns>
                    <Columns className="small-4 end client">
                        <h2>{this.state.client_name}</h2>
                    </Columns>
                </Row>
                <Row className="separator">
                    <Columns className="small-offset-2 small-8 end">
                        <p>The application ({this.state.client_name}) is requesting access to the following information:</p>
                        <ul className="">
                            {
                                Array.prototype.map.call(this._requestedData(this.state.requested_scope), (message) => {
                                  return (<li key={keyNumber++}>{message}</li>);
                                })
                            }
                        </ul>
                        <p>{this.state.client_name} (the application) and Space (the current website) will use this information
                        in accordance with their respective terms of service and privacy policies.</p>
                    </Columns>
                </Row>
                <Row className="separator">
                    <Columns className="small-12 text-center">
                        <Row>
                            <Columns className="small-offset-2 small-4">
                                <form action="" method="post">
                                    <input type="hidden" name="access_denied" value="true" />
                                    <button className="button expand secondary" type="submit">Cancel</button>
                                </form>
                            </Columns>
                            <Columns className="small-4 end">
                                <form action="" method="post">
                                    <button className="button expand" type="submit">Accept</button>
                                </form>
                            </Columns>
                        </Row>
                    </Columns>
                </Row>
            </div>
    );
  }

  _requestedData (scope) {
    switch (scope) {
      case 'public':
        return [
          'Authentication data for that given application'
        ];
      case 'read':
        return [
          'Authentication data for that given application',
          'Your profile data (including e-mail and first and last names)'
        ];
      case 'read_write':
        return [
          'Authentication data for that given application',
          'Your profile data (including e-mail and first and last names)',
          'Update your profile data (including e-mail and first and last names)'
        ];
    }
  }

  _loadData () {
    if (document.getElementById('data')) {
      const data = JSON.parse(document.getElementById('data').innerHTML);
      return data;
    }
    return {};
  }
}
