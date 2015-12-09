import React from 'react';

export default class Song extends React.Component {

  componentDidMount() {
    console.log(this.props.params.songId);
    // check access. if no access, show dialog.
  }

  render() {
    return (
      <div>foo</div>
    );
  }

}
