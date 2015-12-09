import React from 'react';

const ThemeManager = require('material-ui/lib/styles/theme-manager');
const LightRawTheme = require('material-ui/lib/styles/raw-themes/light-raw-theme');


export default class App extends React.Component {

  getChildContext() {
    return {
      muiTheme: ThemeManager.getMuiTheme(LightRawTheme)
    };
  }

  render() {
    return (
      <div>
        <h1>Song smash</h1>
        <div className="master">
          {this.props.children}
        </div>
      </div>
    )
  }

};

App.childContextTypes = {
    muiTheme: React.PropTypes.object
};
