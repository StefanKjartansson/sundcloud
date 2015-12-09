import React from 'react';

import {GridList, GridTile} from 'material-ui';
import history from './history';
import Actions from './Actions';
import Store from './Store';


export default class SongList extends React.Component {

  constructor(props, context) {
    super(props, context);
    this.state = {
    };
  }

  _getState() {
    return {
      items: Store.data,
    };
  }

  _onChange() {
    this.setState(this._getState());
  }

  componentDidMount() {
    this.changeListener = this._onChange.bind(this);
    Store.addChangeListener(this.changeListener);
    Actions.refresh();
    this.setState(this._getState());
  }

  render() {
    if (!this.state.items) {
      return <div>Loading</div>;
    }
    return (
      <GridList>
        {
          this.state.items.map(tile => <GridTile
            key={tile.id}
            title={tile.title}
            subtitle={<span>by <b>{tile.author}</b></span>}
            onClick={() => {
              history.push(`/song/${tile.id}`);
             }}><img src={tile.img} /></GridTile>)
        }
      </GridList>
    );
  }

};
