import React from 'react';
// Needed for React Developer Tools
window.React = React;

import {render} from 'react-dom';
import {Router, Route, Link, IndexRoute} from 'react-router';

import Master from './Master';
import List from './List';
import Song from './song';
import history from './history';

render((
  <Router history={history} >
    <Route path="/" component={Master}>
      <IndexRoute component={List}/>
      <Route path="/song/:songId" component={Song}/>
    </Route>
  </Router>
),
document.getElementById('container'));
