import React from 'react';
import Store from '../Store.js';
import FlatButton from 'material-ui';

export default class Song extends React.Component {

  constructor(props, context) {
    super(props, context);
    this.state = {dialog: false, loading: true};
  }

  componentDidMount() {
    let d = Store.getById(this.props.params.songId);
    d.loading = false;
    this.setState(d);
  }

  get player() {
    if (!this.state.access) {
      return <div/>;
    }
    return (
      <span>{this.state.url}</span>
    );
  }

  get dialogLauncher() {
   if (this.state.access) {
      return <div/>;
    }
    return (
      <FlatButton label="Buy" primary={true} onClick={() => setState({dialog: true})} />
    );
  }

  render() {
    if (this.state.loading) {
      return <div />;
    }
    console.log(this.state);
    return (
      <div>
        <h1>{this.state.title}</h1>
        {this.player}
        {this.dialogLauncher}
      </div>
    );
  }

}
